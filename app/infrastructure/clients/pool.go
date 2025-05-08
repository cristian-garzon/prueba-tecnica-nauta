package clients

import (
	"database/sql"
	"fmt"
	"prueba-tecnica-nauta/app/infrastructure/config"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Pool struct {
	connection *sql.DB
}

func NewPool(config *config.DatabaseConfig) (*Pool, error) {
	connection, err := sql.Open("pgx", getConnectionString(config))

	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(config.MaxIdleConns)
	connection.SetMaxOpenConns(config.MaxOpenConns)
	connection.SetConnMaxLifetime(config.ConnectionTimeout)
	connection.SetConnMaxIdleTime(config.IdleTimeout)

	return &Pool{connection: connection}, nil
}

func getConnectionString(config *config.DatabaseConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}

func (p Pool) FindAllQuery(query string, args ...any) ([]map[string]any, error) {
	result, err := p.connection.Query(query, args...)

	if err != nil {
		return nil, err
	}

	cols, err := result.Columns()

	if err != nil {
		return nil, err
	}

	var response []map[string]any
	for result.Next() {
		scans := make([]interface{}, len(cols))

		row := make(map[string]any)

		for i := range scans {
			scans[i] = &scans[i]
		}

		_ = result.Scan(scans...)

		for i, v := range scans {

			if v != nil {
				row[cols[i]] = v
			}

		}
		response = append(response, row)
	}

	result.Close()

	return response, nil
}

func (p *Pool) queryToMap(result *sql.Rows) (map[string]any, error) {
	cols, err := result.Columns()
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, sql.ErrNoRows
	}

	scans := make([]interface{}, len(cols))
	for i := range scans {
		scans[i] = &scans[i]
	}

	err = result.Scan(scans...)
	if err != nil {
		return nil, err
	}

	row := make(map[string]any)
	for i, v := range scans {
		if v != nil {
			row[cols[i]] = v
		}
	}

	return row, nil
}

func (p *Pool) FindSingleQuery(query string, args ...any) (map[string]any, error) {
	result, err := p.connection.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	return p.queryToMap(result)
}

func (p *Pool) ExecuteQuery(query string, args ...any) (map[string]any, error) {
	result, err := p.connection.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	return p.queryToMap(result)
}

func (p *Pool) Close() error {
	return p.connection.Close()
}
