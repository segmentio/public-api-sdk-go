# UsersPerSourceSnapshotV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The Source id. | 
**PeriodStart** | **float32** | The start of the usage period, in unix time (seconds since epoch). | 
**PeriodEnd** | **float32** | The end of the usage period, in unix time (seconds since epoch). | 
**Anonymous** | **string** | A bigint of the number of anonymous users in this snapshot. | 
**AnonymousIdentified** | **string** | A bigint of the number of anonymous identified users in this snapshot. | 
**Identified** | **string** | A bigint of the number of identified users in this snapshot. | 
**NeverIdentified** | **string** | A bigint of the number of never identified users in this snapshot. | 
**Timestamp** | **string** | The timestamp of this snapshot within the billing cycle, in the ISO-8601 format. | 

## Methods

### NewUsersPerSourceSnapshotV1

`func NewUsersPerSourceSnapshotV1(sourceId string, periodStart float32, periodEnd float32, anonymous string, anonymousIdentified string, identified string, neverIdentified string, timestamp string, ) *UsersPerSourceSnapshotV1`

NewUsersPerSourceSnapshotV1 instantiates a new UsersPerSourceSnapshotV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPerSourceSnapshotV1WithDefaults

`func NewUsersPerSourceSnapshotV1WithDefaults() *UsersPerSourceSnapshotV1`

NewUsersPerSourceSnapshotV1WithDefaults instantiates a new UsersPerSourceSnapshotV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *UsersPerSourceSnapshotV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UsersPerSourceSnapshotV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UsersPerSourceSnapshotV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetPeriodStart

`func (o *UsersPerSourceSnapshotV1) GetPeriodStart() float32`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UsersPerSourceSnapshotV1) GetPeriodStartOk() (*float32, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UsersPerSourceSnapshotV1) SetPeriodStart(v float32)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *UsersPerSourceSnapshotV1) GetPeriodEnd() float32`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UsersPerSourceSnapshotV1) GetPeriodEndOk() (*float32, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UsersPerSourceSnapshotV1) SetPeriodEnd(v float32)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetAnonymous

`func (o *UsersPerSourceSnapshotV1) GetAnonymous() string`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *UsersPerSourceSnapshotV1) GetAnonymousOk() (*string, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *UsersPerSourceSnapshotV1) SetAnonymous(v string)`

SetAnonymous sets Anonymous field to given value.


### GetAnonymousIdentified

`func (o *UsersPerSourceSnapshotV1) GetAnonymousIdentified() string`

GetAnonymousIdentified returns the AnonymousIdentified field if non-nil, zero value otherwise.

### GetAnonymousIdentifiedOk

`func (o *UsersPerSourceSnapshotV1) GetAnonymousIdentifiedOk() (*string, bool)`

GetAnonymousIdentifiedOk returns a tuple with the AnonymousIdentified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousIdentified

`func (o *UsersPerSourceSnapshotV1) SetAnonymousIdentified(v string)`

SetAnonymousIdentified sets AnonymousIdentified field to given value.


### GetIdentified

`func (o *UsersPerSourceSnapshotV1) GetIdentified() string`

GetIdentified returns the Identified field if non-nil, zero value otherwise.

### GetIdentifiedOk

`func (o *UsersPerSourceSnapshotV1) GetIdentifiedOk() (*string, bool)`

GetIdentifiedOk returns a tuple with the Identified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentified

`func (o *UsersPerSourceSnapshotV1) SetIdentified(v string)`

SetIdentified sets Identified field to given value.


### GetNeverIdentified

`func (o *UsersPerSourceSnapshotV1) GetNeverIdentified() string`

GetNeverIdentified returns the NeverIdentified field if non-nil, zero value otherwise.

### GetNeverIdentifiedOk

`func (o *UsersPerSourceSnapshotV1) GetNeverIdentifiedOk() (*string, bool)`

GetNeverIdentifiedOk returns a tuple with the NeverIdentified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverIdentified

`func (o *UsersPerSourceSnapshotV1) SetNeverIdentified(v string)`

SetNeverIdentified sets NeverIdentified field to given value.


### GetTimestamp

`func (o *UsersPerSourceSnapshotV1) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UsersPerSourceSnapshotV1) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UsersPerSourceSnapshotV1) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


