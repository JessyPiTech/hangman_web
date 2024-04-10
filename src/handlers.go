package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// la c'est pour le index donc l'acceuill
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Page = "la"
	Ms = ""
	data := createData(Statue, "WordHide", "allguess", "Essay", score, Ms, Page, "pseudo")
	tmpl.Execute(w, data)
}

// connection
func CompteHandler(w http.ResponseWriter, r *http.Request) {
	Page = "Connection"
	//deja on remet tout a 0 dès la connection pour pouvoir changer de compte au besoin
	Ms = ""
	pseudo = ""
	password = ""
	Statue = false
	WordHide = ""
	allguess = ""
	Essay = ""
	score = 0
	attempts = 0
	//si la page envoie un post
	if r.Method == http.MethodPost {
		pseudo = strings.ToLower(r.FormValue("pseudo"))
		password = strings.ToLower(r.FormValue("password"))
		//on recup le pseudo et mots de pass
		contenu := password + "\n"
		filename := "players/" + pseudo + ".txt"
		//si le fichier si dessu n'existe pas on le cree avec comme nom le pseudo et avec comme comptenue notre mots de passe
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			creerFichierJoueur(pseudo, contenu)
		}
		//la on recupe tout nos donner du fichier text
		lastinfo, err := ReadWords2(filename) //la on recupe la tout le text du fichier wo.txt
		if err != nil {
			fmt.Println(err)
			return
		}
		// Vérifiez si le mot de passe correspond
		if strings.TrimSpace(lastinfo[0]) == password {
			//on verif si il a de la donnée enregistre dans le fichier .txt
			if lastinfo[1] != "" {
				//si oui on recupe les données
				score, err = strconv.Atoi(strings.TrimSpace(lastinfo[4]))
				if err != nil {
					fmt.Println("Erreur de conversion en entier:", err)
					Ms = "Erreur de conversion en entier:"
					score = 0
				}
			}
			//et on met statue de connection a true pour ne plus a avoir a se connecter jusqu'a la deconexion
			Statue = true
			//puis on redirige a l'acceuil
			http.Redirect(w, r, "/Ac", http.StatusSeeOther)
		} else {
			//sinon on garde le statue a false
			//et on affiche l'erreur sur la page avec ms
			Statue = false
			Ms = "Mot de passe incorrect."
		}
		//on cree notre data
		data := createData(Statue, "WordHide", "allguess", "Essay", score, Ms, Page, pseudo)
		//puis on execute notre page avec le contenue assoscier
		tmpl.Execute(w, data)
	}
}

// simplement la page d'acceil
func AcHandler(w http.ResponseWriter, r *http.Request) {
	Page = "Ac"

	data := createData(Statue, "WordHide", "allguess", "Essay", score, Ms, Page, "pseudo")
	tmpl.Execute(w, data)
}

// ici le code de tout le jeu
func HangmanHandler(w http.ResponseWriter, r *http.Request) {
	verif = true
	Page = "la"
	//ici on verifi que l'on a bien reçu une donner
	if r.Method == http.MethodPost || Test {
		hangman(w, r)
		if verif == true {
			Ms = "attantion: ta fais une betise ne reguess pas la meme lettre !!"
			data := createData(Statue, WordHide, allguess, Essay, score, Ms, Page, pseudo)
			tmpl.Execute(w, data)
			//j'aurais bien rajouter un shutdown /s si on reguessai la meme chose
			//mais vue que je suis casi sur que tu va tester je l'ai pas fais ;)
		}
	}
	//du coup si notre code a du faire un break sa va comme meme executer la page
	//et pas simplement faire une page blanche

}

// code pour tout relancer a 0 au niveau du jeu
func ReplayHandler(w http.ResponseWriter, r *http.Request) {
	Page = "la"
	Replay(w, r)
}

// win
func WinHandler(w http.ResponseWriter, r *http.Request) {
	Page = "HasWon"
	data := createData(Statue, "WordHide", "allguess", "Essay", score, "Ms", Page, "pseudo")
	tmpl.Execute(w, data)
}

// defeat
func DefeatHandler(w http.ResponseWriter, r *http.Request) {
	Page = "HasDefeat"
	data := createData(Statue, "WordHide", "allguess", "Essay", score, "Ms", Page, "pseudo")
	tmpl.Execute(w, data)
}

// -----------------different mode de jeu--------------------//
func EasyModeHandler(w http.ResponseWriter, r *http.Request) {
	Page = "la"
	//avec un word par mode different
	words, err := ReadWords("words/words.txt")
	if err != nil {
		fmt.Println("Erreur: Need a 'words.txt' file.")
		return
	}
	fmt.Println("you Choose EasyMode")
	//on choisi le mots dans dans les mots recupe
	ChooseWord(words)
	if r.Method == http.MethodGet {
		//15 essay
		attempts = 15
		//et on affiche deux lettre aleatoir
		letterRandom := rand.Intn(len(word))
		guessedLetters[letterRandom] = true
		letterRandom2 := rand.Intn(len(word))
		guessedLetters[letterRandom2] = true
		//puis on chache le mots avec nos deux lettre afficher
		WordHide = ReWriting(word, guessedLetters)
		//conversion des essay
		Essay := strconv.Itoa(attempts)
		//...
		data := createData(Statue, WordHide, allguess, Essay, score, Ms, Page, pseudo)
		tmpl.Execute(w, data)
	}
}

// pareill mais en medieum
func MediumModeHandler(w http.ResponseWriter, r *http.Request) {
	Page = "la"
	words, err := ReadWords("words/words2.txt")
	if err != nil {
		fmt.Println("Erreur: Need a 'words2.txt' file.")
		return
	}
	fmt.Println("you Choose MediumMode")
	ChooseWord(words)
	if r.Method == http.MethodGet {
		attempts = 10
		letterRandom := rand.Intn(len(word))
		guessedLetters[letterRandom] = true
		WordHide = ReWriting(word, guessedLetters)
		Essay := strconv.Itoa(attempts)
		data := createData(Statue, WordHide, allguess, Essay, score, Ms, Page, pseudo)
		tmpl.Execute(w, data)
	}
}

// et encore en hard
func HardModeHandler(w http.ResponseWriter, r *http.Request) {
	Page = "la"
	words, err := ReadWords("words/words3.txt")
	if err != nil {
		fmt.Println("Erreur: Need a 'words3.txt' file.")
		return
	}
	fmt.Println("you Choose HardMode")
	ChooseWord(words)
	if r.Method == http.MethodGet {
		//ici pas de lettre aleatoir
		attempts = 5
		WordHide = ReWriting(word, guessedLetters)
		Essay := strconv.Itoa(attempts)
		data := createData(Statue, WordHide, allguess, Essay, score, Ms, Page, pseudo)
		tmpl.Execute(w, data)
	}
}

// ----------------------------------------------------------//
func DevHandler(w http.ResponseWriter, r *http.Request) {
	//page dev pour qui a cree la page
	Page = "Dev"
	data := createData(Statue, "WordHide", "allguess", "Essay", score, "Ms", Page, pseudo)
	tmpl.Execute(w, data)
}

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	//le profil jeuour
	Page = "Profil"
	data := createData(Statue, "WordHide", "allguess", "Essay", score, "Ms", Page, pseudo)
	tmpl.Execute(w, data)
}

// quand on leave
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//enregistrement de donnée

	filename := "players/" + pseudo + ".txt"
	// Ouvrez le fichier en mode lecture
	lastinfo, err := ReadWords2(filename) //la on recupe la tout le text du fichier wo.txt
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}

	if lastinfo[1] == "" && WordHide == "" {
		Ms = "you didn't have play"
	} else if WordHide == "" {
		recup()
		mise()
	} else {
		mise()
	}

	Statue = false
	//puis on repart a la page d'acceuil
	http.Redirect(w, r, "/", http.StatusSeeOther)
	b = true
	data := createData(Statue, "WordHide", "allguess", "Essay", score, "Ms", Page, pseudo)
	tmpl.Execute(w, data)
}

// quand on enregistre
func MisHandler(w http.ResponseWriter, r *http.Request) {
	mise()
	http.Redirect(w, r, "/Ac", http.StatusSeeOther)
	data := createData(Statue, "WordHide", "allguess", "Essay", score, "Ms", Page, pseudo)
	tmpl.Execute(w, data)
}

func RegoHandler(w http.ResponseWriter, r *http.Request) {
	//ici c'est quand on relance une partie en cours
	Page = "la"
	filename := "players/" + pseudo + ".txt"
	// Ouvrez le fichier en mode lecture
	lastinfo, err := ReadWords2(filename) //la on recupe la tout le text du fichier wo.txt
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}

	if lastinfo[1] == "" {
		//si y a pas de donnée c'est qu'il n'y pas de sauvegard
		// pour explication de se que je met la
		//le probleme est quand je fais save and leave et que je cree un autre compte sa va garder les info du premier compte avec lequel on c'est connecter
		//donc faut tout remetre a 0
		Ms = "Il n'y a pas de sauvegarde"
		fmt.Println("Il n'y a pas de sauvegarde")
		Statue = true
		WordHide = ""
		allguess = ""
		Essay = ""
		score = 0
		attempts = 0
	} else {
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
	data := createData(Statue, WordHide, allguess, Essay, score, Ms, Page, pseudo)
	tmpl.Execute(w, data)
	b = true
}
