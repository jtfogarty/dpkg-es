package ESmgr

import (
	"fmt"

	"golang.org/x/net/context"

	elastic "gopkg.in/olivere/elastic.v5"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//CreateIndex comment
func CreateIndex() {
	ctx := context.Background()

	client, err := elastic.NewClient()
	check(err)

	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	check(err)
	fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	exists, err := client.IndexExists("twitter").Do(ctx)
	check(err)

	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("twitter").Do(ctx)
		check(err)
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}
