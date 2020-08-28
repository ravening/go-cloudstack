package main

import (
	"fmt"
	"github.com/xanzy/go-cloudstack/v2/cloudstack"
	Constants "github.com/xanzy/go-cloudstack/v2/constants"
)

func main() {
	server := Constants.Server
	apikey := Constants.Apikey
	secretkey := Constants.Secretkey

	cs := cloudstack.NewAsyncClient(server, apikey, secretkey, false)

	params := &cloudstack.ListRoutersParams{}
	params.SetListall(true)

	var response *cloudstack.ListRoutersResponse
	response, _ = cs.Router.ListRouters(params)

	for i := 0; i < response.Count; i++ {
		fmt.Println(response.Routers[i].Name)
		// add more fields if you need
	}
}
