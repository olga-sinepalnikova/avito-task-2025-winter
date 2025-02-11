# InfoResponseCoinHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | Pointer to [**[]InfoResponseCoinHistoryReceivedInner**](InfoResponseCoinHistoryReceivedInner.md) |  | [optional] 
**Sent** | Pointer to [**[]InfoResponseCoinHistorySentInner**](InfoResponseCoinHistorySentInner.md) |  | [optional] 

## Methods

### NewInfoResponseCoinHistory

`func NewInfoResponseCoinHistory() *InfoResponseCoinHistory`

NewInfoResponseCoinHistory instantiates a new InfoResponseCoinHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoResponseCoinHistoryWithDefaults

`func NewInfoResponseCoinHistoryWithDefaults() *InfoResponseCoinHistory`

NewInfoResponseCoinHistoryWithDefaults instantiates a new InfoResponseCoinHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *InfoResponseCoinHistory) GetReceived() []InfoResponseCoinHistoryReceivedInner`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *InfoResponseCoinHistory) GetReceivedOk() (*[]InfoResponseCoinHistoryReceivedInner, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *InfoResponseCoinHistory) SetReceived(v []InfoResponseCoinHistoryReceivedInner)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *InfoResponseCoinHistory) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSent

`func (o *InfoResponseCoinHistory) GetSent() []InfoResponseCoinHistorySentInner`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InfoResponseCoinHistory) GetSentOk() (*[]InfoResponseCoinHistorySentInner, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InfoResponseCoinHistory) SetSent(v []InfoResponseCoinHistorySentInner)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InfoResponseCoinHistory) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


