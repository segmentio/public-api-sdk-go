# DbtModelSyncTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the dbt model sync. | 
**SourceId** | Pointer to **string** | The Source id that was triggered. | [optional] 
**Status** | **string** | The status of the trigger. | 

## Methods

### NewDbtModelSyncTrigger

`func NewDbtModelSyncTrigger(id string, status string, ) *DbtModelSyncTrigger`

NewDbtModelSyncTrigger instantiates a new DbtModelSyncTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtModelSyncTriggerWithDefaults

`func NewDbtModelSyncTriggerWithDefaults() *DbtModelSyncTrigger`

NewDbtModelSyncTriggerWithDefaults instantiates a new DbtModelSyncTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbtModelSyncTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbtModelSyncTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbtModelSyncTrigger) SetId(v string)`

SetId sets Id field to given value.


### GetSourceId

`func (o *DbtModelSyncTrigger) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DbtModelSyncTrigger) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DbtModelSyncTrigger) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DbtModelSyncTrigger) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetStatus

`func (o *DbtModelSyncTrigger) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DbtModelSyncTrigger) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DbtModelSyncTrigger) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


