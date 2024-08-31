package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (app *application) ValidUser(next http.Handler) http.Handler {
	// next

	return next
}

func (app *application) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")

		// Check if the header is present and starts with "Bearer "
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			// Extract the JWT token
			token := authHeader[7:]

			// Verify the JWT token
			if token == "valid_token" {
				// Handle the request
				fmt.Fprint(w, "Hello, World!")
			} else {
				// Return an error
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
		} else {
			// Return an error
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
		}

		ctx := context.Background()
		err := app.Redis.Set(ctx, "foo", "bar", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := app.Redis.Get(ctx, "foo").Result()
		if err != nil {
			panic(err)
		}

		fmt.Println("foo", val)
		next.ServeHTTP(w, r)
	})
}

func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil {
		return nil, errors.New("failed to decode private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
