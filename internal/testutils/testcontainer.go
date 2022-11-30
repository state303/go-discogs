package testutils

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"strconv"
	"time"
)

type DatabaseType int

const (
	MySQL DatabaseType = iota + 1
	Postgres
	SQLServer
)

type Database struct {
	Username  string
	Password  string
	Hostname  string
	DBName    string
	Type      DatabaseType
	Port      string
	Container testcontainers.Container
}

func GetDsn(dt DatabaseType, db Database) string {
	var n string
	var suffix string
	switch dt {
	case MySQL:
		return fmt.Sprintf("%+v:%+v@tcp(%+v:%+v)/%+v?charset=utf8&parseTime=True&loc=Local&tls=false", db.Username, db.Password, db.Hostname, db.Port, db.DBName)
	case Postgres:
		return fmt.Sprintf("host=%+v user=%+v password=%+v dbname=%+v port=%+v sslmode=disable", db.Hostname, db.Username, db.Password, db.DBName, db.Port)
	case SQLServer:
		n = "sqlserver"
		suffix = "?database=" + db.DBName
	case 4:
		n = "mongo"
	}
	dsn := fmt.Sprintf("%+v://%+v:%+v@%+v:%+v", n, db.Username, db.Password, db.Hostname, db.Port)
	return dsn + suffix
}

func GetDatabase(db DatabaseType) Database {
	if db == MySQL {
		return setupMySQL()
	}
	if db == Postgres {
		return setupPostgres()
	}
	panic("unknown server types index: " + strconv.Itoa(int(db)))
}

func setupMySQL() Database {
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{},
		Image:          "mysql:latest",
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "mysql",
			"MYSQL_DATABASE":      "test_db",
		},
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		WaitingFor:   wait.ForHealthCheck().WithPollInterval(time.Second),
	}
	c, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	h, _ := c.Host(context.Background())
	p, _ := c.MappedPort(context.Background(), "3306")
	return Database{
		Username:  "root",
		Password:  "mysql",
		Hostname:  h,
		DBName:    "test_db",
		Type:      MySQL,
		Port:      p.Port(),
		Container: c,
	}
}

func setupPostgres() Database {
	rootDir := GetProjectRoot()
	mountFrom := fmt.Sprintf("%s/scripts/postgres/schema.sql", rootDir)
	fmt.Println("mount from", mountFrom)
	mountTo := "/docker-entrypoint-initdb.d/init.sql"

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{},
		Image:          "postgres:latest",
		Entrypoint:     nil,
		Env: map[string]string{
			"POSTGRES_DB":       "test_db",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
		ExposedPorts: []string{"5432/tcp"},
		Mounts: testcontainers.Mounts(testcontainers.ContainerMount{
			Source:   testcontainers.GenericBindMountSource{HostPath: mountFrom},
			Target:   testcontainers.ContainerMountTarget(mountTo),
			ReadOnly: false,
		}),
		WaitingFor: wait.ForHealthCheck().WithPollInterval(time.Second),
	}
	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})

	if err != nil {
		panic(err)
	}

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")
	return Database{
		Username:  "postgres",
		Password:  "postgres",
		Hostname:  host,
		DBName:    "test_db",
		Type:      Postgres,
		Port:      port.Port(),
		Container: dbContainer,
	}
}
