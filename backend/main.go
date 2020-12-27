package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Global handles
var driver neo4j.Driver

var DB_URL = "bolt://localhost:7687"
var DB_USERNAME = "neo4j"
var DB_PASSWORD = "axiom"

//****************************************** MAIN ********************************************************//
func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Setting up neo4j
	var err error
	driver, err = neo4j.NewDriver(DB_URL, neo4j.BasicAuth(DB_USERNAME, DB_PASSWORD, ""))
	if err != nil {
		fmt.Printf("Error initialising neo4j driver: %s", err)
		return
	}

	defer driver.Close()

	// Check connectivity
	if err := driver.VerifyConnectivity(); err != nil {
		fmt.Printf("Error verifying neo4j driver: %s", err.Error())
		return
	}

	// Adding the routes
	r := mux.NewRouter()
	r.HandleFunc("/", TestHandler)
	r.HandleFunc("/concept/new", CreateConcept)

	// Addings the middlewares
	r.Use(corsMiddleware)

	// Setting Up the Server
	srv := &http.Server{
		Addr: "0.0.0.0:3000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("Knowledge Server up and running at localhost:3000")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

//****************************************** Handlers ********************************************************//
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gorilla!\n"))
}

type createConceptRequest struct {
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Prerequisites []string `json:"prerequisites"`
}

func CreateConcept(w http.ResponseWriter, r *http.Request) {
	var req createConceptRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(req.Title)

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	// Generate UUID
	id, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(id.String())

	c := map[string]string{
		"uuid":    id.String(),
		"title":   req.Title,
		"content": req.Content,
	}

	fmt.Println(c["uuid"])

	// Create the concept and connect prereuisites
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run("CREATE (c:Concept $concept)", map[string]interface{}{"concept": c})
		if err != nil {
			return nil, err
		}

		for _, prereqUuid := range req.Prerequisites {
			_, err = transaction.Run("MATCH (a:Concept {uuid: $uuid1}), (b:Concept {uuid: $uuid2}) MERGE (a)-[r:PREREQ_OF]->(b)",
				map[string]interface{}{"uuid1": prereqUuid, "uuid2": id.String()})
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//---------------------- MIDDLEWARES ------------------------------//
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		next.ServeHTTP(w, r)
	})
}
