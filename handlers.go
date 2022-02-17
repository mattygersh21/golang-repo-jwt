package main

import {
	"net/http"

	"github.com/dgrijalva/jwt-go"
}

var jwtKey = []byte("secret_key")
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password`
}

type Claims struct {
	Username string `jason:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Home(w http.ResponseWriter, r *http.Request) {

}

func Refresh(w http.ResponseWriter, r *http.Request) {

}
