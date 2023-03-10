package brand

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/vageeshabr/brands-api/internal/models"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) Exists(ctx *gofr.Context, name string) (exists bool, err error) {
	row := ctx.DB().QueryRowContext(ctx, "select count(*) from brands where name = ?", name)
	count := 0
	if err := row.Scan(&count); err != nil {
		return false, errors.DB{Err: err}
	}
	return count > 0, nil
}

func (s *Store) Find(ctx *gofr.Context, name string) (res []*models.Brand, err error) {
	rows, err := ctx.DB().QueryContext(ctx, "select id, name from brands where name like ?", name)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var m models.Brand
		if err := rows.Scan(&m.Id, &m.Name); err != nil {
			return nil, errors.DB{Err: err}
		}

		res = append(res, &m)
	}

	return res, nil
}

func (s *Store) Create(ctx *gofr.Context, name string) (*models.Brand, error) {
	res, err := ctx.DB().ExecContext(ctx, "insert into brands(name) values(?)", name)
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return &models.Brand{
		Id:   int(id),
		Name: name,
	}, err
}
