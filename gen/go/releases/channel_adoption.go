/* 
 * Vendor API V1
 *
 * Create, list, promote, update and archive releases.
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// ChannelAdoption represents the versions that licenses are on in the channel
type ChannelAdoption struct {

	CurrentVersionCountActive map[string]int64 `json:"current_version_count_active,omitempty"`

	CurrentVersionCountAll map[string]int64 `json:"current_version_count_all,omitempty"`

	OtherVersionCountActive map[string]int64 `json:"other_version_count_active,omitempty"`

	OtherVersionCountAll map[string]int64 `json:"other_version_count_all,omitempty"`

	PreviousVersionCountActive map[string]int64 `json:"previous_version_count_active,omitempty"`

	PreviousVersionCountAll map[string]int64 `json:"previous_version_count_all,omitempty"`
}
