package q2

import (
	"errors"
	"regexp"
	"strings"
)

// Um aplicativo de processamento de texto precisa de uma função para contar a média de letras por
//palavra em um texto. A função deve receber um texto e retornar a média de letras por palavra.
//A função deve retornar um erro caso o texto não possua nenhuma palavra. O erro deve conter a mensagem "texto
//vazio".

func AverageLettersPerWord(text string) (float64, error) {
	regexp := regexp.MustCompile(`[^a-zA-Z ]+`)

	text = regexp.ReplaceAllString(text, "")

	words := strings.Fields(text)
	if len(words) == 0 {
		return 0, errors.New("texto vazio")
	}

	var totalLetters int
	for _, word := range words {
		totalLetters += len(word)
	}

	return float64(totalLetters) / float64(len(words)), nil
}
