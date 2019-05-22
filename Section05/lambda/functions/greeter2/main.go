package main

import (
	"encoding/json"
	"fmt"

	"github.com/apex/go-apex"
	// on the client win machine, inside %GOPATH% i.e. c:\go-work we install go libraries with:
	// go get github.com/apex/go-apex
	// note that apex.exe executable is also necessary with AWS CLI and aws-credentials.csc downloaded from AWS portal
	// for newly created user: apex-user, which will deploy go serverless functions to aws portal.
)

// Version 2
type message struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		// 0.0 pass aws-credetials.cvs from AWS portal into AWS CLI:   aws configure <enter>
		// Apex expects you enter Access Key D ad Secret Access Key for the newly created apex-user on AWS Cloud Portal
		// 0.1 apex init <enter>
		// 0.2 supply project name, and description for the Apex project
		// 0.3 apex deploy <enter>
		// cloud OUT: project.json and functions empty dir
		// 1. local client test of function Hello, ime prezime
		// echo '{"event":{"name":"test"}}' | go run main.go
		// output: { "value" :{"greeting":"Hello, test"}}
		// 2. apex deploy <enter>
		// output: creating function		env=function=greeter
		//         creating alias current	env=function=greeter version=1
		//         function created			env=function=greeter name=go-solutions_greter version=1
		// 3. local client test:
		// echo '{"name":"test"}' | apex invoke greeter {"greeting":"Hello, test"}
		// OUTPUT: {"greeting":"Hello, test"}
		// 4. create new version of hello world with first and second name. this is version=2 repate 1-3 steps
		// 5. check logs with:   apex logs greeter <enter>
		// START RequestId: .... version=1
		// END   RequestId: same
		// REPORT: Memory size: 128 MB, Max Memory Used: 20MB, Duration 26 sec, follows report START/END/REPRORT for v2.
		// 6. clean up - of funtcions:  apex delete <enter>


		resp := map[string]string{
			"greeting": fmt.Sprintf("Hello, %s %s", m.FirstName, m.LastName),
		}

		return resp, nil
	})
}
