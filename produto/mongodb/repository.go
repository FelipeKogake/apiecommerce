package banco

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"context"
	"apiecommerce2/produto"
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

func (r *Repository) Atualizar(ctx context.Context, produto produto.Produto) error {return nil}

func (r *Repository) Deletar(ctx context.Context, id string) error {return nil}

func (r *Repository) Buscar(ctx context.Context, id string) (produto.Produto, error) {return produto.Produto{}, nil}

func (r *Repository) Listar(ctx context.Context) ([]produto.Produto, error) {return []produto.Produto{}, nil}

func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 

