package driver

import (
	"errors"
    "database/sql"
    _ "github.com/lib/pq"
)

type PostgresqlDriver struct {
	baseRDB
	*sql.DB
}

func (driver *PostgresqlDriver) Open(connectionString string) error {

	if driver != nil {
		var err error
		driver.DB, err = sql.Open("postgres", connectionString)
		return err
	}
	return errors.New("driver object is")

}
