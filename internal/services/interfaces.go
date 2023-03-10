package services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/vageeshabr/brands-api/internal/models"
)

type Brand interface {
	GetAll(ctx *gofr.Context, name string) ([]*models.Brand, error)
	Add(ctx *gofr.Context, name string) (*models.Brand, error)
}
