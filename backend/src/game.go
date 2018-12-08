package main

import (
	"encoding/json"
	"log"
)

const waitChallenge = "Waiting to be challenged"
const startGame = "Someone has accepted your challenge!"
const draw = "No winners yet!"
const resetWaitPaired = "Opponent has been disconnected... Waiting to get paired again"
const winner = " wins! Congratulations!"

type gameState struct {
	StatusMessage string   `json:"statusMessage"`
	Fields        []field  `json:"fields"`
	PlayerSymbols []string `json:"playerSymbols"`
	Started       bool     `json:"started"`
	Over          bool     `json:"over"`

	numberOfPlayers int
	PlayersTurn     int `json:"playersTurn"`
	numberOfMoves   int
}

type field struct {
	Set    bool   `json:"set"`
	Symbol string `json:"symbol"`
}

func newGameState() gameState {
	gs := gameState{
		StatusMessage:   waitChallenge,
		Fields:          emptyFields(),
		PlayerSymbols:   []string{0: "X", 1: "O"},
		Started:         false,
		numberOfPlayers: 0,
		PlayersTurn:     0,
	}
	return gs
}

func emptyFields() []field {
	return []field{
		{}, {}, {},
		{}, {}, {},
		{}, {}, {},
	}
}

func (gs *gameState) addPlayer() {
	gs.numberOfPlayers++
	switch gs.numberOfPlayers {
	case 1:
		gs.StatusMessage = waitChallenge
	case 2:
		gs.StatusMessage = startGame
		gs.Started = true
	}
}

func (gs *gameState) makeMove(playerNum int, moveNum int) {
	if moveNum <= 9 {
		if gs.isPlayersTurn(playerNum) {
			if gs.isLegalMove(moveNum) {
				gs.Fields[moveNum].Set = true
				gs.Fields[moveNum].Symbol = gs.PlayerSymbols[playerNum]
				gs.switchPlayersTurn(playerNum)
				gs.numberOfMoves++
				if won, symbol := gs.checkForWin(); won {
					gs.setWinner(symbol)
				} else {
					gs.checkForDraw()
				}
			}
		}
	} else {
		gs.specialMove(moveNum)
	}
}

func (gs *gameState) specialMove(moveNum int) {
	switch moveNum {
		case 10:
			gs.restartGame()
	}
}

func (gs *gameState) restartGame() {
	gs.StatusMessage = startGame
	gs.Fields = emptyFields()
	gs.Over = false
	gs.numberOfMoves = 0

}

func (gs *gameState) resetGame() {
	gs.restartGame()
	gs.Started = false
	gs.StatusMessage = resetWaitPaired
}

func (gs *gameState) checkLine(lineComponent1, lineComponent2, lineComponent3 int) bool {
	if gs.Fields[lineComponent1].Symbol  == gs.Fields[lineComponent2].Symbol &&
		gs.Fields[lineComponent2].Symbol == gs.Fields[lineComponent3].Symbol &&
		gs.Fields[lineComponent3].Symbol != "" {
		return true
	}
	return false
}

func (gs *gameState) checkForWin() (bool, string) {

	for i := 0; i < 3; i++ {
		if gs.checkLine(3*i, 3*i+1, 3*i+2) {
			return true, gs.Fields[3*i].Symbol
		}
	}

	for i := 0; i < 3; i++ {
		if gs.checkLine(i, i+3, i+6) {
			return true, gs.Fields[i].Symbol
		}
	}

	if gs.checkLine(0, 4, 8) {
		return true, gs.Fields[0].Symbol
	}

	if gs.checkLine(2, 4, 6) {
		return true, gs.Fields[2].Symbol
	}

	return false, ""
}

func (gs *gameState) setWinner(symbol string) {
	gs.StatusMessage = symbol + winner
	gs.Over = true
}

func (gs *gameState) checkForDraw() {
	if gs.numberOfMoves == 9 {
		gs.StatusMessage = draw
		gs.Over = true
	}
}

func (gs *gameState) isLegalMove(field int) bool {
	return !gs.Fields[field].Set
}

func (gs *gameState) isPlayersTurn(playerNum int) bool {
	return playerNum == gs.PlayersTurn
}

func (gs *gameState) switchPlayersTurn(playerNum int) {
	switch playerNum {
		case 0:
			gs.PlayersTurn = 1
		case 1:
			gs.PlayersTurn = 0
	}
}

func (gs *gameState) gameStateToJSON() []byte {
	gameState, err := json.Marshal(gs)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}
	return gameState
}
