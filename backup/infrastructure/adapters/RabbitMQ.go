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
    var conn *amqp.Connection
    var err error
    
    url := os.Getenv("URL_RABBIT")

    // --- INICIO DE LÓGICA DE REINTENTOS ---
    maxRetries := 5 // Intentaremos 5 veces
    retryDelay := 2 * time.Second // Esperaremos 2 segundos entre intentos

    for i := 0; i < maxRetries; i++ {
        conn, err = amqp.Dial(url)
        if err == nil {
            // Si no hay error, salimos del bucle
            log.Println("RabbitMQ: Conexión exitosa.")
            break
        }

        // Si hubo error, logueamos y esperamos
        log.Printf("RabbitMQ: Fallo al conectar (Intento %d/%d): %v. Reintentando en %v...", i+1, maxRetries, err, retryDelay)
        time.Sleep(retryDelay)
    }
    // --- FIN DE LÓGICA DE REINTENTOS ---

    // Si después de los 5 intentos sigue habiendo error, devolvemos el error al main
    if err != nil {
        log.Printf("RabbitMQ: No se pudo conectar después de %d intentos: %v", maxRetries, err)
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