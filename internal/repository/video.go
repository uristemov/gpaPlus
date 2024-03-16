package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"strings"
)

func (p *Postgres) CreateVideo(ctx context.Context, req *api.CreateVideoRequest) (string, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var videoId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
							video_path,
							name,
							description,
							module_id
			                )
			VALUES ($1, $2, $3, $4) RETURNING id
			`, videosTable)

	err = p.Pool.QueryRow(ctx, query, req.VideoPath, req.Name, req.Description, req.ModuleId).Scan(&videoId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return videoId, tx.Commit(ctx)
}

func (p *Postgres) GetAllVideos(ctx context.Context) ([]entity.Video, error) {
	var videos []entity.Video

	query := fmt.Sprintf("SELECT id, name, video_path, description, module_id, created_at FROM %s", videosTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		video := entity.Video{}
		err = rows.Scan(&video.Id, &video.Name, &video.VideoPath, &video.Description, &video.ModuleId, &video.CreatedAt)
		videos = append(videos, video)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (p *Postgres) GetVideoById(ctx context.Context, id string) (*entity.Video, error) {
	video := new(entity.Video)

	query := fmt.Sprintf("SELECT id, name, video_path, description, module_id, created_at FROM %s WHERE id=$1", videosTable)

	err := pgxscan.Get(ctx, p.Pool, video, query, id)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (p *Postgres) DeleteVideoById(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", videosTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateVideoById(ctx context.Context, req *api.UpdateVideoRequest, id string) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.VideoPath != "" {
		values = append(values, fmt.Sprintf("video_path=$%d", paramCount))
		params = append(params, req.VideoPath)
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
	setQuery = fmt.Sprintf("UPDATE %s SET ", videosTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil

}
