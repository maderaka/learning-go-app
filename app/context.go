package app

import (
	"github.com/maderaka/goapp/packages/core"
	"github.com/maderaka/goapp/packages/core/email"
	"database/sql"
	"fmt"
	"golang.org/x/net/context"
)

func Database() *sql.DB {
	source := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		core.DbUsername,
		core.DbPassword,
		core.DbHost,
		core.DbName,
		core.SslMode)

	db, _ := core.NewDB(source)
	return db
}

func EmailSetup() *email.Engine {
	return email.NewEmailEngine(
	"rakatejaa@gmail.com",
	"passwordcuy",
	"smtp.gmail.com",
	587)
}

func ContextValues() (ctx context.Context) {
	ctx = context.Background()
	ctx = context.WithValue(ctx, "db", Database())
	ctx = context.WithValue(ctx, "email", EmailSetup())

	return ctx
}
