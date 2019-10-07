package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"github.com/KMPolakowski/diploradar/webcrawler/db"
)

func main() {
	query := "SELECT website FROM foreign_ministries WHERE is_english = 1;"
	rows, err :=  db.Db.Query(query)
	
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	var users []string

	for rows.Next() {
		var user string

		err := rows.Scan(&user)

		if err != nil {
			fmt.Println(err.Error())
		}

		users = append(users, user)
	}

	fmt.Println(users)

	colly.NewCollector(colly.AllowURLRevisit())
}