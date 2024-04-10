package main

//ici on est vraiment dans la partie fonctionnement du jeu

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

// relance du jeu
func Replay(w http.ResponseWriter, r *http.Request) {
	//on meyt tout a 0
	for i := range guessedLetters {
		guessedLetters[i] = false
	}
	Ms = ""
	Page = "la"
	WordHide = ""
	guess = "0"
	attempts = 10
	allguess = ""
	guessedLetters = make([]bool, len(word))
	//pusi on redirige a l'accuille
	http.Redirect(w, r, "/Ac", http.StatusSeeOther)
}

// choisi un mot aleatoir et enleve le /n si il en a
func ChooseWord(words []string) {
	word = ChooseRandomWord(words)
	if strings.HasSuffix(word, "\n") {
		word = strings.TrimRight(word, "\n")
	}
	word = strings.TrimSpace(word)
	fmt.Println("word to guess :", word)
	attempts = 10
	allguess = ""
	guessedLetters = make([]bool, len(word))
}

// verif si la lettre est dans le mots
func ContainsGuess(allguess string, guess string) bool {
	// separez "allguess" en mots en utilisant le tiret comme séparateur
	mots := strings.Split(allguess, "-")
	// parcoure les mots pour trouver une correspondance
	for _, mot := range mots {
		if mot == guess {
			return true
		}
	}
	return false
}

// recris le mots en chacahnt les lettre non trouver
func ReWriting(word string, guessedLetters []bool) string {
	WordHide := ""
	for i, char := range word {
		if guessedLetters[i] {
			WordHide += string(char)
		} else {
			WordHide += "_ "
		}
	}
	return WordHide
}

// choisi un mot aleatoir
func ChooseRandomWord(words []string) string {
	if len(words) == 1 {
		fmt.Println("Erreur: le fichier words ne contient pas de mots.")
	}
	randIndex := rand.Intn(len(words))
	word := words[randIndex]
	return word
}

// split notre une data en un nombre non defini
func ReadWords(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(data), "\n")
	return words, nil
}

// split en nombre defini
func ReadWords2(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename) //la on recupe la tout le text du fichier wo.txt
	if err != nil {
		return nil, err
	}
	infos := strings.SplitN(string(data), "\n", 7) //ici on fais en sorte que notre data devienne un string avec string(data) puis on le split a chaque . en 8 separation
	//puis on renvoie notre tableau
	return infos, nil
}

// verif si sa a deja ete guess
func AllLettersGuessed(guessedLetters []bool) bool {
	for _, guessed := range guessedLetters {
		if !guessed {
			return false
		}
	}
	return true
}

// met un score quand on gagne
func Scoring(word string, score *int) {
	sc := 100 * len(word)
	sco := *score
	*score = sco + sc
	fmt.Println("Score:", *score)
}

// verif si c'est bien une lettre
func IsAlpha(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func hangman(w http.ResponseWriter, r *http.Request) {
	for attempts >= 1 {
		//ici on vien recuper se que renvoie la page web par un post
		r.ParseForm()
		guess = strings.ToLower(r.FormValue("letter"))
		//
		//verificaation que se que on a gess n'a pas deja été guess car sinon sa fais un guess en boucle genre a a a a a a a a a a a a
		// mais sa fesait un petit beug qui t'enpechait de guess deux fois la meme lettre a la suite san
		//et fesait crash le jeu mais je l'es regler avec les derniere ligne du HangmanHandler
		if lastGuess == guess || guess == "" || guess == "\n" {
			break
		} else {
			verif = false
		}

		Ms = ""
		//on met a jour notre lastGuess
		lastGuess = guess
		//ensuit on fais differente verification et on revoie si il a un probleme Ms qui decris notre probleme
		//et parfois si on veut skipe une etape de la verification on met une valeur tres peu probable et specifique au guess
		if !IsAlpha(guess) || guess == " " {
			Ms = " seules les lettres alphabétiques sont autorisées."
			guess = "rerere"
		}
		guess = strings.ToLower(strings.TrimSpace(guess))
		//longeur
		if len(guess) != 1 && guess != word && guess != "rerere" {
			Ms = "essaye lettre par lettre vous aurez moins de chance de vous trompez"
			attempts--
			attempts--
			guess = "rerere"
			if attempts <= 0 {
				//pour evite le -1
				attempts = 0
				Page = "HasDefeat"
				http.Redirect(w, r, "/defeat", http.StatusSeeOther)
			}
		}
		//deja guess
		if ContainsGuess(allguess, guess) {
			Ms = "lettre already guess"
			guess = "rerere"
		}
		//verif si le mots que tu a gess  est bon
		if guess == word {
			Scoring(word, &score)
			Page = "HasWon"
			http.Redirect(w, r, "/win", http.StatusSeeOther)
		}
		//c'est pas un caratere special

		//puis on fais la verif de si c'est dans le mots
		if guess != "rerere" {
			found := false
			for i, char := range word {
				if guessedLetters[i] {
					continue
				}
				if guess == string(char) {
					guessedLetters[i] = true
					found = true
				}
			}
			//on cree un un text genre ---h-l en comparent un tableau de vrai faux a une liste de caratere genre la sa ferais F F F V F V
			WordHide = ReWriting(word, guessedLetters)
			//verif si win, sinon verif si la lettre est fausse
			if AllLettersGuessed(guessedLetters) || word == WordHide {
				Scoring(word, &score)
				Page = "HasWon"
				http.Redirect(w, r, "/win", http.StatusSeeOther)
			} else if !found && Page != "HasWon" {
				attempts--
				if attempts <= 0 {
					Page = "HasDefeat"
					http.Redirect(w, r, "/defeat", http.StatusSeeOther)
				}
			}
			//on cree une liste de proposition deja guess
			allguess = allguess + guess
			allguess = allguess + "-"
		}
		//on convertie un int en string
		Essay := strconv.Itoa(attempts)
		data := createData(Statue, WordHide, allguess, Essay, score, Ms, Page, pseudo)
		tmpl.Execute(w, data)
	}
}
