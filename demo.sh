#!/bin/bash

TARGET_BASE_URL='0.0.0.0:8080'
AUTH_URL='0.0.0.0:8181/discharge'
CURL='curl -s'
U='steve'
P=
M=
D=

function verify_auth {
	V=$($CURL "$TARGET_BASE_URL/verify?u=$U&m=$M&d=$D")

	if [ "$V" == "Valid" ]
	then
		echo "valid"
	else
		echo "not valid"
	fi

	echo
}

function acquire_macaroon {
	$CURL $TARGET_BASE_URL/issue?u=$U
}

function check_macaroon {
	if [ -s $M ]
	then
		echo "not acquired"
	else
		echo "acquired"
	fi

	echo
}

function discharge_macaroon {
	$CURL "$AUTH_URL?u=$U&p=$P&m=$M"
}

function check_discharge {
	if [ -s $D ]
	then
		echo "not acquired"
	else
		echo "acquired"
	fi

	echo
}

echo "> Verify auth - no macaroon"
verify_auth

echo "> Acquire macaroon"
M=$(acquire_macaroon)
check_macaroon

echo "> Discharge macaroon - incorrect password"
D=$(discharge_macaroon)
check_discharge

echo "> Discharge macaroon - correct password"
P='abc'
D=$(discharge_macaroon)
check_discharge

echo "> Verify auth"
verify_auth
