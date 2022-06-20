package core

import (
	database "app_connector/databases"
	"app_connector/databases/api"
	"log"
)

func CreateAction(d QueueData) {

	/**
	 * Ici, nous allons vérifier les informations de la table et des datas envoyé.
	 * Cette fonction a pour but simple de convertire les données reçut dans le bon model.
	 * Puis envoyé dans la fonction de traitement avec les différentes données
	 */

	switch d.Table {
	case "User":
		user := database.UserModelConvertion(d.Data)
		api.CreateUser(d.ToApp, user)

	case "Client":
		client := database.ClientModelConvertion(d.Data)
		api.CreateClient(d.ToApp, client)

	case "Site":
		site := database.SiteModelConvertion(d.Data)
		api.CreateSite(d.ToApp, site)

	default:
		log.Printf("Unable to handle table %s", d.Table)
	}

}
