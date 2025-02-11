# InfoResponseCoinHistorySentInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToUser** | Pointer to **string** | Имя пользователя, которому отправлены монеты. | [optional] 
**Amount** | Pointer to **int32** | Количество отправленных монет. | [optional] 

## Methods

### NewInfoResponseCoinHistorySentInner

`func NewInfoResponseCoinHistorySentInner() *InfoResponseCoinHistorySentInner`

NewInfoResponseCoinHistorySentInner instantiates a new InfoResponseCoinHistorySentInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoResponseCoinHistorySentInnerWithDefaults

`func NewInfoResponseCoinHistorySentInnerWithDefaults() *InfoResponseCoinHistorySentInner`

NewInfoResponseCoinHistorySentInnerWithDefaults instantiates a new InfoResponseCoinHistorySentInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToUser

`func (o *InfoResponseCoinHistorySentInner) GetToUser() string`

GetToUser returns the ToUser field if non-nil, zero value otherwise.

### GetToUserOk

`func (o *InfoResponseCoinHistorySentInner) GetToUserOk() (*string, bool)`

GetToUserOk returns a tuple with the ToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUser

`func (o *InfoResponseCoinHistorySentInner) SetToUser(v string)`

SetToUser sets ToUser field to given value.

### HasToUser

`func (o *InfoResponseCoinHistorySentInner) HasToUser() bool`

HasToUser returns a boolean if a field has been set.

### GetAmount

`func (o *InfoResponseCoinHistorySentInner) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InfoResponseCoinHistorySentInner) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InfoResponseCoinHistorySentInner) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InfoResponseCoinHistorySentInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


