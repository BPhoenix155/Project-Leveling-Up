package main

import "fmt"

func main() {
	var hp, atk, def int
	var iv float64

	fmt.Scan(&hp, &atk, &def)
	iv = ((float64(hp+atk+def) / 45) * 100)
	if (hp >= 0 && hp <= 15) && (atk >= 0 && atk <= 15) && (def >= 0 && def <= 15) {
	} else if (hp < 0 || hp > 15) || (atk < 0 || atk > 15) || (def < 0 || def > 15) {
		fmt.Println("Not a Pokemon")
	} else if iv < 0 || iv > 100 {
		fmt.Println("Not a Pokemon")
	}
	if iv >= 0 && iv <= 48.90 {
		fmt.Println("Overall, your pokémon has room for improvement as far as battling goes")
		fmt.Println("0 Stars")
	} else if iv >= 51.10 && iv <= 64.45 {
		fmt.Println("Overall, your pokémon is pretty decent!")
		fmt.Println("1 Stars")
	} else if iv >= 66.60 && iv <= 80 {
		fmt.Println("Overall, your pokémon is really strong!")
		fmt.Println("2 Stars")
	} else if iv >= 82.22 && iv <= 97.8 {
		fmt.Println("Overall, your pokémon looks like it can really battle with the best of them!")
		fmt.Println("3 Stars")
	} else if iv == 100 {
		fmt.Println("Overall, your pokémon looks like it can really battle with the best of them!")
		fmt.Println("4 Stars")
	}
}
