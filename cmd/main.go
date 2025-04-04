package main

import (
	"context"
	"encoding/gob"
	"fmt"
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
)

func main() {
	gob.Register(uuid.UUID{})
	logger.Info("Starting the application...")

	godotenv.Load()

	ctx := context.Background()
	pool := db.InitPool(ctx)

	err := db.CreateAdmin(ctx, pool)
	if err != nil {
		logger.Error("Error creating admin", fmt.Errorf("error creating admin: %w", err))
	}

	defer pool.Close()

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: *services.NewUserService(pool),
		Sessions:    s,
	}

	api.BindRoutes()

	fmt.Println("Server is running on port 3080")
	if err := http.ListenAndServe("localhost:3080", api.Router); err != nil {
		panic(err)
	}
}
