package http

import (
	"github.com/julienschmidt/httprouter"
)

// Router holds all HTTP handlers and creates routes.
type Router struct {
	bookHandler  *BookHandler
	userHandler  *UserHandler
	orderHandler *OrderHandler
}

// NewRouter creates a new Router with all handlers.
func NewRouter(
	bookHandler *BookHandler,
	userHandler *UserHandler,
	orderHandler *OrderHandler,
) *Router {
	return &Router{
		bookHandler:  bookHandler,
		userHandler:  userHandler,
		orderHandler: orderHandler,
	}
}

// Setup registers all routes and returns the router.
func (r *Router) Setup() *httprouter.Router {
	router := httprouter.New()

	// Auth routes
	router.POST("/register", r.userHandler.Register)
	router.POST("/login", r.userHandler.Login)

	// Book routes
	router.POST("/books", r.bookHandler.Create)
	router.GET("/books", r.bookHandler.FindAll)
	router.GET("/books/:id", r.bookHandler.FindByID)
	router.PUT("/books/:id", r.bookHandler.Update)
	router.DELETE("/books/:id", r.bookHandler.Delete)

	// Order routes
	router.POST("/orders", r.orderHandler.Create)
	router.GET("/orders", r.orderHandler.FindAll)
	router.GET("/orders/:id", r.orderHandler.FindByID)
	router.GET("/users/:userId/orders", r.orderHandler.FindByUserID)

	return router
}
