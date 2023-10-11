package database

import (
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"zychimne/instant/config"
	"zychimne/instant/pkg/model"
)

var PostgresDB *gorm.DB

func ConnectPostgres() {
	conn := strings.Join([]string{"host=" + config.Conf.Postgres.Host, "user=" + config.Conf.Postgres.User, "password=" + config.Conf.Postgres.Password, "dbname=" + config.Conf.Postgres.Database, config.Conf.Postgres.Extras}, " ")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Following{})
	PostgresDB = db
}
