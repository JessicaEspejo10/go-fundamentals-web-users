package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

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

	sqlQuery := "SELECT id, first_name, last_name, email FROM users WHERE id = ?"
	var u domain.User
	if err := r.db.QueryRow(sqlQuery, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
		r.log.Println(err.Error())
		if err == sql.ErrNoRows {
			return nil, ErrorNotFound{id}
		}
		return nil, err
	}
	r.log.Println("get user with id: ", id)

	return &u, nil
}

func (r *repo) Update(ctx context.Context, id uint64, firstName, lastName, email *string) error {
	var fields []string
	var values []interface{}

	if firstName != nil {
		fields = append(fields, "first_name = ?")
		values = append(values, *firstName)
	}

	if lastName != nil {
		fields = append(fields, "last_name = ?")
		values = append(values, *lastName)

	}
	if email != nil {
		fields = append(fields, "email = ?")
		values = append(values, *email)
	}

	if len(fields) == 0 {
		r.log.Println(ErrNotFields.Error())
		return ErrNotFields
	}
	values = append(values, id)

	sqlQuery := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(fields, ","))
	res, err := r.db.Exec(sqlQuery, values...)
	if err != nil {
		r.log.Println(err.Error)
		return err
	}

	row, err := res.RowsAffected()

	if err != nil {
		r.log.Println(err.Error)
		return err
	}

	if row == 0 {
		err := ErrorNotFound{id}
		r.log.Println(err.Error())
		return err
	}
	r.log.Println("user updated with id ", id)

	return nil
}
