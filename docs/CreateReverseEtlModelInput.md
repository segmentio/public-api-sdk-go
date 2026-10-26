# CreateReverseEtlModelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | Indicates which Source to attach this model to. | 
**Name** | **string** | A short, human-readable description of the Model. | 
**Description** | **string** | A longer, more descriptive explanation of the Model. | 
**Enabled** | **bool** | Indicates whether the Model should have syncs enabled. When disabled, no syncs will be triggered, regardless of the enabled status of the attached destinations/subscriptions. | 
**Query** | **string** | The SQL query that will be executed to extract data from the connected Source. | 
**QueryIdentifierColumn** | **string** | Indicates the column named in &#x60;query&#x60; that should be used to uniquely identify the extracted records. | 

## Methods

### NewCreateReverseEtlModelInput

`func NewCreateReverseEtlModelInput(sourceId string, name string, description string, enabled bool, query string, queryIdentifierColumn string, ) *CreateReverseEtlModelInput`

NewCreateReverseEtlModelInput instantiates a new CreateReverseEtlModelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReverseEtlModelInputWithDefaults

`func NewCreateReverseEtlModelInputWithDefaults() *CreateReverseEtlModelInput`

NewCreateReverseEtlModelInputWithDefaults instantiates a new CreateReverseEtlModelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *CreateReverseEtlModelInput) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateReverseEtlModelInput) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateReverseEtlModelInput) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetName

`func (o *CreateReverseEtlModelInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReverseEtlModelInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReverseEtlModelInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateReverseEtlModelInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateReverseEtlModelInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateReverseEtlModelInput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *CreateReverseEtlModelInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateReverseEtlModelInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateReverseEtlModelInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetQuery

`func (o *CreateReverseEtlModelInput) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateReverseEtlModelInput) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateReverseEtlModelInput) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQueryIdentifierColumn

`func (o *CreateReverseEtlModelInput) GetQueryIdentifierColumn() string`

GetQueryIdentifierColumn returns the QueryIdentifierColumn field if non-nil, zero value otherwise.

### GetQueryIdentifierColumnOk

`func (o *CreateReverseEtlModelInput) GetQueryIdentifierColumnOk() (*string, bool)`

GetQueryIdentifierColumnOk returns a tuple with the QueryIdentifierColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIdentifierColumn

`func (o *CreateReverseEtlModelInput) SetQueryIdentifierColumn(v string)`

SetQueryIdentifierColumn sets QueryIdentifierColumn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


