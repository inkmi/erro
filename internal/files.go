package internal

import (
	"github.com/spf13/afero"
	"os"
	"strings"
)

var (
	fs = afero.NewOsFs() //fs is at package level because I think it needn't be scoped to loggers
)

var (
	gopath = os.Getenv("GOPATH")
)

func ReadSource(filepath string) []string {
	b, err := afero.ReadFile(fs, filepath)
	if err != nil {
		if LogTo != nil {
			(*LogTo).Debug().Msgf("erro: cannot read file '%s': %s. If sources are not reachable in this environment, you should set getData=false in logger config.", filepath, err)
		}
		return nil
	}
	lines := strings.Split(string(b), "\n")
	return lines
}

func getShorterFilePath(filepath string, remove string) string {
	filepathShort := filepath
	if remove != "" {
		filepathShort = strings.Replace(filepath, remove, "", -1)
	}
	return filepathShort
}

func getShortFilePath(filepath string) string {
	return getShorterFilePath(filepath, gopath+"/src/")
}
