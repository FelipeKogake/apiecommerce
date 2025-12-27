package pedido

import (
	"context"
	// "strings"
	"log"
	// "golang.org/x/text/language"
	// "golang.org/x/text/cases"
	"apiecommerce2/base_carrinho"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UseCase interface {
	ListarPedidos(ctx context.Context, id bson.ObjectID) ([]bson.D, error)
	BuscarPedido(ctx context.Context, id bson.ObjectID) (Pedido, error)
	AtualizarStatus(ctx context.Context, idUsu bson.ObjectID, idPed bson.ObjectID, status string) error
	CriarPedido(ctx context.Context, carrinho carrinho.Carrinho, id bson.ObjectID, endereco string) error
	ProdutosMaisVendidos(ctx context.Context) ([]bson.D, error)
	VendasPorPeriodo(ctx context.Context) ([]bson.D, error)
}

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) ListarPedidos(ctx context.Context, id bson.ObjectID) ([]bson.D, error) {

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"carrinho.usuario._id", id},
			}},
		},
	}

	lista, _ := s.Repo.Listar(ctx, pipeline)

	return lista, nil
}

func (s *Service) BuscarPedido(ctx context.Context, id bson.ObjectID) (Pedido, error) {

	filtro := bson.M{
		"_id": id,
	}

	pedido, _ := s.Repo.BuscarPedido(ctx, filtro)

	return pedido, nil
}

func (s *Service) AtualizarStatus(ctx context.Context, id bson.ObjectID, idPed bson.ObjectID, status string) error {
	filtro := bson.M{
		"_id": idPed,
		"carrinho.usuario._id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}

	if err := s.Repo.AtualizarStatus(ctx, filtro, update); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) CriarPedido(ctx context.Context, carrinho carrinho.Carrinho, id bson.ObjectID, endereco string) error {
	
	pedido := Pedido{
		Data: time.Now().UTC(),
		Carrinho: carrinho,
		Status: "Aberto",
		Endereco: endereco,
	}

	if err := s.Repo.CriarPedido(ctx, carrinho, pedido); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) ProdutosMaisVendidos(ctx context.Context) ([]bson.D, error) {
	pipeline := mongo.Pipeline{
		bson.D{
			{"$unwind", "$carrinho.itens"},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", "$carrinho.itens.produto.nome"},
				{"VendasTotais", bson.D{
					{"$sum", "$carrinho.itens.quantidade"},
				}},
			}},
		},
		bson.D{
			{"$sort", bson.D{
				{"VendasTotais", -1},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_id", 0,},
				{"produto", "$_id"},
				{"VendasTotais", 1},
			}},
		},
	}

	lista, _ := s.Repo.Listar(ctx, pipeline)

	return lista, nil
}

func (s *Service) VendasPorPeriodo(ctx context.Context) ([]bson.D, error) {
	inicioDoAno := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	inicioDoProximoAno := time.Date(time.Now().Year() + 1, 1, 1, 0, 0, 0, 0, time.UTC)

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"data", bson.D{
					{"$gte", inicioDoAno},
					{"$lte", inicioDoProximoAno},
				}},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", bson.D{
					{"$month", "$data"},
				}},
				{"NumVendas", bson.D{
					{"$sum", 1},
				}},
				{"ValorTotal", bson.D{
					{"$sum", "$total"},
				}},
			}},
		},
		bson.D{
			{"$sort", bson.D{
				{"_id", 1},
			}},
		},
	}

	lista, _ := s.Repo.Listar(ctx, pipeline)

	return lista, nil
}