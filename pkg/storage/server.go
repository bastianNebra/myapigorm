package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/bastianNebra/myapigorm.git/pkg/logs"
	"net/http"
)


const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password="1234"
	dbname = "mydb"
	
)
var myLogger = logs.NewLogger()
//Struct Servershop
type ServerShop interface{
	GetHallo() http.HandlerFunc
	openDBCon() error
	Conn() http.HandlerFunc
	GetCustomers() http.HandlerFunc
	GetCustomer() http.HandlerFunc
	CreateCustomers() http.HandlerFunc
	UpdateCustomer() http.HandlerFunc
	DeleteCustomer() http.HandlerFunc
	
}
type serverShop struct {
	DB *gorm.DB
}


func NewServerShop()ServerShop{
	return &serverShop{}
}

// Open a database connexion
func  (shop *serverShop) openDBCon() error{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=disable",host,port,user,password,dbname)

	//DB connexion with GORM
	db, err := gorm.Open(postgres.Open(psqlInfo),&gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("failed to connect database: %v", err)
		myLogger.ErrorLoggerFunc(msg)
	}

	
	shop.DB = db

	// Migrate the schema
	err = db.AutoMigrate(&Customer{},&Product{})
	if err !=nil {
		msg := fmt.Sprintf("failed to connect database: %v", err)
		myLogger.ErrorLoggerFunc(msg)
	}
	
	return nil
}
