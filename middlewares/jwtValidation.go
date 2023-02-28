package middlewares

import (
	"net/http"
	"net/url"
	"time"

	auth0Jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/julienschmidt/httprouter"
)

func JwtValidationMiddleware(nextHandle httprouter.Handle, issuer string, scope string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// Extract the token from header
		token, err := auth0Jwtmiddleware.AuthHeaderTokenExtractor(r)
		if err != nil {
			// We failed to extract the token. Not because missing.
			http.Error(w, "Failed to extract token", http.StatusInternalServerError)
			return
		}

		// Checking the token
		issuerURL, err := url.Parse(issuer)
		if err != nil {
			http.Error(w, "Failed to parse the issuer url", http.StatusInternalServerError)
		}
		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{scope},
			validator.WithAllowedClockSkew(time.Duration(2*time.Minute)),
		)
		if err != nil {
			http.Error(w, "Failed to set up the jwt validator", http.StatusInternalServerError)
			return
		}
		_, err = jwtValidator.ValidateToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Delegate to the next given handle
		nextHandle(w, r, ps)

	}
}
