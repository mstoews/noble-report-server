package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/mstoews/prd-backup-server/db/sqlc"
	"github.com/mstoews/prd-backup-server/token"
	"github.com/mstoews/prd-backup-server/util"
	"github.com/gin-contrib/cors"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
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
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Users
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)
	
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))	
	// Accounts
	authRoutes.POST("/account", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	
	
	// Trades 
	authRoutes.GET("/trades", server.listAllTrades)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	
	// server.router.Use(cors.New(cors.Config{
    //     AllowOrigins:     []string{"*"},
    //     AllowMethods:     []string{"PUT", "PATCH, POST, GET"},
    //     AllowHeaders:     []string{"Origin"},
    //     ExposeHeaders:    []string{"Content-Length"},
    //     AllowCredentials: true,
    //     AllowOriginFunc: func(origin string) bool {
    //         return origin == "http://localhost"
    //     },
    //     MaxAge: 12 * time.Hour,
    // }))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	server.router.Use(cors.New(config))
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
