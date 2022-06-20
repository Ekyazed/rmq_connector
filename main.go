package main

import (
	"app_connector/utils"
	"log"
	"os"

	"github.com/streadway/amqp"
)

/**
 * Auteur: Simon Didier
 * date: 20/06/2022
 * description: Worker permettant la communication entre les différentes applications
 */

func main() {

	utils.LoadEnv()

	username := os.Getenv("AMQP_USERNAME")
	password := os.Getenv("AMQP_PASSWORD")
	host := os.Getenv("AMQP_HOST")
	port := os.Getenv("AMQP_PORT")

	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + host + ":" + port + "/")
	utils.FailOnError(err, "Failed to open dial")
	defer conn.Close()

	//TODO: récupération de toute les applications
	queues := []string{}

	forever := make(chan bool)
	for _, queue := range queues {

		//ouverture des channels
		ch, err := conn.Channel()
		utils.FailOnError(err, "unable to open channel")
		defer ch.Close()

		//création des queues
		q, err := ch.QueueDeclare(
			queue, //name
			true,  //durable
			false, //delete when unused
			false, //exclusive
			false, // no wait
			nil,   //arguments
		)
		utils.FailOnError(err, "unable to declare queue "+queue)

		//récupération des messages
		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		utils.FailOnError(err, "failed to register consumer")

		//lecture et traitement des messages
		go func() {
			for d := range msgs {
				log.Printf("message: %s from %s", d.Body, d.RoutingKey)
			}
		}()

	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
