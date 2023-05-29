package main

import (
	"database/sql"
	"fmt"
	//"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//dbPassword := os.Getenv("DB_PASSWORD")

	for {
		db, err := sql.Open("mysql", fmt.Sprintf("admin:%s@tcp(mysql:3306)/test", "57968"))
		if err != nil {
                	fmt.Printf("Failed to connect to MySQL: %v\n", err)
                	return
        	}
		rows, err := db.Query("SELECT * FROM main")
		if err != nil {
			fmt.Printf("Failed to query data: %v\n", err)
			return
		}

		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			fmt.Printf("Failed to get columns: %v\n", err)
			return
		}

		values := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columns {
			columnPointers[i] = &values[i]
		}

		for rows.Next() {
			if err := rows.Scan(columnPointers...); err != nil {
				fmt.Printf("Failed to scan row: %v\n", err)
				return
			}

			for i, column := range columns {
				value := values[i]
				fmt.Printf("%s: %v\t", column, value)
			}
			fmt.Println()
		}

		if err := rows.Err(); err != nil {
			fmt.Printf("Error occurred during rows iteration: %v\n", err)
			return
		}
		db.Close()
		time.Sleep(2 * time.Second)
	}
}

