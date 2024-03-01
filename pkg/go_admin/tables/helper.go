package tables

import (
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/html"
	tmpl "html/template"
)

func Link(url, content string) tmpl.HTML {
	return html.AEl().
		SetAttr("href", url).
		SetContent(template.HTML(content)).
		Get()
}
