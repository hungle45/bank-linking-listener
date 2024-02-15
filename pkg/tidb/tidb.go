package tidb

import (
	"demo/bank-linking-listener/config"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TiDB struct {
	db *gorm.DB
}

func NewDB(cfg *config.DatabaseConfig) *TiDB {
	db, err := openConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %s", err)
	}

	if err := createDatabase(db, cfg.DBName); err != nil {
		log.Fatalf("Error creating database: %s", err)
	}

	db, err = connectToDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	return &TiDB{db: db}
}

func openConnection(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := getDSN(cfg, "")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func createDatabase(db *gorm.DB, dbName string) error {
	return db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)).Error
}

func connectToDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := getDSN(cfg, cfg.DBName)
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger})
}

func getDSN(cfg *config.DatabaseConfig, dbName string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		dbName,
		url.QueryEscape(cfg.Timezone),
	)
}

func (m *TiDB) GetConn() *gorm.DB {
	return m.db
}
