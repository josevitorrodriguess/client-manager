package api

import (
	"context"
	"fmt"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/josevitorrodriguess/client-manager/internal/services"
)

type Api struct {
	Router          *chi.Mux
	UserService     services.UserService
	CustomerService services.CustomerService
	Sessions        *scs.SessionManager
}

func GetAuthenticatedUserID(ctx context.Context, session *scs.SessionManager) (uuid.UUID, error) {
	val := session.GetString(ctx, "AuthenticatedUserId") // já faz type assertion p/ string
	if val == "" {
		return uuid.Nil, fmt.Errorf("AuthenticatedUserId not found in session")
	}

	id, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID in session: %w", err)
	}

	return id, nil
}
