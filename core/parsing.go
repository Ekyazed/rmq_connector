package core

import (
	"app_connector/utils"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type QueueData struct {
	ToApp  string      `json:"to_app"`
	Table  string      `json:"table"`
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func ParsingStruct(d amqp.Delivery) {

	var delivery QueueData
	err := json.Unmarshal([]byte(d.Body), &delivery)
	utils.FailOnError(err, "Cannot Unmarshal JSON "+string(d.Body))

	/**
	 * L'objectif a partir d'ici est de décanté le json reçut pour traité les données.
	 * On vérifie en premier lieu l'action a effectuer
	 * On se connecte ensuite à la base de données de l'application visé
	 * on parse les datas obtenu pour effectuer la requete sans erreurs
	 * On effectue la requête sur la table avec les données que nous avons dans notre possession
	 *
	 * ! Le but est de limité au maximum les actions humaines sur les bases de données
	 */

	switch delivery.Action {
	case "CREATE":
		CreateAction(delivery)

	case "UPDATE":
		//action de mise a jour

	case "DELETE":
		//action de suppression

	default:
		fmt.Printf("Unable to do the following action: %s", delivery.Action)
	}
}
