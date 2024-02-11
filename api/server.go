package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	// "github.com/gin-gonic/gin/binding"
	db "github.com/mstoews/prd-backup-server/db/sqlc"
	pbauth "github.com/mstoews/prd-backup-server/fbauth"
	"github.com/mstoews/prd-backup-server/token"
	"github.com/mstoews/prd-backup-server/util"
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

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	client, err := pbauth.InitAuth()

	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Users
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)
	router.GET("/accts", server.GLAccounts)

	// Middleware
	authRoutes := router.Group("/").Use(pbauth.AuthJWT(client))

	authRoutes.GET("/journal", server.KBTasks)
	authRoutes.GET("/dist", server.ListDistLedger)
	authRoutes.GET("/accounts", server.GLAccounts)
	authRoutes.GET("/types", server.GLAccountTypes)
	authRoutes.GET("/journals", server.GlJournalDetail)
	authRoutes.GET("/period", server.GlPeriod)
	authRoutes.GET("/dist_by_prd/:id", server.ListDistributionLedgerByPeriod)
	authRoutes.GET("/funds", server.ListFunds)
	authRoutes.GET("/tasks", server.KBTasks)
	authRoutes.GET("/tasklist", server.GetTasks)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	server.router.Use(cors.Default())
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
