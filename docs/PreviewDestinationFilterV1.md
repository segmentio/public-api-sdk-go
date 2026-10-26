# PreviewDestinationFilterV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**If** | **string** | A FQL statement which determines if the provided filter&#39;s actions will apply to the provided JSON payload. The literal string \&quot;all\&quot; will result in this filter to all events. For guidance on using FQL, see the Segment documentation site. | 
**Actions** | [**[]DestinationFilterActionV1**](DestinationFilterActionV1.md) | The filtering action to take on events that match the \&quot;if\&quot; statement. Action types must be one of: \&quot;drop\&quot;, \&quot;allow_properties\&quot;, \&quot;drop_properties\&quot; or \&quot;sample\&quot;. | 

## Methods

### NewPreviewDestinationFilterV1

`func NewPreviewDestinationFilterV1(if_ string, actions []DestinationFilterActionV1, ) *PreviewDestinationFilterV1`

NewPreviewDestinationFilterV1 instantiates a new PreviewDestinationFilterV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewDestinationFilterV1WithDefaults

`func NewPreviewDestinationFilterV1WithDefaults() *PreviewDestinationFilterV1`

NewPreviewDestinationFilterV1WithDefaults instantiates a new PreviewDestinationFilterV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIf

`func (o *PreviewDestinationFilterV1) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *PreviewDestinationFilterV1) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *PreviewDestinationFilterV1) SetIf(v string)`

SetIf sets If field to given value.


### GetActions

`func (o *PreviewDestinationFilterV1) GetActions() []DestinationFilterActionV1`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PreviewDestinationFilterV1) GetActionsOk() (*[]DestinationFilterActionV1, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PreviewDestinationFilterV1) SetActions(v []DestinationFilterActionV1)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


