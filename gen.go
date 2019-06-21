package golayout

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

var (
	projName string
)

func Generate(projectName string) error {
	log.Infof("Generate project %s", projectName)
	projName = projectName

	if err := setup(); err != nil {
		log.Fatal(err)
	}

	err := travel()
	tearDownIfError(err)
	return err
}

func setup() error {
	log.Infof("Create project directory")
	return os.Mkdir(projName, 0751)
}

func tearDownIfError(err error) {
	if err != nil {
		os.RemoveAll(projName)
		log.Fatal(err)
	}
}

func travel() error {
	err := TemplateBox.Walk(func(s string, f packr.File) error {
		log.Infof("walk to %s", s)
		return toFile(s, f.String())
	})

	return err
}

func toFile(relativeFilePath, content string) error {
	// remove .tpl
	rp := strings.TrimSuffix(relativeFilePath, TemplateExt)
	// open output file
	fp := filepath.Join(projName, rp)
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
