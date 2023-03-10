package stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/vageeshabr/brands-api/internal/models"
)

type BrandStorer interface {
	Find(ctx *gofr.Context, name string) ([]*models.Brand, error)
	Create(ctx *gofr.Context, name string) (*models.Brand, error)
	Exists(ctx *gofr.Context, name string) (bool, error)
}
