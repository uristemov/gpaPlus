package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"strings"
)

func (p *Postgres) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	var courses []entity.Course

	query := fmt.Sprintf("SELECT id, name, image_path, description, created_at FROM %s", courseTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		course := entity.Course{}
		err = rows.Scan(&course.Id, &course.Name, &course.ImagePath, &course.Description, &course.CreatedAt)
		courses = append(courses, course)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (p *Postgres) GetCourseById(ctx context.Context, id string) (*entity.Course, error) {
	course := new(entity.Course)

	query := fmt.Sprintf("SELECT id, name, image_path, description, created_at FROM %s WHERE id=$1", courseTable)

	err := pgxscan.Get(ctx, p.Pool, course, query, id)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (p *Postgres) DeleteCourseById(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", courseTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateCourseById(ctx context.Context, req *api.UpdateCourseRequest, id string) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.Name != "" {
		values = append(values, fmt.Sprintf("name=$%d", paramCount))
		params = append(params, req.Name)
		paramCount++
	}
	if req.ImagePath != "" {
		values = append(values, fmt.Sprintf("image_path=$%d", paramCount))
		params = append(params, req.ImagePath)
		paramCount++
	}
	if req.Description != "" {
		values = append(values, fmt.Sprintf("description=$%d", paramCount))
		params = append(params, req.Description)
	}

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", courseTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil

}
