# DestinationMetadataFeaturesV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudModeInstances** | Pointer to **string** | This Destination&#39;s support level for cloud mode instances. The values &#39;0&#39; and &#39;NONE&#39;, and &#39;1&#39; and &#39;SINGLE&#39; are equivalent. | [optional] 
**DeviceModeInstances** | Pointer to **string** | This Destination&#39;s support level for device mode instances. Support for multiple device mode instances is currently not planned. The values &#39;0&#39; and &#39;NONE&#39;, and &#39;1&#39; and &#39;SINGLE&#39; are equivalent. | [optional] 
**Replay** | Pointer to **bool** | Whether this Destination supports replays. | [optional] 
**BrowserUnbundling** | Pointer to **bool** | Whether this Destination supports browser unbundling. | [optional] 
**BrowserUnbundlingPublic** | Pointer to **bool** | Whether this Destination supports public browser unbundling. | [optional] 

## Methods

### NewDestinationMetadataFeaturesV1

`func NewDestinationMetadataFeaturesV1() *DestinationMetadataFeaturesV1`

NewDestinationMetadataFeaturesV1 instantiates a new DestinationMetadataFeaturesV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationMetadataFeaturesV1WithDefaults

`func NewDestinationMetadataFeaturesV1WithDefaults() *DestinationMetadataFeaturesV1`

NewDestinationMetadataFeaturesV1WithDefaults instantiates a new DestinationMetadataFeaturesV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudModeInstances

`func (o *DestinationMetadataFeaturesV1) GetCloudModeInstances() string`

GetCloudModeInstances returns the CloudModeInstances field if non-nil, zero value otherwise.

### GetCloudModeInstancesOk

`func (o *DestinationMetadataFeaturesV1) GetCloudModeInstancesOk() (*string, bool)`

GetCloudModeInstancesOk returns a tuple with the CloudModeInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudModeInstances

`func (o *DestinationMetadataFeaturesV1) SetCloudModeInstances(v string)`

SetCloudModeInstances sets CloudModeInstances field to given value.

### HasCloudModeInstances

`func (o *DestinationMetadataFeaturesV1) HasCloudModeInstances() bool`

HasCloudModeInstances returns a boolean if a field has been set.

### GetDeviceModeInstances

`func (o *DestinationMetadataFeaturesV1) GetDeviceModeInstances() string`

GetDeviceModeInstances returns the DeviceModeInstances field if non-nil, zero value otherwise.

### GetDeviceModeInstancesOk

`func (o *DestinationMetadataFeaturesV1) GetDeviceModeInstancesOk() (*string, bool)`

GetDeviceModeInstancesOk returns a tuple with the DeviceModeInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModeInstances

`func (o *DestinationMetadataFeaturesV1) SetDeviceModeInstances(v string)`

SetDeviceModeInstances sets DeviceModeInstances field to given value.

### HasDeviceModeInstances

`func (o *DestinationMetadataFeaturesV1) HasDeviceModeInstances() bool`

HasDeviceModeInstances returns a boolean if a field has been set.

### GetReplay

`func (o *DestinationMetadataFeaturesV1) GetReplay() bool`

GetReplay returns the Replay field if non-nil, zero value otherwise.

### GetReplayOk

`func (o *DestinationMetadataFeaturesV1) GetReplayOk() (*bool, bool)`

GetReplayOk returns a tuple with the Replay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplay

`func (o *DestinationMetadataFeaturesV1) SetReplay(v bool)`

SetReplay sets Replay field to given value.

### HasReplay

`func (o *DestinationMetadataFeaturesV1) HasReplay() bool`

HasReplay returns a boolean if a field has been set.

### GetBrowserUnbundling

`func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundling() bool`

GetBrowserUnbundling returns the BrowserUnbundling field if non-nil, zero value otherwise.

### GetBrowserUnbundlingOk

`func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundlingOk() (*bool, bool)`

GetBrowserUnbundlingOk returns a tuple with the BrowserUnbundling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUnbundling

`func (o *DestinationMetadataFeaturesV1) SetBrowserUnbundling(v bool)`

SetBrowserUnbundling sets BrowserUnbundling field to given value.

### HasBrowserUnbundling

`func (o *DestinationMetadataFeaturesV1) HasBrowserUnbundling() bool`

HasBrowserUnbundling returns a boolean if a field has been set.

### GetBrowserUnbundlingPublic

`func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundlingPublic() bool`

GetBrowserUnbundlingPublic returns the BrowserUnbundlingPublic field if non-nil, zero value otherwise.

### GetBrowserUnbundlingPublicOk

`func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundlingPublicOk() (*bool, bool)`

GetBrowserUnbundlingPublicOk returns a tuple with the BrowserUnbundlingPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUnbundlingPublic

`func (o *DestinationMetadataFeaturesV1) SetBrowserUnbundlingPublic(v bool)`

SetBrowserUnbundlingPublic sets BrowserUnbundlingPublic field to given value.

### HasBrowserUnbundlingPublic

`func (o *DestinationMetadataFeaturesV1) HasBrowserUnbundlingPublic() bool`

HasBrowserUnbundlingPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


