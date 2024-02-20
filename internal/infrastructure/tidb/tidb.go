package tidb

import (
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/internal/repository/tidb_repo/tidb_dto"
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
	cli *gorm.DB
}

func NewTiDB(cfg *config.Config) *TiDB {
	cli, err := openConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %s", err)
	}

	if err := createDatabase(cli, cfg.Database.DBName); err != nil {
		log.Fatalf("Error creating database: %s", err)
	}

	cli, err = connectToDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	// migrate models
	if err := cli.AutoMigrate(&tidb_dto.UserModel{}); err != nil {
		log.Fatalf("Error migrating user model: %s", err)
	}

	if err := cli.AutoMigrate(&tidb_dto.BankModel{}); err != nil {
		log.Fatalf("Error migrating bank model: %s", err)
	}

	return &TiDB{cli: cli}
}

func openConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := getDSN(cfg, "")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func createDatabase(cli *gorm.DB, dbName string) error {
	return cli.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)).Error
}

func connectToDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := getDSN(cfg, cfg.Database.DBName)
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

func getDSN(cfg *config.Config, dbName string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		dbName,
		url.QueryEscape(cfg.Database.Timezone),
	)
}

func (m *TiDB) GetClient() *gorm.DB {
	return m.cli
}
