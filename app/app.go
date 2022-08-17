package app

import (
	"fmt"
	"go-echo/logger"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("ERROR ENV!!!")
	} else {
		logger.Info("ENV RUN SMOOTHLY!!!")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connection")
	}

	fmt.Println("DB", db)

	// initialize router echo
	e := echo.New()
	api := e.Group("/api/v1")
	api.GET("/users", TestConnectionFunctionUsers)
	fmt.Println("api", api)

	// start server echo
	start := fmt.Sprintf(":%v", serverPort)
	startServer := e.Start(start)
	if startServer != nil {
		logger.Fatal("SERVER CONNECTION FAILED!")
	}

}

func TestConnectionFunctionUsers(c echo.Context) error {
	return c.String(http.StatusOK, "HEllo users")
}
