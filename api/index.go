package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Todo struct {
	Name     string    `json:"name"`
	Computed bool      `json:"computed"`
	Due      time.Time `json:"due"`
}

type Todos []Todo

func Handler(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
	fmt.Println("Encode JSON")
}
