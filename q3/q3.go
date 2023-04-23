package q3

import "errors"

// Você está desenvolvendo um programa que recebe uma lista de números inteiros como entrada e precisa encontrar o maior
//valor e o menor valor da lista. Além disso, você precisa calcular a média dos valores da lista, desconsiderando o maior
//e o menor valor na média.
//
//Escreva uma função que recebe uma lista de números inteiros como parâmetro e
//retorna uma tupla contendo o maior valor, o menor valor e a média dos valores da lista, desconsiderando o maior e o
//menor valor.
//
//A função deve retornar um erro caso a lista esteja vazia. O erro deve conter a mensagem "lista vazia".

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}

	var min, max, sum int
	min = numbers[0]
	max = numbers[0]
	for _, number := range numbers {
		if min == 0 || number < min {
			min = number
		}

		if max == 0 || number > max {
			max = number
		}

		sum += number
	}

	return max, min, float64(sum-min-max) / float64(len(numbers)-2), nil
}
