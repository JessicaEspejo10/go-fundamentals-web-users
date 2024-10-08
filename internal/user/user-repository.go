package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/domain"
)

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
		Update(ctx context.Context, id uint64, firstName, lastName, email *string) error
	}

	/*estructura de llamada repo que se utiliza para implementar la interfz Repository,
	almacena una instancia de DB (usuarios y contador maximo de ID)
	y un puntero a un objeto log.logger */
	repo struct {
		db  *sql.DB
		log *log.Logger
	}
)

func NewRepo(db *sql.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

// la funcion opera sobre una instancia de repo por lo cual modifica sus datos
func (r *repo) Create(ctx context.Context, user *domain.User) error {

	sqlQuery := "INSERT INTO users(first_name, last_name, email) VALUES(?, ?, ?)"
	res, err := r.db.Exec(sqlQuery, user.FirstName, user.LastName, user.Email)
	if err != nil {
		r.log.Println(err.Error())
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		r.log.Println(err.Error())
		return err
	}
	user.ID = uint64(id)
	r.log.Println("user created with id: ", id)
	return nil
}

// la funcion opera sobre una instancia de repo por lo cual modifica sus datos
func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	//devuelve el slice de los usuarios de la base de datos del repositorio y el error
	var users []domain.User
	sqlQuery := "SELECT id, first_name, last_name, email FROM users"
	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		r.log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			r.log.Println(err.Error())
			return nil, err
		}
		users = append(users, u)
	}

	r.log.Println("user get all: ", len(users))
	return users, nil

}

func (r *repo) Get(ctx context.Context, id uint64) (*domain.User, error) {
	/*index := slices.IndexFunc(r.db.Users, func(v domain.User) bool {
		return v.ID == id
	})

	if index < 0 {
		return nil, ErrorNotFound{id}
	}
	return &r.db.Users[index], nil
	*/
	return nil, nil
}

func (r *repo) Update(ctx context.Context, id uint64, firstName, lastName, email *string) error {
	/*
		user, err := r.Get(ctx, id)
		if err != nil {
			return err
		}

		if firstName != nil {
			user.FirstName = *firstName
		}

		if lastName != nil {
			user.LastName = *lastName
		}
		if email != nil {
			user.Email = *email
		}
	*/
	return nil

}
