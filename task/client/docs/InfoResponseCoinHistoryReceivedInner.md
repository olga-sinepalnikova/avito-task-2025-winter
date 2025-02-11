# InfoResponseCoinHistoryReceivedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromUser** | Pointer to **string** | Имя пользователя, который отправил монеты. | [optional] 
**Amount** | Pointer to **int32** | Количество полученных монет. | [optional] 

## Methods

### NewInfoResponseCoinHistoryReceivedInner

`func NewInfoResponseCoinHistoryReceivedInner() *InfoResponseCoinHistoryReceivedInner`

NewInfoResponseCoinHistoryReceivedInner instantiates a new InfoResponseCoinHistoryReceivedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoResponseCoinHistoryReceivedInnerWithDefaults

`func NewInfoResponseCoinHistoryReceivedInnerWithDefaults() *InfoResponseCoinHistoryReceivedInner`

NewInfoResponseCoinHistoryReceivedInnerWithDefaults instantiates a new InfoResponseCoinHistoryReceivedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromUser

`func (o *InfoResponseCoinHistoryReceivedInner) GetFromUser() string`

GetFromUser returns the FromUser field if non-nil, zero value otherwise.

### GetFromUserOk

`func (o *InfoResponseCoinHistoryReceivedInner) GetFromUserOk() (*string, bool)`

GetFromUserOk returns a tuple with the FromUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUser

`func (o *InfoResponseCoinHistoryReceivedInner) SetFromUser(v string)`

SetFromUser sets FromUser field to given value.

### HasFromUser

`func (o *InfoResponseCoinHistoryReceivedInner) HasFromUser() bool`

HasFromUser returns a boolean if a field has been set.

### GetAmount

`func (o *InfoResponseCoinHistoryReceivedInner) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InfoResponseCoinHistoryReceivedInner) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InfoResponseCoinHistoryReceivedInner) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InfoResponseCoinHistoryReceivedInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


