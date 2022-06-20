package api

import (
	database "app_connector/databases"
	"app_connector/model"
	"app_connector/utils"
	"fmt"
)

func UpdateUser(dbName string, updatedUser model.User) {

	/**
	 * Connexion à la base de données visé
	 * récupération de l'utilisateur existant
	 * mettre à jour les données
	 *
	 * ! si l'utilisateur est introuvable notifier par envoi de mail le client
	 * ! et notifier les administrateur de l'application
	 */
	database.ConnectToDatabase(dbName)

	var user model.User
	database.DB.Where("email = ?", updatedUser.Email).Find(&user)
	if user.ID == 0 {
		fmt.Println("The user identified by email " + updatedUser.Email + " cannot be found")
		//TODO: envoi d'un email à l'utilisateur
		//TODO: envoi d'une notification aux administrateurs (CRM)
		return
	}

	result := database.DB.Model(user).Updates(&updatedUser)
	if result.RowsAffected == 0 {
		fmt.Println("update failled for user identified by email " + updatedUser.Email)
		return
	} else if result.Error != nil {
		utils.FailOnError(result.Error, "error while updating user identified by email "+updatedUser.Email)
		fmt.Println("update failled for user identified by email " + updatedUser.Email)
		return
	}
}

func UpdateSite(dbName string, updatedSite model.Site) {

	/**
	 * Connexion à la base de données visé
	 * récupération du site existant
	 * mettre à jour les données
	 *
	 * ! si l'utilisateur est introuvable notifier par envoi de mail le client
	 * ! et notifier les administrateur de l'application
	 */
	database.ConnectToDatabase(dbName)

	var site model.Site
	database.DB.Where("id = ?", updatedSite.ID).Find(&site)
	if site.ID == 0 {
		fmt.Println("The site with address " + updatedSite.Address + " cannot be found")
		//TODO: envoi d'un email à l'utilisateur
		//TODO: envoi d'une notification aux administrateurs (CRM)
		return
	}

	result := database.DB.Model(site).Updates(&updatedSite)
	if result.RowsAffected == 0 {
		fmt.Println("update failled for Site with address " + updatedSite.Address)
		return
	} else if result.Error != nil {
		utils.FailOnError(result.Error, "error while updating Site with address "+updatedSite.Address)
		fmt.Println("update failled for Site with address " + updatedSite.Address)
		return
	}
}

func UpdateClient(dbName string, updatedClient model.Client) {

	/**
	 * Connexion à la base de données visé
	 * récupération du site existant
	 * mettre à jour les données
	 *
	 * ! si l'utilisateur est introuvable notifier par envoi de mail le client
	 * ! et notifier les administrateur de l'application
	 */
	database.ConnectToDatabase(dbName)

	var site model.Site
	database.DB.Where("id = ?", updatedClient.ID).Find(&site)
	if site.ID == 0 {
		fmt.Println("The Client with name " + updatedClient.Name + " cannot be found")
		//TODO: envoi d'un email à l'utilisateur
		//TODO: envoi d'une notification aux administrateurs (CRM)
		return
	}

	result := database.DB.Model(site).Updates(&updatedClient)
	if result.RowsAffected == 0 {
		fmt.Println("update failled for Client with name " + updatedClient.Name)
		return
	} else if result.Error != nil {
		utils.FailOnError(result.Error, "error while updating Client with name "+updatedClient.Name)
		fmt.Println("update failled for Client with name " + updatedClient.Name)
		return
	}
}
