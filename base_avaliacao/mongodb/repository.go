package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"context"
	"apiecommerce2/base_avaliacao"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repository struct {
	collection *mongo.Collection
	Client *mongo.Client
}

func NewRepository(database, collection, uri string) (*Repository, error){
	r := &Repository{}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Deu errado na conexao com o servidor", err)
	}	

	r.collection = client.Database(database).Collection(collection)
	r.Client = client

	return r, nil 
}	

func (r *Repository) Listar(ctx context.Context, filtro mongo.Pipeline) ([]bson.D, error) {
	var resultado []bson.D

	cursor, err := r.collection.Aggregate(ctx, filtro)
	if err != nil {
		log.Println(err)
		return resultado, err
	}

	for cursor.Next(ctx) {
		var avaliacao bson.D
		cursor.Decode(&avaliacao)
		resultado = append(resultado, avaliacao)
	}
	defer cursor.Close(ctx)

	return resultado, nil
}

func (r *Repository)AdicionarAvaliacao(ctx context.Context, avaliacao avaliacao.Avaliacao) error {

	_, err := r.collection.InsertOne(ctx, avaliacao)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 