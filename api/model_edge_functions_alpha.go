/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the EdgeFunctionsAlpha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFunctionsAlpha{}

// EdgeFunctionsAlpha Represents an Edge Function bundle.
type EdgeFunctionsAlpha struct {
	// The Edge Function id.
	Id string `json:"id"`
	// The Source id.
	SourceId string `json:"sourceId"`
	// Creation date.
	CreatedAt string `json:"createdAt"`
	// Creating user's id.
	CreatedBy string `json:"createdBy"`
	// The CDN URL that can be used to fetch your latest EdgeFunctions bundle.
	DownloadURL string `json:"downloadURL"`
	// Revision number associated with the latest Edge Function.
	Version float32 `json:"version"`
}

// NewEdgeFunctionsAlpha instantiates a new EdgeFunctionsAlpha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFunctionsAlpha(
	id string,
	sourceId string,
	createdAt string,
	createdBy string,
	downloadURL string,
	version float32,
) *EdgeFunctionsAlpha {
	this := EdgeFunctionsAlpha{}
	this.Id = id
	this.SourceId = sourceId
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.DownloadURL = downloadURL
	this.Version = version
	return &this
}

// NewEdgeFunctionsAlphaWithDefaults instantiates a new EdgeFunctionsAlpha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFunctionsAlphaWithDefaults() *EdgeFunctionsAlpha {
	this := EdgeFunctionsAlpha{}
	return &this
}

// GetId returns the Id field value
func (o *EdgeFunctionsAlpha) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsAlpha) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeFunctionsAlpha) SetId(v string) {
	o.Id = v
}

// GetSourceId returns the SourceId field value
func (o *EdgeFunctionsAlpha) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsAlpha) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *EdgeFunctionsAlpha) SetSourceId(v string) {
	o.SourceId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EdgeFunctionsAlpha) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsAlpha) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EdgeFunctionsAlpha) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *EdgeFunctionsAlpha) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsAlpha) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *EdgeFunctionsAlpha) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDownloadURL returns the DownloadURL field value
func (o *EdgeFunctionsAlpha) GetDownloadURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadURL
}

// GetDownloadURLOk returns a tuple with the DownloadURL field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsAlpha) GetDownloadURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadURL, true
}

// SetDownloadURL sets field value
func (o *EdgeFunctionsAlpha) SetDownloadURL(v string) {
	o.DownloadURL = v
}

// GetVersion returns the Version field value
func (o *EdgeFunctionsAlpha) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsAlpha) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EdgeFunctionsAlpha) SetVersion(v float32) {
	o.Version = v
}

func (o EdgeFunctionsAlpha) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFunctionsAlpha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["sourceId"] = o.SourceId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["downloadURL"] = o.DownloadURL
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableEdgeFunctionsAlpha struct {
	value *EdgeFunctionsAlpha
	isSet bool
}

func (v NullableEdgeFunctionsAlpha) Get() *EdgeFunctionsAlpha {
	return v.value
}

func (v *NullableEdgeFunctionsAlpha) Set(val *EdgeFunctionsAlpha) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFunctionsAlpha) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFunctionsAlpha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFunctionsAlpha(val *EdgeFunctionsAlpha) *NullableEdgeFunctionsAlpha {
	return &NullableEdgeFunctionsAlpha{value: val, isSet: true}
}

func (v NullableEdgeFunctionsAlpha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFunctionsAlpha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
