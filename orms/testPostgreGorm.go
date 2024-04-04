package orms

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgreGorm() {

	dsn := "host=localhost user=user password=password dbname=database_name port=5432 sslmode=disable"

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	rows := db.Create(User{Nome: "Teste", Email: "rodrigo@email.com", Cpf: "123654789"})

	fmt.Println(rows.RowsAffected)

}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Nome  string // A regular string field
	Email string // A pointer to a string, allowing for null values
	Cpf   string // An unsigned 8-bit integer

}
