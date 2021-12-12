package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb struct
type Mongodb struct {
	config *Config
	db     *mongo.Client
}

type QueryGet struct {
	Collection string                 `json:"collection"`
	Query      map[string]interface{} `json:"query"`
}

type QueryInsert struct {
	Collection string                   `json:"collection"`
	Query      []map[string]interface{} `json:"query"`
}

// New returns new *Mongodb struct with passed config
func New(c *Config) *Mongodb {
	return &Mongodb{
		config: c,
	}
}

// Open initialises new mongodb client
func (m *Mongodb) Open() error {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(m.config.ConnectionURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	m.db = client
	return nil
}

// Close calls func disconnect for mongo
func (m *Mongodb) Close() error {
	return m.db.Disconnect(context.TODO())
}

func (m *Mongodb) GetRecords(q QueryGet) ([]map[string]interface{}, error) {
	collection := m.db.Database(m.config.Database).Collection(q.Collection)

	result := make([]map[string]interface{}, 0)

	filter := bson.M{}

	if q.Query != nil {
		for k, v := range q.Query {
			if k == "_id" {
				oid, err := primitive.ObjectIDFromHex(v.(string))
				if err != nil {
					return nil, err
				}
				filter[k] = oid
			} else {
				filter[k] = v
			}
		}
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		doc := make(map[string]interface{})

		err := cur.Decode(&doc)
		if err != nil {
			return nil, err
		}

		result = append(result, doc)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *Mongodb) CreateRecords(q QueryInsert) (ok bool, err error) {
	collection := m.db.Database(m.config.Database).Collection(q.Collection)
	docs := make([]interface{}, len(q.Query))
	for i, item := range q.Query {
		docs[i] = item
	}

	_, err = collection.InsertMany(context.TODO(), docs)
	if err != nil {
		return false, err
	}

	return true, nil
}
