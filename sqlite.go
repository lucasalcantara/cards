package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
)

type card struct {
	Id       int
	Question string
	Answer   string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	createTable()
}

func createTable() {
	db := database()
	defer db.Close()

	sqlStmt := `
	create table if not exists card (id integer not null primary key autoincrement, question text, answer text);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func database() *sql.DB {
	db, err := sql.Open("sqlite3", "./cards.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func insertCards(cards []card) {
	for _, c := range cards {
		insertCard(c)
	}
}

func insertCard(c card) {
	command := fmt.Sprintf("insert into card(question, answer) values('%s', '%s')", c.Question, c.Answer)

	db := database()
	defer db.Close()

	_, err := db.Exec(command)
	if err != nil {
		fmt.Println(command)
		log.Fatal(err)
	}
}

func deleteCard(id string) {
	db := database()
	defer db.Close()

	_, err := db.Exec("delete from card where id = " + id)
	if err != nil {
		log.Fatal(err)
	}
}

func allCards() []card {

	rows := rowResults("select id, question, answer from card")
	defer rows.Close()

	cards := parseResultToCard(rows)

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return cards
}

func shuffleCards() []card {
	cards := allCards()

	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

func rowResults(query string) *sql.Rows {
	db := database()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func parseResultToCard(rows *sql.Rows) []card {
	cards := make([]card, 0)

	for rows.Next() {
		c := card{}
		err := rows.Scan(&c.Id, &c.Question, &c.Answer)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, c)
	}

	return cards
}
