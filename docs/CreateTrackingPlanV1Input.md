# CreateTrackingPlanV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Tracking Plan&#39;s name.  Config API note: equal to &#x60;displayName&#x60;. | 
**Description** | Pointer to **string** | The Tracking Plan&#39;s description. | [optional] 
**Type** | **string** | The Tracking Plan&#39;s type. | 

## Methods

### NewCreateTrackingPlanV1Input

`func NewCreateTrackingPlanV1Input(name string, type_ string, ) *CreateTrackingPlanV1Input`

NewCreateTrackingPlanV1Input instantiates a new CreateTrackingPlanV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTrackingPlanV1InputWithDefaults

`func NewCreateTrackingPlanV1InputWithDefaults() *CreateTrackingPlanV1Input`

NewCreateTrackingPlanV1InputWithDefaults instantiates a new CreateTrackingPlanV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTrackingPlanV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTrackingPlanV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTrackingPlanV1Input) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTrackingPlanV1Input) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTrackingPlanV1Input) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTrackingPlanV1Input) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTrackingPlanV1Input) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateTrackingPlanV1Input) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTrackingPlanV1Input) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTrackingPlanV1Input) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


