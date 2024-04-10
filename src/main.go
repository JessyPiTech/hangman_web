package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	tmpl           *template.Template
	Page           = "la"
	Statue         bool
	Ms             = ""
	pseudo         = ""
	password       = ""
	score          int
	attempts       int
	guess          = "0"
	lastGuess      = "8"
	allguess       string
	word           string
	WordHide       = ""
	guessedLetters []bool
	Test           = false
	Essay          string
	b              = false
	verif          = false
)

// Fonction principale
func main() {
	//on charge notre page html
	var err error
	tmpl, err = loadTemplate("static/html/index.html")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}

	//la on gere se qui il a dans l'url pour dire genre si t'es dans cet url fais se code...

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/Compte", CompteHandler)
	http.HandleFunc("/Ac", AcHandler)
	http.HandleFunc("/hangman", HangmanHandler)
	http.HandleFunc("/replay", ReplayHandler)
	http.HandleFunc("/win", WinHandler)
	http.HandleFunc("/defeat", DefeatHandler)
	http.HandleFunc("/EasyMode", EasyModeHandler)
	http.HandleFunc("/MediumMode", MediumModeHandler)
	http.HandleFunc("/HardMode", HardModeHandler)
	http.HandleFunc("/dev", DevHandler)
	http.HandleFunc("/profil", ProfilHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/mise", MisHandler)
	http.HandleFunc("/rego", RegoHandler)
	//start du serveur
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server started on localhost:8901")
	http.ListenAndServe(":8901", nil)
}

//comme tu a du le lire dans le hangman c'est le meme code
//voir en un peu mieux par se que j'ai fais
// hangman classic puis hangman web puis hangman classic puis hangman web donc globalement je dois avoir poncée le truc
//j'ai pas remis les explications mais c'est normalement les mêmes  que pour le hangman classic
//comme la dernier fois si jamis
//il y a un index out of rang c'est que le fichier joueur est cassez sa arrive aleatoirement meme si une petit mise a jour est cense l'avoir corrige
//car le probleme étais que le programme recris sur le fichier text au lieu de le remplacer mais normalement c'est corriger
//mais si sa arrive supprime le fichier joueur dans le dossier player
//mrc bonne chance
