package adapters

import (
	"PyBot-DataServer/backup/domian/models"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
    conn, err := amqp.Dial(os.Getenv("URL_RABBIT"))
    if err != nil {
        // sólo logueamos, pero devolvemos el error para que el main siga vivo
        log.Printf("RabbitMQ: no se pudo conectar: %v", err)
        return nil, err
    }
    ch, err := conn.Channel()
    if err != nil {
        conn.Close()
        log.Printf("RabbitMQ: no se pudo abrir canal: %v", err)
        return nil, err
    }
    return &RabbitMQ{conn: conn, ch: ch}, nil
}

func (r *RabbitMQ) SendDataTables(d []models.DataTable) {
	payload, err := json.Marshal(d)
	failOnError(err, "Error al serializar Loan a JSON")
	r.prepareToMessage(payload)
}

func (r *RabbitMQ) prepareToMessage(body []byte) {
	// Declaración del exchange (intercambiador):
	r.ch.ExchangeDeclare(
		"Inserciones",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	  
	r.ch.PublishWithContext(ctx,
		"Inserciones",     // exchange
		"quainsbackup", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "application/json",
		  Body:        body,
		})
	log.Printf(" [x] Sent %s\n", body)
}


func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}