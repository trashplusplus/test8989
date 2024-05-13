// router/router.go
package router

import (
	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func Init() {
	Router = httprouter.New()
	Router.POST("/calculate", PositiveNumbersMiddleware(CalculateHandler))
}
