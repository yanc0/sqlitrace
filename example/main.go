package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/sync/errgroup"
)

func main() {
	sql.Register("sqlite3-with-trace", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			err := conn.LoadExtension("./trace", "sqlite3_extension_init")
			if err != nil {
				return err
			}

			err = conn.LoadExtension("./stats", "sqlite3_stats_init")
			if err != nil {
				return err
			}

			return nil
		},
	})

	db, err := sql.Open("sqlite3-with-trace", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	generateSerie, err := db.Prepare("select value from generate_series(5,1000,5)")
	if err != nil {
		log.Fatal(err)
	}

	errg, _ := errgroup.WithContext(context.Background())
	errg.SetLimit(5)

	for i := 0; i < 1_000_000; i++ {
		errg.Go(func() error {
			return q(generateSerie)
		})
	}

	err = errg.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

func q(stmt *sql.Stmt, args ...any) error {
	res, err := stmt.Query(args...)
	if err != nil {
		return err
	}

	var i int
	for res.Next() {
		err = res.Scan(&i)
		if err != nil {
			return err
		}
	}
	return nil
}
