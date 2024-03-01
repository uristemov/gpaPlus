package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"strings"
)

func (p *Postgres) GetAllCourseModules(ctx context.Context, courseId string) ([]entity.Module, error) {
	var modules []entity.Module

	query := fmt.Sprintf("SELECT id, name, course_id FROM %s WHERE course_id=$1", modulesTable)

	rows, err := p.Pool.Query(ctx, query, courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		module := entity.Module{}
		err = rows.Scan(&module.Id, &module.Name, &module.CourseId)
		modules = append(modules, module)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (p *Postgres) GetModuleById(ctx context.Context, id string) (*entity.Module, error) {
	module := new(entity.Module)

	query := fmt.Sprintf("SELECT id, name, course_id FROM %s WHERE id=$1", modulesTable)

	err := pgxscan.Get(ctx, p.Pool, module, query, id)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (p *Postgres) DeleteModuleById(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", modulesTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateModuleById(ctx context.Context, req *api.UpdateModuleRequest, id string) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.Name != "" {
		values = append(values, fmt.Sprintf("name=$%d", paramCount))
		params = append(params, req.Name)
		paramCount++
	}

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", modulesTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil

}
