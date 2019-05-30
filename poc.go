package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type rssParsed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func main() {

	http.HandleFunc("/parse", func(w http.ResponseWriter, r *http.Request) {

		var res []rssParsed

		var req struct {
			URL          string    `json:"url"`
			StartingFrom time.Time `json:"from_time"`
		}

		decoder := json.NewDecoder(r.Body)

		if decoder.Decode(&req) != nil {
			return
		}

		fmt.Println(req.StartingFrom)

		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(req.URL)

		if err != nil {
			log.Print(err)
			return
		}

		for _, i := range feed.Items {
			if i.PublishedParsed.Sub(req.StartingFrom) > 0 {
				res = append(res, rssParsed{
					Title:       i.Title,
					Description: i.Description,
					URL:         i.Link,
				})
			}
		}

		json.NewEncoder(w).Encode(res)
	})

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
