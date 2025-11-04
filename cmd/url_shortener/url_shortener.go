// Package main - This is the entrypoint of the main executable of
// this project.
package main

import (
	"os"
	"time"

	"github.com/timkral5/url_shortener/internal/auth"
	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/database"
	"github.com/timkral5/url_shortener/internal/jwt"
	"github.com/timkral5/url_shortener/internal/log"
	"github.com/timkral5/url_shortener/internal/server"
)

const URLShortenerVersion string = "v0.1.0"
const urlShortenerDefaultAddress string = ":3000"
const databaseTimeout time.Duration = 2 * time.Second

type environment struct {
	Address                   string
	Authentication            string
	Database                  string
	Cache                     string
	StaticToken               string
	JWTSigningKey             string
	MemcachedConnectionString string
	MongoDBConnectionString   string
}

func main() {
	log.Info("Launching URL Shortener version", URLShortenerVersion, "...")

	env := loadEnvironment()
	server := server.NewServer()
	server.APIVersion = URLShortenerVersion

	success := setupConnections(server, env)
	if !success {
		return
	}

	address := urlShortenerDefaultAddress
	if env.Address != "" {
		address = env.Address
	}

	server.JWTHandler = jwt.NewHandler("HelloWorld")

	log.Log("Listening on", address)
	server.Listen(address)
}

func loadEnvironment() environment {
	return environment{
		Address:                   os.Getenv("SHORTENER_ADDRESS"),
		Authentication:            os.Getenv("SHORTENER_AUTH"),
		Database:                  os.Getenv("SHORTENER_DATABASE"),
		Cache:                     os.Getenv("SHORTENER_CACHE"),
		StaticToken:               os.Getenv("SHORTENER_STATIC_TOKEN"),
		JWTSigningKey:             os.Getenv("SHORTENER_JWT_KEY"),
		MemcachedConnectionString: os.Getenv("SHORTENER_MEMCACHED_URL"),
		MongoDBConnectionString:   os.Getenv("SHORTENER_MONGODB_URL"),
	}
}

func setupConnections(server *server.Server, env environment) bool {
	success := setupAuthentication(server, env)
	if !success {
		return success
	}

	success = connectToCache(server, env)
	if !success {
		return success
	}

	success = connectToDatabase(server, env)
	if !success {
		return success
	}

	return true
}

func setupAuthentication(server *server.Server, env environment) bool {
	// var err error
	switch env.Authentication {
	case "static":
		server.Auth = auth.NewStaticAuth(env.StaticToken)
	default:
		server.Auth = nil
	}

	return true
}

func connectToCache(server *server.Server, env environment) bool {
	var err error

	switch env.Cache {
	case "memcached":
		log.Log("Connecting to Memcached...")

		server.Cache, err = cache.NewMemcachedConnection(env.MemcachedConnectionString)
		if err != nil {
			log.Error(err)

			return false
		}

		log.Log("Connection established.")
	case "fake":
		log.Log("Setting up fake cache...")

		server.Cache = cache.NewFakeCacheConnection()

		log.Log("Cache set up.")
	default:
		log.Log(
			"No cache configured (environment variable SHORTENER_CACHE). Disabling cache.")

		server.Cache = nil
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
