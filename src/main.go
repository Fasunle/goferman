package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func init() {
	// load environment variables from .env file
	err := loadEnv()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if port == "" {
		port = ":8080"
		fmt.Printf("PORT environment variable not set. Defaulting to port %s.\n", strings.Replace(port, ":", "", 1))
	}

	fmt.Printf("Starting server on http://localhost%s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Welcome to my Go web server.")
	})

	http.ListenAndServe(port, nil)
}

func loadEnv() error {
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return fmt.Errorf(".env file does not exist")
	}

	// Open the .env file
	file, err := os.Open(".env")
	if err != nil {
		return fmt.Errorf("failed to open .env file: %v", err)
	}
	defer file.Close()

	// Read the .env file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue // Skip empty lines and comments
		}
		keyValue := strings.SplitN(line, "=", 2)
		if len(keyValue) != 2 {
			return fmt.Errorf("invalid line in .env file: %s", line)
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %v", err)
	}

	return nil
}
