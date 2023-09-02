# Yu-Gi-Oh Simulator

Stack: Go + MySQL + JS



## Features

- User Registration/Login: Users can create an account by providing a unique username, password, and email. Passwords are encrypted via one-way hashing for storage.
- Card Management: Upload image files of custom cards (cardmaker.net is one place you can make them) and organize your cards into different decks
- Single-User Duel: Simulate a Yu-Gi-Oh Duel using your custom cards as your deck. Select a deck and extra deck and begin the duel with 8000 Life Points and a blank Yu-Gi-Oh field. Perform card operations such as drawing cards from the deck, playing cards onto the field, flipping and changing card positions, adding cards to graveyard, etc. (You can simulate a duel against another player by opening another window as a second duel interface)

## Getting Started

1. Install dependencies: Ensure you have Go and MySQL installed.
2. Set up MySQL: See the steps below to make sure your MySQL environment is set up properly.
3. Clone the repository: `git clone https://github.com/kanetempleton/yugioh-sim.git`
4. Navigate to the project directory: `cd yugioh-sim`
5. Get dependencies: `go mod tidy` (If you get a version error, edit the go.mod file to change required go version to the one you have installed)
6. Compile the Go application: `go build`
7. Run tests: `go test`
8. Run the server: `./yugiohgo`
9. Access the app: Open your web browser and visit `http://localhost:8080` to interact with the app.

## Usage

1. Register a new account by clicking the "Register" link on the homepage and filling in the required information.
2. Log in with your registered username and password.
3. Once logged in, you'll be redirected to your account page where you can manage your deck and cards (coming soon).


## Setting Up MySQL Database

To run the YugiohGo web application, you need to set up a MySQL database. Follow these steps:

1. **Install MySQL**: If you don't have MySQL installed, download and install it from the official website or a package manager. On mac, I recommend the homebrew command `brew install mysql`

2. **Create the Admin and Database**: Log in to the MySQL command-line interface using a user with administrative privileges (NOTE: it is highly recommended you change these credentials):

   ```bash
   mysql -u root -p
   ```

   Then create a user to manage the application, and create a database for the application to use.

   ```MySQL
   CREATE DATABASE yugiohgo;
   CREATE DATABASE yugiohgo_test;
   CREATE USER 'admin'@'localhost' IDENTIFIED BY 'obviouspassword';
   GRANT ALL PRIVILEGES ON yugiohgo.* TO 'admin'@'localhost';
   GRANT ALL PRIVILEGES ON yugiohgo_test.* TO 'admin'@'localhost';
    FLUSH PRIVILEGES;
    
    ```

    You do not need to create any tables. The application uses XORM to manage this automatically.

3. **Modify Connection Details**: If you changed any of the database creation values, swap them out in `main.go` and `user.go`


## TODO
### Dueling:
- have more than 6 cards in your hand
- add cards back into deck & extra deck
- prevent graveyard slot from turning horizontal when adding DEF position cards
- player vs player dueling
### Cards:
- upload multiple of an individual card at once
- public decks
### Users:
- changing passwords
- email verification for registration


## Note

- This project is for educational purposes and demonstrates basic web application concepts.
- The app is a simplified example and may lack certain security and error handling features for production use.
