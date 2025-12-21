package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"ayee-portal-backend/config"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the token from Keycloak
func AuthMiddleware() gin.HandlerFunc {
	// Initialize OIDC provider
	issuer := config.AppConfig.KeycloakIssuer
	var verifier *oidc.IDTokenVerifier

	if issuer != "" {
		provider, err := oidc.NewProvider(context.Background(), issuer)
		if err != nil {
			log.Printf("WARNING: Failed to initialize OIDC provider: %v", err)
		} else {
			oidcConfig := &oidc.Config{
				ClientID: config.AppConfig.KeycloakClientID,
			}
			verifier = provider.Verifier(oidcConfig)
		}
	} else {
		log.Println("WARNING: KEYCLOAK_ISSUER not set, token verification will be skipped or limited.")
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		if verifier != nil {
			idToken, err := verifier.Verify(c.Request.Context(), tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
				c.Abort()
				return
			}

			// Extract claims
			var claims map[string]interface{}
			if err := idToken.Claims(&claims); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse claims"})
				c.Abort()
				return
			}

			c.Set("claims", claims)
			c.Set("token", tokenString)
			
			// Optional: Set user ID if available in claims (e.g., 'sub')
			if sub, ok := claims["sub"].(string); ok {
				c.Set("userID", sub)
			}

		} else {
			// Fallback behavior if OIDC is not configured (prevent locking out if config is missing during dev)
			// In strict mode, we should fail here.
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
				c.Abort()
				return
			}
			// Just pass through if no verifier configured (mock/dev mode)
			c.Set("token", tokenString)
		}

		c.Next()
	}
}
