package avaliacao

import (
	"context"
	// "strings"
	"log"
	// "golang.org/x/text/language"
	// "golang.org/x/text/cases"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

type UseCase interface {
	ListarAvalicoesProduto(ctx context.Context, idProduto bson.ObjectID) ([]bson.D, error)
	AdicionarAvaliacao(ctx context.Context, avaliacao Avaliacao) error
	MediaTempoReal(ctx context.Context) ([]bson.D, error)
}

func (s *Service) ListarAvalicoesProduto(ctx context.Context, idProduto bson.ObjectID) ([]bson.D, error) {


	filtro2 := mongo.Pipeline{
		bson.D{
			{"$match", 
				bson.D{
					{"produto._id", idProduto},
				},
			},
		},
	}


	lista, _ := s.Repo.Listar(ctx, filtro2)


	return lista, nil
}

func (s *Service) AdicionarAvaliacao(ctx context.Context, avaliacao Avaliacao) error {

	if err := s.Repo.AdicionarAvaliacao(ctx, avaliacao); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) MediaTempoReal(ctx context.Context) ([]bson.D, error) {
	pipeline := mongo.Pipeline{
		bson.D{
			{"$group", bson.D{
				{"_id", "$produto.nome"},
				{"mediaPontuacao", bson.D{
					{"$avg", "$pontos"},
				}},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"categoria", "$_id"},
				{"mediaPontuacao", 1},
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