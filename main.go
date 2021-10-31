package main

import (
	_ "embed"
	"flag"
	"log"
	"net/http"
	"runtime"

	"github.com/NYTimes/gziphandler"
	"github.com/kiyor/devkit/controllers"
	"github.com/kiyor/devkit/routers"
)

var (
	intf    string
	port    string
	addr    string
	rootDir string
	App     string
	host    string
)

func init() {
	flag.StringVar(&intf, "i", "0.0.0.0", "http service interface address")
	flag.StringVar(&port, "l", ":8080", "http service listen port")
	flag.StringVar(&host, "host", "", "Host with scheme")
	flag.StringVar(&rootDir, "root", ".", "root dir")
}

func main() {
	flag.Parse()
	addr = intf + port
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf := routers.NewRouterConfig()
	conf.RootDir = rootDir
	conf.App = App
	conf.Host = host
	r := routers.NewRoutes(conf)
	handler := controllers.NewLogHandler().Handler(r)
	handler = gziphandler.GzipHandler(handler)
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
