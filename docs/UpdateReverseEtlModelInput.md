# UpdateReverseEtlModelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A short, human-readable description of the Model. | [optional] 
**Description** | Pointer to **string** | A longer, more descriptive explanation of the Model. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Model should have syncs enabled. When disabled, no syncs will be triggered, regardless of the enabled status of the attached destinations/subscriptions. | [optional] 
**Query** | Pointer to **string** | The SQL query that will be executed to extract data from the connected Source. | [optional] 
**QueryIdentifierColumn** | Pointer to **string** | Indicates the column named in &#x60;query&#x60; that should be used to uniquely identify the extracted records. | [optional] 

## Methods

### NewUpdateReverseEtlModelInput

`func NewUpdateReverseEtlModelInput() *UpdateReverseEtlModelInput`

NewUpdateReverseEtlModelInput instantiates a new UpdateReverseEtlModelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReverseEtlModelInputWithDefaults

`func NewUpdateReverseEtlModelInputWithDefaults() *UpdateReverseEtlModelInput`

NewUpdateReverseEtlModelInputWithDefaults instantiates a new UpdateReverseEtlModelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateReverseEtlModelInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateReverseEtlModelInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateReverseEtlModelInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateReverseEtlModelInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateReverseEtlModelInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateReverseEtlModelInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateReverseEtlModelInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateReverseEtlModelInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateReverseEtlModelInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateReverseEtlModelInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateReverseEtlModelInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateReverseEtlModelInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetQuery

`func (o *UpdateReverseEtlModelInput) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UpdateReverseEtlModelInput) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UpdateReverseEtlModelInput) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *UpdateReverseEtlModelInput) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryIdentifierColumn

`func (o *UpdateReverseEtlModelInput) GetQueryIdentifierColumn() string`

GetQueryIdentifierColumn returns the QueryIdentifierColumn field if non-nil, zero value otherwise.

### GetQueryIdentifierColumnOk

`func (o *UpdateReverseEtlModelInput) GetQueryIdentifierColumnOk() (*string, bool)`

GetQueryIdentifierColumnOk returns a tuple with the QueryIdentifierColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIdentifierColumn

`func (o *UpdateReverseEtlModelInput) SetQueryIdentifierColumn(v string)`

SetQueryIdentifierColumn sets QueryIdentifierColumn field to given value.

### HasQueryIdentifierColumn

`func (o *UpdateReverseEtlModelInput) HasQueryIdentifierColumn() bool`

HasQueryIdentifierColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


