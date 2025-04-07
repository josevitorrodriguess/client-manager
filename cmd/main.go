package main

import (
	"context"
	"encoding/gob"
	"net/http"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/client-manager/internal/api"
	"github.com/josevitorrodriguess/client-manager/internal/config/db"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/services"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	gob.Register(uuid.UUID{})
	logger.Info("Starting the application...")

	godotenv.Load()
	logger.Debug("Environment variables loaded")

	ctx := context.Background()
	logger.Info("Initializing database connection...")
	pool := db.InitPool(ctx)
	logger.Info("Database connection established successfully")

	err := db.CreateAdmin(ctx, pool)
	if err != nil {
		logger.Error("Error creating admin", err)
	} else {
		logger.Info("Admin user created successfully")
	}

	defer pool.Close()

	logger.Info("Initializing session manager...")
	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	logger.Info("Session manager initialized successfully")

	api := api.Api{
		Router:          chi.NewMux(),
		UserService:     *services.NewUserService(pool),
		CustomerService: *services.NewCustomerService(pool),
		Sessions:        s,
	}

	logger.Info("Binding routes...")
	api.BindRoutes()
	logger.Info("Routes bound successfully")

	port := "3080"
	logger.Info("Server is starting", zap.String("port", port))
	if err := http.ListenAndServe("localhost:"+port, api.Router); err != nil {
		logger.Error("Server failed to start", err)
		panic(err)
	}
}
