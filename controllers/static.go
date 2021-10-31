package controllers

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

type File struct {
	Path string
	os.FileInfo
}

type Files []*File

func (c *Controller) StaticHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[1:]
	info, err := os.Stat(p)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	if info.IsDir() {
		var files Files
		walk := func(path string, fi os.FileInfo, err error) error {
			if err == nil && fi.Mode().IsRegular() {
				files = append(files, &File{Path: path, FileInfo: fi})
			}
			return nil
		}

		err := filepath.Walk(p, walk)
		if err == nil {
			sort.Slice(files, func(i, j int) bool {
				return files[i].FileInfo.ModTime().After(files[j].FileInfo.ModTime())
			})
			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte("<ul>\n"))
			for _, v := range files {
				w.Write([]byte(fmt.Sprintf(`<li><a href="/%s">%s</a></li>
`, v.Path, v.Path)))
			}
			w.Write([]byte("</ul>"))
		} else {
			log.Println(err)
		}
	}
	f, err := os.Open(p)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	defer f.Close()
	w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(p)))
	io.Copy(w, f)
}
