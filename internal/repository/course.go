package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"strings"
)

func (p *Postgres) CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (string, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var courseId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                name,
							image_path,
							description,
							user_id,
							price
			                )
			VALUES ($1, $2, $3, $4, $5) RETURNING id
			`, courseTable)

	err = p.Pool.QueryRow(ctx, query, req.Name, req.ImagePath, req.Description, req.UserId, req.Price).Scan(&courseId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return courseId, tx.Commit(ctx)
}

func (p *Postgres) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	var courses []entity.Course

	query := fmt.Sprintf("SELECT id, name, image_path, description, price, rating, user_id, created_at FROM %s", courseTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		course := entity.Course{}
		err = rows.Scan(&course.Id, &course.Name, &course.ImagePath, &course.Description, &course.Price, &course.Rating, &course.UserId, &course.CreatedAt)
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

	query := fmt.Sprintf("SELECT id, name, image_path, description, price, rating, user_id, created_at FROM %s WHERE id=$1", courseTable)

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
		paramCount++
	}
	if req.Price != 0 {
		values = append(values, fmt.Sprintf("price=$%d", paramCount))
		params = append(params, req.Price)
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

func (p *Postgres) GetAllTeacherCourses(ctx context.Context, id string) ([]entity.Course, error) {
	var courses []entity.Course

	query := fmt.Sprintf("SELECT id, name, image_path, description, price, rating, user_id, created_at FROM %s WHERE user_id=$1", courseTable)

	rows, err := p.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		course := entity.Course{}
		err = rows.Scan(&course.Id, &course.Name, &course.ImagePath, &course.Description, &course.Price, &course.Rating, &course.UserId, &course.CreatedAt)
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

func (p *Postgres) AddStudentToCourse(ctx context.Context, userId, courseId string) error {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
			INSERT INTO %s (
							user_id,
							course_id
			                )
			VALUES ($1, $2)
			`, usersCoursesTable)

	_, err = p.Pool.Exec(ctx, query, userId, courseId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (p *Postgres) GetAllCourseStudents(ctx context.Context, userId, courseId string) ([]api.GetStudentsResponse, error) {
	var students []api.GetStudentsResponse

	query := fmt.Sprintf("SELECT users.id, users.email, users.first_name, users.last_name, roles.role_name, users.image_path FROM %s JOIN goadmin_roles roles ON users.role_id = roles.id WHERE users.user_id IN (SELECT user_id FROM users_courses WHERE course_id=$1);", usersTable)

	rows, err := p.Pool.Query(ctx, query, courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		student := api.GetStudentsResponse{}
		err = rows.Scan(&student.Id, &student.Email, &student.FirstName, &student.LastName, &student.Role, &student.ImagePath)
		students = append(students, student)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return students, nil
}
