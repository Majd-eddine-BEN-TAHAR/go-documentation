package auth

import (
	"errors"
	"net/http"
	"os" // Used for accessing environment variables
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtKey is the secret key used for signing JWT tokens. It's fetched from an environment variable.
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Claims struct represents the data embedded in the JWT token.
// In this case, we're storing the user's ID.
type Claims struct {
    UserID int `json:"user_id"`
    jwt.StandardClaims
}

// GenerateToken creates a JWT token with the user's ID.
func GenerateToken(userID int) (string, error) {
    // Token expiration time is set to 1 hour from now.
    expirationTime := time.Now().Add(1 * time.Hour)

    // Creating JWT claims including user ID and expiration time.
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Creating a new token with the provided claims.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Signing the token with the secret key.
	tokenString, err := token.SignedString(jwtKey)
    return tokenString, err
}

// ValidateToken checks the validity of a JWT token and returns the user ID if valid.
func ValidateToken(tokenString string) (int, error) {
    claims := &Claims{}

    // Parsing and validating the token. The key function provides the signing key for verification.
    // ParseWithClaims is a function from the JWT package in Go. It's used to parse a JWT token and validate it.
    // It takes the token string, a claims object where the data from the token will be stored, and a function to retrieve the signing key for verification. The function then parses the token, validates it, and fills the claims object with the token's data.
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return 0, err
    }

    // Verifying that the token is valid.
    if !token.Valid {
        return 0, errors.New("invalid token")
    }

    // Returning the user ID from the token's claims.
    return claims.UserID, nil
}


// GetUserFromToken extracts the user ID from the JWT token provided in the request.
func GetUserFromToken(r *http.Request) (int, error) {
    // Extract the token from the Authorization header
    tokenHeader := r.Header.Get("Authorization")

    // Strip the "Bearer " prefix from the token
    token := strings.TrimPrefix(tokenHeader, "Bearer ")
    if token == tokenHeader {
        return 0, errors.New("no token found")
    }

    // Validate the token and extract the claims
    userID, err := ValidateToken(token)
    if err != nil {
        return 0, err
    }

    return userID, nil
}