package q1

import "errors"

//Uma loja de roupas precisa de uma função para verificar se um cliente é elegível para um desconto com base em seu
//histórico de compras. A função deve receber o valor da compra atual e o histórico de compras do cliente.
//A função deve retornar o valor do desconto e um erro. O erro deve ser nulo (ou seja, nil) se não houver nenhum erro.
//Caso o valor da compra atual seja menor ou igual a zero, a função deve retornar um erro com a mensagem "valor da compra
//inválido".
//
//Regras de desconto:
//
//* Se o valor total de compras do cliente for maior que R$ 1000,00, o desconto é de 10% do valor da compra atual.
//* Se o valor total de compras do cliente for menor ou igual a R$ 1000,00, o desconto é de 5% do valor da compra atual.
//* Se o valor total de compras do cliente for menor ou igual a R$ 500,00, o desconto é de 2% do valor da compra atual.
//* Se a compra atual for a primeira do cliente, o desconto é de 10% do valor da compra atual.
//* Se a média de compras do cliente for maior que R$ 1000,00, o desconto é de 20% do valor da compra atual.
//
//Os descontos não são cumulativos. Ou seja, se o cliente tiver direito a mais de um desconto, deve ser aplicado o maior
//desconto.

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, errors.New("valor da compra inválido")
	}

	var totalPurchases float64
	for _, purchase := range purchaseHistory {
		totalPurchases += purchase
	}

	var discount float64
	if totalPurchases > 1000 {
		discount = 0.1
	} else if totalPurchases > 500 {
		discount = 0.05
	} else {
		discount = 0.02
	}

	if len(purchaseHistory) == 0 {
		discount = 0.1
	}

	var averagePurchase float64
	if len(purchaseHistory) > 0 {
		averagePurchase = totalPurchases / float64(len(purchaseHistory))
	}

	if averagePurchase > 1000 {
		discount = 0.2
	}

	return currentPurchase * discount, nil
}
