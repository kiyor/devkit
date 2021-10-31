package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Controller struct {
	viewsRoot  string
	staticRoot string
	temp       *template.Template
	App        string
	Host       string
}

func NewController(views, static string) *Controller {
	return &Controller{
		viewsRoot:  views,
		staticRoot: static,
	}
}

type m struct {
	Path    string
	Query   url.Values
	Host    string
	Uri     string
	IsPhone bool
	Map     map[string][]string
	App     string
	List    []string
}

func newM() *m {
	return &m{
		Query: make(url.Values),
		Map:   make(map[string][]string),
	}
}

func (c *Controller) ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/index"
	}
	file := filepath.Join(c.viewsRoot, r.URL.Path+".html")
	if _, err := os.Stat(file); err != nil {
		w.WriteHeader(404)
		return
	}

	tmpl, list, err := c.parseTemplates()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	ua := r.Header.Get("User-Agent")
	rePhone := regexp.MustCompile(`(iPhone|iPad|Android).*Mobile`)

	m := newM()
	m.Path = r.URL.Path
	m.Query = r.URL.Query()
	m.Host = "http://" + r.Host
	if len(c.Host) > 0 {
		m.Host = c.Host
	}
	m.Uri = r.URL.RequestURI()
	m.App = c.App
	m.List = list

	if rePhone.MatchString(ua) {
		m.IsPhone = true
	}

	err = tmpl.ExecuteTemplate(w, r.URL.Path[1:]+".html", m)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func isIgnore(path string) bool {
	for _, v := range []string{
		"empty",
		"header",
		"footer",
		"index",
		"nav",
	} {
		if path == v {
			return true
		}
	}
	return false
}

func (c *Controller) parseTemplates() (*template.Template, []string, error) {
	funcMap := make(template.FuncMap)
	t := template.New("").Funcs(funcMap).Delims("[[", "]]")

	var list []string

	err := filepath.Walk(c.viewsRoot, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, ".html") {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			s := string(b)
			dir := filepath.Base(c.viewsRoot)
			name := path[len(dir)+1:]
			path := strings.TrimSuffix(name, ".html")
			if !isIgnore(path) {
				list = append(list, path)
			}
			var tmpl *template.Template
			tmpl = t.New(name)
			_, err = tmpl.Parse(s)
			if err != nil {
				return err
			}
		}

		return err
	})
	c.temp = t
	return t, list, err

}
