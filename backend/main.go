// package main

// import (
// 	"backend/db"
// 	"backend/routes"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	db.Connect()
// 	r := gin.Default()
// 	routes.SetupRoutes(r)
// 	r.Run(":8080")
// }

package main

import (
	"backend/db"
	"backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	// ✅ CORS Middleware (Allow Frontend Access)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Frontend ka URL
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(r)
	r.Run(":8080")
}
