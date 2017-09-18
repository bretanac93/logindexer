package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

//ElasticClient Wrapper for operations with Elastic Search.
type ElasticClient struct {
	Ctx    context.Context
	Client *elastic.Client
}

//Init Initializes Elastic search client
func (c *ElasticClient) Init() (err error) {
	c.Ctx = context.Background()
	c.Client, err = elastic.NewClient()
	return
}

//AddLogEntry Inserts a log entry to the elastic search server.
func (c *ElasticClient) AddLogEntry(logEntry Log) (err error) {
	_, err = c.Client.Index().
		Index("logs").
		Type("log").
		Id(generateGUID()).
		BodyJson(logEntry).
		Refresh("true").
		Do(c.Ctx)
	return
}

//GetEntriesByUser Retrieve a logs list from an username
func (c *ElasticClient) GetEntriesByUser(username string) (res []Log, err error) {
	termQuery := elastic.NewTermQuery("user", username)

	searchResult, err := c.Client.Search().
		Index("logs").
		Query(termQuery).
		Do(c.Ctx)

	if searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var l Log
			err = json.Unmarshal(*hit.Source, &l)
			res = append(res, l)
		}
	}

	return
}

func generateGUID() string {
	f, _ := os.Open("/dev/urandom")
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	return fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
