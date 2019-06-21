package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IAPIResponse struct {
	Error   bool        `json:"error"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

type IMongoBaseInputParams struct {
	Db         *mongo.Database
	Collection string
}

type IMongoFindInputParams struct {
	IMongoBaseInputParams
	Filter     interface{}
	Projection []string
	Limit      uint
	Skip       uint
	Sort       map[string]int
}
