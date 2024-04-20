package main

import (
	"database/sql"
	"fmt"
	"log"
	docs "my-api/docs"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func getCategories(g *gin.Context) {
	user := "2024-sre-prep-usr"
	pass := "sqlpass"
	dbname := "2024-sre-prep-db"
	host := os.Getenv("pg-container") // Retrieve the host from environment variable

	// Construct connection string
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, dbname)

	// Open the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Define a slice to hold the categories
	var categories []string

	// Query the database
	rows, err := db.Query("SELECT category FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the results as JSON
	g.JSON(http.StatusOK, gin.H{"categories": categories})
}

func getCategory(g *gin.Context) {
	g.JSON(http.StatusOK, "getCategory")
}

func getLists(g *gin.Context) {
	g.JSON(http.StatusOK, "getLists")
}

func getList(g *gin.Context) {
	g.JSON(http.StatusOK, "getList")
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		// Group for flashcards
		flashcards := v1.Group("/flashcards")
		{
			// Route for getting all categories
			flashcards.GET("/categories", getCategories)

			// Route for getting a specific category by name
			flashcards.GET("/categories/:category", getCategory)
		}

		// Group for checklist
		checklist := v1.Group("/checklist")
		{
			// Route for getting all lists
			checklist.GET("/lists", getLists)

			// Route for getting a specific list by name
			checklist.GET("/lists/:list", getList)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8000")
}
