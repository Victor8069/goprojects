package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Bien hecho!",
	"¡Buen trabajo!",
	"Deberías comprar un boleto de lotería",
}

var loseMessages = []string{
	"¡Qué lástima!",
	"¡Inténtalo de nuevo!",
	"Hoy simplemente no es tu día",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual.",
	"Oh no. Inténtalo de nuevo.",
	"Nadie gana, pero puedes intentarlo de nuevo.",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió TIJERA"
	}

	var message string

	switch playerValue {
	case computerValue:
		roundResult = "Es un empate"
		message = drawMessages[rand.Intn(3)]
	case (computerValue + 1) % 3:
		PlayerScore++
		roundResult = "¡El jugador gana!"
		message = winMessages[rand.Intn(3)]
	default:
		ComputerScore++
		roundResult = "¡La computadora gana!"
		message = loseMessages[rand.Intn(3)]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
