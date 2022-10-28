package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var tabligne [50]string
var y int
var tabbyte []byte
var tirettab []string
var tentatives int
var allerte string

func illustrator() {
	switch tentatives {
	case 9:
		fmt.Print("\n \n \n \n \n \n=========\n")
	case 8:
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")
	case 7:

		fmt.Println("   +---+ ")
		fmt.Println("       | ")
		fmt.Println("       | ")
		fmt.Println("       | ")
		fmt.Println("       | ")
		fmt.Println("       | ")
		fmt.Println("=========")
	case 6:
		fmt.Println("  +---+ ")
		fmt.Println("  |   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")

	case 5:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")

	case 4:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println("  |   | ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")

	case 3:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|   |  ")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")

	case 2:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|", string(92), "|")
		fmt.Println("      |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")

	case 1:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|", string(92), "|")
		fmt.Println(" /    |  ")
		fmt.Println("      |  ")
		fmt.Println("=========")

	case 0:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("  O   |  ")
		fmt.Println(" /|", string(92), "|")
		fmt.Println(" / ", string(92), "| ")
		fmt.Println("      | ")
		fmt.Println("=========")
	}

}
func lecture() {

	f, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0

	for scanner.Scan() {

		tabligne[i] = scanner.Text()
		i = i + 1

	}

	// ca affiche le mot de la position y

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func decouper() {
	tabbyte = []byte(tabligne[y])
	tiret()
	connu := (len(tirettab) / 2) - 1
	tirettab[connu] = string(tabbyte[connu])
	consigne()
}
func afficher() {
	for i := 0; i < len(tirettab); i++ {
		fmt.Printf(tirettab[i])

	}

}
func tiret() {
	tirettab = make([]string, len(tabbyte))
	//make pour cree un tableau
	for i := 0; i < len(tirettab); i++ {
		tirettab[i] = "_"

	}

}

func consigne() {
	fmt.Println("completez les mots  ci dessous : ", tentatives, "tentatives")
	afficher()
	fmt.Println("")
	fmt.Scanln(&allerte)
	//il va recuperer ce que lutilisateur a saisie

	comparaison()

}
func comparaison() {
	for i := 0; i < len(tabbyte); i++ {
		if allerte != string(tabbyte[i]) && i == len(tabbyte)-1 {
			fmt.Print("veuiller reessayer \n")
			tentatives--
		} else if allerte == string(tabbyte[i]) {
			for y := 0; y < len(tabbyte); y++ {
				if string(tabbyte[y]) == allerte {

					tirettab[y] = allerte
					continue
				}

			}
			break
		}

	}

	verfication()
}
func verfication() {
	illustrator()
	if tentatives == 0 {
		fmt.Println("_|         _|_|         _|_|_|     _|_|_|_|  ")
		fmt.Println("_|       _|    _|      _|          _|        ")
		fmt.Println("_|       _|    _|        _|_|      _|_|_|    ")
		fmt.Println("_|       _|    _|           _|    _|        ")
		fmt.Println("_|_|_|_|   _|_|         _|_|_|     _|_|_|_|  ")

		main()
	} else {
		for i := 0; i < len(tirettab); i++ {
			if tirettab[i] == "_" {
				consigne()
			} else if i == len(tirettab)-1 {
				afficher()
				fmt.Println("_|         _| _|_|_|  _|      _|")
				fmt.Println("_|         _|   _|    _|_|    _| ")
				fmt.Println("_|    _|   _|   _|    _|  _|  _|   ")
				fmt.Println("  _|  _|  _|    _|    _|    _|_| ")
				fmt.Println("   _|  _|     _|_|_|  _|      _| ")
				main()
			}

		}
	}

}
func main() {
	tentatives = 10
	y = rand.Intn(37)
	// rand cest pour choisir au hasard,
	lecture()
	decouper()

}
