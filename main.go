package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

type Link struct {
	Link     string `json:"link"`
	LinkIcon string `json:"link_icon"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.Name == params["name"] {
			users = append(users[:index], users[index+1:]...)
			var User User
			_ = json.NewDecoder(r.Body).Decode(&User)
			User.Name = params["name"]
			users = append(users, User)
			json.NewEncoder(w).Encode(User)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.Name == params["Name"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()

	igorArr := make([]Link, 0, 1)
	var link Link = Link{Link: "vk.com/Igor", LinkIcon: "https://cdn0.iconfinder.com/data/icons/social-network-7/50/11-512.png"}
	igorArr = append(igorArr, link)

	meArr := make([]Link, 0, 1)
	link = Link{Link: "vk.com/MeIRL", LinkIcon: "https://cdn0.iconfinder.com/data/icons/social-network-7/50/11-512.png"}
	meArr = append(meArr, link)

	users = append(users, User{Name: "Igor", Links: igorArr})
	users = append(users, User{Name: "ME", Links: meArr})
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{name}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{name}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")

	fmt.Println(users)

	log.Fatal(http.ListenAndServe(":8000", r))
}
