package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Film struct {
	Title    string
	Director string
}

var (
	addr = flag.String("addr", "localhost:8089", "the address to connect to")
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {},
		}
		tmpl.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		//time.Sleep(1 * time.Second)
		conn := NewGRPCClient(":8089")
		defer conn.Close()

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
