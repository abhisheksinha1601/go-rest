package mongo

import (
	"context"
	"log"
	"rest/interfaces"

	"go.mongodb.org/mongo-driver/bson"
)

func Find(params interfaces.IMongoFindInputParams) ([]bson.M, error) {
	projection := map[string]int{}
	for i := 0; i < len(params.Projection); i++ {
		projection[params.Projection[i]] = 1
	}
	var results []bson.M
	cursor, err := params.Db.Collection(params.Collection).Find(context.Background(), params.Filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	cursor.All(context.Background(), &results)
	return results, nil
}
