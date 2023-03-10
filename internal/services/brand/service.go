package brand

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/vageeshabr/brands-api/internal/models"
	"github.com/vageeshabr/brands-api/internal/stores"
	"net/http"
)

type Service struct {
	store stores.BrandStorer
}

func (s *Service) GetAll(ctx *gofr.Context, name string) ([]*models.Brand, error) {
	if len(name) < 3 {
		return nil, errors.InvalidParam{Param: []string{"name"}}
	}

	return s.store.Find(ctx, name)
}

func (s *Service) Add(ctx *gofr.Context, name string) (*models.Brand, error) {

	if len(name) < 3 {
		return nil, errors.InvalidParam{Param: []string{"name"}}
	}

	if v, err := s.store.Exists(ctx, name); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Reason: "failed to check if brand exists"}
	} else if v {
		return nil, errors.InvalidParam{Param: []string{"name"}}
	}

	return s.store.Create(ctx, name)
}

func New(storer stores.BrandStorer) *Service {
	return &Service{store: storer}
}
