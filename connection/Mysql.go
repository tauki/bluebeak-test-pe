package connection

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tauki/bluebeak-test-pe/models"
)

type MySqlService struct {
	Conn *sql.DB
}

func GetMySqlService(cfg *models.Config) (*MySqlService, error) {

	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", src)
	if err != nil {
		msg := fmt.Sprintf("DBConnection :: MySQL :: conn :: Error : %s", err.Error())
		return nil, errors.New(msg)
	}

	fmt.Println("Connected to MySQL\n")

	return &MySqlService{
		Conn: db,
	}, nil
}

func (m *MySqlService) Ping() error {
	return m.Conn.Ping()
}
