// Using cookies, track how many times a user has been to your website domain.

package main

import (
	"log"
	"io"
	"net/http"
	"strconv"
)

const keyCounter = "counter"

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(keyCounter)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: keyCounter,
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}
