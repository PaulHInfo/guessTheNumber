package main

	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())
	nombreSecret := rand.Intn(100) + 1 // Nombre entre 1 et 100

	fmt.Println("Bienvenue dans le jeu de devinette de nombres !")
	fmt.Println("Je pense à un nombre entre 1 et 100. Pouvez-vous le deviner ?")

	scanner := bufio.NewScanner(os.Stdin)
	tentatives := 0

	for {
		fmt.Print("Entrez votre devinette : ")
		scanner.Scan()
		entree := scanner.Text()
		devinette, err := strconv.Atoi(strings.TrimSpace(entree))

		if err != nil {
			fmt.Println("Veuillez entrer un nombre valide.")
			continue
		}

		tentatives++

		if devinette < nombreSecret {
			fmt.Println("C'est plus grand.")
		} else if devinette > nombreSecret {
			fmt.Println("C'est plus petit.")
		} else {
			fmt.Printf("Félicitations ! Vous avez trouvé le nombre en %d tentatives.\n", tentatives)
			break
		}
	}

	fmt.Println("Merci d'avoir joué !")
}
