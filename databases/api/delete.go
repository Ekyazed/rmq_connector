package api

import (
	database "app_connector/databases"
	"app_connector/model"
	"app_connector/utils"
	"fmt"
)

func DeleteUser(dbName string, deleteingUser model.User) {

	/**
	 * Connexion à la base de données
	 * vérification de l'existance de l'utilisateur.
	 * supprimé l'utilisateur
	 */

	database.ConnectToDatabase(dbName)

	var user model.User
	database.DB.Where("email = ?", deleteingUser.Email).Find(&user)
	if user.ID == 0 {
		fmt.Println("The user identified by email " + deleteingUser.Email + " cannot be found")
		//TODO: envoi d'un email à l'utilisateur
		//TODO: envoi d'une notification aux administrateurs (CRM)
		return
	}

	result := database.DB.Where("email = ?", deleteingUser.Email).Delete(&user)
	if result.RowsAffected == 0 {
		fmt.Println("Deletion failled for user identified by email " + deleteingUser.Email)
		return
	} else if result.Error != nil {
		utils.FailOnError(result.Error, "error while deleting user identified by email "+deleteingUser.Email)
		fmt.Println("Deletion failled for user identified by email " + deleteingUser.Email)
		return
	}
}

func DeleteClient(dbName string, deletingClient model.Client) {

	/**
	 * Connexion à la base de données
	 * Vérification de l'existance du client
	 * supprimé le client
	 */

	database.ConnectToDatabase(dbName)

	var site model.Client
	database.DB.Where("id = ?", deletingClient.ID).Find(&site)
	if site.ID == 0 {
		fmt.Println("The client with name " + deletingClient.Name + " cannot be found")
		//TODO: envoi d'un email à l'utilisateur
		//TODO: envoi d'une notification aux administrateurs (CRM)
		return
	}

	result := database.DB.Where("id = ?", deletingClient.ID).Delete(&site)
	if result.RowsAffected == 0 {
		fmt.Println("Deletion failled for client with name " + deletingClient.Name)
		return
	} else if result.Error != nil {
		utils.FailOnError(result.Error, "error while deleting client with name "+deletingClient.Name)
		fmt.Println("Deletion failled for client with name " + deletingClient.Name)
		return
	}
}

func DeleteSite(dbName string, deletingSite model.Site) {

	/**
	 * Connexion à la base de données
	 * Vérification de l'existance du site
	 * supprimé le site
	 */

	database.ConnectToDatabase(dbName)

	var site model.Site
	database.DB.Where("id = ?", deletingSite.ID).Find(&site)
	if site.ID == 0 {
		fmt.Println("The site with address " + deletingSite.Address + " cannot be found")
		//TODO: envoi d'un email à l'utilisateur
		//TODO: envoi d'une notification aux administrateurs (CRM)
		return
	}

	result := database.DB.Where("id = ?", deletingSite.ID).Delete(&site)
	if result.RowsAffected == 0 {
		fmt.Println("Deletion failled for site with address " + deletingSite.Address)
		return
	} else if result.Error != nil {
		utils.FailOnError(result.Error, "error while deleting site with address "+deletingSite.Address)
		fmt.Println("Deletion failled for site with address " + deletingSite.Address)
		return
	}
}
