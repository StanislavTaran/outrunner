package mongodb

// Config provides options to establish connection to mongo db
type Config struct {
	ConnectionURL string `json:"connectionUrl"`
	Database      string `json:"database"`
}
