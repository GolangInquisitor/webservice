package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"strings"
)

type ConnectionData struct {
	Login,
	Password,
	Host,
	Port,
	BaseNane string
}

func Connect(ctx context.Context, conndata ConnectionData) (*pgx.Conn, error) {

	//connStr := "postgres://username:password@localhost:5432/database_name"
	connStr := buildConnStr(conndata)

	return pgx.Connect(ctx, connStr)
}

func buildConnStr(conndata ConnectionData) string {
	return strings.Join([]string{"postgres://", conndata.Login, ":", conndata.Password, "@", conndata.Host, ":", conndata.Port, "/", conndata.BaseNane}, "")

}
