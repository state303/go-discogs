package testutils

import (
	"os"
	"regexp"
)

var ProjectRootPattern = regexp.MustCompile(`^(.*go-discogs).*`)

func GetProjectRoot() string {
	d, _ := os.Getwd()
	return ProjectRootPattern.FindStringSubmatch(d)[1]
}
