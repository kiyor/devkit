package routers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/kiyor/devkit/controllers"
)

type RouterConfig struct {
	RootDir string
	App     string
	Host    string
}

func NewRouterConfig() *RouterConfig {
	return &RouterConfig{}
}

func NewRoutes(conf *RouterConfig) http.Handler {
	ctrl := controllers.NewController(filepath.Join(conf.RootDir, "views"), filepath.Join(conf.RootDir, "static"))
	ctrl.App = conf.App
	ctrl.Host = conf.Host
	// 	fileServer := http.FileServer(http.Dir(conf.RootDir))
	r := mux.NewRouter()
	// 	r.PathPrefix("/static").Handler(fileServer)
	r.PathPrefix("/static").HandlerFunc(ctrl.StaticHandler)

	r.Path("/api").HandlerFunc(ctrl.ApiHandler)

	r.PathPrefix("/").HandlerFunc(ctrl.ViewHandler)
	return r
}
