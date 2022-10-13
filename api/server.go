package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/mr-od/Asusu-Igbo/db/sqlc"
	"github.com/mr-od/Asusu-Igbo/token"
	"github.com/mr-od/Asusu-Igbo/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()

	return server, err
}

func (server *Server) setupRouter() {

	router := gin.Default()

	// User Route
	router.POST("/api/v1/users", server.createUser)
	router.POST("/api/v1/users/login", server.loginUser)
	router.POST("/api/v1/tokens/refresh_token", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// User Route
	authRoutes.GET("api/v1/users/me", server.getUser)

	// Accounts Route
	authRoutes.POST("api/v1/accounts", server.createAccount)
	authRoutes.GET("api/v1/accounts/:id", server.getAccount)
	authRoutes.GET("api/v1/accounts", server.listAccounts)

	// Products Route
	authRoutes.POST("api/v1/products", server.addProduct)
	authRoutes.GET("api/v1/products/:id", server.getProduct)
	authRoutes.GET("api/v1/products", server.listProducts)
	authRoutes.DELETE("api/v1/products/:id", server.deleteProduct)

	// Order Route
	authRoutes.POST("api/v1/orders/place", server.placeOrder)

	// Transfers Route
	authRoutes.POST("api/v1/transfers", server.createTransfer)

	authRoutes.GET("api/v1/search/:q", server.searchProduct)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
