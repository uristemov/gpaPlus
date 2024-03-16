package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (p *Postgres) GetAllTeacherRequests(ctx context.Context, id string) ([]entity.Request, error) {
	var requests []entity.Request

	query := fmt.Sprintf("SELECT id, course_id, user_id, active, created_at FROM %s WHERE course_id in (SELECT id FROM courses WHERE user_id=$1)", requestsTable)

	rows, err := p.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		request := entity.Request{}
		err = rows.Scan(&request.Id, &request.CourseId, &request.UserId, &request.IsActive, &request.CreatedAt)
		requests = append(requests, request)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (p *Postgres) CreateRequest(ctx context.Context, req *api.CreateRequest) (string, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var courseId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                user_id,
							course_id,
							active
			                )
			VALUES ($1, $2, $3) RETURNING id
			`, requestsTable)

	err = p.Pool.QueryRow(ctx, query, req.UserId, req.CourseId, req.IsActive).Scan(&courseId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return courseId, tx.Commit(ctx)
}

func (p *Postgres) DeleteRequestById(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", requestsTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetRequestById(ctx context.Context, id string) (*entity.Request, error) {
	request := new(entity.Request)

	query := fmt.Sprintf("SELECT id, user_id, course_id, created_at FROM %s WHERE id=$1", requestsTable)

	err := pgxscan.Get(ctx, p.Pool, request, query, id)
	if err != nil {
		return nil, err
	}

	return request, nil
}
