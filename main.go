package main

import (
	"bst/http"
	"bst/tree"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"os"
)

var initFile = "initData.json"
var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
}
func main() {
	values := getInitValues(initFile)
	bstTree := tree.InitBST(logger, values)

	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	indexRoute := e.Group("/")
	http.InitAndBindBSTHandler(logger, indexRoute, bstTree)

	if err := e.Start(":8585"); err != nil {
		logger.Panicf("Произошла ошибка при старте веб-сервера! %v", err)
	}
}
