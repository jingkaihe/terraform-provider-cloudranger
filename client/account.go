/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type Account struct {
	Id string `json:"id,omitempty"`

	OrganizationId string `json:"organization_id,omitempty"`

	Name string `json:"name,omitempty"`

	Initials string `json:"initials,omitempty"`

	Color string `json:"color,omitempty"`

	AvatarImage string `json:"avatar_image,omitempty"`

	Type_ string `json:"type,omitempty"`

	DefaultTimezoneLocation string `json:"default_timezone_location,omitempty"`

	IsActive bool `json:"is_active,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`

	DeletedAt string `json:"deleted_at,omitempty"`

	Settings *Setting `json:"settings,omitempty"`

	Credential *Credential `json:"credential,omitempty"`

	Members *Member `json:"members,omitempty"`
}