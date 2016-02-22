package main

import (
	"fmt"
	"net/http"

	"gopkg.in/macaroon.v1"

	"github.com/stevenwilkin/coconut/biscuit"
	"github.com/stevenwilkin/coconut/crypt"
)

const (
	location = "http://auth/"
)

var auth = map[string]string{
	"steve": "abc",
}

func authenticate(user, pass string) bool {
	p, ok := auth[user]
	if !ok {
		return false
	}

	return  p == pass
}

func dischargeHandler(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")
	if !authenticate(u, r.FormValue("p")) {
		panic("not authenticated")
	}

	m := biscuit.MustDecodeFromString(r.FormValue("m"))

	var caveat macaroon.Caveat
	for _, c := range m.Caveats() {
		if c.Location == location {
			caveat = c
			break
		}
	}

	randomKeyFromTarget := crypt.Decrypt([]byte(caveat.Id))	// uses SSO priv key

	d := biscuit.MustNew(randomKeyFromTarget, caveat.Id, location)
	biscuit.MustAddFirstPartyCaveat(d, fmt.Sprintf("username = %s", u))
	s := biscuit.MustEncodeToString(d)

	fmt.Fprintf(w, "%s\n", s)
}

func main() {
	http.HandleFunc("/discharge", dischargeHandler)

	println("Starting target on port 8181")

	if err := http.ListenAndServe(":8181", nil); err != nil {
		panic(err)
	}
}
