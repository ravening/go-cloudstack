package main

import (
	"fmt"
	"github.com/xanzy/go-cloudstack/v2/cloudstack"
	Constants "github.com/xanzy/go-cloudstack/v2/constants"
	"strconv"
	"time"
)

func main() {
	server := Constants.Server
	apikey := Constants.Apikey
	secretkey := Constants.Secretkey
	BytesReceived := 5
	BytesSent := 4

	cs := cloudstack.NewAsyncClient(server, apikey, secretkey, false)
	t := time.Now()
	now := t.Format(time.RFC3339)
	// adjust the parameters to get desired start date
	yesterday := t.AddDate(0, 0, -1).Format(time.RFC3339)
	params := &cloudstack.ListUsageRecordsParams{}
	params.SetStartdate(yesterday)
	params.SetEnddate(now)
	params.SetType(int64(BytesSent))
	var response *cloudstack.ListUsageRecordsResponse
	response, _ = cs.Usage.ListUsageRecords(params)

	displayResult(response, "Total bytes sent for domain-account ")

	params.SetType(int64(BytesReceived))
	response, _ = cs.Usage.ListUsageRecords(params)
	displayResult(response, "Total bytes received for domain-account ")
}

func displayResult(response *cloudstack.ListUsageRecordsResponse, display string) {
	usage := make(map[string]uint64)
	for i := 0; i < response.Count; i++ {
		key := response.UsageRecords[i].Domain + "-" + response.UsageRecords[i].Account
		current, _ := strconv.ParseUint(response.UsageRecords[i].Rawusage, 10, 64)
		if value, found := usage[key]; found {
			usage[key] = value + current
		} else {
			usage[key] = current
		}
	}
	for k, v := range usage {
		fmt.Printf(display + k + " is " + strconv.FormatUint(v, 10) + " bytes\n")
	}

}
