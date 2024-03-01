package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetCoursesTable return the model of table course.
func GetCourseTable(ctx *context.Context) (coursesTable table.Table) {

	coursesTable = table.NewDefaultTable(table.DefaultConfigWithDriver(db.DriverPostgresql))

	// connect your custom connection

	info := coursesTable.GetInfo()
	info.AddField("ID", "id", db.UUID)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("Image_path", "image_path", db.Text)
	//info.AddField("Email", "email", db.Varchar)
	//info.AddField("Birthdate", "birthdate", db.Date)
	//info.AddField("Added", "added", , db.Timestamp, form.Datetime)

	info.AddButton("Articles", icon.Tv, action.PopUpWithIframe("/authors/list", "文章",
		action.IframeData{Src: "/admin/info/posts"}, "900px", "560px"))
	info.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	formList := coursesTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("First Name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last Name", "last_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Text)
	formList.AddField("Birthdate", "birthdate", db.Date, form.Text)
	formList.AddField("Added", "added", db.Timestamp, form.Text)

	formList.SetTable("authors").SetTitle("Authors").SetDescription("Authors")

	return
}
