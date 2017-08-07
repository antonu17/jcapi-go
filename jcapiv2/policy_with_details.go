/* 
 * JumpCloud APIs
 *
 * V1 and V2 versions of JumpCloud's API. The next version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings. The most recent version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings.
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package jcapiv2

// An instance of a policty template.
type PolicyWithDetails struct {

	// ObjectId uniquely indetifying a Policy.
	Id string `json:"id,omitempty"`

	Template PolicyTemplate `json:"template,omitempty"`

	// The description for this specific Policy.
	Name string `json:"name,omitempty"`

	Values []PolicyValue `json:"values,omitempty"`
}
