package orms

import (
	"context"
	"database/sql"
	"fmt"
	"main/orms/sqlc/orms"

	_ "github.com/lib/pq"
)

func TestPostgreSqlc() {

	ctx := context.Background()

	db, _ := sql.Open("postgres", "host=localhost user=user password=password dbname=database_name port=5432 sslmode=disable")

	qr := orms.New(db)

	users, _ := qr.ListUsers(ctx)

	for i := range users {
		fmt.Println(users[i])
	}

}
