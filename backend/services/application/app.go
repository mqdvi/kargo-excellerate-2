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
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/databases"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/api/driver"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/api/shipment"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/api/truck"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/api/truck_type"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/config"
)

type App struct {
	config    *config.Config
	E         *gin.Engine
	DBManager *databases.Manager
}

var validate *validator.Validate

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

	// Init validator
	validate = validator.New()

	// Repository
	truckRepo := truck.NewRepository(app.config.DBName)
	truckTypeRepo := truck_type.NewRepository(app.config.DBName)
	driverRepo := driver.NewRepository(app.config.DBName)
	shipmentRepo := shipment.NewRepository(app.config.DBName)

	// Service
	truckSvc := truck.NewService(app.DBManager.DB, truckRepo)
	truckTypeSvc := truck_type.NewService(app.DBManager.DB, truckTypeRepo)
	driverSvc := driver.NewService(app.DBManager.DB, driverRepo)
	shipmentSvc := shipment.NewService(app.DBManager.DB, shipmentRepo, truckRepo, driverRepo)

	// Controller
	truckCtrl := truck.NewController(truckSvc, validate)
	truckTypeCtrl := truck_type.NewController(truckTypeSvc)
	driverCtrl := driver.NewController(driverSvc, validate)
	shipmentCtrl := shipment.NewController(shipmentSvc, validate)

	router.GET("/trucks", truckCtrl.HandlerGetTrucks)
	router.GET("/drivers", driverCtrl.HandlerGetDrivers)

	transporters := router.Group("/transporters")
	transporters.GET("/trucks", truckCtrl.HandlerGetTrucks)
	transporters.GET("/trucks/:id", truckCtrl.HandlerGetTruckByID)
	transporters.POST("/trucks", truckCtrl.HandlerCreateTruck)
	transporters.PUT("/trucks/:id", truckCtrl.HandlerUpdateTruck)
	transporters.PATCH("/trucks/:id/deactivate", truckCtrl.HandlerDeactivateTruck)

	transporters.GET("/truck-types", truckTypeCtrl.HandlerGetTruckTypes)

	transporters.GET("/drivers", driverCtrl.HandlerGetDrivers)
	transporters.POST("/drivers", driverCtrl.HandlerCreateDriver)

	shipper := router.Group("/shipper")
	shipper.GET("/shipments", shipmentCtrl.HandlerGetShipments)
	shipper.POST("/shipments", shipmentCtrl.HandlerCreateShipment)
	shipper.POST("/shipments/:id/allocate", shipmentCtrl.HandlerAllocateShipment)
	shipper.PUT("/shipments/:id/status", shipmentCtrl.HandlerUpdateShipmentStatus)
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
