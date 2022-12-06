package main

import (
	"fmt"
	"log"

	"github.com/rqlite/gorqlite"
)

func main() {
	conn, err := gorqlite.Open("http://localhost:4001/") // same only explicitly
	if err != nil {
		log.Fatalf("get conn fail err: %v", err)
	}
	conn.SetConsistencyLevel("none")

	if _, err := conn.WriteOne("CREATE TABLE demo (id integer, name text)"); err != nil {
		log.Fatalf("create table fail err: %v", err)
	}

	statements := make([]string, 0)
	pattern := "INSERT INTO demo(id, name) VALUES (%d, '%s')"
	statements = append(statements, fmt.Sprintf(pattern, 125718, "Speed Gibson"))
	statements = append(statements, fmt.Sprintf(pattern, 209166, "Clint Barlow"))
	statements = append(statements, fmt.Sprintf(pattern, 44107, "Barney Dunlap"))
	WriteResult, err := conn.Write(statements)

	// now we have an array of []WriteResult
	for n, v := range WriteResult {
		fmt.Printf("for result %d, %d rows were affected\n", n, v.RowsAffected)
		if v.Err != nil {
			fmt.Printf("   we have this error: %s\n", v.Err.Error())
		}
	}

	defer conn.Close()
}
