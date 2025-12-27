package usuario

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type UseCase interface {
	Adicionar(ctx context.Context, usuario Usuario) error
	Atualizar(ctx context.Context, usuario Usuario) error
	Deletar(ctx context.Context, id bson.ObjectID) error
	ListarEnderecos(ctx context.Context, id bson.ObjectID) ([]bson.D, error)
	AdicionarEndereco(ctx context.Context, id bson.ObjectID ,endereco string) error
	AtualizarEndereco(ctx context.Context, id bson.ObjectID, endereco Endereco) error
	RemoverEndereco(ctx context.Context, id bson.ObjectID, indice string) error
}

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service{
	return &Service{
		Repo: repo,
	}
}

func (s *Service) Adicionar(ctx context.Context, usuario Usuario) error {
	//Ajustando nomes para entrar no banco
		//Tirando os espacos desnecessarios do nome
	nomeSemEspaco := strings.TrimSpace(usuario.Nome)

		//Colocando as letras em maiuculas
	cases := cases.Title(language.BrazilianPortuguese)
	usuario.Nome = cases.String(nomeSemEspaco)

	//DEVIA VALIDAR SE JA EXISTE NO BANCO DE DADOS
	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"$text", bson.D{
					{"$search", usuario.Nome},
				}},
			}},
		},
	}

	lista, err := s.Repo.Listar(ctx, pipeline)
	if err != nil {
		return err
	} else if lista != nil {
		return errors.New("Usuario já existente")
	}

	//Colocando no banco
	err3 := s.Repo.CriarUsuario(ctx, usuario)
	if err3 != nil {
		log.Println(err3)
		return err3
	}

	return nil
}

func (s *Service) Atualizar(ctx context.Context, usuario Usuario) error {

	filtro := bson.D{
		{"_id", usuario.ID},
	}

	update := bson.D{
		{"$set", bson.D{
			{"nome", usuario.Nome},
			{"senha", usuario.Senha},
		}},
	}


	s.Repo.Atualizar(ctx, filtro, update)

	return nil
}

func (s *Service) Deletar(ctx context.Context, id bson.ObjectID) error {

	filtro := bson.M{
		"_id": id,
	}

	if err := s.Repo.DeletarUsuario(ctx, filtro); err != nil {
		return err
	}

	return nil
}

func (s *Service) ListarEnderecos(ctx context.Context, id bson.ObjectID) ([]bson.D, error) {

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"_id", id},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"enderecos", 1},
			}},
		},
	}

	resultado, err := s.Repo.Listar(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return resultado, err
	}

	return resultado, nil
}

func (s *Service) AdicionarEndereco(ctx context.Context, id bson.ObjectID ,endereco string) error {

	filtro := bson.D{
		{"_id", id},	
	}

	update := bson.D{
		{"$push", bson.D{
			{"enderecos", endereco},
		}},
	}

	

	if err := s.Repo.Atualizar(ctx, filtro, update); err != nil {
		return err
	}

	return nil
}


func (s *Service) AtualizarEndereco(ctx context.Context, id bson.ObjectID, endereco Endereco) error {
	filtro := bson.D{
		{"_id", id},
	}	

	update := bson.D{
		{"$set", bson.D{
			{fmt.Sprintf("enderecos.%s", strconv.Itoa(endereco.Indice)), endereco.Endereco},
		}},
	}

	if err := s.Repo.Atualizar(ctx, filtro, update); err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoverEndereco(ctx context.Context, id bson.ObjectID, indice string) error {
	filtro := bson.D{
		{"_id", id},
	}

	update := bson.D{
		{"$unset", bson.D{
			{fmt.Sprintf("enderecos.%s", indice), 1},
		}},
	}

	if err := s.Repo.Atualizar(ctx, filtro, update); err != nil {
		return err
	}

	update2 := bson.D{
		{"$pull", bson.D{
			{"enderecos", nil},
		}},
	}

	if err2 := s.Repo.Atualizar(ctx, filtro, update2); err2 != nil {
		return err2
	}

	return nil
}
