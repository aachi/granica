/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package bootstrap

import (
	"context"
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	// UserCtxKey - Package standard User context key.
	UserCtxKey = ContextKey("user")
)

// ContextKey - Package standard context key.
type ContextKey string

func (c ContextKey) String() string {
	return "gitlab.com/mikrowezel/granica-" + string(c)
}

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "resources/keys/granica.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "resources/keys/granica.rsa.pub"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func initKeys(th *TestHandler) {
	fullPrivKeyPath := path.Join(th.Volume.BasePath, privKeyPath)
	signBytes, err := ioutil.ReadFile(fullPrivKeyPath)
	if err != nil {
		log.Fatalf("[ERROR]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[ERROR]: %s\n", err)
	}

	fullPubKeyPath := path.Join(th.Volume.BasePath, pubKeyPath)
	verifyBytes, err := ioutil.ReadFile(fullPubKeyPath)
	if err != nil {
		log.Fatalf("[ERROR]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[ERROR]: %s\n", err)
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userID string, username, role string) (string, error) {
	// Create the Claims
	claims := AppClaims{
		userID,
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 20).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// Authorize Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Get token from request
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				// app.ShowError(w, app.ErrTokenExpired, err, 401)
				return

			default:
				// app.ShowError(w, app.ErrTokenParsing, err, 401)
				return
			}
		default:
			// app.ShowError(w, app.ErrTokenParsing, err, 401)
			return
		}
	}
	if token.Valid {
		// Set user name to HTTP context
		claims := token.Claims.(*AppClaims)
		//claims := AppClaims{UserID: "5958b185-8150-4aae-b53f-0c44771ddec5", Username: "admin", Role: "admin"}
		ctx := context.WithValue(r.Context(), UserCtxKey, *claims)
		r = r.WithContext(ctx)
		next(w, r)
	} else {
		// app.ShowError(w, app.ErrTokenInvalid, err, 401)
	}
}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}
