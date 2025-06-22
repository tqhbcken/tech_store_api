package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"api_techstore/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnection struct {
	DB *gorm.DB
}

var (
	dbInstance *DBConnection
	once       sync.Once
)

// AutoMigrate thực hiện tự động migrate các model
func (dbc *DBConnection) AutoMigrate(models ...interface{}) error {
	if err := dbc.DB.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}
	log.Println("✅ Database migration completed successfully")
	return nil
}
//
func InitDB() (*DBConnection, error) {
	var err error
	once.Do(func() {
		dbInstance, err = createDBConnection()
	})
	return dbInstance, err
}
//
func createDBConnection() (*DBConnection, error) {

	// Load environment variables
	config.LoadEnvVar()

	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", v)
		}
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable connect_timeout=5",
		host, user, password, dbname, port,
	)

	//open db conn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// config connection pool
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Connected to PostgreSQL database successfully.")
	return &DBConnection{DB: db}, nil
}

// close connection
func (dbc *DBConnection) Close() error {
	if dbc.DB != nil {
		sqlDB, err := dbc.DB.DB()
		if err != nil {
			return fmt.Errorf("failed to get database instance: %w", err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %w", err)
		}
		log.Println("Database connection closed successfully")
	}
	return nil
}
