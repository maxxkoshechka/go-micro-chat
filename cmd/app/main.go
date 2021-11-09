package main

import (
	"chat/controllers"
	"chat/internal"
	"chat/repository"
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	pgx "github.com/jackc/pgx/v4"
	godotenv "github.com/joho/godotenv"
)

var addr = flag.String("addr", ":8081", "http service address")

type DBConfig struct {
	Username string
	Password string
	DBName string
}

type Config struct {
	DB DBConfig
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	conf := New()
	// Open up our database connection.
	db, _ := pgx.Connect(context.Background(), "postgres://" + conf.DB.Username + ":" + conf.DB.Password + "@localhost:5432/" + conf.DB.DBName)

	// defer the close till after the main function has finished
	// executing
	defer db.Close(context.Background())
	/*var greeting string
	//
	conn.QueryRow(context.Background(), "select email from \"user\"").Scan(&greeting)
	fmt.Println(greeting)*/

	userRepo := repository.NewUserRepo(db)

	h := controllers.NewBaseHandler(userRepo)

	flag.Parse()
	hub := internal.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	//http.Handle("/login", corsHandler(new(controllers.LoginHandler)))
	http.HandleFunc("/login", h.LoginAction)
	//http.Handle("/login", corsHandler(new(controllers.LoginHandler)))
	//http.Handle("/login", baseHandlerRouter(h))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		internal.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		DB: DBConfig{
			Username: getEnv("POSTGRES_USER", ""),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			DBName: getEnv("POSTGRES_DB", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "../../resources/home.html")
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		} else {
			h.ServeHTTP(w,r)
		}
	}
}

func baseHandlerRouter (h *controllers.BaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		} else {
			h.LoginAction(w, r)
		}
	}
}
