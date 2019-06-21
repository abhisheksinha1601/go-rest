package bl

import (
	"rest/interfaces"
	"rest/modules/common"
	"rest/modules/mongo"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(username, password string) (interface{}, error) {
	params := interfaces.IMongoFindInputParams{
		IMongoBaseInputParams: interfaces.IMongoBaseInputParams{
			Db:         mongo.GetMongoDb(),
			Collection: "users",
		},
		Filter: bson.M{"username": strings.ToLower(username), "password": common.Pbkdf2Encrypt(password, username)},
		Limit:  1,
	}
	users, err := mongo.Find(params)
	if err != nil || len(users) == 0 {
		return nil, err
	}

	miniUser := map[string]interface{}{
		"_id":        users[0]["_id"].(primitive.ObjectID).Hex(),
		"role":       users[0]["role"],
		"name":       users[0]["name"],
		"client":     users[0]["client"],
		"username":   users[0]["username"],
		"email":      users[0]["email"],
		"accessList": users[0]["accessList"],
	}

	result := map[string]interface{}{
		"_id":    miniUser["_id"],
		"role":   miniUser["role"],
		"name":   miniUser["name"],
		"client": miniUser["client"],
		"token":  common.GetJWTToken(miniUser),
	}
	return result, nil
}
