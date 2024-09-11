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

// checks if the Version type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Version{}

// Version Represents a Function Version in a list.
type Version struct {
	// An identifier for this version.
	Id string `json:"id"`
	// Author of this version.
	Author *string `json:"author,omitempty"`
	// Source code of this version.
	Code string `json:"code"`
	// The deployed status of this version.
	IsDeployed *bool `json:"isDeployed,omitempty"`
	// The time of this Version's creation.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time of this Version's latest update.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The time of this Version's last deployment.
	DeployedAt *string `json:"deployedAt,omitempty"`
}

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion(id string, code string) *Version {
	this := Version{}
	this.Id = id
	this.Code = code
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetId returns the Id field value
func (o *Version) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Version) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Version) SetId(v string) {
	o.Id = v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Version) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Version) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *Version) SetAuthor(v string) {
	o.Author = &v
}

// GetCode returns the Code field value
func (o *Version) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Version) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Version) SetCode(v string) {
	o.Code = v
}

// GetIsDeployed returns the IsDeployed field value if set, zero value otherwise.
func (o *Version) GetIsDeployed() bool {
	if o == nil || IsNil(o.IsDeployed) {
		var ret bool
		return ret
	}
	return *o.IsDeployed
}

// GetIsDeployedOk returns a tuple with the IsDeployed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetIsDeployedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeployed) {
		return nil, false
	}
	return o.IsDeployed, true
}

// HasIsDeployed returns a boolean if a field has been set.
func (o *Version) HasIsDeployed() bool {
	if o != nil && !IsNil(o.IsDeployed) {
		return true
	}

	return false
}

// SetIsDeployed gets a reference to the given bool and assigns it to the IsDeployed field.
func (o *Version) SetIsDeployed(v bool) {
	o.IsDeployed = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Version) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Version) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Version) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Version) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Version) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Version) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeployedAt returns the DeployedAt field value if set, zero value otherwise.
func (o *Version) GetDeployedAt() string {
	if o == nil || IsNil(o.DeployedAt) {
		var ret string
		return ret
	}
	return *o.DeployedAt
}

// GetDeployedAtOk returns a tuple with the DeployedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetDeployedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeployedAt) {
		return nil, false
	}
	return o.DeployedAt, true
}

// HasDeployedAt returns a boolean if a field has been set.
func (o *Version) HasDeployedAt() bool {
	if o != nil && !IsNil(o.DeployedAt) {
		return true
	}

	return false
}

// SetDeployedAt gets a reference to the given string and assigns it to the DeployedAt field.
func (o *Version) SetDeployedAt(v string) {
	o.DeployedAt = &v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Version) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.IsDeployed) {
		toSerialize["isDeployed"] = o.IsDeployed
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.DeployedAt) {
		toSerialize["deployedAt"] = o.DeployedAt
	}
	return toSerialize, nil
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
