package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	log "social-network/utils/log"

	"github.com/olivere/elastic/v7"
)

type ElasticClient struct {
	esclient *elastic.Client
}

var l = log.New()

func NewElasticClient(urls []string) (*ElasticClient, error) {
	client, err := elastic.NewClient(elastic.SetURL(urls...),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	return &ElasticClient{
		esclient: client,
	}, err
}

func (e *ElasticClient) Index(ctx context.Context, index string, body interface{}) error {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		l.Fatalf("Json error: %v", err)
	}

	_, err = e.esclient.Index().
		Index(index).
		BodyJson(string(bodyJson)).
		Do(ctx)

	return err
}

type SearchParams struct {
	Query  elastic.Query
	SortBy []elastic.Sorter
	Index  string
	Limit  int
	Offset int
}

func (e *ElasticClient) Search(ctx context.Context, searchParams *SearchParams) (*elastic.SearchResult, error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(searchParams.Query)

	queryStr, err := searchSource.Source()
	if err != nil {
		return nil, fmt.Errorf("searchSource.Source: %v", err)
	}

	queryJs, err := json.Marshal(queryStr)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %v", err)
	}

	l.Infof("QueryJson: %v", string(queryJs))

	searchService := e.esclient.Search().
		Index(searchParams.Index).
		SortBy(searchParams.SortBy...).
		SearchSource(searchSource).
		Size(searchParams.Limit).
		From(searchParams.Offset)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		return nil, err
	}

	return searchResult, err
}
