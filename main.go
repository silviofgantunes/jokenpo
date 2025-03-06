package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var moves = []string{"rock", "paper", "scissors"}

type Match struct {
	MovePlayer   string `json:"movePlayer"`
	MoveComputer string `json:"moveComputer"`
	Result       string `json:"result"`
}

func getWinner(movePlayer string, moveComputer string) string {
	switch {
	case movePlayer == moveComputer:
		return "draw"
	case (movePlayer == "rock" && moveComputer == "scissors") ||
		(movePlayer == "paper" && moveComputer == "rock") ||
		(movePlayer == "scissors" && moveComputer == "paper"):
		return "player"
	default:
		return "computer"
	}
}

func handlerPlay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not allowed method", http.StatusMethodNotAllowed)
	}

	var input struct {
		Move string `json:"move"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validates the player's move
	valid := false
	for _, c := range moves {
		if input.Move == c {
			valid = true
			break
		}
	}
	if !valid {
		http.Error(w, "Invalid move. Use 'rock', 'paper' ou 'scissors'", http.StatusBadRequest)
		return
	}

	// Random move done by computer
	rand.Seed(time.Now().UnixNano())
	computerChoice := moves[rand.Intn(len(moves))]

	// Determines who the winner is
	result := getWinner(input.Move, computerChoice)

	response := Match{
		MovePlayer:   input.Move,
		MoveComputer: computerChoice,
		Result:       result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/play", handlerPlay)
	http.ListenAndServe(":8080", nil)
}
