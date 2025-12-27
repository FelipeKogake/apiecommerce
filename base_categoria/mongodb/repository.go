package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"apiecommerce2/base_categoria"
)

type CategoriaDB struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	Nome string `bson:"nome"`
}

func (c *CategoriaDB) fromDomain(categoria categoria.Categoria) error {
	c.ID = categoria.ID
	c.Nome = categoria.Nome
	return nil
}

func (c CategoriaDB) toDomain(categoriaDomain *categoria.Categoria) error {
	categoriaDomain.ID = c.ID
	categoriaDomain.Nome = c.Nome
	return nil
}

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
	var categoriaDB CategoriaDB
	categoriaDB.fromDomain(categoria)

	if _,err := r.collection.InsertOne(ctx, categoriaDB); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Atualizar(ctx context.Context, produto categoria.Categoria) error {return nil}

func (r *Repository) Deletar(ctx context.Context, id string) error {return nil}

func (r *Repository) Buscar(ctx context.Context, idIn string) (categoria.Categoria, error) {
	var resultado CategoriaDB

	id, err := bson.ObjectIDFromHex(idIn)
	if err != nil {
		log.Println("tamo aqui", err)
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": id}) 
	if err := result.Err(); err != nil {
		log.Println(err)
	}
	result.Decode(&resultado)

	var categoria categoria.Categoria

	if err := resultado.toDomain(&categoria); err != nil {
		return categoria, err
	}

	return categoria, nil
}


func (r *Repository) Listar(ctx context.Context, filtro mongo.Pipeline) ([]categoria.Categoria, error) {
	var resultado []categoria.Categoria
	cursor, err := r.collection.Aggregate(ctx, filtro)
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(ctx) {
		var categoriaDB CategoriaDB
		var categoria categoria.Categoria
		if err2 := cursor.Decode(&categoriaDB); err2 != nil {
			return resultado, err
		}

		if err := categoriaDB.toDomain(&categoria); err != nil {
			return resultado, err
		}

		resultado = append(resultado, categoria)
	}
	defer cursor.Close(ctx)

	return resultado, nil
}

func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 