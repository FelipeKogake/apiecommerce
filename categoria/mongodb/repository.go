package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"context"
	"apiecommerce2/categoria"
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

func (r *Repository) Adicionar(ctx context.Context, categoria categoria.Categoria) error {
	if _,err := r.collection.InsertOne(ctx, categoria); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Atualizar(ctx context.Context, produto categoria.Categoria) error {return nil}

func (r *Repository) Deletar(ctx context.Context, id string) error {return nil}

func (r *Repository) Buscar(ctx context.Context, id string) (categoria.Categoria, error) {return categoria.Categoria{}, nil}

func (r *Repository) Listar(ctx context.Context) ([]categoria.Categoria, error) {return []categoria.Categoria{}, nil}

func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 