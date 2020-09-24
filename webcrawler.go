package main

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"fmt"
	"github.com/KMPolakowski/diploradar/webcrawler/db"
	"strings"
)

func main() {
	query := "SELECT website FROM foreign_ministries WHERE is_english = 1 AND website IS NOT NULL;"
	rows, err :=  db.Db.Query(query)
	
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	var urls []string

	for rows.Next() {
		var url string

		err := rows.Scan(&url)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		urls = append(urls, url)
	}

	q, _ := queue.New(
						len(urls),
						&queue.InMemoryQueueStorage{MaxSize:len(urls)},
				)

	for _, url := range urls {
		q.AddURL(url)
	}

	c := colly.NewCollector(
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		isInteresting(e)
	})

	c.OnHTML("span", func(e *colly.HTMLElement) {
		isInteresting(e) 
	})

	q.Run(c)
}

func isInteresting(e *colly.HTMLElement) bool {
	text := strings.ToUpper(e.Text)
	words := strings.Fields(text)

	for _, v := range words {
		if v == "PRESIDENT" || v == "MINISTER" {
			fmt.Println(e.Text)
			return true
		}
	}

	return false
}