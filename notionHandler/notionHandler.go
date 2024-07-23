package notionHandler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jomei/notionapi"

	"github.com/Codesmith28/botCore/internal"
)

var (
	secret     = internal.NotionSecret
	databaseID = internal.DatabaseId
	Tasklist   []*Task
)

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func QueryDatabase(client *notionapi.Client) {
	res, err := client.Database.Query(
		context.Background(),
		notionapi.DatabaseID(databaseID),
		&notionapi.DatabaseQueryRequest{},
	)
	checkNilErr(err)

	for _, page := range res.Results {
		props, err := json.MarshalIndent(page.Properties, "", " ")
		checkNilErr(err)

		var propMap map[string]interface{}
		err = json.Unmarshal(props, &propMap)
		checkNilErr(err)

		task := Formatter(propMap)

		// Skip tasks marked as done
		if task == nil {
			continue
		}

		Tasklist = append(Tasklist, task)
	}
}

func NotionConnect() {
	// clear tasklist:
	Tasklist = nil
	client := notionapi.NewClient(notionapi.Token(secret))
	QueryDatabase(client)
}
