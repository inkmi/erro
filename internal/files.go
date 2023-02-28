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

func ReadSourceFs(filepath string, afs afero.Fs) []string {
	b, err := afero.ReadFile(afs, filepath)
	if err != nil {
		printf("erro: cannot read file '%s': %s. If sources are not reachable in this environment, you should set getData=false in logger config.", filepath, err)

	}
	lines := strings.Split(string(b), "\n")
	return lines
}

func ReadSource(filepath string) []string {
	return ReadSourceFs(filepath, fs)
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
