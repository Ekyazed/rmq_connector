package database

import (
	"app_connector/model"
	"app_connector/utils"
	"encoding/json"
)

func UserModelConvertion(data interface{}) model.User {
	var user model.User

	mapBytes, err := json.Marshal(data)
	utils.FailOnError(err, "Failed to marshal data for user model")

	json.Unmarshal(mapBytes, &user)
	return user
}

func SiteModelConvertion(data interface{}) model.Site {
	var site model.Site

	mapBytes, err := json.Marshal(data)
	utils.FailOnError(err, "Failed to marshal data for site model")

	json.Unmarshal(mapBytes, &site)
	return site
}

func ClientModelConvertion(data interface{}) model.Client {
	var client model.Client

	mapBytes, err := json.Marshal(data)
	utils.FailOnError(err, "Failed to marshal data for site model")

	json.Unmarshal(mapBytes, &client)
	return client
}
