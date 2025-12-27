package carrinho

import (
	"apiecommerce2/internal/usuario"
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
	AdicionarItem(ctx context.Context, item ItemCarrinho, id bson.ObjectID) error
	ListarItens(ctx context.Context, id bson.ObjectID) ([]ItemCarrinho, error)
	AtualizarQuantidade(ctx context.Context, valor int, qtd int, id bson.ObjectID) error
	RemoverItem(ctx context.Context, valor int, id bson.ObjectID) error
	CriarCarrinho(ctx context.Context, id bson.ObjectID) error 
	RemoverCarrinho(ctx context.Context, id bson.ObjectID) (Carrinho, error)
	ValorTotal(ctx context.Context, id bson.ObjectID) (bson.D, error)
}

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) AdicionarItem(ctx context.Context, item ItemCarrinho, id bson.ObjectID) error {

	//Ajustando nomes para entrar no banco
		//Tirando os espacos desnecessarios do nome
	nomeSemEspaco := strings.TrimSpace(item.Produto.Nome)

		//Colocando as letras em maiuculas
	cases := cases.Title(language.BrazilianPortuguese)
	item.Produto.Nome = cases.String(nomeSemEspaco)

	//DEVIA VALIDAR SE JA EXISTE NO BANCO DE DADOS
	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"usuario._id", id},
			}},
		},
	}

	carrinho, _ := s.Existe(ctx, id)

	if carrinho == nil {
		s.CriarCarrinho(ctx, id)
	}

	lista, _ := s.Repo.ListarItens(ctx, pipeline);

	filtro := bson.M{"usuario._id": id}

	update := bson.M{
		"$push": bson.M{
			"itens": bson.M{
				"produto": item.Produto,
				"quantidade": item.Quantidade,
			},
		},
	}

	//Colocando no banco
	if lista != nil {
		for i, elemento := range lista {
			j := i + 1
			if elemento.Produto.ID == item.Produto.ID{
				log.Println(elemento.Quantidade + item.Quantidade)
				err2 := s.AtualizarQuantidade(ctx, j,elemento.Quantidade + item.Quantidade,  id)
				if err2 != nil {
					log.Println(err2)
				}
				return nil
			}
		}
	}
	

	if err := s.Repo.Atualizar(ctx, filtro, update); err != nil{
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) ListarItens(ctx context.Context, id bson.ObjectID) ([]ItemCarrinho, error) {
	var resultado []ItemCarrinho

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"usuario._id", id},
			}},
		},
	}

	itens, err2 := s.Repo.ListarItens(ctx, pipeline)
	if err2 != nil {
		log.Println(err2)
	}

	for _, item := range itens {
		resultado = append(resultado, item)
	}

	return resultado, nil
}

func (s *Service) AtualizarQuantidade(ctx context.Context, valor int, qtd int, id bson.ObjectID) error {
	var itemCar ItemCarrinho

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{}},
		},
	}

	lista, err := s.Repo.ListarItens(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return err
	}
	
	for i, item := range lista {
		j := i+1
		if valor == j {
			itemCar = ItemCarrinho{
				Produto: item.Produto,
				Quantidade: item.Quantidade,
			}
		}
	}


	filtro := bson.M{"itens.produto._id": itemCar.Produto.ID}

	update := bson.M{
    "$set": bson.M{
        // A sintaxe chave: "array.placeholder.campo"
        "itens.$.quantidade": qtd,
    },
}

	if err := s.Repo.Atualizar(ctx, filtro, update); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) RemoverItem(ctx context.Context, valor int, id bson.ObjectID) error {
	var itemCar ItemCarrinho

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"usuario._id", id},
			}},
		},
	}

	lista, err := s.Repo.ListarItens(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return err
	}
	
	for i, item := range lista {
		j := i+1
		if valor == j {
			itemCar = ItemCarrinho{
				Produto: item.Produto,
				Quantidade: item.Quantidade,
			}
		}
	}

	if itemCar.Produto.ID.IsZero() {
		return errors.New("O produto não foi encontrado.")
	}


	filtro := bson.M{"itens.produto._id": itemCar.Produto.ID}

	update := bson.M{
		"$pull": bson.M{
			"itens": bson.M{
				"produto._id": itemCar.Produto.ID,
			} ,
		},
	}

	if err2 := s.Repo.Atualizar(ctx, filtro, update); err2 != nil {
		log.Println(err2)
		return err2
	}

	return nil
}

func (s *Service) CriarCarrinho(ctx context.Context, id bson.ObjectID) error {

	carrinho := Carrinho{
		Itens: []ItemCarrinho{},
		Usuario: usuario.Usuario{
			ID: id,
		},
	}

	s.Repo.CriarCarrinho(ctx, carrinho)

	return nil
}

func (s *Service) RemoverCarrinho(ctx context.Context, id bson.ObjectID) (Carrinho, error) {
	filtro := bson.M{"usuario._id": id}

	carrinho, _ := s.Repo.RemoverCarrinho(ctx, filtro)
	

	return carrinho, nil
}

func (s *Service) ValorTotal(ctx context.Context, id bson.ObjectID) (bson.D, error) {	
	
	pipeline := mongo.Pipeline{
		
		bson.D{
			{"$match", bson.D{
				{"usuario._id", id},
			}},
		},
		bson.D{
			{"$unwind", "$itens"},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"ValorTotal", bson.D{
					{"$sum", bson.D{
						{"$multiply", bson.A{
							"$itens.produto.preco",
							"$itens.quantidade",
						}},
					}},
				}},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_id", 0},

			}},
		},
	}

	resultado, err := s.Repo.Listar(ctx, pipeline)
	if err != nil {
		log.Println(err)
	}
	
	if len(resultado) == 0 {
		return bson.D{{"ValorTotal", 0}}, err
	}

	return resultado[0], nil
}

func (s *Service) Existe(ctx context.Context, id bson.ObjectID) ([]bson.D, error) {
	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"usuario._id", id},
			}},
		},
	}

	lista, err := s.Repo.Listar(ctx, pipeline)
	if err != nil {
		log.Println(err)
	}

	return lista, nil
}