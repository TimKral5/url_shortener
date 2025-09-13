// Package main - This is the entrypoint of the main executable of
// this project.
package main

import (
	"os"
	"time"

	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/database"
	"github.com/timkral5/url_shortener/internal/log"
	"github.com/timkral5/url_shortener/internal/server"
)

const URLShortenerVersion string = "v0.1.0"
const urlShortenerDefaultAddress string = ":3000"
const databaseTimeout time.Duration = 2 * time.Second

type environment struct {
	Address                 string
	Database                string
	Cache                   string
	MongoDBConnectionString string
}

func main() {
	log.Info("Launching URL Shortener version", URLShortenerVersion, "...")

	env := loadEnvironment()
	server := server.NewServer()

	success := setupConnections(server, env)
	if !success {
		return
	}

	address := urlShortenerDefaultAddress
	if env.Address != "" {
		address = env.Address
	}

	log.Log("Listening on", address)
	server.Listen(address)
}

func loadEnvironment() environment {
	return environment{
		Address:                 os.Getenv("SHORTENER_ADDRESS"),
		Database:                os.Getenv("SHORTENER_DATABASE"),
		Cache:                   os.Getenv("SHORTENER_CACHE"),
		MongoDBConnectionString: os.Getenv("SHORTENER_MONGODB_URL"),
	}
}

func setupConnections(server *server.Server, env environment) bool {
	success := connectToCache(server, env)
	if !success {
		return success
	}

	success = connectToDatabase(server, env)
	if !success {
		return success
	}

	return true
}

func connectToCache(server *server.Server, env environment) bool {
	switch env.Cache {
	case "fake":
		log.Log("Setting up fake cache...")

		server.Cache = cache.NewFakeCacheConnection()

		log.Log("Cache set up.")
	default:
		log.Error("No cache configured (environment variable SHORTENER_CACHE).")

		return false
	}

	return true
}

func connectToDatabase(server *server.Server, env environment) bool {
	var err error

	switch env.Database {
	case "mongodb", "mongo":
		log.Log("Connecting to MongoDB...")

		server.Database, err = database.NewMongoDBConnection(env.MongoDBConnectionString, databaseTimeout)
		if err != nil {
			log.Error(err)

			return false
		}

		log.Log("Connection established.")
	case "fake":
		log.Log("Setting up fake database...")

		server.Database = database.NewFakeDatabaseConnection()

		log.Log("Database set up.")
	default:
		log.Error("No database configured (environment variable SHORTENER_DATABASE).")

		return false
	}

	return true
}
