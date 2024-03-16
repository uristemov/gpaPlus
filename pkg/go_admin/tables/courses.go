package tables

import (
	contx "context"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/service"
	"log"
)

type GoAdmin struct {
	*plugins.Base
	//conn    db.Connection
	service service.Service
}

func New(service service.Service) *GoAdmin {
	return &GoAdmin{service: service}
}

func interfaces(arr []string) []interface{} {
	var iarr = make([]interface{}, len(arr))

	for key, v := range arr {
		iarr[key] = v
	}

	return iarr
}

func (g *GoAdmin) GetCoursesTable(ctx *context.Context) table.Table {

	cfg := table.DefaultConfigWithDriver("postgresql")
	cfg.PrimaryKey.Type = db.UUID

	courses := table.NewDefaultTable(cfg)

	info := courses.GetInfo().HideFilterArea()
	info.AddField("ID", "id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("Image_path", "image_path", db.Text)
	info.AddField("Price", "price", db.Int)
	info.AddField("Rating", "rating", db.Float4)
	info.AddField("AuthorID", "user_id", db.UUID)
	info.AddField("Created_at", "created_at", db.Datetime).FieldFilterable().FieldSortable()

	info.SetTable("courses").SetTitle("Courses").SetDescription("Courses").
		SetDeleteFn(func(idArr []string) error {

			err := g.service.DeleteCourseById(contx.Background(), idArr[0])
			if err != nil {
				log.Printf("Delete course by id error: %v", err)
				return err
			}

			return nil
		})

	formList := courses.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Text, form.Text)
	formList.AddField("Image_path", "image_path", db.Text, form.Text)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()

	formList.SetTable("courses").SetTitle("Courses").SetDescription("Courses")

	formList.SetUpdateFn(func(values form2.Values) error {

		var req api.UpdateCourseRequest

		fmt.Println("Name values get from: ", values.Get("name"))
		if values.Get("name") != "" {
			req.Name = values.Get("name")
		}
		if values.Get("image_path") != "" {
			req.ImagePath = values.Get("image_path")
		}
		if values.Get("description") != "" {
			req.Description = values.Get("description")
		}

		err := g.service.UpdateCourseById(contx.Background(), &req, values.Get("id"))
		if err != nil {
			log.Printf("update course error: %v", err)
			return err
		}

		return nil
	})

	return courses
}
