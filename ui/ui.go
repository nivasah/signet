package ui

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed dist/*
var spaFiles embed.FS

func SPAHandler() http.HandlerFunc {
	spaFS, err := fs.Sub(spaFiles, "dist")
	if err != nil {
		panic(fmt.Errorf("failed getting the subtree for the site files: %w", err))
	}
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := spaFS.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		if err == nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		http.FileServer(http.FS(spaFS)).ServeHTTP(w, r)
	}
}
