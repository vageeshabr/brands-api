package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	brandService "github.com/vageeshabr/brands-api/internal/services/brand"

	"github.com/vageeshabr/brands-api/internal/http/brand"

	brandStore "github.com/vageeshabr/brands-api/internal/stores/brand"
)

func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false

	brandStore := brandStore.New()
	brandSvc := brandService.New(brandStore)
	brandHTTP := brand.New(brandSvc)

	app.REST("brand", brandHTTP)

	app.Start()
}
