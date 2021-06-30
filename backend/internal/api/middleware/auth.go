package middleware

import (
	"context"
	"net/http"

	c "npm/internal/api/context"
	h "npm/internal/api/http"
	"npm/internal/config"
	"npm/internal/entity/user"
	njwt "npm/internal/jwt"
	"npm/internal/logger"

	"github.com/go-chi/jwtauth"
)

// DecodeAuth decodes an auth header
func DecodeAuth() func(http.Handler) http.Handler {
	privateKey, privateKeyParseErr := njwt.GetPrivateKey()
	if privateKeyParseErr != nil && privateKey == nil {
		logger.Error("PrivateKeyParseError", privateKeyParseErr)
	}

	publicKey, publicKeyParseErr := njwt.GetPublicKey()
	if publicKeyParseErr != nil && publicKey == nil {
		logger.Error("PublicKeyParseError", publicKeyParseErr)
	}

	tokenAuth := jwtauth.New("RS256", privateKey, publicKey)
	return jwtauth.Verifier(tokenAuth)
}

// Enforce is a authentication middleware to enforce access from the
// jwtauth.Verifier middleware request context values. The Authenticator sends a 401 Unauthorised
// response for any unverified tokens and passes the good ones through.
func Enforce() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if config.IsSetup {
				token, claims, err := jwtauth.FromContext(ctx)

				if err != nil {
					h.ResultErrorJSON(w, r, http.StatusUnauthorized, err.Error(), nil)
					return
				}

				userID := int(claims["uid"].(float64))
				_, enabled := user.IsEnabled(userID)
				if token == nil || !token.Valid || !enabled {
					h.ResultErrorJSON(w, r, http.StatusUnauthorized, "Unauthorised", nil)
					return
				}

				// Add claims to context
				ctx = context.WithValue(ctx, c.UserIDCtxKey, userID)
			}

			// Token is authenticated, continue as normal
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
