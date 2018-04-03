package clustertemplates

import (
	"fmt"

	"github.com/gophercloud/gophercloud/pagination"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/containerorchestration/v1/common"
)

// List returns a Pager which allows you to iterate over a collection of
// clustertemplates.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return ClusterTemplatePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific clustertemplate based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	ro := &gophercloud.RequestOpts{ErrorContext: &common.ErrorResponse{}}
	_, r.Err = c.Get(getURL(c, id), &r.Body, ro)
	return
}

// CreateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Create operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type CreateOptsBuilder interface {
	ToClusterTemplateMap() (map[string]interface{}, error)
}

// CreateOpts satisfies the CreateOptsBuilder interface
type CreateOpts struct {
	KeyPair             string `json:"keypair_id" required:"true"`
	TLSDisabled         bool   `json:"tls_disabled" required:"true"`
	Public              bool   `json:"public" required:"true"`
	DockerVolumeSize    int    `json:"docker_volume_size" required:"true"`
	ServerType          string `json:"server_type" required:"true"`
	ExternalNetworkID   string `json:"external_network_id" required:"true"`
	ImageID             string `json:"image_id" required:"true"`
	VolumeDriver        string `json:"volume_driver" required:"true"`
	DockerStorageDriver string `json:"docker_storage_driver" required:"true"`
	Name                string `json:"name" required:"true"`
	NetworkDriver       string `json:"network_driver" required:"true"`
	COE                 string `json:"coe" required:"true"`
	FlavorID            string `json:"flavor_id" required:"true"`
	MasterLBEnabled     bool   `json:"master_lb_enabled" required:"true"`
	DNSNameserver       string `json:"dns_nameserver" required:"true"`
}

// ClusterCreateTemplateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToClusterTemplateMap() (map[string]interface{}, error) {
	// zz, err := gophercloud.BuildRequestBody(opts, "")
	// fmt.Println("KKK", zz)
	// fmt.Println("KKK", err)
	req, err := gophercloud.BuildRequestBody(opts, "")
	fmt.Println("ZZZ", req)
	fmt.Println("ZZZ", err)
	return req, err
	// return map[string]interface{}{
	// 	"keypair_id":         opts.KeyPair,
	// 	"tls_disabled":       opts.TLSDisabled,
	// 	"public":             opts.Public,
	// 	"docker_volume_size": opts.DockerVolumeSize,
	// 	"image_id":           opts.ImageID,
	// 	"coe":                opts.COE,
	// }, nil
}

// Create accepts a CreateOpts struct and creates a new cluster using the values
// provided.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToClusterTemplateMap()
	fmt.Println("ClusterTemplateCreateMap:", b)
	if err != nil {
		r.Err = err
		return
	}

	ro := &gophercloud.RequestOpts{ErrorContext: &common.ErrorResponse{}}
	fmt.Println(createURL(c), b, ro)
	_, r.Err = c.Post(createURL(c), b, &r.Body, ro)
	return
}

// Delete accepts a unique ID and deletes the cluster associated with it.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	ro := &gophercloud.RequestOpts{ErrorContext: &common.ErrorResponse{}}
	_, r.Err = c.Delete(deleteURL(c, id), ro)
	return
}
