package app

import (
	"github.com/gin-gonic/gin"
	db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.POST("/recipes", server.createRecipe)
	router.GET("/recipes/:id", server.getRecipe)
	router.GET("/recipes", server.listRecipes)
	router.PUT("/recipes/:recipe_id", server.updateRecipe)
	router.DELETE("/recipes/:recipe_id", server.deleteRecipe)

	server.router = router
	return server
}

// CORSMiddleware adds headers to the api server
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}


