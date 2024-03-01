package go_admin

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	go_config "github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/sword"
	"github.com/gin-gonic/gin"
	"github.com/uristemov/repeatPro/internal/config"
	"github.com/uristemov/repeatPro/pkg/go_admin/tables"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func InitGoAdmin(router *gin.Engine, cfg config.Database, logger *zap.SugaredLogger, goAdmin *tables.GoAdmin) *engine.Engine {

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	// PostgreSQL database configuration
	goAdminDbConfig := go_config.Database{
		Host:            cfg.Host,
		Port:            strconv.Itoa(cfg.Port),
		User:            cfg.Username,
		Pwd:             cfg.Password,
		Name:            cfg.DBName,
		Driver:          go_config.DriverPostgresql,
		MaxIdleConns:    50,
		MaxOpenConns:    150,
		ConnMaxLifetime: time.Hour,
		//File:   "./test.db", // SQLite3 file path, only for sqlite
	}

	// GoAdmin global configuration, can also be imported as a json file.
	goAdminCfg := go_config.Config{
		Databases: go_config.DatabaseList{
			"default": goAdminDbConfig,
		},
		UrlPrefix: "admin", // The url prefix of the website.
		// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
		Store: go_config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		ColorScheme: adminlte.ColorschemeSkinBlack,
		Language:    language.EN,
	}

	//adminPlugin := admin.NewAdmin(datamodel.Generators)
	// Enable debug mode

	if err := eng.AddConfig(&goAdminCfg).
		//AddPlugins(adminPlugin).
		AddGenerator("courses", goAdmin.GetCoursesTable).
		AddGenerator("users", goAdmin.GetUsersTable).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		//AddGenerator("user", datamodel.GetUserTable).
		Use(router); err != nil {
		logger.Panicf("go-admin engine add config error: %v", err)
	}
	//eng.AddPlugins(adminPlugin)
	// customize your pages
	eng.HTML("GET", "/admin", datamodel.GetContent)
	return eng
}
