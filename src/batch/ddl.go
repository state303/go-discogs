package batch

import (
	"github.com/state303/go-discogs/src/database"
	"gorm.io/gorm"
	"os"
	"path"
	"regexp"
	"strings"
)

func RunDDL(db *gorm.DB) error {
	if sql, err := ReadScript(); err != nil {
		return err
	} else {
		parts := strings.Split(sql, ";")
		for i := range parts {
			q := strings.TrimSpace(parts[i])
			if len(q) == 0 {
				continue
			}
			if err := db.Session(&gorm.Session{}).Exec(q).Error; err != nil {
				return err
			}
		}
		return nil
	}
}

func ReadScript() (string, error) {
	v, err := os.ReadFile(GetScriptPath())
	return string(v), err
}

func GetScriptPath() string {
	return path.Join(GetProjectRoot(), "scripts", GetScriptFolderName(), "schema.sql")
}

func GetScriptFolderName() string {
	switch database.Kind {
	case database.MySQL:
		return "mysql"
	case database.Postgres:
		return "postgres"
	default:
		return "unknown"
	}
}

var ProjectRootPattern = regexp.MustCompile("^(.*go-discogs).*")

func GetProjectRoot() string {
	d, _ := os.Getwd()
	return ProjectRootPattern.FindStringSubmatch(d)[1]
}
