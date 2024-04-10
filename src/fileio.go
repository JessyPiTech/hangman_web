package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fonctions liées à la lecture et à l'écriture de fichiers

// suppression
func supprimerFichier(pseudo string) {
	filename := "players/" + pseudo + ".txt"

	// Vérifiez si le fichier existe
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Le fichier n'existe pas :", filename)
		return
	}
	//puis supprime
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier :", err)
		return
	}
	fmt.Println("Fichier supprimé avec succès :", filename)
}

// creation
func creerFichierJoueur(pseudo string, contenu string) {
	filename := "players/" + pseudo + ".txt"

	//la on le cree le fichier
	fichier, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}

	//on cree un writer (tampon)
	writer := bufio.NewWriter(fichier)
	//et on lui donne un comptenu a ecrire
	_, err = writer.WriteString(contenu)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
	// on s'assure que toutes les données tamponnées sont écrites dans le fichier
	err = writer.Flush()
	if err != nil {
		fmt.Println("Erreur lors du vidage du tampon dans le fichier :", err)
		return
	}
	//ici le b c'est pour faire une difference de quand on cree un ficheier de quand on fais simplemnt un mise a jour
	if !b {
		fmt.Println("Compte créé avec succès :", pseudo, ",", password)
	}
	// on s'assure ensuite de ferme le fichier a la fin du main
	defer fichier.Close()
}

// mise a jour
func mise() {
	b = true
	//donc d'abord on supprime l'anciene sauvegarde
	supprimerFichier(pseudo)
	Essay := strconv.Itoa(attempts)
	Sco := strconv.Itoa(score)
	resultat := ""
	for _, valeur := range guessedLetters {
		resultat = resultat + strconv.FormatBool(valeur) + " "
	}
	//on convertie tout en string pour le metre en contenu a metre dans le .txt
	contenu := password + "\n" + WordHide + "\n" + resultat + "\n" + Essay + "\n" + Sco + "\n" + allguess + "\n" + word + "\n"
	creerFichierJoueur(pseudo, contenu)
	fmt.Println("Sauvegarde créé avec succès :", pseudo, ",", password)
}

// recuperation
func recup() {
	filename := "players/" + pseudo + ".txt"
	// Ouvrez le fichier en mode lecture
	lastinfo, err := ReadWords2(filename) //la on recupe la tout le text du fichier wo.txt
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	//sinon on recupe le mots on enleve le \n
	ty := strings.TrimSpace(lastinfo[1])

	//puis pour chauqe caratere du mots on va soit atribué un true ou un false
	//si c'est un _ on met false si c'est une lettre on met vrai
	for _, caractere := range ty {
		if caractere == '_' {
			guessedLetters = append(guessedLetters, false)
		} else if IsAlpha(string(caractere)) {
			guessedLetters = append(guessedLetters, true)
		}
	}
	//la variable test permet de skipe la condition du post au niveau du hangman
	Test = true
	//puis on recupe tout les info avec de la guestion d'erreure
	//j'ai laisser les prints pour une gestion d'erreur terminal pour que se soit plus facile
	attempts, err = strconv.Atoi(strings.TrimSpace(lastinfo[3]))
	if err != nil {
		fmt.Println("Erreur de conversion en entier:", err)
		Ms = "Erreur de conversion en entier:"
		attempts = 5
	}
	score, err = strconv.Atoi(strings.TrimSpace(lastinfo[4]))
	if err != nil {
		fmt.Println("Erreur de conversion en entier:", err)
		Ms = "Erreur de conversion en entier:"
		score = 0
	}
	allguess = lastinfo[5]
	word = strings.TrimSpace(lastinfo[6])
	word = strings.TrimRight(lastinfo[6], "\n")
	Page = "la"
	WordHide = ReWriting(word, guessedLetters)
	Essay = strconv.Itoa(attempts)
}
