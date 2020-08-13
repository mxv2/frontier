package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/db")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
	defer db.Close()

	selectWithTimeout("A", db, 0)

	selectWithTimeout("B", db, 1*time.Second)

	selectWithTimeout("C", db, 1*time.Second)
}

func selectWithTimeout(sid string, db *sql.DB, timeout time.Duration) {
	log.Printf("[%s] Start query with timeout %v\n", sid, timeout)
	ctx := context.Background()
	if timeout != 0 {
		tctx, cancelFunc := context.WithTimeout(ctx, timeout)
		ctx = tctx
		defer cancelFunc()
	}

	rows, err := db.QueryContext(ctx, "SELECT id, name, age FROM users WHERE age < 25 OR sleep(1)")
	if err != nil {
		log.Printf("[%s] db.Query, err: %s\n", sid, err)
		return
	}
	defer rows.Close()

	var (
		id   int
		name string
		age  int
	)
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Printf("[%s] rows.Scan, err: %s\n", sid, err)
			break
		}

		log.Printf("[%s] row: id=%d, name=%s, age=%d\n", sid, id, name, age)
		break
	}
	err = rows.Err()
	if err != nil {
		log.Printf("[%s] rows.Err, err: %s\n", sid, err)
	}
}
