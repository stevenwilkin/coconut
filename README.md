# Macaroon demo

A simple demonstration of a target system with access authenticated via [macaroons](http://hackingdistributed.com/2014/05/16/macaroons-are-better-than-cookies/).

Macaroons minted by the target system are caveated upon an external authentication system which must issue a discharge macaroon in order for the minted macaroon to be validated. Trust between these two systems is ensured via a shared secret which in this case is the public key of the authentication system.

## Acquire dependencies

	go get gopkg.in/macaroon.v1

##  Start services

	go run ./cmd/target/main.go
	go run ./cmd/auth/main.go

# Run demo

	./demo.sh

## Sample output

	> Verify auth - no macaroon
	not valid
	
	> Acquire macaroon
	acquired
	
	> Discharge macaroon - incorrect password
	not acquired
	
	> Discharge macaroon - correct password
	acquired
	
	> Verify auth
	valid
