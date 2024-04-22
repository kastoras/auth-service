package server

import (
	"auth-service/helpers"
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *APIServer) ClientSetUp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		domain := "localhost"
		if helpers.GetEnvParam("AUTHSERVER_TYPE", "local") == "origin" {
			domain = getRequestOrigin(r)
		}

		authServerInfo, err := getDomainData(domain)
		if err != nil {
			//Provide internal server error
			panic("No domain found")
		}

		s.keycloack, _ = setUpKeycloakClient(authServerInfo)

		next.ServeHTTP(w, r)
	})
}

func getDomainData(domain string) (Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		panic("No db found")
	}

	collection := client.Database("auth-service").Collection("domains")

	filer := bson.M{"domain": domain, "active": true}

	var result Domain
	err = collection.FindOne(ctx, filer).Decode(&result)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func setUpKeycloakClient(domain Domain) (*KeycloackClient, error) {
	return &KeycloackClient{
		Host:            domain.Host,
		ClientID:        domain.Client_id,
		ClientSecret:    domain.Client_secret,
		ClientGrandType: domain.Grand_type,
		Realm:           domain.Realm,
	}, nil
}

func getRequestOrigin(r *http.Request) string {
	return r.Header.Get("Origin")
}
