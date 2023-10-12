package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/user"

	// drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/sijms/go-ora/v2"

	"github.com/flpgst/go-reportgen-producer/models"
	"github.com/xo/dburl"
	"github.com/xo/dburl/passfile"
)

func main() {
	verbose := flag.Bool("v", false, "verbose")
	dsn := flag.String("dsn", "postgres://postgres:password@localhost:5434/test", "dsn")
	flag.Parse()
	if err := run(context.Background(), *verbose, *dsn); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, verbose bool, dsn string) error {
	if verbose {
		logger := func(s string, v ...interface{}) {
			fmt.Printf("-------------------------------------\nQUERY: %s\n  VAL: %v\n\n", s, v)
		}

		models.SetLogger(logger)
	}
	v, err := user.Current()

	if err != nil {
		return err
	}
	// parse url
	u, err := parse(dsn)
	if err != nil {
		return err
	}

	// open database
	db, err := passfile.OpenURL(u, v.HomeDir, "xopass")
	if err != nil {
		return err
	}

	var f func(context.Context, *sql.DB) error
	switch u.Driver {
	case "mysql":
		f = runMysql
	case "oracle":
		f = runOracle
	case "postgres":
		f = runPostgres
	case "sqlite3":
		f = runSqlite3
	case "sqlserver":
		f = runSqlserver
	}
	return f(ctx, db)

}

func parse(dsn string) (*dburl.URL, error) {
	v, err := dburl.Parse(dsn)
	if err != nil {
		return nil, err
	}
	switch v.Driver {
	case "mysql":
		q := v.Query()
		q.Set("parseTime", "true")
		v.RawQuery = q.Encode()
		return dburl.Parse(v.String())
	case "sqlite3":
		q := v.Query()
		q.Set("_loc", "auto")
		v.RawQuery = q.Encode()
		return dburl.Parse(v.String())
	}
	return v, nil
}
