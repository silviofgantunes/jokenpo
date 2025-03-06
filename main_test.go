package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWinner(t *testing.T) {
	tests := []struct {
		playerMove   string
		computerMove string
		result       string
	}{
		{"rock", "scissors", "player"},
		{"scissors", "paper", "player"},
		{"paper", "rock", "player"},
		{"rock", "paper", "computer"},
		{"scissors", "rock", "computer"},
		{"paper", "scissors", "computer"},
		{"rock", "rock", "draw"},
		{"paper", "paper", "draw"},
		{"scissors", "scissors", "draw"},
	}

	for _, test := range tests {
		result := getWinner(test.playerMove, test.computerMove)
		if result != test.result {
			t.Errorf("For moves (%s vs %s), expected %s but got %s", test.playerMove, test.computerMove, test.result, result)
		}
	}
}

func TestHandlerPlay(t *testing.T) {
	tests := []struct {
		input         string
		expectedCode  int
		expectedError string
	}{
		{"{\"move\": \"rock\"}", http.StatusOK, ""},
		{"{\"move\": \"invalid\"}", http.StatusBadRequest, "Invalid move"},
		{"{\"wrong_field\": \"rock\"}", http.StatusBadRequest, "Invalid move. Use 'rock', 'paper' ou 'scissors"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/play", bytes.NewBuffer([]byte(test.input)))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlerPlay)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedCode {
			t.Errorf("For input %s, expected status %d but got %d", test.input, test.expectedCode, rr.Code)
		}

		if test.expectedError != "" && !bytes.Contains(rr.Body.Bytes(), []byte(test.expectedError)) {
			t.Errorf("For input %s, expected error message containing '%s' but got '%s'", test.input, test.expectedError, rr.Body.String())
		}
	}
}
