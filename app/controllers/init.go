package controllers

import(
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/lib/pq"
)

var Db *gorp.DbMap

func init() {
  db, err := sql.Open("postgres", "dbname=jobs_app sslmode=disable")
  if err != nil {
    panic(err)
  }

  Db = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
}
