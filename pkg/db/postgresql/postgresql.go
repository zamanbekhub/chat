package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	//fmt.Println("dsn =", dsn)
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
