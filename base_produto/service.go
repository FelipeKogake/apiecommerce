package produto

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var validador = validator.New()

type UseCase interface {
	Adicionar(ctx context.Context, produto Produto) error
	Atualizar(ctx context.Context, produto Produto) error
	Deletar(ctx context.Context, id string) error
	Buscar(ctx context.Context, id string) (Produto, error)
	Listar(ctx context.Context, skip int64) ([]bson.D, error)
	ListarQuantidadesPorCategoria(ctx context.Context) ([]bson.D, error)
	BuscarPorNome(ctx context.Context, nome string) ([]bson.D, error)
	BuscaPorFaixaPreco(ctx context.Context, faixa FaixaPreco) ([]bson.D, error)
}



type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}	
}

func (s *Service) Adicionar(ctx context.Context, produto Produto) error{

	//Usando a struct de validacao para verificar os campos
	if err := validador.Struct(produto); err != nil {
		log.Println(err)
		return err
	} 
	
	//Ajustando nomes para entrar no banco
		//Tirando os espacos desnecessarios do nome
	nomeSemEspaco := strings.TrimSpace(produto.Nome)

		//Colocando as letras em maiuculas
	cases := cases.Title(language.BrazilianPortuguese)
	produto.Nome = cases.String(nomeSemEspaco)

	//DEVIA VALIDAR SE JA EXISTE NO BANCO DE DADOS
		filtro := mongo.Pipeline{
			bson.D{
				{"$match", bson.D{
					{"$text", bson.D{
						{"$search", produto.Nome},
					}},
				}},
			},
		}

	lista, err := s.Repo.Listar(ctx, filtro)
	if err != nil {
		return err
	} else if lista != nil {
		return errors.New("Already Exist")
	}

	//Colocando no banco
	err2 := s.Repo.Adicionar(ctx, produto)
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	return nil
} 

func (s *Service) Atualizar(ctx context.Context, produto Produto) error {
	filtro := bson.M{"_id": produto.ID}

	if err := s.Repo.Atualizar(ctx, produto, filtro); err != nil {
		return err
	}

	return nil
}

func (s *Service) Deletar(ctx context.Context, idIn string) error {

	id, err := bson.ObjectIDFromHex(idIn)
	if err != nil {
		log.Println(err)
	}
	filtro := bson.M{"_id": id}

	if err := s.Repo.Deletar(ctx, filtro); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) Buscar(ctx context.Context, id string) (Produto, error) {
	produto, err := s.Repo.Buscar(ctx, id)

	if err != nil {
		log.Println(err)
	}

	return produto, nil
}
	
func (s *Service) Listar(ctx context.Context, skip int64) ([]bson.D, error) {

	filtro := mongo.Pipeline{
		bson.D{
			{"$sort", bson.D{
				{"nome", 1},
			}},
		},
		bson.D{
			{"$skip", skip},
		},
		bson.D{
			{"$limit", 4},
		},
	}

	result, _ := s.Repo.Listar(ctx, filtro)
	
	return result, nil
}

func (s *Service) ListarQuantidadesPorCategoria(ctx context.Context) ([]bson.D, error) {
	filtro := mongo.Pipeline{
		bson.D{
			{"$group", bson.D{
					{"_id", "$categoria.nome"},
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

func (s *Service) BuscarPorNome(ctx context.Context, nome string) ([]bson.D, error) {
	filtro := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"$text", bson.D{
					{"$search", nome},
				}},
			}},
		},
	}

	result, _ := s.Repo.Listar(ctx, filtro)
	
	return result, nil
}

func (s *Service) BuscaPorFaixaPreco(ctx context.Context, faixa FaixaPreco) ([]bson.D, error) {
	filtro := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"preco", bson.D{
					{"$gte", faixa.Min},
					{"$lte", faixa.Max},
				}},
			}},
		},
	}

	result, _ := s.Repo.Listar(ctx, filtro)
	
	return result, nil
}
