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
			fmt.Printf("Starting the vm with id %s\n", text)
		}
	}

	p := cs.VirtualMachine.NewStartVirtualMachineParams(text)
	r, err := cs.VirtualMachine.StartVirtualMachine(p)

	if err != nil {
		fmt.Printf("Error Starting the vm %s:", err)
	} else {
		fmt.Printf("Successfully started vm %s", r.Name)
	}
}
