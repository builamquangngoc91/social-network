package main

import (
	"context"
	"encoding/json"
	"fmt"

	"social-network/utils/elasticsearch"
	log "social-network/utils/log"

	"github.com/olivere/elastic/v7"
)

type Feed struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var l = log.New()

func main() {
	ctx := context.Background()

	elasticClient, err := elasticsearch.NewElasticClient([]string{"http://localhost:9200"})
	if err != nil {
		l.Info(err.Error())
	}

	// if err := elasticClient.Index(ctx, "feed", newFeed); err != nil {
	// 	l.Info(err.Error())
	// }

	searchResult, err := elasticClient.Search(ctx, &elasticsearch.SearchParams{
		Query: elastic.NewMatchQuery("name", "random"),
		SortBy: []elastic.Sorter{
			elastic.NewFieldSort("name").Desc().Sorter,
		},
		Index:  "feed",
		Limit:  20,
		Offset: 0,
	})
	if err != nil {
		l.Fatalf("elasticClient.Search: %v", err.Error())
	}

	var (
		feed  Feed
		feeds []Feed
	)

	for _, hit := range searchResult.Hits.Hits {
		err := json.Unmarshal(hit.Source, &feed)
		if err != nil {
			fmt.Println("[Getting Feeds][Unmarshal] Err=", err)
		}

		feeds = append(feeds, feed)
	}

	if err != nil {
		fmt.Println("Fetching feed fail: ", err)
	} else {
		for _, s := range feeds {
			fmt.Printf("Feed found Name: %s, ID: %s \n", s.Name, s.ID)
		}
	}
}
