package story

import (
	"html/template"
	"log"
	"path"
	"runtime"

	"golang.org/x/xerrors"
)

var (
	layoutTpl *template.Template
)

func init() {
	_, filename, _, _ := runtime.Caller(0)

	var err error
	layoutTpl, err = template.ParseFiles(path.Dir(filename) + "/templates/layout.html")
	if err != nil {
		log.Fatal(xerrors.Errorf("could not load template : %v", err))
	}
}
