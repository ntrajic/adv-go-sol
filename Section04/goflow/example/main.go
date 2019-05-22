package main

import (
	"fmt"

	".."
	flow "github.com/trustmaster/goflow"
)

func main() {

	net := goflow.NewEncodingApp()

	in := make(chan string)
	net.SetInPort("In", in)

	wait := flow.Run(net)

	for i := 0; i < 20; i++ {
		in <- fmt.Sprint("Message", i)
	}

	close(in)
	<-wait
	//<-net.Wait()
}
