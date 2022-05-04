package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/go-redis/redis/v8"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	findImpl "github.com/Elderly-AI/findface/internal/app/find"
	findFacade "github.com/Elderly-AI/findface/internal/pkg/find"
	pbFind "github.com/Elderly-AI/findface/pkg/proto/find"
)

func registerServices(opts Options, s *grpc.Server) {
	findF := findFacade.New(opts.RedisConnection)
	findI := findImpl.New(findF)
	pbFind.RegisterFindServer(s, &findI)
}

func newGateway(ctx context.Context, conn *grpc.ClientConn, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)

	for _, f := range []func(ctx context.Context, mux *gwruntime.ServeMux, conn *grpc.ClientConn) error{
		pbFind.RegisterFindHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}

type Options struct {
	Addr            string
	RedisConnection *redis.Client
	Mux             []gwruntime.ServeMuxOption
}

func createInitialOptions() Options {
	opts := Options{}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	opts.RedisConnection = redisClient

	opts.Addr = "0.0.0.0:8080"
	return opts
}

func main() {
	opts := createInitialOptions()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	registerServices(opts, s)
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	// register services

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		opts.Addr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gw, err := newGateway(context.Background(), conn, opts.Mux)
	if err != nil {
		log.Fatalln("Failed to init gateway:", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", gw)

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
