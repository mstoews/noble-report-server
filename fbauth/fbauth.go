package fbauth

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"os"
)

var (
	firebaseConfiguration = os.Getenv("FIREBASE_CONFIG_FILE")
)

func InitAuth() (*auth.Client, error) {
	opt := option.WithCredentialsFile("../firebase-config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "error initializating firebase auth (creating client app)")
	}

	client, errAuth := app.Auth(context.Background())
	if errAuth != nil {
		return nil, errors.Wrap(err, "error initializing firebase auth (creating client app)")
	}

	return client, nil

}
