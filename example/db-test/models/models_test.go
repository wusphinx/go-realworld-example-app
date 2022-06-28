package models

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestModels(t *testing.T) {
	var sqlDB *sql.DB
	var err error
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		// notice: go test with M1 shoud set platform to "linux/amd64"
		Platform:   "linux/amd64",
		Repository: "mysql",
		Tag:        "5.7",
		Env:        []string{"MYSQL_ROOT_PASSWORD=secret"}},
		func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		})
	// autoRemove after 60s
	resource.Expire(60) //nolint

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	var dsn string
	if err = pool.Retry(func() error {
		var err error
		dsn = fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true", resource.GetPort("3306/tcp"))
		sqlDB, err = sql.Open("mysql", dsn)
		if err != nil {
			return err
		}
		return sqlDB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()
	// 迁移 schema
	if err := db.AutoMigrate(&Product{}); err != nil {
		panic("failed to migrate product")
	}

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	p := &Product{}
	ins, err := p.GetByID(db, 1)
	assert.Nil(t, err)
	assert.Equal(t, "D42", ins.Code)
}
