package api

import (
	database "app_connector/databases"
	"app_connector/model"
	"app_connector/utils"
	"fmt"
	"strconv"
)

func CreateUser(dbName string, newUser model.User) {

	/**
	 * Connection à la base de données visé
	 * On verifie si le compte n'existe pas déjà en tant que soft delete
	 * S'il existe, alors on lui enleve sa suppression
	 * Dans le cas où il n'existe pas, on le crée.
	 */

	database.ConnectToDatabase(dbName)

	var existingUser model.User

	database.DB.Unscoped().Where("email = ?", newUser.Email).Find(&existingUser)
	if existingUser.ID != 0 && existingUser.DeletedAt.Valid {
		database.DB.Unscoped().Where("email = ?", newUser).Update("deleted_at", nil)
		fmt.Println("restored user: " + existingUser.Email + " on app " + dbName)
		return
	}

	result := database.DB.Create(&newUser)
	utils.FailOnError(result.Error, "Error While creating User for the "+dbName+" application")
	fmt.Println("Created user: " + newUser.Email + " on app " + dbName)
}

func CreateSite(dbName string, newSite model.Site) {

	/**
	 * Connexion à la base de données visé
	 * verification de si le site n'est pas déjà existant en soft delete
	 * s'il existe, on supprime le soft delete
	 * Dans le cas où il n'existe pas on le créer
	 */

	database.ConnectToDatabase(dbName)

	var existingSite model.Site

	database.DB.Unscoped().Where("client_id = ?", newSite.ClientID).Where("address = ?", newSite.Address).Find(&existingSite)
	if existingSite.ID != 0 && existingSite.DeletedAt.Valid {
		database.DB.Unscoped().Where("client_id = ?", newSite.ClientID).Where("address = ?", newSite.Address).Update("deleted_at", nil)
		fmt.Println("restored site of client : " + strconv.Itoa(int(existingSite.ClientID)) + " on app " + dbName)
		return
	}

	result := database.DB.Create(&newSite)
	utils.FailOnError(result.Error, "Error while creating Site for the "+dbName+"database")
	fmt.Println("Created site for client: " + strconv.Itoa(int(newSite.ClientID)) + " on app " + dbName)
}

func CreateClient(dbName string, newClient model.Client) {

	/**
	 * Connexion à la base de données visé
	 * verification de si le client n'est pas déjà existant en soft delete
	 * s'il existe, on supprime le soft delete
	 * Dans le cas où il n'existe pas on le créer
	 */

	database.ConnectToDatabase(dbName)

	var existingClient model.Client

	database.DB.Unscoped().Where("name = ?", newClient.Name).Where("uuid = ?", newClient.UUID).Find(&existingClient)
	if existingClient.ID != 0 && existingClient.DeletedAt.Valid {
		database.DB.Unscoped().Where("client_id = ?", newClient.Name).Where("uuid = ?", newClient.UUID).Update("deleted_at", nil)
		fmt.Println("Restored client: " + existingClient.Name + " on app " + dbName)

		return
	}

	result := database.DB.Create(&newClient)
	utils.FailOnError(result.Error, "Error while creating Site for the "+dbName+"database")
	fmt.Println("Created client: " + newClient.Name + " on app " + dbName)

}
