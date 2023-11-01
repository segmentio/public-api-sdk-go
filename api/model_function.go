/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Function A Function object.
type Function struct {
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

// NewFunction instantiates a new Function object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunction() *Function {
	this := Function{}
	return &this
}

// NewFunctionWithDefaults instantiates a new Function object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionWithDefaults() *Function {
	this := Function{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Function) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Function) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Function) SetId(v string) {
	o.Id = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Function) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Function) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *Function) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Function) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Function) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Function) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Function) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Function) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Function) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Function) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Function) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Function) SetCode(v string) {
	o.Code = &v
}

// GetDeployedAt returns the DeployedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Function) GetDeployedAt() string {
	if o == nil || o.DeployedAt.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeployedAt.Get()
}

// GetDeployedAtOk returns a tuple with the DeployedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Function) GetDeployedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeployedAt.Get(), o.DeployedAt.IsSet()
}

// HasDeployedAt returns a boolean if a field has been set.
func (o *Function) HasDeployedAt() bool {
	if o != nil && o.DeployedAt.IsSet() {
		return true
	}

	return false
}

// SetDeployedAt gets a reference to the given NullableString and assigns it to the DeployedAt field.
func (o *Function) SetDeployedAt(v string) {
	o.DeployedAt.Set(&v)
}

// SetDeployedAtNil sets the value for DeployedAt to be an explicit nil
func (o *Function) SetDeployedAtNil() {
	o.DeployedAt.Set(nil)
}

// UnsetDeployedAt ensures that no value is present for DeployedAt, not even an explicit nil
func (o *Function) UnsetDeployedAt() {
	o.DeployedAt.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Function) GetSettings() []FunctionSettingV1 {
	if o == nil || o.Settings == nil {
		var ret []FunctionSettingV1
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetSettingsOk() ([]FunctionSettingV1, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Function) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []FunctionSettingV1 and assigns it to the Settings field.
func (o *Function) SetSettings(v []FunctionSettingV1) {
	o.Settings = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Function) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Function) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Function) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Function) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Function) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Function) SetDescription(v string) {
	o.Description = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *Function) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *Function) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *Function) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetPreviewWebhookUrl returns the PreviewWebhookUrl field value if set, zero value otherwise.
func (o *Function) GetPreviewWebhookUrl() string {
	if o == nil || o.PreviewWebhookUrl == nil {
		var ret string
		return ret
	}
	return *o.PreviewWebhookUrl
}

// GetPreviewWebhookUrlOk returns a tuple with the PreviewWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetPreviewWebhookUrlOk() (*string, bool) {
	if o == nil || o.PreviewWebhookUrl == nil {
		return nil, false
	}
	return o.PreviewWebhookUrl, true
}

// HasPreviewWebhookUrl returns a boolean if a field has been set.
func (o *Function) HasPreviewWebhookUrl() bool {
	if o != nil && o.PreviewWebhookUrl != nil {
		return true
	}

	return false
}

// SetPreviewWebhookUrl gets a reference to the given string and assigns it to the PreviewWebhookUrl field.
func (o *Function) SetPreviewWebhookUrl(v string) {
	o.PreviewWebhookUrl = &v
}

// GetBatchMaxCount returns the BatchMaxCount field value if set, zero value otherwise.
func (o *Function) GetBatchMaxCount() float32 {
	if o == nil || o.BatchMaxCount == nil {
		var ret float32
		return ret
	}
	return *o.BatchMaxCount
}

// GetBatchMaxCountOk returns a tuple with the BatchMaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetBatchMaxCountOk() (*float32, bool) {
	if o == nil || o.BatchMaxCount == nil {
		return nil, false
	}
	return o.BatchMaxCount, true
}

// HasBatchMaxCount returns a boolean if a field has been set.
func (o *Function) HasBatchMaxCount() bool {
	if o != nil && o.BatchMaxCount != nil {
		return true
	}

	return false
}

// SetBatchMaxCount gets a reference to the given float32 and assigns it to the BatchMaxCount field.
func (o *Function) SetBatchMaxCount(v float32) {
	o.BatchMaxCount = &v
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *Function) GetCatalogId() string {
	if o == nil || o.CatalogId == nil {
		var ret string
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetCatalogIdOk() (*string, bool) {
	if o == nil || o.CatalogId == nil {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *Function) HasCatalogId() bool {
	if o != nil && o.CatalogId != nil {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given string and assigns it to the CatalogId field.
func (o *Function) SetCatalogId(v string) {
	o.CatalogId = &v
}

// GetIsLatestVersion returns the IsLatestVersion field value if set, zero value otherwise.
func (o *Function) GetIsLatestVersion() bool {
	if o == nil || o.IsLatestVersion == nil {
		var ret bool
		return ret
	}
	return *o.IsLatestVersion
}

// GetIsLatestVersionOk returns a tuple with the IsLatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Function) GetIsLatestVersionOk() (*bool, bool) {
	if o == nil || o.IsLatestVersion == nil {
		return nil, false
	}
	return o.IsLatestVersion, true
}

// HasIsLatestVersion returns a boolean if a field has been set.
func (o *Function) HasIsLatestVersion() bool {
	if o != nil && o.IsLatestVersion != nil {
		return true
	}

	return false
}

// SetIsLatestVersion gets a reference to the given bool and assigns it to the IsLatestVersion field.
func (o *Function) SetIsLatestVersion(v bool) {
	o.IsLatestVersion = &v
}

func (o Function) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.DeployedAt.IsSet() {
		toSerialize["deployedAt"] = o.DeployedAt.Get()
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.LogoUrl != nil {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if o.PreviewWebhookUrl != nil {
		toSerialize["previewWebhookUrl"] = o.PreviewWebhookUrl
	}
	if o.BatchMaxCount != nil {
		toSerialize["batchMaxCount"] = o.BatchMaxCount
	}
	if o.CatalogId != nil {
		toSerialize["catalogId"] = o.CatalogId
	}
	if o.IsLatestVersion != nil {
		toSerialize["isLatestVersion"] = o.IsLatestVersion
	}
	return json.Marshal(toSerialize)
}

type NullableFunction struct {
	value *Function
	isSet bool
}

func (v NullableFunction) Get() *Function {
	return v.value
}

func (v *NullableFunction) Set(val *Function) {
	v.value = val
	v.isSet = true
}

func (v NullableFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunction(val *Function) *NullableFunction {
	return &NullableFunction{value: val, isSet: true}
}

func (v NullableFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
