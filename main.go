package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vsualzm/restapi-gin-golang/controllers"
	"github.com/vsualzm/restapi-gin-golang/database"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "1234"
// 	dbname   = "practice"
// )

var (
	DB  *sql.DB
	err error
)

func main() {

	// ENV SETTING
	err = godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("failed load file ENV")
	} else {
		fmt.Println("Success read file enviroment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()

	if err != nil {
		fmt.Println("DB CONNECTION FAILED")
	} else {
		fmt.Println("DB CONNECTION SUCCESS")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// router GIN

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PORT"))

}
