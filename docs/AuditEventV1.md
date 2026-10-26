# AuditEventV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this audit trail event. | 
**Timestamp** | **string** | The timestamp of this event in ISO-8601 format. | 
**Type** | **string** | The type of this event. | 
**Actor** | **string** | The user or API token that triggered this event. | 
**ActorEmail** | Pointer to **string** | The email of the user that triggered this event. | [optional] 
**ResourceId** | **string** | The identifier of the resource affected by this event. | 
**ResourceType** | **string** | The kind of resource affected by this event. | 
**ResourceName** | **string** | The name of the resource affected by this event. | 

## Methods

### NewAuditEventV1

`func NewAuditEventV1(id string, timestamp string, type_ string, actor string, resourceId string, resourceType string, resourceName string, ) *AuditEventV1`

NewAuditEventV1 instantiates a new AuditEventV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventV1WithDefaults

`func NewAuditEventV1WithDefaults() *AuditEventV1`

NewAuditEventV1WithDefaults instantiates a new AuditEventV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditEventV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditEventV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditEventV1) SetId(v string)`

SetId sets Id field to given value.


### GetTimestamp

`func (o *AuditEventV1) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditEventV1) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditEventV1) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *AuditEventV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditEventV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditEventV1) SetType(v string)`

SetType sets Type field to given value.


### GetActor

`func (o *AuditEventV1) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AuditEventV1) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AuditEventV1) SetActor(v string)`

SetActor sets Actor field to given value.


### GetActorEmail

`func (o *AuditEventV1) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *AuditEventV1) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *AuditEventV1) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.

### HasActorEmail

`func (o *AuditEventV1) HasActorEmail() bool`

HasActorEmail returns a boolean if a field has been set.

### GetResourceId

`func (o *AuditEventV1) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuditEventV1) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuditEventV1) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *AuditEventV1) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AuditEventV1) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AuditEventV1) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceName

`func (o *AuditEventV1) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AuditEventV1) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AuditEventV1) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


