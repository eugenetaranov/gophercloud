package clustertemplates

import "github.com/gophercloud/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("clustertemplates")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("clustertemplates", id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("clustertemplates")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("clustertemplates", id)
}
