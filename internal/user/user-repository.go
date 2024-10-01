package user

import (
	"context"
	"errors"
	"log"
	"slices"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/domain"
)

// base de datos creada como estructura con campos users(del package domain) y maxUserID
type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

// se genera una interfaz Repository y una estructura
type (
	//la interface tiene dos metodos, create y get all
	Repository interface {
		//Añade un nuevo usuario al repositorio y permite gestionar el contexto y los errores durante la operaciòn
		//las request que ingresan deben crear un contexto, y las response deben aceptar un contexto.
		//el contexto se usa para controlar el tiempo de vida de una operaciòn
		//el puntero al objeto user permite modificar el objeto original que se pasa a la funciòn
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id uint64) (*domain.User, error)
	}

	/*estructura de llamada repo que se utiliza para implementar la interfz Repository,
	almacena una instancia de DB (usuarios y contador maximo de ID)
	y un puntero a un objeto log.logger */
	repo struct {
		db  DB
		log *log.Logger
	}
)

func NewRepo(db DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

// la funcion opera sobre una instancia de repo por lo cual modifica sus datos
func (r *repo) Create(ctx context.Context, user *domain.User) error {
	//incrementa en uno el valor maximo de el id
	r.db.MaxUserID++
	//crea un ID del usuario con el valor maximo
	user.ID = r.db.MaxUserID
	//agrega a los ususarios de la bd del repositorio el usuario creado
	r.db.Users = append(r.db.Users, *user)
	//el logger imprime la fecha y hora de el mensaje junto con el texto del mensaje
	r.log.Println("repository created")

	return nil
}

// la funcion opera sobre una instancia de repo por lo cual modifica sus datos
func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	//el logger imprime la fecha y hora de el mensaje junto con el texto del mensaje
	r.log.Println("repository get all")
	//devuelve el slice de los usuarios de la base de datos del repositorio y el error
	return r.db.Users, nil
}

func (r *repo) Get(ctx context.Context, id uint64) (*domain.User, error) {
	index := slices.IndexFunc(r.db.Users, func(v domain.User) bool {
		return v.ID == id
	})

	if index < 0 {
		return nil, errors.New("User not found, doesn't exist")
	}
	return &r.db.Users[index], nil
}
