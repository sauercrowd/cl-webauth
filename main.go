package main

/*
#cgo CFLAGS: -I.
#include "webauthn.h"
*/
import "C"

import (
	"encoding/json"
	// "errors"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/go-webauthn/webauthn/protocol"
)

type User struct {
	ID string
	Credentials []webauthn.Credential
	Name string
}


//export ValidateLogin
func ValidateLogin (cId, cName, cCredentialJSONs, cSessionDataJson *C.char) C.int {
	var car protocol.CredentialAssertionResponse

	var credentials []webauthn.Credential

	id := C.GoString(cId)
	name := C.GoString(cName)
	sessionDataJson := C.GoString(cSessionDataJson)

	err := json.Unmarshal([]byte(C.GoString(cCredentialJSONs)), &credentials)

	if err != nil {
		fmt.Println("unable to parse credentials")
		return 1
	}

	var sessionData webauthn.SessionData

	err = json.Unmarshal([]byte(sessionDataJson), &sessionData)

	if err != nil{
		fmt.Println("Unable to parse session JSON");
		return 1
	}

	par, err := car.Parse()

	if err != nil{
		fmt.Println("Unable to parse session JSON into credential assertion data");
		return 1
	}
	
	user :=  User {ID: id, Credentials: credentials, Name: name}

	wconfig := &webauthn.Config{
		RPDisplayName: "Go Webauthn", // Display Name for your site
		RPID: "go-webauthn.local", // Generally the FQDN for your site
		RPOrigins: []string{"https://login.go-webauthn.local"}, // The origin URLs allowed for WebAuthn requests
	}

	wa, err := webauthn.New(wconfig)

	if err != nil {
	    fmt.Println(err)
	    return 1
	}

	_, err = wa.ValidateLogin(user, sessionData, par)

	if err != nil {
		return 1
	}

	return 0
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u User)  WebAuthnName() string {
	return u.Name
}

func (u User) WebAuthnID() []byte {
	return []byte(u.ID)
}

func (u User) WebAuthnDisplayName() string {
	return u.Name
}

func (u User) WebAuthnIcon() string {
        return "https://pics.com/avatar.png"
}

func main(){}
