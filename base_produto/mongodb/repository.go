package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"context"
	"apiecommerce2/base_produto"
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

func (r *Repository) Adicionar(ctx context.Context, produto produto.Produto) error {
	if _,err := r.collection.InsertOne(ctx, produto); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Atualizar(ctx context.Context, produtoUp produto.Produto, filtro bson.M) error {
	_, err := r.collection.UpdateOne(ctx, filtro, bson.M{
		"$set": bson.M{
			"nome": produtoUp.Nome,
			"preco": produtoUp.Preco,
			"categoria": produtoUp.Categoria,
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Deletar(ctx context.Context, filtro bson.M) error {
	
	if _, err := r.collection.DeleteOne(ctx, filtro); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Buscar(ctx context.Context, idIn string) (produto.Produto, error) {
var resultado produto.Produto

	id, err := bson.ObjectIDFromHex(idIn)
	if err != nil {
		log.Println("tamo aqui", err)
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": id}) 
	if err := result.Err(); err != nil {
		log.Println(err)
	}
	result.Decode(&resultado)

	return resultado, nil
}

func (r *Repository) Listar(ctx context.Context, filtro mongo.Pipeline) ([]bson.D, error) {
	var resultado []bson.D
	cursor, err := r.collection.Aggregate(ctx, filtro)
	if err != nil {
		log.Println(err)
		return resultado, err
	}
	for cursor.Next(ctx) {
		var produto bson.D
		cursor.Decode(&produto)
		resultado = append(resultado, produto)
	}
	defer cursor.Close(ctx)

	return resultado, nil
}

func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 

