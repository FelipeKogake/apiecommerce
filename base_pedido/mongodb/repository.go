package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	// "apiecommerce2/internal/usuario"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	// "time"
	"apiecommerce2/base_carrinho"
	"apiecommerce2/base_pedido"
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
	var lista []bson.D

	cursor, err := r.collection.Aggregate(ctx, filtro)
	if err != nil {
		log.Println(err)
		return lista, err
	}

	for cursor.Next(ctx) {
		var pedido bson.D
		cursor.Decode(&pedido)
		lista = append(lista, pedido)
	}
	defer cursor.Close(ctx)

	return lista, nil
}

func (r *Repository) BuscarPedido(ctx context.Context, filtro bson.M)  (pedido.Pedido, error) {
	var pedido pedido.Pedido

	result := r.collection.FindOne(ctx, filtro)
	if err := result.Err(); err != nil {
		log.Println(err)
		return pedido, err
	}

	result.Decode(&pedido)

	return pedido, nil
}

func (r *Repository) CriarPedido(ctx context.Context, carrinho carrinho.Carrinho, pedido pedido.Pedido) error {

	_, err2 := r.collection.InsertOne(ctx, pedido)
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	return nil
}

func (r *Repository) AtualizarStatus(ctx context.Context, filtro bson.M, update bson.M) error {
	if _, err := r.collection.UpdateOne(ctx, filtro, update); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 