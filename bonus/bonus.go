package bonus

import "errors"

//Um aplicativo de jogos precisa de uma função para calcular o dano causado por um personagem em um inimigo. A função
//deve receber o nível do personagem e o nível do inimigo e retornar o dano causado. A função deve retornar um erro
//caso o nível do personagem ou do inimigo seja inválido. O erro deve conter a mensagem "nível inválido".
//
//O dano causado é calculado da seguinte forma:
//
//* Se o nível do personagem for maior que o nível do inimigo, o dano causado é igual ao nível do personagem multiplicado
//  por 10.
//* Se o nível do personagem for menor que o nível do inimigo, o dano causado é igual ao nível do personagem multiplicado
//  por 5.
//* Se o nível do personagem for igual ao nível do inimigo, o dano causado é igual ao nível do
//  personagem multiplicado por 7.
//* Se o nível do personagem for maior que 100, o dano causado é igual ao nível do personagem multiplicado por 20.
//* Se o nível do inimigo for maior que 100, o dano causado é igual ao nível do personagem multiplicado por 2.
//* Se a diferença entre o nível do personagem e o nível do inimigo for maior que 20, o dano causado é multiplicado por 5.
//* Se a diferença entre o nível do personagem e o nível do inimigo for menor que 20, o dano causado é por 2.

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, errors.New("nível inválido")
	}

	var damage int

	if characterLevel > 100 {
		damage = characterLevel * 20
	} else if enemyLevel > 100 {
		damage = characterLevel * 2
	} else if characterLevel > enemyLevel {
		damage = characterLevel * 10
	} else if characterLevel < enemyLevel {
		damage = characterLevel * 5
	} else {
		damage = characterLevel * 7
	}

	if diff := characterLevel - enemyLevel; diff > 20 {
		damage *= 5
	} else if diff < -20 {
		damage *= 2
	}

	return damage, nil
}
