package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
    "github.com/gomodule/redigo/redis"
	"time"
	_"os"
	"github.com/bastianNebra/myapigorm.git/pkg/storage"
)
var store *sessions.CookieStore
var redisPool *redis.Pool


func main() {
	fmt.Println("Hallo and welcome to my new API")
	//Connection a la db
	shop := storage.NewServerShop()
	
	fmt.Println("Connecte avec success")
	//Router
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", getHallo()).Methods("GET")
	r.HandleFunc("/api/jwt/conn",shop.Conn()).Methods("POST")
	r.HandleFunc("/customers", shop.GetCustomers()).Methods("GET")
	r.HandleFunc("/customers/{id}", shop.GetCustomer()).Methods("GET")
	r.HandleFunc("/customers", shop.CreateCustomers()).Methods("POST")
	r.HandleFunc("/customers/{id}", shop.UpdateCustomer()).Methods("PUT")
	r.HandleFunc("/customers/{id}", shop.DeleteCustomer()).Methods("DELETE")

	//Start a server
	log.Fatal(http.ListenAndServe(":9000", r))

}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("INFO: %s %s\n", r.Method, r.RequestURI)
        next.ServeHTTP(w, r)
    })
}

func init() {
	// Connexion à Redis
	redisPool = &redis.Pool{
		MaxIdle:     80,
		MaxActive:   12000,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	// Configuration du CookieStore
	store = sessions.NewCookieStore([]byte("my-secret-key"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	}
}

func getHallo() http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		if session.IsNew {
			session.Values["foo"] = "bar New"
			session.Save(r, w)
			fmt.Fprintf(w, "Nouvelle session créée. Valeur de foo: %s\n", session.Values["foo"])
		} else {
			fmt.Fprintf(w, "Session existante. Valeur de foo: %s\n", session.Values["foo"])
		}
	storage.NewServerShop().GetHallo()
	}
}
