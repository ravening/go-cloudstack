package main

import (
    "fmt"
    "github.com/xanzy/go-cloudstack/v2/cloudstack"
    Constants "github.com/xanzy/go-cloudstack/v2/constants"
)

func main()  {
    server := Constants.Server
    apikey := Constants.Apikey
    secretkey := Constants.Secretkey

    cs := cloudstack.NewAsyncClient(server, apikey, secretkey, false)

    params := &cloudstack.ListRoutersParams{}
    params.SetListall(true)
    var listRoutersResponse *cloudstack.ListRoutersResponse
    listRoutersResponse, _ = cs.Router.ListRouters(params)

    if listRoutersResponse.Count > 0 {
        for i := 0; i < listRoutersResponse.Count; i++ {
            routerHealthCheckParams := &cloudstack.GetRouterHealthCheckResultsParams{}
            routerHealthCheckParams.SetRouterid(listRoutersResponse.Routers[i].Id)
            routerHealthCheckParams.SetPerformfreshchecks(false)

            var routerHealthCheckResponse *cloudstack.GetRouterHealthCheckResultsResponse
            routerHealthCheckResponse, _ = cs.Router.GetRouterHealthCheckResults(routerHealthCheckParams)
            fmt.Printf("Router Name: %s\n", listRoutersResponse.Routers[i].Name)
            for j := 0; j < len(routerHealthCheckResponse.Healthcheck); j++ {
                fmt.Printf("\tCheck name: %s, status: %t",
                    routerHealthCheckResponse.Healthcheck[j].Checkname,
                    routerHealthCheckResponse.Healthcheck[j].Success)
                fmt.Println()
            }
        }
    }
}
