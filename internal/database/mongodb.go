package database

import (
	"context"
	"time"

	"github.com/timkral5/url_shortener/internal/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type urlEntry struct {
	Hash string `bson:"hash"`
	FullURL string `bson:"url"`
}

// MongoDBConnection is a wrapper for a connection to a MongoDB database.
type MongoDBConnection struct {
	ConnectionString  string
	Database          string
	URLCollectionName string

	ConnectTimeout time.Duration
	AddURLTimeout time.Duration
	GetURLTimeout time.Duration

	client            *mongo.Client
	urlCollection     *mongo.Collection
}

var _ Connection = (*MongoDBConnection)(nil)

// NewMongoDBConnection constructs a new instance of a MongoDB connection.
func NewMongoDBConnection(connStr string, timeout time.Duration) (*MongoDBConnection, error) {
	conn := &MongoDBConnection{
		ConnectionString:  connStr,
		Database:          "url_shortener",
		URLCollectionName: "urls",

		ConnectTimeout: timeout,
		AddURLTimeout: timeout,
		GetURLTimeout: timeout,

		client: nil,
		urlCollection: nil,
	}

	err := conn.Connect(connStr)
	if err != nil {
		return conn, err
	}

	err = conn.SetupDatabase()
	if err != nil {
		return conn, err
	}

	return conn, nil
}



// Connect establishes a new connection using the provided connection
// string.
func (conn *MongoDBConnection) Connect(connStr string) error {
	conn.ConnectionString = connStr

	log.Log(connStr)

	client, err := mongo.Connect(options.Client().
		ApplyURI(connStr).
		SetConnectTimeout(conn.ConnectTimeout))
	if err != nil {
		return Error{
			Message: err.Error(),
			Code:    FakeConnectError,
		}
	}

	conn.client = client
	conn.urlCollection = conn.client.
		Database(conn.Database).
		Collection(conn.URLCollectionName)

	return nil
}

// SetupDatabase initializes the database, the collections and the
// indeces.
func (conn *MongoDBConnection) SetupDatabase() error {
	urlHashIndex := mongo.IndexModel{
		Keys: bson.D{ bson.E{ Key: "hash", Value: 1 } },
	}

	_, err := conn.urlCollection.Indexes().CreateOne(context.TODO(), urlHashIndex)
	if err != nil {
		return err
	}

	return nil
}

// Disconnect terminates an active connection.
func (conn *MongoDBConnection) Disconnect() error {
	return conn.client.Disconnect(context.Background())
}

// AddURL inserts a new URL entry into the database.
func (conn *MongoDBConnection) AddURL(hash string, url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), conn.AddURLTimeout)
	defer cancel()

	result, err := conn.urlCollection.InsertOne(ctx, urlEntry{
		Hash: hash,
		FullURL: url,
	})
	if err != nil {
		return Error{
			Message: err.Error(),
			Code: FakeConnectError,
		}
	}

	if !result.Acknowledged {
		return Error{
			Message: "Failed to add URL (not aknowledged)",
			Code: FakeConnectError,
		}
	}

	return nil
}

// GetURL fetches a URL entry from the database.
func (conn *MongoDBConnection) GetURL(hash string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), conn.GetURLTimeout)
	defer cancel()

	filter := bson.D{
		bson.E{Key: "hash", Value: hash},
	}

	result := conn.urlCollection.FindOne(ctx, filter)

	err := result.Err()
	if err != nil {
		return "", err
	}

	var entry urlEntry
	result.Decode(&entry)

	return entry.FullURL, nil
}
