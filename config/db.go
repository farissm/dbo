package config

import (
	"fmt"
	"os"
	"dbo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

type TableInfo struct {
    TableName string `gorm:"column:TABLE_NAME"`
}

type Table struct {
    Name string
}

type DatabaseCredential struct {
	Username	string
	Password	string
	Host		string
	Port		string
	Name		string
}

func DBCredential() *DatabaseCredential {
	return &DatabaseCredential{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}
}

func setupDBConnection(DBCredential *DatabaseCredential) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBCredential.Username,
		DBCredential.Password,
		DBCredential.Host,
		DBCredential.Port,
		DBCredential.Name,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func ConnectDB() *gorm.DB {
	dbCredential := DBCredential()
	db, err := setupDBConnection(dbCredential)
	if err != nil {
		panic(err)
	}

	var tables []TableInfo
    db.Raw("SELECT TABLE_NAME FROM information_schema.tables WHERE TABLE_SCHEMA = ?", dbCredential.Name).Scan(&tables)

    if len(tables) == 0 {
		db.AutoMigrate(&model.Customer{}, &model.Order{})
	} else {
		log.Info("table customer and order already exist")
	}

	return db
}