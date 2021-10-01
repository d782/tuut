package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/d782/tuut/middlew"
	"github.com/d782/tuut/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* handlers */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
