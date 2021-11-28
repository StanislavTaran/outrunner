package server

// MySQL struct...
type MySQL struct {
	Port     string `json:"port"`
	Host     string `json:"host"`
	Password string `json:"password"`
	User     string `json:"user"`
}
