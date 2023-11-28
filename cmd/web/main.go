package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Base Config for the server
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	// Middleware Setup
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Http Handler Setup
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Add Endpoints
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
