package categoria

import (
	"context"
	"strings"
	"log"
	"golang.org/x/text/language"
	"golang.org/x/text/cases"
)

type UseCase interface {
	Adicionar(ctx context.Context, categoria Categoria) error
	Atualizar(ctx context.Context, categoria Categoria) error
	Deletar(ctx context.Context, id string) error
	Buscar(ctx context.Context, id string) (Categoria, error)
	Listar(ctx context.Context) ([]Categoria, error)
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

	//Colocando no banco
	err := s.Repo.Adicionar(ctx, categoria)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) Atualizar(ctx context.Context, categoria Categoria) error {return nil}

func (s *Service) Deletar(ctx context.Context, id string) error {return nil}

func (s *Service) Buscar(ctx context.Context, id string) (Categoria, error) {return Categoria{}, nil}

func (s *Service) Listar(ctx context.Context) ([]Categoria, error) {return []Categoria{}, nil}
