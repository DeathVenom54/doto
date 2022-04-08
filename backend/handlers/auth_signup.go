package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DeathVenom54/doto-backend/db"
	"github.com/DeathVenom54/doto-backend/token"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	// check if already logged in
	if alreadyLoggedIn := token.VerifyToken(r.Header.Get("user")); alreadyLoggedIn {
		w.WriteHeader(200)
		_, err := w.Write([]byte("success"))
		if err != nil {
			log.Printf("Error while writing response at /auth/signup\n%s", err)
		}
		return
	}

	var userData db.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	err = validateUserSignup(&userData)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	// create user in database
	var createdUser *db.User
	createdUser, err = db.CreateUser(&userData)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	// write jwt token
	claims := &token.AuthClaims{
		ID:       createdUser.ID,
		Username: createdUser.Username,
	}
	jwtToken, err := token.CreateToken(claims)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	cookie := http.Cookie{
		Name:     "user",
		Value:    jwtToken,
		HttpOnly: true,
	}
	r.AddCookie(&cookie)

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Printf("Error while writing response at /auth/signup\n%s", err)
	}
}

func validateUserSignup(user *db.User) error {
	// required fields
	var missingFields []string

	if user.Username == "" {
		missingFields = append(missingFields, "username")
	}
	if user.Email == "" {
		missingFields = append(missingFields, "email")
	}
	if user.Password == "" {
		missingFields = append(missingFields, "password")
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing required field(s) %s", strings.Join(missingFields, ", "))
	}

	// email format
	if match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$", user.Email); !match {
		return fmt.Errorf("invalid email %s", user.Email)
	}
	// username format
	if match, _ := regexp.MatchString("^[a-z0-9_-]{1,30}$", user.Username); !match {
		return fmt.Errorf("invalid username %s", user.Username)
	}
	return nil
}
