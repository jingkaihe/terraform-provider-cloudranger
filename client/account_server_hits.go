/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type AccountServerHits struct {
	Type_ string `json:"Type,omitempty"`

	InstanceId string `json:"InstanceId,omitempty"`

	Name string `json:"Name,omitempty"`

	TagCount float32 `json:"TagCount,omitempty"`

	Tags *Tag `json:"Tags,omitempty"`

	InstanceType string `json:"InstanceType,omitempty"`

	State *AccountServerState `json:"State,omitempty"`

	Region string `json:"Region,omitempty"`

	AvailabilityZone string `json:"AvailabilityZone,omitempty"`

	Schedules *Schedule `json:"Schedules,omitempty"`

	Policies *Policy `json:"Policies,omitempty"`
}
