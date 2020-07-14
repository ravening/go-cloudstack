package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xanzy/go-cloudstack/v2/cloudstack"
	Constants "github.com/xanzy/go-cloudstack/v2/constants"
)

func main() {
	server := Constants.Server
	apikey := Constants.Apikey
	secretkey := Constants.Secretkey

	cs := cloudstack.NewAsyncClient(server, apikey, secretkey, false)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err == nil {
		text = strings.Trim(text, "\n")
		if text != "" {
			fmt.Printf("Stopping the vm with id %s\n", text)
		}

	}

	p := cs.VirtualMachine.NewStopVirtualMachineParams(text)
	r, err := cs.VirtualMachine.StopVirtualMachine(p)

	if err != nil {
		fmt.Printf("Error stopping the vm %s:", err)
	} else {
		fmt.Printf("Successfully Stopped vm %s", r.Name)
	}

}
