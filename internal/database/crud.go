package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DB *sql.DB
}

// Create inserts a new record into the specified table
func (d *Database) Create(table string, columns []string, values []interface{}) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, joinColumns(columns), placeholders(len(values)))
	var id int64
	err := d.DB.QueryRow(query, values...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Read retrieves records from the specified table with optional conditions
func (d *Database) Read(table string, columns []string, condition string, args ...interface{}) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", joinColumns(columns), table)
	if condition != "" {
		query += " WHERE " + condition
	}
	return d.DB.Query(query, args...)
}

// Update modifies existing records in the specified table
func (d *Database) Update(table string, updates map[string]interface{}, condition string, args ...interface{}) (int64, error) {
	setClause, values := setClauses(updates)
	query := fmt.Sprintf("UPDATE %s SET %s", table, setClause)
	if condition != "" {
		query += " WHERE " + condition
	}
	values = append(values, args...)
	result, err := d.DB.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Delete removes records from the specified table
func (d *Database) Delete(table string, condition string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s", table)
	if condition != "" {
		query += " WHERE " + condition
	}
	result, err := d.DB.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Helper functions
func joinColumns(columns []string) string {
	return "`" + sqlJoin(columns, "`, `") + "`"
}

func placeholders(n int) string {
	return sqlJoin(make([]string, n), ", ")
}

func setClauses(updates map[string]interface{}) (string, []interface{}) {
	clauses := make([]string, 0, len(updates))
	values := make([]interface{}, 0, len(updates))
	for col, val := range updates {
		clauses = append(clauses, fmt.Sprintf("`%s` = ?", col))
		values = append(values, val)
	}
	return sqlJoin(clauses, ", "), values
}

func sqlJoin(elements []string, sep string) string {
	if len(elements) == 0 {
		return ""
	}
	joined := elements[0]
	for _, elem := range elements[1:] {
		joined += sep + elem
	}
	return joined
}
