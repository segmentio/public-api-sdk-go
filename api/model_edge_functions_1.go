/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.7
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EdgeFunctions1 The latest version of Edge Function bundle.
type EdgeFunctions1 struct {
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

// NewEdgeFunctions1 instantiates a new EdgeFunctions1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFunctions1(id string, sourceId string, createdAt string, createdBy string, downloadURL string, version float32) *EdgeFunctions1 {
	this := EdgeFunctions1{}
	this.Id = id
	this.SourceId = sourceId
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.DownloadURL = downloadURL
	this.Version = version
	return &this
}

// NewEdgeFunctions1WithDefaults instantiates a new EdgeFunctions1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFunctions1WithDefaults() *EdgeFunctions1 {
	this := EdgeFunctions1{}
	return &this
}

// GetId returns the Id field value
func (o *EdgeFunctions1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeFunctions1) SetId(v string) {
	o.Id = v
}

// GetSourceId returns the SourceId field value
func (o *EdgeFunctions1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *EdgeFunctions1) SetSourceId(v string) {
	o.SourceId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EdgeFunctions1) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions1) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EdgeFunctions1) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *EdgeFunctions1) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions1) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *EdgeFunctions1) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDownloadURL returns the DownloadURL field value
func (o *EdgeFunctions1) GetDownloadURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadURL
}

// GetDownloadURLOk returns a tuple with the DownloadURL field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions1) GetDownloadURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadURL, true
}

// SetDownloadURL sets field value
func (o *EdgeFunctions1) SetDownloadURL(v string) {
	o.DownloadURL = v
}

// GetVersion returns the Version field value
func (o *EdgeFunctions1) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions1) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EdgeFunctions1) SetVersion(v float32) {
	o.Version = v
}

func (o EdgeFunctions1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["downloadURL"] = o.DownloadURL
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableEdgeFunctions1 struct {
	value *EdgeFunctions1
	isSet bool
}

func (v NullableEdgeFunctions1) Get() *EdgeFunctions1 {
	return v.value
}

func (v *NullableEdgeFunctions1) Set(val *EdgeFunctions1) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFunctions1) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFunctions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFunctions1(val *EdgeFunctions1) *NullableEdgeFunctions1 {
	return &NullableEdgeFunctions1{value: val, isSet: true}
}

func (v NullableEdgeFunctions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFunctions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


