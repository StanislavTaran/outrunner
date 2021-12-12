package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// MySQL provides a way to work with MySQL database
type MySQL struct {
	config *Config
	db     *sql.DB
}

type QueryInfo struct {
	Query string `json:"query"`
}

// New - initialize new MySQL struct with config
func New(config *Config) *MySQL {
	return &MySQL{
		config: config,
	}
}

// Open new MySQL connection using passed to New func config
func (m *MySQL) Open() error {
	db, err := sql.Open("mysql", m.config.ConnectionURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	m.db = db

	return nil
}

// Close current MySQL connection
func (m *MySQL) Close() error {
	return m.db.Close()
}

// GetRecords return records from table passed in url
// and by query passed in body
func (m *MySQL) GetRecords(q QueryInfo) ([]map[string]interface{}, error) {
	rows, err := m.db.Query(q.Query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	// It can be slow but necessary when we work with interface type, In my humble opinion.
	// If you want to get more info - look at
	//         https://github.com/go-sql-driver/mysql/pull/1281
	//         https://github.com/golang/go/issues/22544
	// If you have better idea how to deal with it, let me know via GitHub issue.
	res, err := rowsToJSON(rows)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateRecord create records by passed query
func (m *MySQL) CreateRecord(q QueryInfo) (ok bool, err error) {
	if res, err := m.db.Exec(q.Query); err != nil {
		fmt.Print(res)
		return false, err
	}

	return true, nil
}
