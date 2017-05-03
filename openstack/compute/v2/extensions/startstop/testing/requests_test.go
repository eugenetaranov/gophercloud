package testing

import (
	"testing"

	"github.com/eugenetaranov/gophercloud/openstack/compute/v2/extensions/startstop"
	th "github.com/eugenetaranov/gophercloud/testhelper"
	"github.com/eugenetaranov/gophercloud/testhelper/client"
)

const serverID = "{serverId}"

func TestStart(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStartServerResponse(t, serverID)

	err := startstop.Start(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestStop(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStopServerResponse(t, serverID)

	err := startstop.Stop(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
