package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/server"
)

func main() {
	server.Run()
}
