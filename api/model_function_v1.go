/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 42.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the FunctionV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionV1{}

// FunctionV1 Represents a Function.
type FunctionV1 struct {
	// An identifier for this Function.
	Id *string `json:"id,omitempty"`
	// The Function type.  Config API note: equal to `type`.
	ResourceType *string `json:"resourceType,omitempty"`
	// The time this Function was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The id of the user who created this Function.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The Function code.
	Code *string `json:"code,omitempty"`
	// The time of this Function's last deployment.
	DeployedAt NullableString `json:"deployedAt,omitempty"`
	// The list of settings for this Function.
	Settings []FunctionSettingV1 `json:"settings,omitempty"`
	// A display name for this Function.
	DisplayName *string `json:"displayName,omitempty"`
	// A description for this Function.
	Description *string `json:"description,omitempty"`
	// The URL of the logo for this Function.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// The preview webhook URL for this Function.
	PreviewWebhookUrl *string `json:"previewWebhookUrl,omitempty"`
	// The max count of the batch for this Function.
	BatchMaxCount *float32 `json:"batchMaxCount,omitempty"`
	// The catalog id of this Function.
	CatalogId *string `json:"catalogId,omitempty"`
	// Whether the deployment of this Function is the latest version.
	IsLatestVersion *bool `json:"isLatestVersion,omitempty"`
}

// NewFunctionV1 instantiates a new FunctionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionV1() *FunctionV1 {
	this := FunctionV1{}
	return &this
}

// NewFunctionV1WithDefaults instantiates a new FunctionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionV1WithDefaults() *FunctionV1 {
	this := FunctionV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FunctionV1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FunctionV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FunctionV1) SetId(v string) {
	o.Id = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *FunctionV1) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *FunctionV1) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *FunctionV1) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FunctionV1) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FunctionV1) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FunctionV1) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FunctionV1) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FunctionV1) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *FunctionV1) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FunctionV1) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FunctionV1) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *FunctionV1) SetCode(v string) {
	o.Code = &v
}

// GetDeployedAt returns the DeployedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionV1) GetDeployedAt() string {
	if o == nil || IsNil(o.DeployedAt.Get()) {
		var ret string
		return ret
	}
	return *o.DeployedAt.Get()
}

// GetDeployedAtOk returns a tuple with the DeployedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionV1) GetDeployedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeployedAt.Get(), o.DeployedAt.IsSet()
}

// HasDeployedAt returns a boolean if a field has been set.
func (o *FunctionV1) HasDeployedAt() bool {
	if o != nil && o.DeployedAt.IsSet() {
		return true
	}

	return false
}

// SetDeployedAt gets a reference to the given NullableString and assigns it to the DeployedAt field.
func (o *FunctionV1) SetDeployedAt(v string) {
	o.DeployedAt.Set(&v)
}

// SetDeployedAtNil sets the value for DeployedAt to be an explicit nil
func (o *FunctionV1) SetDeployedAtNil() {
	o.DeployedAt.Set(nil)
}

// UnsetDeployedAt ensures that no value is present for DeployedAt, not even an explicit nil
func (o *FunctionV1) UnsetDeployedAt() {
	o.DeployedAt.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *FunctionV1) GetSettings() []FunctionSettingV1 {
	if o == nil || IsNil(o.Settings) {
		var ret []FunctionSettingV1
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetSettingsOk() ([]FunctionSettingV1, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *FunctionV1) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []FunctionSettingV1 and assigns it to the Settings field.
func (o *FunctionV1) SetSettings(v []FunctionSettingV1) {
	o.Settings = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FunctionV1) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FunctionV1) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FunctionV1) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FunctionV1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FunctionV1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FunctionV1) SetDescription(v string) {
	o.Description = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *FunctionV1) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *FunctionV1) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *FunctionV1) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetPreviewWebhookUrl returns the PreviewWebhookUrl field value if set, zero value otherwise.
func (o *FunctionV1) GetPreviewWebhookUrl() string {
	if o == nil || IsNil(o.PreviewWebhookUrl) {
		var ret string
		return ret
	}
	return *o.PreviewWebhookUrl
}

// GetPreviewWebhookUrlOk returns a tuple with the PreviewWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetPreviewWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewWebhookUrl) {
		return nil, false
	}
	return o.PreviewWebhookUrl, true
}

// HasPreviewWebhookUrl returns a boolean if a field has been set.
func (o *FunctionV1) HasPreviewWebhookUrl() bool {
	if o != nil && !IsNil(o.PreviewWebhookUrl) {
		return true
	}

	return false
}

// SetPreviewWebhookUrl gets a reference to the given string and assigns it to the PreviewWebhookUrl field.
func (o *FunctionV1) SetPreviewWebhookUrl(v string) {
	o.PreviewWebhookUrl = &v
}

// GetBatchMaxCount returns the BatchMaxCount field value if set, zero value otherwise.
func (o *FunctionV1) GetBatchMaxCount() float32 {
	if o == nil || IsNil(o.BatchMaxCount) {
		var ret float32
		return ret
	}
	return *o.BatchMaxCount
}

// GetBatchMaxCountOk returns a tuple with the BatchMaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetBatchMaxCountOk() (*float32, bool) {
	if o == nil || IsNil(o.BatchMaxCount) {
		return nil, false
	}
	return o.BatchMaxCount, true
}

// HasBatchMaxCount returns a boolean if a field has been set.
func (o *FunctionV1) HasBatchMaxCount() bool {
	if o != nil && !IsNil(o.BatchMaxCount) {
		return true
	}

	return false
}

// SetBatchMaxCount gets a reference to the given float32 and assigns it to the BatchMaxCount field.
func (o *FunctionV1) SetBatchMaxCount(v float32) {
	o.BatchMaxCount = &v
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *FunctionV1) GetCatalogId() string {
	if o == nil || IsNil(o.CatalogId) {
		var ret string
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetCatalogIdOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogId) {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *FunctionV1) HasCatalogId() bool {
	if o != nil && !IsNil(o.CatalogId) {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given string and assigns it to the CatalogId field.
func (o *FunctionV1) SetCatalogId(v string) {
	o.CatalogId = &v
}

// GetIsLatestVersion returns the IsLatestVersion field value if set, zero value otherwise.
func (o *FunctionV1) GetIsLatestVersion() bool {
	if o == nil || IsNil(o.IsLatestVersion) {
		var ret bool
		return ret
	}
	return *o.IsLatestVersion
}

// GetIsLatestVersionOk returns a tuple with the IsLatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionV1) GetIsLatestVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLatestVersion) {
		return nil, false
	}
	return o.IsLatestVersion, true
}

// HasIsLatestVersion returns a boolean if a field has been set.
func (o *FunctionV1) HasIsLatestVersion() bool {
	if o != nil && !IsNil(o.IsLatestVersion) {
		return true
	}

	return false
}

// SetIsLatestVersion gets a reference to the given bool and assigns it to the IsLatestVersion field.
func (o *FunctionV1) SetIsLatestVersion(v bool) {
	o.IsLatestVersion = &v
}

func (o FunctionV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionV1) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.DeployedAt.IsSet() {
		toSerialize["deployedAt"] = o.DeployedAt.Get()
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
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
	if !IsNil(o.PreviewWebhookUrl) {
		toSerialize["previewWebhookUrl"] = o.PreviewWebhookUrl
	}
	if !IsNil(o.BatchMaxCount) {
		toSerialize["batchMaxCount"] = o.BatchMaxCount
	}
	if !IsNil(o.CatalogId) {
		toSerialize["catalogId"] = o.CatalogId
	}
	if !IsNil(o.IsLatestVersion) {
		toSerialize["isLatestVersion"] = o.IsLatestVersion
	}
	return toSerialize, nil
}

type NullableFunctionV1 struct {
	value *FunctionV1
	isSet bool
}

func (v NullableFunctionV1) Get() *FunctionV1 {
	return v.value
}

func (v *NullableFunctionV1) Set(val *FunctionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionV1(val *FunctionV1) *NullableFunctionV1 {
	return &NullableFunctionV1{value: val, isSet: true}
}

func (v NullableFunctionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
