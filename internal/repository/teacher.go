package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (p *Postgres) GetAllTeachers(ctx context.Context) ([]entity.Teacher, error) {
	var teachers []entity.Teacher

	query := fmt.Sprintf("SELECT id, name, image_path, description, phone, created_at FROM %s", teacherTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		teacher := entity.Teacher{}
		err = rows.Scan(&teacher.Id, &teacher.Name, &teacher.ImagePath, &teacher.Description, &teacher.Phone, &teacher.CreatedAt)
		teachers = append(teachers, teacher)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (p *Postgres) GetTeacherById(ctx context.Context, id string) (*entity.Teacher, error) {
	teacher := new(entity.Teacher)

	query := fmt.Sprintf("SELECT id, name, image_path, description, phone, created_at FROM %s WHERE id=$1", teacherTable)

	err := pgxscan.Get(ctx, p.Pool, teacher, query, id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}
