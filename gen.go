package golayout

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

var (
	projOverall ProjectOverall
	tpl         *template.Template
)

const (
	AppNamePlaceholder = "appname"
)

type ProjectOverall struct {
	ProjName string
	ModName  string
}

func Generate(project ProjectOverall) error {
	log.Infof("Generate project %s", project.ProjName)
	projOverall = project

	if err := setup(); err != nil {
		log.Fatal(err)
	}

	err := travel()
	tearDownIfError(err)

	log.Infof("Finish to generate project %s", project.ProjName)

	return err
}

func setup() error {
	log.Infof("Create project directory")
	return os.Mkdir(projOverall.ProjName, 0751)
}

func tearDownIfError(err error) {
	if err != nil {
		os.RemoveAll(projOverall.ProjName)
		log.Fatal(err)
	}
}

func travel() error {
	err := TemplateBox.Walk(func(s string, f packr.File) error {
		s = strings.ReplaceAll(s, AppNamePlaceholder, projOverall.ProjName)
		log.Infof("walk to %s", s)

		d, e := compileTpl(f.String())
		if e != nil {
			return e
		}

		return toFile(s, d)
	})

	return err
}

func compileTpl(tplString string) (string, error) {
	buf := &bytes.Buffer{}
	log.Infof("tplString %s", tplString)
	tpl = template.Must(template.New("letter").Parse(tplString))

	err := tpl.Execute(buf, projOverall)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func toFile(relativeFilePath, content string) error {
	// remove .tpl
	rp := strings.TrimSuffix(relativeFilePath, TemplateExt)
	// open output file
	fp := filepath.Join(projOverall.ProjName, rp)
	fo, err := CreateFileIncludeDir(fp)
	if err != nil {
		return err
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			log.WithError(err).Errorf("fail to open file %s", fp)
		}
	}()

	if _, err := fo.Write([]byte(content)); err != nil {
		log.WithError(err).Errorf("fail to write file %s", fp)
		return err
	}

	return nil
}
