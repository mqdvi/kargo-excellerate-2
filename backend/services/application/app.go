package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/databases"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/config"
)

type App struct {
	config    *config.Config
	E         *gin.Engine
	DBManager *databases.Manager
}

func NewApp(config *config.Config) *App {
	app := &App{
		config:    config,
		E:         gin.New(),
		DBManager: &databases.Manager{},
	}

	app.initDatabase()
	app.initRoutes()

	return app
}

func (app *App) initDatabase() {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?parseTime=true",
		app.config.DBUser,
		app.config.DBPassword,
		app.config.DBHost,
		app.config.DBPort,
		app.config.DBName,
	)

	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error when creating DB connection: %v", err)
	}

	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxOpenConns(app.config.DBMaxOpenConns)
	db.SetMaxIdleConns(app.config.DBMaxIdleConns)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error during ping database: %v", err)
	}

	app.DBManager.DB = db
}

func (app *App) initRoutes() {
	router := app.E
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
}

func (app *App) Start() {
	srv := &http.Server{
		Addr:         ":" + app.config.AppPort,
		Handler:      app.E,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func (app *App) PreStop() {
	log.Println("Clossing Connection")

	_ = app.DBManager.DB.Close()
}
