package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type LoginHandler struct {
	Name string
	Password string
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	decoder.UseNumber()
	err := decoder.Decode(&h)
	if err != nil {
		panic(err)
	}
	log.Println(h.Name)
}

func (h *BaseHandler) LoginAction(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		log.Println(r.URL)

		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if user, err := h.UserRepo.Select("1"); err != nil {
			fmt.Println("Error", user)
		}

		fmt.Println(name, weight)

		w.Write([]byte("Hello, World"))
	}




	/*decoder := json.NewDecoder(r.Body)

	decoder.UseNumber()
	err := decoder.Decode(&h.UserRepo)
	if err != nil {
		panic(err)
	}
	log.Println(h.UserRepo.Name)*/

}
