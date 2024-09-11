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

// checks if the ListFunctionItemV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFunctionItemV1{}

// ListFunctionItemV1 Represents a Function in a list.
type ListFunctionItemV1 struct {
	// An identifier for this Function.
	Id *string `json:"id,omitempty"`
	// The Function type.  Config API note: equal to `type`.
	ResourceType *string `json:"resourceType,omitempty"`
	// The time this Function was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The id of the user who created this Function.
	CreatedBy *string `json:"createdBy,omitempty"`
	// A display name for this Function.
	DisplayName *string `json:"displayName,omitempty"`
	// A description for this Function.
	Description *string `json:"description,omitempty"`
	// The URL of the logo for this Function.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// The catalog id of this Function.
	CatalogId *string `json:"catalogId,omitempty"`
}

// NewListFunctionItemV1 instantiates a new ListFunctionItemV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFunctionItemV1() *ListFunctionItemV1 {
	this := ListFunctionItemV1{}
	return &this
}

// NewListFunctionItemV1WithDefaults instantiates a new ListFunctionItemV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFunctionItemV1WithDefaults() *ListFunctionItemV1 {
	this := ListFunctionItemV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListFunctionItemV1) SetId(v string) {
	o.Id = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ListFunctionItemV1) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ListFunctionItemV1) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ListFunctionItemV1) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListFunctionItemV1) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListFunctionItemV1) SetDescription(v string) {
	o.Description = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *ListFunctionItemV1) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *ListFunctionItemV1) GetCatalogId() string {
	if o == nil || IsNil(o.CatalogId) {
		var ret string
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionItemV1) GetCatalogIdOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogId) {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *ListFunctionItemV1) HasCatalogId() bool {
	if o != nil && !IsNil(o.CatalogId) {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given string and assigns it to the CatalogId field.
func (o *ListFunctionItemV1) SetCatalogId(v string) {
	o.CatalogId = &v
}

func (o ListFunctionItemV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFunctionItemV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if !IsNil(o.CatalogId) {
		toSerialize["catalogId"] = o.CatalogId
	}
	return toSerialize, nil
}

type NullableListFunctionItemV1 struct {
	value *ListFunctionItemV1
	isSet bool
}

func (v NullableListFunctionItemV1) Get() *ListFunctionItemV1 {
	return v.value
}

func (v *NullableListFunctionItemV1) Set(val *ListFunctionItemV1) {
	v.value = val
	v.isSet = true
}

func (v NullableListFunctionItemV1) IsSet() bool {
	return v.isSet
}

func (v *NullableListFunctionItemV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFunctionItemV1(val *ListFunctionItemV1) *NullableListFunctionItemV1 {
	return &NullableListFunctionItemV1{value: val, isSet: true}
}

func (v NullableListFunctionItemV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFunctionItemV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
