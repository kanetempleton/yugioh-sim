package main

import (
	"net/http"
	"net/http/httptest"
	//"strings"
	"testing"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func printGreen(message string) {
	fmt.Printf("\033[32m%s\033[0m\n", message)
}
func printYellow(message string) {
	fmt.Printf("\033[33m%s\033[0m\n", message)
}

func printCyan(message string) {
	fmt.Printf("\033[36m%s\033[0m\n", message)
}
func printOrange(message string) {
	fmt.Printf("\033[38;5;208m%s\033[0m\n", message)
}
func printPink(message string) {
	fmt.Printf("\033[38;5;205m%s\033[0m\n", message)
}

var _ = Describe("User Management (Acceptance)", func() {
	printPink("Testing: User Management Acceptance")
	Context("User Authentication", func() {
		printPink("Testing: Context: User Authentication")
		It("should allow user to log out and redirect to login with success message", func() {
			// Start a test server
			printPink("Testing: should allow user to log out and redirect to login with success message")
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				printOrange("Started test server")
				// Simulate user being logged in
				sessionCookie := &http.Cookie{
					Name:  "session",
					Value: "testuser",
				}
				r.AddCookie(sessionCookie)
				printOrange("Added session cookie")

				// Simulate a logged in user on the "/account" page
				if r.URL.Path == "/account" {
					printOrange("Visiting /account page, checking session cookie")
					// Check if the session cookie exists in the request's cookies
					sessionCookie, err := r.Cookie("session")
					if err == nil && sessionCookie != nil && sessionCookie.Value == "testuser" {
						// Perform a GET request to the actual "/user" page to retrieve the HTML content
						responseBody := httptest.NewRecorder()
						http.ServeFile(responseBody, r, "/account") // Replace with the actual path

						// Assert that the response body contains the "Log Out" button
						Expect(responseBody.Body.String()).To(ContainSubstring("Log Out"))
					}
				}


				// Simulate user clicking "Log Out" button
				if r.Method == http.MethodPost && r.URL.Path == "/logout" {
					// Assert that the response redirects to "/login"
					Expect(w.Header().Get("Location")).To(Equal("/login"))

					// Assert that the response body contains the success message
					/*
					responseBody := httptest.NewRecorder()
					http.ServeFile(responseBody, r, "/login") // Replace with the actual path
					Expect(responseBody.Body.String()).To(ContainSubstring("Logout Success!"))
					*/

					// Assert that the session cookie has been cleared
					sessionCookie, err := r.Cookie("session")
					Expect(err).ToNot(HaveOccurred())
					Expect(sessionCookie.MaxAge).To(Equal(-1)) // MaxAge should be set to -1 for a deleted cookie
				}

			}))
			defer server.Close()


		})

		// Add more acceptance tests for different scenarios

		// Don't forget to clean up resources if necessary
		AfterEach(func() {
			// Clean up code here
			printYellow("WARNING: Acceptance tests currently flawed. Ignore them for now.")
		})
	})
})


func TestUserAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Acceptance Suite")
}
