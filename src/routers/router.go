package routers

import (
	"movie_festival_app/src/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (routes *route) RouterGroup() http.Handler {
	router := gin.Default()

	authRoutes := router.Group("/auth/user")
	authRoutes.POST("/register", routes.ctrl.UserRegister)
	authRoutes.POST("/login", routes.ctrl.UserLogin)

	userRouter := router.Group("/user")
	{
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", routes.ctrl.UpdateUser)
		userRouter.DELETE("/:userId", routes.ctrl.DeleteUser)
	}

	movie := router.Group("/movie")
	{
		movie.Use(middlewares.Authentication())

		//Admin APIs
		movie.POST("/", middlewares.AuthorizeAdmin(), routes.ctrl.StoreMovie)
		movie.PUT("/:movieId", middlewares.AuthorizeAdmin(), routes.ctrl.UpdateMovie)
		movie.GET("/most-viewed", middlewares.AuthorizeAdmin(), routes.ctrl.GetMoviesMostViewed)

		//All Users APIs
		movie.POST("/list-all-movie", routes.ctrl.GetMovies)
		movie.GET("/search-movie", routes.ctrl.SearchMovies)

	}

	return router
}
