package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"strings"
)

func (p *Postgres) CreateImage(ctx context.Context, req *api.CreateImageRequest) (string, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var videoId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
							image_path,
							name,
							description,
							module_id
			                )
			VALUES ($1, $2, $3, $4) RETURNING id
			`, imagesTable)

	err = p.Pool.QueryRow(ctx, query, req.ImagePath, req.Name, req.Description, req.ModuleId).Scan(&videoId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return videoId, tx.Commit(ctx)
}

func (p *Postgres) GetAllImages(ctx context.Context) ([]entity.Image, error) {
	var images []entity.Image

	query := fmt.Sprintf("SELECT id, name, image_path, description, module_id, created_at FROM %s", imagesTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		image := entity.Image{}
		err = rows.Scan(&image.Id, &image.Name, &image.ImagePath, &image.Description, &image.ModuleId, &image.CreatedAt)
		images = append(images, image)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (p *Postgres) GetImageById(ctx context.Context, id string) (*entity.Image, error) {
	video := new(entity.Image)

	query := fmt.Sprintf("SELECT id, name, image_path, description, module_id, created_at FROM %s WHERE id=$1", imagesTable)

	err := pgxscan.Get(ctx, p.Pool, video, query, id)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (p *Postgres) DeleteImageById(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", imagesTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateImageById(ctx context.Context, req *api.UpdateImageRequest, id string) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.ImagePath != "" {
		values = append(values, fmt.Sprintf("image_path=$%d", paramCount))
		params = append(params, req.ImagePath)
		paramCount++
	}
	//if req.LogoImage != "" {
	//	values = append(values, fmt.Sprintf("logo_image=$%d", paramCount))
	//	params = append(params, req.LogoImage)
	//	paramCount++
	//}
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
	setQuery = fmt.Sprintf("UPDATE %s SET ", imagesTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil

}
