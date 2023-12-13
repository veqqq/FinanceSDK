package main

import (
	"FinanceSDK/e"
	"encoding/json"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// for new readers
// the global "structType" drives JsonToPostgres() and QueryBuilder() which are
// the main workhorses here. WhatDoesUserWantToDo() is the CLI and invokes everything

func main() {

	for {
		WhatDoesUserWantToDo() // in UI.go, houses the CLI and all invocations
	}
}

// build the initial url, strongly coupled to alphavantage #todo

var apiKey string

// build a base url, fetching the apikey from .env file
// lazy implementation from godotenv to reduce dependencies
func buildBaseURL() string {
	f, err := (os.Open("e/.env")) // e for err, but hide from docker
	e.Check(err)
	defer f.Close()

	var envMap map[string]string
	err = json.NewDecoder(f).Decode(&envMap)
	e.Check(err)

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}
	for key, value := range envMap {
		if !currentEnv[key] {
			_ = os.Setenv(key, value)
		}
	}
	apiKey = os.Getenv("X-RapidAPI-Key")
	// apiKey, ok := os.LookupEnv("APIKEY")
	// if !ok {
	// 	log.Fatalf("Add API Key to .env")
	// }
	return "https://alpha-vantage.p.rapidapi.com/query?" + "function="
}

var baseUrl string    // global
var structType string // global, is specified in query builder,
// then used to marshal json, manage sql inserts...

func init() {
	baseUrl = buildBaseURL()
}
