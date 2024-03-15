package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// command line flag -addr for address, defaults to :4000
	// e.g. go run ./cmd/web/ -addr=":8000"
	// by using the flag library, we also get the -help command
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view/", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
