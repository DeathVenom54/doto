package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DeathVenom54/doto-backend/db"
	"github.com/DeathVenom54/doto-backend/token"
	"log"
	"net/http"
	"regexp"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// check if already logged in
	if alreadyLoggedIn := token.VerifyToken(r.Header.Get("user")); alreadyLoggedIn {
		w.WriteHeader(200)
		_, err := w.Write([]byte("success"))
		if err != nil {
			log.Printf("Error while writing response at /auth/signup\n%s", err)
		}
		return
	}

	// get user data from body
	var userData db.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	err = validateUserLogin(&userData)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	correct, user := db.VerifyUserPassword(&userData, userData.Password)
	if !correct {
		handleError(fmt.Errorf("invalid username or password"), w, http.StatusBadRequest)
		return
	}

	claims := &token.AuthClaims{
		ID:       user.ID,
		Username: user.Username,
	}
	cookie, err := token.CreateJwtHttpOnlyCookie(claims)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	http.SetCookie(w, cookie)

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Printf("Error while writing response at /auth/signup\n%s", err)
	}
}

func validateUserLogin(user *db.User) error {
	// required fields
	if user.Username == "" && user.Email == "" {
		return fmt.Errorf("any one of email and username must be provided")
	}
	if user.Password == "" {
		return fmt.Errorf("missing required field(s) password")
	}

	// email format
	if match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$", user.Email); !match && user.Email != "" {
		return fmt.Errorf("invalid email %s", user.Email)
	}
	// username format
	if match, _ := regexp.MatchString("^[a-z0-9_-]{1,30}$", user.Username); !match && user.Username != "" {
		return fmt.Errorf("invalid username %s", user.Username)
	}
	return nil
}
