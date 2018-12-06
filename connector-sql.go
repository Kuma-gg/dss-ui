package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
