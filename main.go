package main

/*
#cgo CFLAGS: -I.

struct ValidateLoginResult {
    int isValid;
    char* error;
};
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

func createError(msg string) *C.struct_ValidateLoginResult {
	result := (*C.struct_ValidateLoginResult)(C.malloc(C.sizeof_struct_ValidateLoginResult))

	result.isValid = 0
	result.error = C.CString(msg)
	return result
}



//export ValidateLogin
func ValidateLogin (cId, cName, cCredentialJSONs, cSessionDataJson *C.char) *C.struct_ValidateLoginResult {
	var car protocol.CredentialAssertionResponse

	var credentials []webauthn.Credential

	id := C.GoString(cId)
	name := C.GoString(cName)
	sessionDataJson := C.GoString(cSessionDataJson)

	err := json.Unmarshal([]byte(C.GoString(cCredentialJSONs)), &credentials)

	if err != nil {
		return createError("unable to parse credentials")
	}

	var sessionData webauthn.SessionData

	err = json.Unmarshal([]byte(sessionDataJson), &sessionData)

	if err != nil{
		return createError("Unable to parse session JSON");
	}

	par, err := car.Parse()

	if err != nil{
		return createError("Unable to parse session JSON into credential assertion data");
	}
	
	user :=  User {ID: id, Credentials: credentials, Name: name}

	wconfig := &webauthn.Config{
		RPDisplayName: "Go Webauthn", // Display Name for your site
		RPID: "go-webauthn.local", // Generally the FQDN for your site
		RPOrigins: []string{"https://login.go-webauthn.local"}, // The origin URLs allowed for WebAuthn requests
	}

	wa, err := webauthn.New(wconfig)

	if err != nil {
	    return createError(fmt.Sprint(err))
	}

	_, err = wa.ValidateLogin(user, sessionData, par)

	if err != nil {
	    return createError(fmt.Sprint(err))
	}

	result := (*C.struct_ValidateLoginResult)(C.malloc(C.sizeof_struct_ValidateLoginResult))

	result.isValid = 1
	result.error = C.CString("");
	return result
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
