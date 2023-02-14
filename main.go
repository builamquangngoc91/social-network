package main

import (
	"context"
	"fmt"
	elasticsearch "social-network/utils/elasticsearch"
	log "social-network/utils/log"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var l = log.New()

func main() {
	ctx := context.Background()

	newFeed := Feed{
		Name: fmt.Sprintf("random dsa %d", time.Now().Nanosecond()),
		ID:   uuid.New().String(),
	}

	elasticClient, err := elasticsearch.NewElasticClient([]string{"http://localhost:9200"})
	if err != nil {
		l.Info(err.Error())
	}

	if err := elasticClient.Index(ctx, "feed", newFeed); err != nil {
		l.Info(err.Error())
	}

	// var feeds []Feed

	// searchSource := elastic.NewSearchSource()
	// searchSource.Query(elastic.NewMatchQuery("name", "random"))

	// queryStr, err1 := searchSource.Source()
	// queryJs, err2 := json.Marshal(queryStr)

	// if err1 != nil || err2 != nil {
	// 	fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	// }
	// fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	// /* until this block */

	// searchService := esclient.Search().Index("feed").SearchSource(searchSource)

	// searchResult, err := searchService.Do(ctx)
	// if err != nil {
	// 	fmt.Println("[ProductsES][GetPIds]Error=", err)
	// 	return
	// }

	// for _, hit := range searchResult.Hits.Hits {
	// 	var feed Feed
	// 	err := json.Unmarshal(hit.Source, &feed)
	// 	if err != nil {
	// 		fmt.Println("[Getting Feeds][Unmarshal] Err=", err)
	// 	}

	// 	feeds = append(feeds, feed)
	// }

	// if err != nil {
	// 	fmt.Println("Fetching feed fail: ", err)
	// } else {
	// 	for _, s := range feeds {
	// 		fmt.Printf("Feed found Name: %s, ID: %s \n", s.Name, s.ID)
	// 	}
	// }
}
