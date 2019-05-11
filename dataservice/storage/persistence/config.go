package persistence

// DBConfiguration ...
type DBConfiguration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Encoding string `json:"encoding"`
	PoolSize int    `json:"poolsize"`
}
