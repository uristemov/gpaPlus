package tables

import (
	contx "context"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/uristemov/repeatPro/api"
	"log"
	"strconv"
)

func (g *GoAdmin) GetUsersTable(ctx *context.Context) table.Table {

	cfg := table.DefaultConfigWithDriver("postgresql").SetCanAdd(false)
	cfg.PrimaryKey.Type = db.UUID

	//users := table.NewDefaultTable()
	users := table.NewDefaultTable(cfg)

	info := users.GetInfo().HideFilterArea()
	info.AddField("ID", "id", db.UUID)
	info.AddField("FirstName", "first_name", db.Varchar)
	info.AddField("LastName", "last_name", db.Varchar)
	info.AddField("Email", "email", db.Varchar).FieldFilterable()
	info.AddField("Image_path", "image_path", db.Text)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Role", "name", db.Varchar).
		FieldJoin(types.Join{
			Table:     "goadmin_roles",
			JoinField: "id",
			Field:     "role_id",
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labels := template.HTML("")
			labelTpl := template.Get(config.GetTheme()).Label().SetType("success")
			labels += labelTpl.SetContent(template.HTML(model.Value)).GetContent()

			if labels == template.HTML("") {
				return language.Get("no roles")
			}

			return labels
		}).FieldFilterable()
	info.AddField("University", "name", db.Text).
		FieldJoin(types.Join{
			Table:     "universities",
			JoinField: "id",
			Field:     "university_id",
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labels := template.HTML("")
			labelTpl := template.Get(config.GetTheme()).Label().SetType("primary")
			labels += labelTpl.SetContent(template.HTML(model.Value)).GetContent()

			if labels == template.HTML("") {
				return language.Get("no roles")
			}

			return labels
		}).FieldFilterable()
	info.AddField("Verified", "verified", db.Boolean)
	info.AddField("Created_at", "created_at", db.Datetime).FieldFilterable().FieldSortable()

	info.SetTable("users").SetTitle("Users").SetDescription("There are all list of available users.").
		SetDeleteFn(func(idArr []string) error {

			err := g.service.DeleteUser(contx.Background(), idArr[0])
			if err != nil {
				log.Printf("Delete user by id error: %v", err)
				return err
			}

			return nil
		})

	formList := users.GetForm()

	formList.AddField("ID", "id", db.UUID, form.Text).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("FirstName", "first_name", db.Varchar, form.Text).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("LastName", "last_name", db.Varchar, form.Text).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Email", "email", db.Varchar, form.Email).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Image_path", "image_path", db.Text, form.Text).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Phone", "phone", db.Varchar, form.Text).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("RoleID", "role_id", db.Int, form.Select).
		FieldOptionsFromTable("goadmin_roles", "slug", "id").AddLimitFilter(1).
		FieldHelpMsg(template.HTML("no corresponding options?") + Link("/admin/info/roles/new", "Create here."))
	formList.AddField("UniversityID", "university_id", db.UUID, form.Select).
		FieldOptionsFromTable("universities", "name", "id").AddLimitFilter(1)

	formList.AddField("Verified", "verified", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "false", Text: "False"},
			{Value: "true", Text: "True"},
		})

	formList.SetTable("users").SetTitle("Users").SetDescription("There are all list of available users.")

	formList.SetUpdateFn(func(values form2.Values) error {
		var req api.UpdateUserRequest

		if values.Get("verified") != "" {
			if values.Get("verified") == "true" {
				req.Verified = true
			} else {
				req.Verified = false
			}
		}
		if len(values["role_id[]"]) != 0 {

			if len(values["role_id[]"]) > 1 {
				return fmt.Errorf("role id limit is 1 error: Please decrease limit to 1")
			}

			roleId, err := strconv.Atoi(values["role_id[]"][0])
			if err != nil {
				return fmt.Errorf("convert string to int on role_id err: %w", err)
			}
			req.RoleId = int64(roleId)
		}
		if len(values["university_id[]"]) != 0 {
			if len(values["university_id[]"]) > 1 {
				return fmt.Errorf("universities limit is 1 error: Please decrease limit to 1")
			}
			req.UniversityId = values["university_id[]"][0]
		}

		err := g.service.UpgradeUser(contx.Background(), values.Get("id"), &req)
		if err != nil {
			log.Printf("update user error: %v", err)
			return err
		}

		return nil
	})

	//formList.SetInsertFn(func(values form2.Values) error {
	//	var req entity.User
	//
	//	if values.Get("first_name") != "" {
	//		req.FirstName = values.Get("first_name")
	//	}
	//	if values.Get("last_name") != "" {
	//		req.FirstName = values.Get("last_name")
	//	}
	//	if values.Get("image_path") != "" {
	//		req.ImagePath = values.Get("image_path")
	//	}
	//	if values.Get("phone") != "" {
	//		req.Phone = values.Get("phone")
	//	}
	//	if values.Get("password") != "" {
	//		req.Password = values.Get("password")
	//	}
	//
	//	if values.Get("role_id") != "" {
	//		roleId, err := strconv.Atoi(values.Get("role_id"))
	//		if err == nil {
	//			req.RoleId = int64(roleId)
	//			log.Printf("[GoAdmin] Convert string to int on role_id err: %w", err)
	//		} else {
	//			log.Printf("[GoAdmin] Convert string to int on role_id err: %w", err)
	//		}
	//	}
	//	if values.Get("university_id") != "" {
	//		req.UniversityId = values.Get("university_id")
	//	}
	//	if values.Get("email") != "" {
	//		req.Email = values.Get("email")
	//	}
	//
	//	_, err := g.service.CreateUser(contx.Background(), &req)
	//	if err != nil {
	//		log.Printf("create user error: %v", err)
	//		return err
	//	}
	//
	//	return nil
	//})

	return users
}
