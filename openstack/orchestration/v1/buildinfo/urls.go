package buildinfo

import "github.com/eugenetaranov/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
