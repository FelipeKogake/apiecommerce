package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"apiecommerce2/base_carrinho"
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

func (r *Repository) ListarItens(ctx context.Context, pipeline mongo.Pipeline) ([]carrinho.ItemCarrinho, error) {
	var itensCarrinho []carrinho.ItemCarrinho
	var carrinho carrinho.Carrinho
	

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(ctx) {
		cursor.Decode(&carrinho)
	}
	defer cursor.Close(ctx)

	for  _, item := range carrinho.Itens {
		itensCarrinho = append(itensCarrinho, item)
	}


	return itensCarrinho, nil
}

func (r *Repository) Atualizar(ctx context.Context, filtro bson.M, update bson.M) error {
	_, err := r.collection.UpdateOne(ctx, filtro, update)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (r *Repository) CriarCarrinho(ctx context.Context, carrinho carrinho.Carrinho) error {

	_, err := r.collection.InsertOne(ctx, carrinho)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (r *Repository) RemoverCarrinho(ctx context.Context, filtro bson.M) (carrinho.Carrinho, error) {
	var carrinho carrinho.Carrinho

	 r.collection.FindOneAndDelete(ctx, filtro).Decode(&carrinho)
	
	return carrinho, nil
}


func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 

func (r *Repository) Listar(ctx context.Context, pipeline mongo.Pipeline) ([]bson.D, error) {
	var resultado []bson.D
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return resultado, err
	}

	for cursor.Next(ctx) {
		var item bson.D
		cursor.Decode(&item)
		resultado = append(resultado, item)
	}
	defer cursor.Close(ctx)

	return resultado, nil
}	

