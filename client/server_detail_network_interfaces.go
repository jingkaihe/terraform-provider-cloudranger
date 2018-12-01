/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type ServerDetailNetworkInterfaces struct {
	Attachment *ServerDetailAttachment `json:"Attachment,omitempty"`

	Description string `json:"Description,omitempty"`

	Groups []ServerDetailGroups `json:"Groups,omitempty"`

	Ipv6Addresses []Empty `json:"Ipv6Addresses,omitempty"`

	MacAddress string `json:"MacAddress,omitempty"`

	NetworkInterfaceId string `json:"NetworkInterfaceId,omitempty"`

	OwnerId string `json:"OwnerId,omitempty"`

	PrivateDnsName string `json:"PrivateDnsName,omitempty"`

	PrivateIpAddress string `json:"PrivateIpAddress,omitempty"`

	PrivateIpAddresses []ServerDetailPrivateIpAddresses `json:"PrivateIpAddresses,omitempty"`

	SourceDestCheck string `json:"SourceDestCheck,omitempty"`

	Status string `json:"Status,omitempty"`

	SubnetId string `json:"SubnetId,omitempty"`

	VpcId string `json:"VpcId,omitempty"`
}
