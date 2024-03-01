package tables

import (
	contx "context"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/uristemov/repeatPro/api"
	"log"
)

func (g *GoAdmin) GetUniversitiesTable(ctx *context.Context) table.Table {

	universities := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := universities.GetInfo().HideFilterArea()
	info.AddField("ID", "id", db.UUID)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Image_path", "image_path", db.Text)
	info.AddField("Created_at", "created_at", db.Timestamp).FieldFilterable().FieldSortable()

	info.SetTable("universities").SetTitle("Universities").SetDescription("There are all list of available universities.").
		SetDeleteFn(func(idArr []string) error {

			err := g.service.DeleteCourseById(contx.Background(), idArr[0])
			if err != nil {
				log.Printf("Delete course by id error: %v", err)
				return err
			}

			return nil
		})

	formList := universities.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate().FieldDisableWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Text, form.Text)
	formList.AddField("Image_path", "image_path", db.Text, form.Text)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("universities").SetTitle("Courses").SetDescription("Courses")

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

	return universities
}
