package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"strings"
	"time"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) (string, error) {

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var userId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                email,
			                first_name,
			                last_name,
			                password,
							role_id,
							university_id,
							created_at
			                )
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
			`, usersTable)

	err = p.Pool.QueryRow(ctx, query, u.Email, u.FirstName, u.LastName, u.Password, u.RoleId, u.UniversityId, time.Now()).Scan(&userId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return userId, tx.Commit(ctx)
}

func (p *Postgres) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	user := new(entity.User)

	query := fmt.Sprintf("SELECT id, email, first_name, last_name, password FROM %s WHERE email=$1", usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) GetUserById(ctx context.Context, id string) (*entity.User, error) {

	user := new(entity.User)

	query := fmt.Sprintf("SELECT id, email, first_name, last_name, password FROM %s WHERE id=$1", usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, id string, user *api.UpdateUserRequest) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if user.FirstName != "" {
		values = append(values, fmt.Sprintf("first_name=$%d", paramCount))
		params = append(params, user.FirstName)
		paramCount++
	}
	if user.LastName != "" {
		values = append(values, fmt.Sprintf("last_name=$%d", paramCount))
		params = append(params, user.LastName)
		paramCount++
	}
	if user.Email != "" {
		// change to user id for correctness
		values = append(values, fmt.Sprintf("email=$%d", paramCount))
		params = append(params, user.Email)
		paramCount++
	}
	if user.Password != "" {
		values = append(values, fmt.Sprintf("password=$%d", paramCount))
		params = append(params, user.Password)
		paramCount++
	}
	if user.ImagePath != "" {
		values = append(values, fmt.Sprintf("image_path=$%d", paramCount))
		params = append(params, user.ImagePath)
		paramCount++
	}
	if user.Phone != "" {
		values = append(values, fmt.Sprintf("phone=$%d", paramCount))
		params = append(params, user.Phone)
		paramCount++
	}
	if user.UniversityId != "" {
		values = append(values, fmt.Sprintf("university_id=$%d", paramCount))
		params = append(params, user.UniversityId)
		paramCount++
	}

	if user.RoleId != 0 {
		values = append(values, fmt.Sprintf("role_id=$%d", paramCount))
		params = append(params, user.RoleId)
		paramCount++
	}

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", usersTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpgradeUser(ctx context.Context, id string, user *api.UpdateUserRequest) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if user.UniversityId != "" {
		values = append(values, fmt.Sprintf("university_id=$%d", paramCount))
		params = append(params, user.UniversityId)
		paramCount++
	}

	if user.RoleId != 0 {
		values = append(values, fmt.Sprintf("role_id=$%d", paramCount))
		params = append(params, user.RoleId)
		paramCount++
	}
	values = append(values, fmt.Sprintf("verified=$%d", paramCount))
	params = append(params, user.Verified)

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", usersTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
