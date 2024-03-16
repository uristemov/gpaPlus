package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/internal/entity"
)

func (p *Postgres) GetAllTeachers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	query := fmt.Sprintf("SELECT id, first_name, last_name, image_path, phone, created_at FROM %s WHERE role_id=3", usersTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.ImagePath, &user.Phone, &user.CreatedAt)
		users = append(users, user)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *Postgres) GetTeacherById(ctx context.Context, id string) (*entity.User, error) {
	teacher := new(entity.User)

	query := fmt.Sprintf("SELECT id, first_name, last_name, image_path, phone, created_at FROM %s WHERE role_id=3 AND id=$1", usersTable)

	err := pgxscan.Get(ctx, p.Pool, teacher, query, id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}
