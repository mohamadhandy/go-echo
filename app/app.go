package app

import (
	"fmt"
	"go-echo/logger"
	"go-echo/users"
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

	// initialize db
	userRepoDB := users.NewUserRepositoryDB(db)

	// initialize service
	userService := users.NewUserService(userRepoDB)

	// initialize handler
	userHandler := users.NewUserHandler(*userService)

	// initialize router echo
	e := echo.New()
	api := e.Group("/api/v1")
	api.GET("/users", TestConnectionFunctionUsers)
	api.POST("/users", userHandler.RegisterUser)

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
