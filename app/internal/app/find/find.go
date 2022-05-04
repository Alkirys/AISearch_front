package auth

import (
	"context"
	"encoding/base64"
	"github.com/Elderly-AI/findface/internal/pkg/find"
	pb "github.com/Elderly-AI/findface/pkg/proto/find"
)

type Implementation struct {
	facade find.Facade
	pb.UnimplementedFindServer
}

func New(facade find.Facade) Implementation {
	return Implementation{
		facade: facade,
	}
}

type FindResponse struct {
	url       string
	boundings [][]uint64
}

func (i *Implementation) FindHandler(ctx context.Context, req *pb.FindHandlerRequest) (*pb.FindHandlerResponse, error) {
	imgRaw, err := base64.StdEncoding.DecodeString(req.Img)
	if err != nil {
		return nil, err
	}
	imgs, err := i.facade.Find(ctx, imgRaw)
	if err != nil {
		return nil, err
	}
	pbImgs := make([]*pb.FindHandlerResponseImage, len(imgs))
	for j, img := range imgs {
		pbImgs[j] = &pb.FindHandlerResponseImage{
			Name:  img.Name,
			Bound: img.Bound,
		}
	}
	return &pb.FindHandlerResponse{
		Imgs: pbImgs,
	}, nil
}

func (i *Implementation) DetectHandler(ctx context.Context, req *pb.DetectHandlerRequest) (*pb.DetectHandlerResponse, error) {
	urls, err := i.facade.Detect(ctx, req.Img.Name)
	if err != nil {
		return nil, err
	}
	return &pb.DetectHandlerResponse{
		Users: urls,
	}, nil
}
