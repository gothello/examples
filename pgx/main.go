package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Client struct {
	Username string
	Password string
	Email    string
}

var (
	QUERY_CREATE_TABLE = `CREATE TABLE IF NOT EXISTS client (id varchar(120), password varchar(120), email varchar(120));`
)

func CreateTable(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), QUERY_CREATE_TABLE)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postdb")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// if err := CreateTable(conn); err != nil {
	// 	log.Fatalln(err)
	// }

	// queryInsert := `INSERT INTO client (id, password, email) VALUES ($1,$2,$3);`

	// if _, err := conn.Exec(context.Background(), queryInsert, uuid.NewString(), "1234", "adam@gmail.com"); err != nil {
	// 	log.Println(err)
	// }

	// queryUser := `SELECT * FROM client;`

	// rows, err := conn.Query(context.Background(), queryUser)
	// if err != nil {
	// 	log.Println(err)
	// }

	// var cls []Client

	// for rows.Next() {
	// 	var cl Client
	// 	err = rows.Scan(&cl.Username, &cl.Password, &cl.Email)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 		continue
	// 	}

	// 	cls = append(cls, cl)
	// }

	// fmt.Println(cls)

	// queryUser = `SELECT * FROM client where id=$1;`
	// cl := &Client{}

	// _ = conn.QueryRow(context.Background(), queryUser, "de826d5a-cce2-4483-b64a-1813f817828e").Scan(&cl.Username, &cl.Password, &cl.Email)
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(cl)
}
