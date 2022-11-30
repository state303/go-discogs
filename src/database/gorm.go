package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type DBKind int

const (
	MySQL DBKind = iota + 1
	Postgres
)

var (
	m = regexp.MustCompile(`^(mysql)://([^/]+:[^/]+)@([^/]+:\d+)/(.*)?$`)
	p = regexp.MustCompile(`(^(postgres)://.*$|^host=\w+ user=\w+ password=\w+ dbname=\w+ port=\d+ .*$)`)
	x = regexp.MustCompile(`^(.*)://.*$`)
)

var DB *gorm.DB
var Kind DBKind

// Connect to database based on the dsn.
// mysql://user:pass@host:port/db_name?options
// postgres://user:pass@host:port/db_name?options
func Connect(dsn string) (err error) {
	if DB, err = GetConnect(dsn); err != nil {
		return
	} else if sqlDB, sqlDbErr := DB.DB(); sqlDbErr != nil {
		return sqlDbErr
	} else {
		sqlDB.SetConnMaxLifetime(time.Second)
		return
	}
}

func GetConnect(dsn string) (*gorm.DB, error) {
	var dl gorm.Dialector
	if len(dsn) == 0 {
		return nil, errors.New("missing dsn")
	}
	if m.MatchString(dsn) || strings.Contains(dsn, "tcp") {
		match := m.FindStringSubmatch(dsn)
		dd := dsn
		if m.MatchString(dsn) {
			dd = fmt.Sprintf("%+v@tcp(%+v)/%+v", match[2], match[3], match[4])
		}
		Kind = MySQL
		dl = mysql.Open(dd)
	} else if p.MatchString(dsn) {
		Kind = Postgres
		dl = postgres.Open(dsn)
	} else {
		if match := x.FindStringSubmatch(dsn); match != nil {
			return nil, errors.New("unsupported database from dsn: " + match[1])
		} else {
			return nil, errors.New("unsupported dsn. please check again")
		}
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Silent,
		})

	return gorm.Open(dl, &gorm.Config{Logger: newLogger})
}
