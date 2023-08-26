package main


import (
    "database/sql"
    "net/http"
	"fmt"
	"log"
	"strings"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
	"mime/multipart"
	"os"
	"github.com/go-xorm/xorm"
	"io"
	"math/rand"
    "time"
	"encoding/json"
	"path/filepath"
)

var db *sql.DB
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
    rand.Seed(time.Now().UnixNano())
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return string(result)
}

func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func isDuplicateKeyError(err error) bool {
    if err != nil {
        if strings.Contains(err.Error(), "Duplicate entry") {
            return true
        }
    }
    return false
}

func setContentTypeMiddleware(contentType string, h http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", contentType)
        h.ServeHTTP(w, r)
    }
}

func saveImageToFile(file multipart.File, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil
}


func startServer() {

	err := Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo")
    if err != nil {
        log.Fatal("Error initializing database:", err)
    } else {
		fmt.Println("Database initialized successfully!");
	}

	err = InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo")
    if err != nil {
        log.Fatal("Error initializing database:", err)
    } else {
		fmt.Println("Database initialized successfully!");
	}
    r := mux.NewRouter()

	staticDir := "/static/"
    fs := http.FileServer(http.Dir("static"))
    r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs))

	r.HandleFunc("/static/site_style.css", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Setting content type to text/css")
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "static/site_style.css")
	})

	stylesfs := http.FileServer(http.Dir("styles"))
    //http.Handle("/card-images/", http.StripPrefix("/card-images/", cardImagesFS))
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", setContentTypeMiddleware("text/css", stylesfs)))


	// Set up file server for card images
    cardImagesFS := http.FileServer(http.Dir("card-images"))
    //http.Handle("/card-images/", http.StripPrefix("/card-images/", cardImagesFS))
	r.PathPrefix("/card-images/").Handler(http.StripPrefix("/card-images/", setContentTypeMiddleware("image/jpeg", cardImagesFS)))

	//http.Handle("/card-images/", http.StripPrefix("/card-images/", http.FileServer(http.Dir("card-images"))))


	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "templates/index.html")
    })

   r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
		err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form", http.StatusInternalServerError)
            return
        }

        username := r.FormValue("username")
        password := r.FormValue("password")

        // Validate username and password
        user, err := GetUserByUsername(username)
        if err != nil {
            http.Error(w, "Error querying database", http.StatusInternalServerError)
            return
        }
        if user == nil || user.Password != password {
			fmt.Println("Invalid username/password")
			http.Redirect(w, r, "/login?message2=Invalid+username+or+password", http.StatusSeeOther)
            return
        }

        // Set session cookie (example, needs improvement)
		http.SetCookie(w, createSessionCookie(username))

        // Redirect to account page
        http.Redirect(w, r, "/account", http.StatusSeeOther)
		fmt.Println("Login Success")
    } else {
		fmt.Println("HTTP GET /login")
		fmt.Println("Rendering template: templates/login.html")
		http.ServeFile(w, r, "templates/login.html")

    }
})


    r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("HTTP POST /register")
			// Parse form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusInternalServerError)
				return
			}
	
			// Extract form values
			username := r.FormValue("username")
			password := r.FormValue("password")
			email := r.FormValue("email")
			// Other form values
	
			// Create a new user object
			newUser := &User{
				Username: username,
				Password: password,
				Email:    email,
				// Set other fields
			}
	
			// Insert the user into the database
			err = CreateUser(newUser)
			if err != nil {
				if isDuplicateKeyError(err) {
					http.Error(w, "Username already exists. Please go back and try again.", http.StatusBadRequest)
				} else {
					http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
				}
				return
			}
	
			// Redirect to a success page or login page
			http.Redirect(w, r, "/login?message=User+created+successfully%21", http.StatusSeeOther)
		} else {
			fmt.Println("HTTP GET /register")
			// Render the registration form template
			fmt.Println("Rendering template: templates/register.html")
			http.ServeFile(w, r, "templates/register.html")

		}
    })

    r.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
        // Implement your user account page handler here
		handleUser(w,r)
    })

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		// Expire the session cookie by setting MaxAge to a negative value
		handleLogout(w,r)
	})

	r.HandleFunc("/view-deck", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Render the "view-deck.html" template
			http.ServeFile(w, r, "templates/view-deck.html")
		}
	})

	r.HandleFunc("/duel/solo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Render the "view-deck.html" template
			http.ServeFile(w, r, "templates/duel_solo.html")
		}
	})

	r.HandleFunc("/get-user-cards", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /get-user-cards")
		sessionCookie, err := r.Cookie("session")
		if err != nil || sessionCookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther) // Redirect to login page if not authenticated
			return
		}
	
		userID := sessionCookie.Value
		deck := r.URL.Query().Get("deck")
	
		//session := CardEngine.NewSession()
	
		cardIdentifiers, err := GetCardsByUserID(userID,deck)
		if err != nil {
			http.Error(w, "Error fetching user's cards", http.StatusInternalServerError)
			return
		}
	
		jsonResponse, err := json.Marshal(cardIdentifiers)
		if err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})


	r.HandleFunc("/get-user-decks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /get-user-decks")
		
		sessionCookie, err := r.Cookie("session")
		if err != nil || sessionCookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther) // Redirect to login page if not authenticated
			return
		}
		
		userID := sessionCookie.Value
		
		userDecks, err := GetDecksByUserID(userID)
		if err != nil {
			http.Error(w, "Error fetching user's decks", http.StatusInternalServerError)
			return
		}
		
		jsonResponse, err := json.Marshal(userDecks)
		if err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	r.HandleFunc("/cards/set-deck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Testing set-deck")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	
		// Parse the request body
		var requestData struct {
			FileName  string `json:"file_name"`
			DeckName  string `json:"deck"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request data", http.StatusBadRequest)
			return
		}
	
		// Update the card's deck name in the database
		fmt.Println("FileName="+requestData.FileName+" DeckName="+requestData.DeckName+"")
		card := new(Card)
		has, err := CardEngine.Where("file_name = ?", requestData.FileName).Get(card)
		if err != nil || !has {
			fmt.Println("Card not found",err)
			http.Error(w, "Card not found", http.StatusBadRequest)
			return
		}
	
		if err := card.SetDeck(requestData.DeckName); err != nil {
			fmt.Println("Error setting deck",err)
			http.Error(w, "Error setting card deck", http.StatusInternalServerError)
			return
		}

		fmt.Println("Responding with success")
	
		// Respond with success
		response := struct {
			Success bool `json:"success"`
		}{
			Success: true,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error: ", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})
	
	
	
	

	r.HandleFunc("/add-card", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusInternalServerError)
				return
			}
	
			sessionCookie, err := r.Cookie("session")
			if err != nil || sessionCookie.Value == "" {
				http.Redirect(w, r, "/login", http.StatusSeeOther) // Redirect to login page if not authenticated
				return
			}
	
			userID := sessionCookie.Value
			userID = strings.ToLower(userID)
			name := r.FormValue("name")
			deck := r.FormValue("deck") 
			if deck == "" {
				deck = "None"
			}
			file, fileHeader, err := r.FormFile("image")
			if err != nil {
				http.Error(w, "Error retrieving image file", http.StatusInternalServerError)
				return
			}
			defer file.Close()
	
			
			baseFilename := fmt.Sprintf("%s_%s", userID, strings.ReplaceAll(fileHeader.Filename, " ", "_"))
        	fileName := baseFilename
			//fileName = generateRandomString(10)

			// Check if the filename already exists
			i := 1
			for fileExists(fileName) {
				fileName = fmt.Sprintf("%s-%d%s", baseFilename, i, filepath.Ext(fileHeader.Filename))
				i++
			}

	
			// Save the uploaded image to a directory (e.g., "card-images")
			imagePath := "card-images/" + fileName
			err = saveImageToFile(file, imagePath)
			if err != nil {
				http.Error(w, "Error saving image", http.StatusInternalServerError)
				return
			}
	
			// Get or create a database session using XORM
			session, err := xorm.NewEngine("mysql", "admin:obviouspassword@tcp(localhost:3306)/yugiohgo")
			if err != nil {
				http.Error(w, "Error creating database session", http.StatusInternalServerError)
				return
			}
			defer session.Close()
	
			// Create a new Card entry in the database
			NewCard(userID, fileName, name, deck)
			if err != nil {
				http.Error(w, "Error creating card entry", http.StatusInternalServerError)
				return
			}
	
			// Redirect to a success page or user account page
			http.Redirect(w, r, "/account?message=Card+added+successfully%21", http.StatusSeeOther)
		} else {
			// Render the "add-card.html" template
			http.ServeFile(w, r, "templates/add-card.html")
		}
	})

	r.HandleFunc("/cards/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	
		var requestData struct {
			FileName string `json:"file_name"`
		}
	
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
	
		err = DeleteCardByFileName(requestData.FileName)
		if err != nil {
			http.Error(w, "Error deleting card", http.StatusInternalServerError)
			return
		}
	
		// Return a JSON response indicating success
		response := struct {
			Success bool `json:"success"`
		}{
			Success: true,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	
	
	

    http.Handle("/", r)

	fmt.Println("Running server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}



// Example route handler for "/" or "/login" or others
func handleHome(w http.ResponseWriter, r *http.Request) {
    // Your code to handle the route here
}

// Example route handler for "/user" or others
func handleUser(w http.ResponseWriter, r *http.Request) {
    // Your code to handle the route here
	fmt.Println("HTTP GET /account")
		 // Get the session cookie
		 sessionCookie, err := r.Cookie("session")
		 if err != nil || sessionCookie.Value == "" {
			 // Redirect to login page if no session cookie
			 http.Redirect(w, r, "/login", http.StatusSeeOther)
			 return
		 }
		 http.ServeFile(w, r, "templates/account.html")
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
    // Expire the session cookie by setting MaxAge to a negative value
    http.SetCookie(w, createSessionCookie(""))
    
    // Redirect to the login page
    http.Redirect(w, r, "/login?message=Logout+success%21", http.StatusSeeOther)
}


func createSessionCookie(username string) *http.Cookie {
    // Create and return a session cookie
    return &http.Cookie{
        Name:   "session",
        Value:  username,
        MaxAge: 3600,
    }
}