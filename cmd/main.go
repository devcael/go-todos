package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	credential := options.Credential{
		Username: "teste",
		Password: "teste",
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao MongoDB!")
	return client, err
}

func Disconnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Desconectado do MongoDB!")
	return err
}

type Todo struct {
	Description string
	Checked     bool
}

type TodoErrado struct {
	Description string
	Checked     string
}

func main() {
	client, err := Connect()

	ctx := context.TODO()

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("teste").Collection("todos")
	db.DeleteMany(ctx, bson.D{{}})

	pre_todos := []interface{}{
		TodoErrado{
			Description: "Micael",
			Checked:     "Micael",
		},
		Todo{
			Description: "Micael Certo",
			Checked:     false,
		},
	}

	result, err := db.InsertMany(ctx, pre_todos)
	db.InsertOne(ctx, bson.D{{"description", "bar"}, {"checked", "world"}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Documentos inseridos: %v\n", len(result.InsertedIDs))

	cursor, err := db.Find(ctx, bson.D{{}})

	for cursor.Next(ctx) {
		var todo Todo

		err := cursor.Decode(&todo)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Descrição: %s\n", todo.Description)
		fmt.Printf("Status: %v\n", todo.Checked)
		fmt.Println("----------")
	}

	Disconnect(client)
}
