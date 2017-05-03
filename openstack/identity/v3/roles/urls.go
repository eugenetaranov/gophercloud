package roles

import "github.com/eugenetaranov/gophercloud"

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}
