package logs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddLogEntry(t *testing.T) {
	c := new(ElasticClient)
	var err error
	err = c.Init()

	log := Log{
		User:       "lamg",
		Addr:       "twitter.com",
		Meth:       "GET",
		URI:        "/bretanac93/tweets",
		Proto:      "HTTP1.1",
		StatusCode: 200,
		RespSize:   449992,
		Time:       time.Now(),
	}
	err = c.AddLogEntry(log)
	require.NoError(t, err)
}

func TestGetEntriesByUser(t *testing.T) {
	c := new(ElasticClient)
	var (
		err error
		res []Log
	)
	err = c.Init()

	res, err = c.GetEntriesByUser("lamg")

	require.NotEmpty(t, res)
	require.NoError(t, err)
}
