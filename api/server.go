package api

import (
    "github.com/gin-gonic/gin"
    db "github.com/techschool/simplebank/db/sqlc"
    "github.com/go-playground/validator/v10"
    "github.com/gin-gonic/gin/binding"
    "github.com/techschool/simplebank/util"
    "github.com/techschool/simplebank/token"
    "fmt"
)

// Server serves HTTP requests for our banking service
type Server struct {
    config util.Config
    store db.Store
    tokenMaker token.Maker
    router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
    tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
    if err != nil {
        return nil, fmt.Errorf("cannot create token maker: %w", err)
    }

    server := &Server{
        config: config,
        store: store,
        tokenMaker: tokenMaker,
    }
    router := gin.Default()

    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("currency", validCurrency)
    }

    //Users
    router.POST("/users", server.createUser)
    router.POST("/users/login", server.loginUser)

    // Accounts
    router.POST("/accounts", server.createAccount)
    router.GET("/accounts/:id", server.getAccount)
    router.GET("/accounts", server.listAccount)
    router.PUT("/accounts/:id", server.updateAccount)
    router.DELETE("/accounts/:id", server.deleteAccount)

    // Transfers
    router.POST("/transfers", server.createTransfer)

    server.router = router
    return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
    return server.router.Run(address)
}

func errorResponse(err error) gin.H {
    return gin.H{"error": err.Error()}
}