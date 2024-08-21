package routers

import (
	"movie_festival_app/src/controllers"
	"net/http"
)

type route struct {
	ctrl controllers.Controllers
}

type Route interface {
	RouterGroup() http.Handler
}

func NewRouter(ctrl controllers.Controllers) Route {
	return &route{ctrl: ctrl}
}
