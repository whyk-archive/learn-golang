package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
)

type Todo struct {
	Name     string    `json:"name"`
	Computed bool      `json:"computed"`
	Due      time.Time `json:"due"`
}

type Todos []Todo

type Config struct {
	APIKey string `env:"API_KEY"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		panic(err)
	}

	authorization := r.Header.Get("API_KEY")
	if authorization == "" || authorization != conf.APIKey {
		return
	}

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
	fmt.Println("Encode JSON")
}
