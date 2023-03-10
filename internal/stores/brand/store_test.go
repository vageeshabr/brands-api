package brand

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"github.com/vageeshabr/brands-api/internal/models"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	tcs := []struct {
		Name         string
		ExpectedRows []*models.Brand
		ExpectedErr  error
	}{
		{Name: "", ExpectedRows: nil, ExpectedErr: nil},
	}

	app := gofr.New()

	for i, tc := range tcs {
		store := Store{}
		req := httptest.NewRequest("GET", "/", nil)
		ctx := gofr.NewContext(nil, request.NewHTTPRequest(req), app)

		out, err := store.Find(ctx, tc.Name)

		if err != tc.ExpectedErr {
			t.Errorf("tc %d failed, expected err: %v, got %v", i, tc.ExpectedErr, err)
		}

		if !reflect.DeepEqual(out, tc.ExpectedRows) {
			t.Errorf("tc %d failed, expected: %v, got %v", i, tc.ExpectedRows, out)
		}
	}
}
