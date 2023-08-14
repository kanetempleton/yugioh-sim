package main

import (
    "database/sql"
    //"html/template"
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

func main() {

	err := Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo")
    if err != nil {
        log.Fatal("Error initializing database:", err)
    } else {
		fmt.Println("Database initialized successfully!");
	}


    //templates := template.Must(template.ParseFiles("templates/login.html", "templates/register.html"))

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
        http.SetCookie(w, &http.Cookie{
            Name:   "session",
            Value:  username, // Store username in the cookie for session management
            MaxAge: 3600,     // Cookie will expire after 1 hour
        })

        // Redirect to account page
        http.Redirect(w, r, "/account", http.StatusSeeOther)
		fmt.Println("Login Success")
    } else {
        // Render the login form template
        // Implement template rendering logic here
		// Load the login template
		fmt.Println("HTTP GET /login")
       // tmpl, _ := templates.ParseFiles("templates/login.html")
		fmt.Println("Rendering template: templates/login.html")
		//tmpl.Execute(w, nil)
		http.ServeFile(w, r, "templates/login.html")
       /* if err != nil {
            http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			fmt.Println("Error: "+err.Error())
            return
        }

        // Execute the template with no data (nil)
        err = tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
			fmt.Println("Error: "+err.Error())
            return
        }*/
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
			//err = tmpl.Execute(w, nil)
			http.ServeFile(w, r, "templates/register.html")
			
		/*	if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}*/
		}
    })

    r.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
        // Implement your user account page handler here
		fmt.Println("HTTP GET /account")
		 // Get the session cookie
		 sessionCookie, err := r.Cookie("session")
		 if err != nil || sessionCookie.Value == "" {
			 // Redirect to login page if no session cookie
			 http.Redirect(w, r, "/login", http.StatusSeeOther)
			 return
		 }
		 http.ServeFile(w, r, "templates/account.html")
    })

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		// Expire the session cookie by setting MaxAge to a negative value
		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		})
	
		// Redirect to the login page
		http.Redirect(w, r, "/login?message=Logout+success%21", http.StatusSeeOther)
	})

    http.Handle("/", r)

	fmt.Println("Running server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
