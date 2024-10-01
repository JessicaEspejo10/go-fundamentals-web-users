package user

import (
	"context"
	"log"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/domain"
)

type (
	//genera interfaz con sus metodos
	Service interface {
		//recibe el contexto y los campos que devuelve el repositorio como variables
		//su responsabilidad es enviar el usuario al repositorio
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id uint64) (*domain.User, error)
	}
	//genera estructura con logger y la interfaz de la capa repositorio
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	//crea una nueva instancia de una estructura de tipo service y retornando un puntero a esa estructura
	return &service{
		log:  l,
		repo: repo,
	}
}

// crea una puntero a una nueva instancia o estructura de domain.User inicializando los campos y lo devuelve junto con el valor del error
func (s service) Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// obtiene los usuarios del repositorio junto con el valor del error y los devuelve
func (s service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (s service) Get(ctx context.Context, id uint64) (*domain.User, error) {
	user, err := s.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}
