package server

import (
	"database/sql"
	"log"
	"os"

	"github.com/Keysie27/go-backend/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateDB() (*sql.DB, error) {
	dbCon := os.Getenv("DB_CONN")

	log.Println("-- Opening Database Connection --")
	db, err := sql.Open("mysql", dbCon)
	if err != nil {
		return nil, err
	}

	log.Println("-- Pinging Database --")
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateServer() *mux.Router {
	router := mux.NewRouter()

	db, err := CreateDB()
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.CreateNewHandler(db)
	CreateUserHandler(h, router)

	return router
}

func CreateUserHandler(h *handlers.Handler, router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", h.EnableCors(h.CreateUser)).Methods("POST")
	userRouter.HandleFunc("/all", h.EnableCors(h.GetAllUsers)).Methods("GET")
	userRouter.HandleFunc("/{id}", h.EnableCors(h.GetUserById)).Methods("GET")
	userRouter.HandleFunc("/{id}/adresses", h.EnableCors(h.GetUserAddresses)).Methods("GET")
	userRouter.HandleFunc("/{id}/", h.EnableCors(h.UpdateUser)).Methods("POST")
}
