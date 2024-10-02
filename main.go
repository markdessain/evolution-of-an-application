package main

import (
	"net/http"
	"strconv"
)

var counter int

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/increment", increment)
	http.HandleFunc("/decrement", decrement)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>Evolution</title>
				<script src="https://unpkg.com/htmx.org@2.0.2"></script>
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.classless.slate.min.css" />
			</head>
			<body>
				<main>
					<div role="group">
						<button hx-target="#counter" hx-get="/increment">Increment</button>
						<article>Counter: <span  id="counter">` + strconv.Itoa(counter) + `</span></article>
						<button hx-target="#counter" hx-get="/decrement">Decrement</button>
					</div>
				</main>
			</body>
		</html>
	`))
}

func increment(w http.ResponseWriter, r *http.Request) {
	counter += 1
	w.Write([]byte(strconv.Itoa(counter)))
}

func decrement(w http.ResponseWriter, r *http.Request) {
	counter -= 1
	w.Write([]byte(strconv.Itoa(counter)))
}

func name() string {
	return "Evolution"
}
