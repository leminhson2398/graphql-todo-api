package route

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/leminhson2398/todo-api/internal/auth"
	"github.com/leminhson2398/todo-api/internal/utils"
	log "github.com/sirupsen/logrus"
)

type AuthenticationMiddleware struct {
	jwtKey []byte
}

// Middleware for authenticate all the graphql requests
func (m *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearerTokenRaw := r.Header.Get("Authorization")
		if bearerTokenRaw == "" {
			next.ServeHTTP(w, r)
			return
		}

		splitToken := strings.Split(bearerTokenRaw, "Bearer")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		accessTokenString := strings.TrimSpace(splitToken[1])
		accessClaims, err := auth.ValidateAccessToken(accessTokenString, m.jwtKey)
		if err != nil {
			if _, ok := err.(*auth.ErrExpiredToken); ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{
					"data": {},
					"errors": [
						"extensions": {
							"code": "UNAUTHENTICATED"
						}
					]
				}`))
				return
			}
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var userID uuid.UUID
		if accessClaims.Restricted == auth.InstallOnly {
			userID = uuid.New()
		} else {
			userID, err = uuid.Parse(accessClaims.UserID)
			if err != nil {
				log.WithError(err).Error("middleware access token userID parse")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		// set some values to come along with context
		ctx := context.WithValue(r.Context(), utils.UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
