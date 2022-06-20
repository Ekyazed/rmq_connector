package core

import (
	database "app_connector/databases"
	"app_connector/databases/api"
	"log"
)

func DeleteAction(d QueueData) {
	switch d.Table {
	case "User":
		user := database.UserModelConvertion(d.Data)
		api.DeleteUser(d.ToApp, user)

	case "Client":
		client := database.ClientModelConvertion(d.Data)
		api.DeleteClient(d.ToApp, client)

	case "Site":
		site := database.SiteModelConvertion(d.Data)
		api.DeleteSite(d.ToApp, site)

	default:
		log.Printf("Unable to handle table %s", d.Table)
	}
}
