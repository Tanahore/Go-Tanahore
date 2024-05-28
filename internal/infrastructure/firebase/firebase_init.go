package firebase

import (
	"context"
	"tanahore/configs"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func InitFirebase(config *configs.FirebaseConfig) (*db.Client, *firebase.App) {
	credential, err := FirebaseConfigConverter(config)
	if err != nil {
		logrus.Error("failed to encode firebase key")
	}
	// Konfigurasi Firebase Admin SDK
	opt := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Errorf("Failed to create Firebase app: %v", err)
	}
	client, err := app.DatabaseWithURL(context.Background(), config.FirebaseURL)
	if err != nil {
		logrus.Errorf("Failed toconnect to Firebase client: %v", err)
	}
	return client, app
}
