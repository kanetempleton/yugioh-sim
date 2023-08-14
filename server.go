package main


import (
    "database/sql"
    "net/http"
	"fmt"
	"log"
	"strings"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

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

func startServer() {

	err := Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo")
    if err != nil {
        log.Fatal("Error initializing database:", err)
    } else {
		fmt.Println("Database initialized successfully!");
	}
    r := mux.NewRouter()

	staticDir := "/static/"
    fs := http.FileServer(http.Dir("static"))
    r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, setContentTypeMiddleware("application/javascript", fs)))


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