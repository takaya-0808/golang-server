package main

import(
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"time"
)

type UserInfo struct {
	Username string   `json:"username"`
	Email    string   `json:"username"`
	Password string   `json:"username"`
}

type TokenResponse struct {
	ID           int     `json:"id"`
	AccessToken  string  `json:"accesstoken"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/a",handle)
	
	server := &http.Server{
		Handler: r,
		Addr: "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	server.ListenAndServe()
	
}

func handle(w http.ResponseWriter, r *http.Request) {

	userinfo := UserInfo{
		Username: "hoge",
		Email:    "1222@gmail.com",
		Password: "password",
	}

	//req, err := http.NewRequest()

	res, err := json.Marshal(userinfo)
	fmt.Println(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	
	fmt.Fprint(w, userinfo)
    return
}
