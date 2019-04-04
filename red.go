package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/gorilla/mux"
)

// The Person Type (more like an object)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address is here
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// AddRed has comments
func AddRed(w http.ResponseWriter, r *http.Request) {

	// Specify profile for config and region for requests
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		fmt.Println("ADD Red session error")
		log.Fatal(err)
	}

	svc := redshift.New(sess)

	result, err := svc.DescribeClusters(&redshift.DescribeClustersInput{})

	if err != nil {
		fmt.Println("ADDRed describe error")
		fmt.Println(err)
		os.Exit(1)
	}

	for _, n := range result.Clusters {
		fmt.Println(*n)
	}

	// Make connection
	db, err := MakeRedshfitConnection("strategyone", "StrategyOne1!", "strategy-one.chhdcmijmh2c.us-east-1.redshift.amazonaws.com", "5439", "dev")

	if err != nil {
		fmt.Println("ADDRed MakeRedshfitConnection error")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Have db - gonna query", db)

	rows, rerr := db.Query("SELECT * FROM customer")

	if rerr != nil {
		fmt.Println("ADDRed query error")
		fmt.Println(err)
		os.Exit(1)
	}

	var (
		color    int
		shoetype string
	)

	for rows.Next() {
		err := rows.Scan(&color, &shoetype)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(color, shoetype)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(make([]int, 0))
}

// MakeRedshfitConnection used to connect to redshift
func MakeRedshfitConnection(username, password, host, port, dbName string) (*sql.DB, error) {
	fmt.Println("MakeRedshfitConnection")
	url := fmt.Sprintf("sslmode=require user=%v password=%v host=%v port=%v dbname=%v",
		username,
		password,
		host,
		port,
		dbName)

	var err error
	var db *sql.DB
	fmt.Println("MakeRedshfitConnection2")
	if db, err = sql.Open("postgres", url); err != nil {
		return nil, fmt.Errorf("redshift connect error : (%v)", err)
	}

	fmt.Println("MakeRedshfitConnection Ping")
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("redshift ping error : (%v)", err)
	}

	fmt.Println("MakeRedshfitConnection return")
	return db, nil
}

// DeletePerson an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

// main function to boot up everything
func mainer() {

	fmt.Println("api running")

	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	router.HandleFunc("/api/people", GetPeople).Methods("GET")
	router.HandleFunc("/api/addred", AddRed).Methods("GET")
	router.HandleFunc("/api/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/api/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/api/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
