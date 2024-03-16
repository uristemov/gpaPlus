package front

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func InitFront(router *gin.Engine) *gin.Engine {

	router.LoadHTMLGlob("C:/Users/meiro/OneDrive/Desktop/Diploma_Project/*.html")
	//router.LoadHTMLGlob("C:/Users/meiro/OneDrive/Desktop/Diploma_Project/*/*.html")
	router.Static("/assets", "C:/Users/meiro/OneDrive/Desktop/Diploma_Project/assets")
	//	router.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("../../front/static/assets/"))))

	return router
}

func loadTemplates(root string) (files []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			if path != root {
				loadTemplates(path)
			}
		} else {
			files = append(files, path)
		}
		return err
	})
	return files, err
}
