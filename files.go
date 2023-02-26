package erro

import (
	"github.com/spf13/afero"
	"os"
	"strings"
)

var (
	gopath = os.Getenv("GOPATH")
)

func ReadSource(filepath string) []string {
	b, err := afero.ReadFile(fs, filepath)
	if err != nil {
		if LogTo != nil {
			(*LogTo).Debug().Msgf("erro: cannot read file '%s': %s. If sources are not reachable in this environment, you should set PrintSource=false in logger config.", filepath, err)
		}
		return nil
	}
	lines := strings.Split(string(b), "\n")
	return lines
}

func GetShortFilePath(filepath string) string {
	filepathShort := filepath
	if gopath != "" {
		filepathShort = strings.Replace(filepath, gopath+"/src/", "", -1)
	}
	return filepathShort
}
