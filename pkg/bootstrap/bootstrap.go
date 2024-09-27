package bootstrap

import (
	"log"
	"os"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/domain"
	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/user"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "Ana",
			LastName:  "Montes",
			Email:     "amontes@gmail.com",
		}, {
			ID:        2,
			FirstName: "Juan",
			LastName:  "Suarez",
			Email:     "jsuarez@gmail.com",
		}, {
			ID:        3,
			FirstName: "Luis",
			LastName:  "Naranjos",
			Email:     "lnaranjo@gmail.com",
		}, {
			ID:        4,
			FirstName: "Ruth",
			LastName:  "Castro",
			Email:     "rcastro@gmail.com",
		}}, MaxUserID: 4,
	}

}
