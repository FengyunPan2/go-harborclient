/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * OpenAPI spec version: 0.3.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Storage struct {

	// Total volume size.
	Total int64 `json:"total,omitempty"`

	// Free volume size.
	Free int64 `json:"free,omitempty"`
}
