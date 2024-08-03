package database

import (
    "fmt"
    "log"
    "notes/config"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func GetDBConnection(cfg *config.Config) *sqlx.DB {
    dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, cfg.PostgresHost, cfg.PostgresPort)
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalln(err)
    }
    return db
}

func MigrateDB(cfg *config.Config) {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalln(err)
    }
    driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
    if err != nil {
        log.Fatalln(err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://database/migrations",
        "postgres", driver)
    if err != nil {
        log.Fatalln(err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalln(err)
    }
}
