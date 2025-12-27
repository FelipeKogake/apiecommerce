package categoria

import (
	"context"
	"log"
	"strings"
	"errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type UseCase interface {
	Adicionar(ctx context.Context, categoria Categoria) error
	Atualizar(ctx context.Context, categoria Categoria) error
	Deletar(ctx context.Context, id string) error
	Buscar(ctx context.Context, id string) (Categoria, error)
	Listar(ctx context.Context) ([]Categoria, error)
	ListarQuantidadeProdutosCategoria(ctx context.Context) ([]Categoria, error)
}

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
} 

func (s *Service) Adicionar(ctx context.Context, categoria Categoria) error {
	//Ajustando nomes para entrar no banco
		//Tirando os espacos desnecessarios do nome
	nomeSemEspaco := strings.TrimSpace(categoria.Nome)

		//Colocando as letras em maiuculas
	cases := cases.Title(language.BrazilianPortuguese)
	categoria.Nome = cases.String(nomeSemEspaco)

	//DEVIA VALIDAR SE JA EXISTE NO BANCO DE DADOS
	filtro := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"nome", cases.String(nomeSemEspaco)},
			}},
		},
	}

	lista, err := s.Repo.Listar(ctx, filtro)
	if err != nil {
		return err
	} else if lista != nil {
		return errors.New("Already exists")
	}

	//Colocando no banco
	err2 := s.Repo.Adicionar(ctx, categoria)
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	return nil
}

func (s *Service) Atualizar(ctx context.Context, categoria Categoria) error {return nil}

func (s *Service) Deletar(ctx context.Context, id string) error {return nil}

func (s *Service) Buscar(ctx context.Context, id string) (Categoria, error) {

	produto, err := s.Repo.Buscar(ctx, id)
	if err != nil {
		log.Println(err)
	}

	return produto, nil
}

func (s *Service) Listar(ctx context.Context) ([]Categoria, error) {
	filtro := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{}},
		},
	}

	lista, err := s.Repo.Listar(ctx, filtro)
	if err != nil {
		log.Println(err)
	}
	
	return lista, nil
}

func (s *Service) ListarQuantidadeProdutosCategoria(ctx context.Context) ([]Categoria, error){

	filtro := mongo.Pipeline{
		bson.D{
			{"$group", bson.D{
					{"_id", "categoria.nome"},
					{"quantidade", bson.D{{"$sum", 1}}},
				},
			},
		},
	}

	lista, err := s.Repo.Listar(ctx, filtro)
	if err != nil {
		return lista, err
	}

	return lista, nil
}