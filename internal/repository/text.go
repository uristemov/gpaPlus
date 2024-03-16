package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"strings"
)

func (p *Postgres) CreateText(ctx context.Context, req *api.CreateTextRequest) (string, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var textId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
							name,
							description,
							module_id
			                )
			VALUES ($1, $2, $3) RETURNING id
			`, textsTable)

	err = p.Pool.QueryRow(ctx, query, req.Name, req.Description, req.ModuleId).Scan(&textId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return textId, tx.Commit(ctx)
}

func (p *Postgres) GetAllTexts(ctx context.Context) ([]entity.Text, error) {
	var texts []entity.Text

	query := fmt.Sprintf("SELECT id, name, description, module_id, created_at FROM %s", textsTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		text := entity.Text{}
		err = rows.Scan(&text.Id, &text.Name, &text.Description, &text.ModuleId, &text.CreatedAt)
		texts = append(texts, text)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return texts, nil
}

func (p *Postgres) GetTextById(ctx context.Context, id string) (*entity.Text, error) {
	text := new(entity.Text)

	query := fmt.Sprintf("SELECT id, name, description, module_id, created_at FROM %s WHERE id=$1", textsTable)

	err := pgxscan.Get(ctx, p.Pool, text, query, id)
	if err != nil {
		return nil, err
	}

	return text, nil
}

func (p *Postgres) DeleteTextById(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", textsTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateTextById(ctx context.Context, req *api.UpdateTextRequest, id string) error {
	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.Name != "" {
		values = append(values, fmt.Sprintf("name=$%d", paramCount))
		params = append(params, req.Name)
		paramCount++
	}
	if req.Description != "" {
		values = append(values, fmt.Sprintf("description=$%d", paramCount))
		params = append(params, req.Description)
	}

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", textsTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil
}
