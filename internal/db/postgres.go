package database

import (
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"zychimne/instant/config"
	"zychimne/instant/pkg/model"
)

var PostgresDB *gorm.DB

func ConnectPostgres() {
	conn := strings.Join([]string{"host=" + config.Conf.Database.Postgres.Host, "user=" + config.Conf.Database.Postgres.User, "password=" + config.Conf.Database.Postgres.Password, "dbname=" + config.Conf.Database.Postgres.Database, config.Conf.Database.Postgres.Extras}, " ")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                   // Don't include params in the SQL log
			Colorful:                  true,                   // Enable color
		},
	)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{Logger: newLogger, PrepareStmt: true})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Following{}, &model.Instant{}, &model.InboxInstant{}, &model.Conversation{}, &model.InboxMessage{}, &model.Message{})
	PostgresDB = db
}
