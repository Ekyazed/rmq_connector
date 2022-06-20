package core

import (
	database "app_connector/databases"
	"app_connector/databases/api"
	"log"
)

func UpdateAction(d QueueData) {

	switch d.Table {
	case "User":
		user := database.UserModelConvertion(d.Data)
		api.UpdateUser(d.ToApp, user)

	case "Client":
		client := database.ClientModelConvertion(d.Data)
		api.UpdateClient(d.ToApp, client)

	case "Site":
		site := database.SiteModelConvertion(d.Data)
		api.UpdateSite(d.ToApp, site)

	default:
		log.Printf("Unable to handle table %s", d.Table)
	}

}
