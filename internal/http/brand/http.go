package brand

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/vageeshabr/brands-api/internal/models"
	"github.com/vageeshabr/brands-api/internal/services"
)

type Handler struct {
	svc services.Brand
}

func New(svc services.Brand) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Index(c *gofr.Context) (interface{}, error) {
	name := c.Param("name")

	res, err := h.svc.GetAll(c, name)
	if err != nil {
		return nil, err
	}

	var result = &struct {
		Brands []*models.Brand `json:"brands"`
	}{
		Brands: res,
	}
	return result, nil
}

func (h *Handler) Create(c *gofr.Context) (interface{}, error) {
	var body = struct {
		Name string `json:"name"`
	}{}

	if err := c.Bind(&body); err != nil {
		return nil, err
	}

	res, err := h.svc.Add(c, body.Name)
	if err != nil {
		return nil, err
	}

	var result = &struct {
		Brand *models.Brand `json:"brand"`
	}{
		Brand: res,
	}

	return result, nil
}
