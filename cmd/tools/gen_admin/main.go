package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
)

func main() {
    user := flag.String("u", "admin", "username")
    pass := flag.String("p", "admin123", "password")
    name := flag.String("n", "Administrator", "name")
    role := flag.String("r", "admin", "role")
    flag.Parse()

    h, err := util.HashPassword(*pass)
    if err != nil {
        log.Fatalf("hash error: %v", err)
    }
    fmt.Printf("-- Run this SQL in your database to create admin user\n")
    fmt.Printf("INSERT INTO admins (username,password_hash,name,role,created_at,updated_at) VALUES ('%s','%s','%s','%s',NOW(),NOW());\n", *user, h, *name, *role)
}

