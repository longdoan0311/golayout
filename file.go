package golayout

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func CreateFileIncludeDir(fp string) (*os.File, error) {
	dir := filepath.Dir(fp)
	isExist, err := Exists(dir)
	if err != nil {
		return nil, err
	}

	if !isExist {
		log.Infof("Create sub directory %s", dir)
		if err := os.MkdirAll(dir, 0751); err != nil {
			return nil, err
		}
	}

	fo, err := os.Create(fp)
	if err != nil {
		return nil, err
	}

	return fo, nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
