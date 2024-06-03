package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"

	"github.com/SalemLU/GoServerTest/getProducer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Film struct {
	Title    string
	Year     int64
	Director string
}

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {},
		}
		tmpl.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		title := r.PostFormValue("title")
		year, err := strconv.ParseInt(r.PostFormValue("year"), 10, 64)
		if err != nil {
			panic(err)
		}
		//time.Sleep(1 * time.Second)
		conn := NewGRPCClient(":8089")
		defer conn.Close()
		cc := getProducer.NewGetProducerClient(conn)
		cr := &getProducer.CreateRequest{
			Film: &getProducer.Film{
				Title: title,
				Year:  year,
			},
		}
		director, err := cc.Create(context.Background(), cr)
		if err != nil {
			log.Fatalf("Could not retrieve value: %v", err)
		}

		ds := director.String()
		re := regexp.MustCompile(`"(.*)"`)
		match := re.FindStringSubmatch(ds)
		if !(len(match) > 1) {
			fmt.Println("match not found")
		} else {
			if match[1] == "not found" {
				w.Header().Set("HX-Retarget", "#film-list-error")
				w.Header().Set("HX-Reswap", "innerHTML")
				tmpl.ExecuteTemplate(w, "film-not-found", "This film could not be found in our database")
			} else {
				tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Year: year, Director: match[1]})
			}
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
