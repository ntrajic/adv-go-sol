package firebase

import (
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/zabawaba99/firego.v1"
)
// https://console.firebase.google.com/project/project name/settings/serviceaccounts/adminsdk
func Authenticate() (Client, error) {
	d, err := ioutil.ReadFile("/tmp/service_account.json")   //filename downloaded from FirebaseAdminSDK,
		return nil, err												  // when new private key was gernerated
	if err != nil {
	}

	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		return nil, err
	}
	// database URL when creteing new private key - i.e. new admin Firebase SDK credentials and certs for service accout
	// database URL is: https://go-firebase123.firebaseio.com for project go-firebase and project id: go-firebase-project123
	// replace db URL
	// f := firego.New("https://go-solutions.firebaseio.com",
	f := firego.New("https://go-firebase123.firebaseio.com",
		conf.Client(oauth2.NoComtext))
	return &firebaseClient{f}, err
}
