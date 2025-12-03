package produto

import (
	"context"
	"log"
	"strings"
	"golang.org/x/text/cases"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
)

var validador = validator.New()

type UseCase interface {
	Adicionar(ctx context.Context, produto Produto) error
	Atualizar(ctx context.Context, produto Produto) error
	Deletar(ctx context.Context, id string) error
	Buscar(ctx context.Context, id string) (Produto, error)
	Listar(ctx context.Context) ([]Produto, error)
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

	//Colocando no banco
	err := s.Repo.Adicionar(ctx, produto)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
} 

func (s *Service) Atualizar(ctx context.Context, produto Produto) error {return nil}

func (s *Service) Deletar(ctx context.Context, id string) error {return nil}

func (s *Service) Buscar(ctx context.Context, id string) (Produto, error) {return Produto{}, nil}
	
func (s *Service) Listar(ctx context.Context) ([]Produto, error) {return []Produto{}, nil}

