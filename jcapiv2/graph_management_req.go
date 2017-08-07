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

type GraphManagementReq struct {

	// How to modify the graph connection.
	Op string `json:"op"`

	Type_ GraphType `json:"type"`

	// The ObjectID of graph object being added or removed as an association.
	Id string `json:"id"`
}
