package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base de datos*/
var MongoCN = connectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://edgar9694:Edgar_9694@redsocial.iln0q.mongodb.net/twitter?retryWrites=true&w=majority")

/* connectDB() funcion que me permite conectar la base de datos */
func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa a la BD")
	return client
}

/* CheckConnection() es el ping a la base de datos*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return 1
}
