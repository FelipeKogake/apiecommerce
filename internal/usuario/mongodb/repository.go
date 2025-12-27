package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"apiecommerce2/internal/usuario"
	"errors"
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

func (r *Repository) BuscarUsuarioPorId(ctx context.Context, idIn string) (usuario.Usuario, error) {
	var usu usuario.Usuario

	id, err := bson.ObjectIDFromHex(idIn)
	if err != nil {
		log.Println(err)
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": id})
	if err := result.Decode(&usu); err != nil {
		log.Println("felipe",err)
	}

	usuFinal := usuario.Usuario{
		ID: usu.ID,
		Nome: usu.Nome,
		Senha: usu.Senha,
	}
	return usuFinal, nil
}

func (r *Repository) BuscarUsuarioPorNome(ctx context.Context, nome string) (usuario.Usuario, error) {
	var usu usuario.Usuario
	result := r.collection.FindOne(ctx, bson.M{"nome": nome})
	if err := result.Decode(&usu); err != nil {
		log.Println(err)
	}
	usuFinal := usuario.Usuario{
		ID: usu.ID,
		Nome: usu.Nome,
		Senha: usu.Senha,
	}
	return usuFinal, nil
}

func (r *Repository) ValidarUsuario(ctx context.Context, credenciais usuario.Usuario) (usuario.Usuario, error) {
	usu, err := r.BuscarUsuarioPorNome(ctx, credenciais.Nome)

	if usu.Nome == "" {
		log.Println(err)
		return usu, err
	} else if usu.Senha != credenciais.Senha {
		return usu, errors.New("senhas incorretas")
	}

	return usu, nil
}

func (r *Repository) Close(ctx context.Context) error {
	r.Client.Disconnect(ctx)
	return nil
} 

func (r *Repository) CriarUsuario(ctx context.Context, usuario usuario.Usuario) error {
	_, err := r.collection.InsertOne(ctx, usuario)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (r *Repository) Atualizar(ctx context.Context, filtro bson.D, update bson.D) error {

	_, err := r.collection.UpdateOne(ctx, filtro, update)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) DeletarUsuario(ctx context.Context, filtro bson.M) error {

	if _, err := r.collection.DeleteOne(ctx, filtro); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Listar(ctx context.Context, pipeline mongo.Pipeline) ([]bson.D, error) {
	var resultado []bson.D
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(ctx) {
		var elemento bson.D
		err := cursor.Decode(&elemento)
		if err != nil {
			return resultado, err
		}
		resultado = append(resultado, elemento)
	}
	defer cursor.Close(ctx)

	return resultado, nil
}