package main

import (
	"fmt"
	"html/template"
)

// chargement de la page
func loadTemplate(filename string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Error loading template:", err)
	}
	return tmpl, err
}

// sa c'est une fonction qui fais passer mon code de a peut pret 1500 ligne a 1000 quand tout le code est regrouper
// en gros c'est la data que je fais envoiller a ma page
func createData(Statue bool, WordHide string, allguess string, Essay string, score int, Ms string, Page string, pseudo string) struct {
	Statue        bool
	Word          string
	GuessedLetter string
	Essay         string
	Score         int
	Ms            string
	Page          string
	Pseudo        string
} {
	return struct {
		Statue        bool
		Word          string
		GuessedLetter string
		Essay         string
		Score         int
		Ms            string
		Page          string
		Pseudo        string
	}{
		Statue:        Statue,
		Word:          WordHide,
		GuessedLetter: allguess,
		Essay:         Essay,
		Score:         score,
		Ms:            Ms,
		Page:          Page,
		Pseudo:        pseudo,
	}
}
