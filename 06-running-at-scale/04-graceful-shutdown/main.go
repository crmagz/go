package main

import "fmt"

func main() {
	// TODO: create an *http.ServeMux with a handler on "/" that prints
	// "request received, working for 3s...", sleeps 3 seconds, prints
	// "request completed", then writes "done" to the response.

	// TODO: create srv := &http.Server{Addr: "localhost:8090", Handler: mux}.

	// TODO: declare serverErr := make(chan error, 1). In a goroutine, print
	// "listening on http://localhost:8090", then send srv.ListenAndServe()'s
	// return value on serverErr.

	// TODO: declare sigCh := make(chan os.Signal, 1) and call
	// signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM).

	// TODO: select on serverErr and sigCh. On serverErr, print
	// "server error: <err>" and return. On sigCh, print
	// "received signal: <sig>, shutting down gracefully...".

	// TODO: create a context.WithTimeout of 5 seconds (defer its cancel), call
	// srv.Shutdown(ctx), and print "graceful shutdown complete" or
	// "graceful shutdown failed: <err>" depending on the result.
	fmt.Println("implement me")
}
