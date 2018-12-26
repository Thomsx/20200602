package main

import (
	"enhance/pgx"
	"log"
)

func createConnPool() *pgx.ConnPool {
	dbconfig := pgx.ConnConfig{
		Host:     "127.0.0.1",
		User:     "postgres",
		Password: "123",
		Database: "shipcrew",
	}
	dbpoolconfig := pgx.ConnPoolConfig{
		ConnConfig:     dbconfig,
		MaxConnections: 3,
		AfterConnect:   nil,
	}
	dbpool, err := pgx.NewConnPool(dbpoolconfig)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	return dbpool
}
