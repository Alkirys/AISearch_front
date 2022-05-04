package find

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"github.com/Elderly-AI/findface/internal/pkg/model"
)

type Facade struct {
	db *redis.Client
}

func New(db *redis.Client) Facade {
	return Facade{
		db: db,
	}
}

func (f *Facade) Find(ctx context.Context, image []byte) ([]model.Img, error) {
	reqBody := bytes.NewBuffer(image)
	resp, err := http.Post("https://search4faces.com/upload.php", "image/jpeg", reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	find := model.Find{}
	err = json.Unmarshal(body, &find)
	if err != nil {
		return nil, err
	}
	imgs := make([]model.Img, len(find.Boundings))
	for i, bound := range find.Boundings {
		imgs[i] = model.Img{
			Name:  find.Url,
			Bound: bound,
		}
	}
	return f.Save(ctx, imgs)
}

func (f *Facade) Detect(ctx context.Context, name string) ([]string, error) {
	img, err := f.Get(ctx, name)
	if err != nil {
		return nil, err
	}
	req := model.Detect{
		Query:     "vkok",
		Filename:  img.Name,
		Boundings: img.Bound,
	}
	raw, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(raw)
	resp, err := http.Post("https://search4faces.com/detect.php", "text/plain;charset=UTF-8", reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	faces := model.Faces{}
	err = json.Unmarshal(body, &faces)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, face := range faces.Faces {
		if len(face) < 2 {
			continue
		}
		faceString, ok := face[1].(string)
		if ok && strings.HasPrefix(faceString, `https://vk.com`) {
			res = append(res, faceString)
		}
	}
	return res, nil
}

func (f *Facade) Save(ctx context.Context, imgs []model.Img) ([]model.Img, error) {
	var out []model.Img
	for _, img := range imgs {
		if len(img.Bound) >= 4 {
			uimg := model.Img{
				Name:  uuid.New().String(),
				Bound: img.Bound[0:4],
			}
			raw, err := json.Marshal(img)
			if err != nil {
				return nil, err
			}
			f.db.Append(ctx, uimg.Name, string(raw))
			out = append(out, uimg)
		}
	}
	return out, nil
}

func (f *Facade) Get(ctx context.Context, name string) (img model.Img, err error) {
	raw, err := f.db.Get(ctx, name).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(raw), &img)
	if err != nil {
		return
	}
	return
}
