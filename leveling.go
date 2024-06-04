package main

import "fmt"

func main() {
	var i, playerHp, playerAtk, playerStamina int
	var EnemyAtk, EnemyHp, EnemyStamina int

	fmt.Scan(&playerHp, &playerAtk, &playerStamina)
	i = 1
	for playerHp > 100 && playerAtk > 20 && playerStamina > 100 {
		fmt.Println("Level", i)
		fmt.Scan(&EnemyHp, &EnemyAtk, &EnemyStamina)
		if EnemyAtk < playerAtk && EnemyHp < playerHp && EnemyStamina < playerStamina {
			playerAtk = playerAtk - EnemyAtk
			playerHp = playerHp - EnemyHp
			playerStamina = playerStamina - EnemyStamina
			i = i + 1
		} else {
			break
		}
	}
	fmt.Println("Congratulations!!")
	fmt.Println("You Finished", i-1, "Level!!")
	fmt.Println("Atk Remaining:", playerAtk)
	fmt.Println("Hp Remaining:", playerHp)
	fmt.Println("Stamina remaining:", playerStamina)
}
