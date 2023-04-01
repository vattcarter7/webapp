package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "2ddflsdkslksdlksfkljsdflkjsdflkjaszkskieghafkSecret"
	os.Exit(m.Run())
}
