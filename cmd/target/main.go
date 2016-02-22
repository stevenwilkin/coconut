package main

import (
	"fmt"
	"net/http"

	"github.com/stevenwilkin/coconut/biscuit"

	"gopkg.in/macaroon.v1"
)

const (
	key = "target-service-key"
	location = "http://target/"
)

func issueHandler(w http.ResponseWriter, r *http.Request) {
	identity := r.FormValue("u")
	m := biscuit.MustNew([]byte(key), identity, location)

	biscuit.MustAddFirstPartyCaveat(m, "can_install_snaps")
	biscuit.MustAddFirstPartyCaveat(m, fmt.Sprintf("username = %s", identity))
	biscuit.MustAddThirdPartyCaveat(m, "http://auth/")

	s := biscuit.MustEncodeToString(m)

	fmt.Fprintf(w, "%s\n", s)
}

func newCaveatChecker(identity string) func(caveat string) error {
	return func(caveat string) error {
		fmt.Printf("%v\n", caveat)
		if caveat == "can_install_snaps" {
			return nil
		}
		if caveat == fmt.Sprintf("username = %s", identity) {
			return nil
		}
		return fmt.Errorf("condition %q not met", caveat)
	}
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	m := biscuit.MustDecodeFromString(r.FormValue("m"))
	d := biscuit.MustDecodeFromString(r.FormValue("d"))
	identity := r.FormValue("u")

	d.Bind(m.Signature())

	err := m.Verify([]byte(key), newCaveatChecker(identity), []*macaroon.Macaroon{d})
	if err != nil {
		fmt.Fprint(w, "Invalid!\n")
	} else {
		fmt.Fprint(w, "Valid\n")
	}
}

func main() {
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/verify", verifyHandler)

	println("Starting target on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
