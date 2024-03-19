package storage

import (
	"database/sql"
	"fmt"
)

const (
	MigrateTable = `
		CREATE TABLE IF NOT EXISTS table_products(
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(40) NOT NULL,
			price FLOAT NOT NULL,
			quantity INT NOT NULL,
			description TEXT NOT NULL,
			create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			update_at TIMESTAMP
		)ENGINE = InnoDB;
	`
)

type MySqlProduct struct {
	db *sql.DB
}

func NewMySql(db *sql.DB) *MySqlProduct {
	return &MySqlProduct{db}
}

func (m *MySqlProduct) Migrate() error {
	stm, err := m.db.Prepare(MigrateTable)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Procees Create table was successfully!!")
	return nil
}
