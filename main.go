package main

import (
	"fmt"
	"log"
	"movie_festival_app/packages"

	"movie_festival_app/packages/database"
	"movie_festival_app/src/controllers"
	"movie_festival_app/src/repositories"
	"movie_festival_app/src/routers"
	"movie_festival_app/src/usecases"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.LoadDB()
}

func main() {
	db := database.GetDB()
	repository := repositories.NewRepo(db)
	usecase := usecases.NewUC(repository)
	ctrl := controllers.NewCtrl(usecase)
	route := routers.NewRouter(ctrl).RouterGroup()
	log.Printf("[SERVER] starting at port : %v", os.Getenv("SERVER_PORT"))
	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")),
			packages.AllowCORS(route),
		),
	)
}
