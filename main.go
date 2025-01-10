package main

/*

Petit projet pour apprendre les bases de Go
But : deviner le nombre choisi aléatoirement
Date : 10.01.25
comm : code surrement pas opti

*/
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())   // init du gen
	nombreSecret := rand.Intn(100) + 1 // Nombre entre 1 et 100

	fmt.Println("Bienvenue dans le jeu de devinette de nombres !")
	fmt.Println("Je pense à un nombre entre 1 et 100. le deviner !")

	scanner := bufio.NewScanner(os.Stdin)
	tentatives := 0

	for {
		fmt.Print("Entre ta devinette : ")
		scanner.Scan()
		entree := scanner.Text()
		devinette, err := strconv.Atoi(strings.TrimSpace(entree))

		if err != nil {
			fmt.Println("Entre un nombre valide.")
			continue
		}

		tentatives++

		if devinette < nombreSecret {
			fmt.Println("C'est plus grand ")
		} else if devinette > nombreSecret {
			fmt.Println("C'est plus petit ")
		} else {
			fmt.Printf("Bravo ! Tu as trouvé le nombre en %d tentatives.\n", tentatives)
			break
		}
	}

	fmt.Println("Merci d'avoir joué !")
}
