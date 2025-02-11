# InfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coins** | Pointer to **int32** | Количество доступных монет. | [optional] 
**Inventory** | Pointer to [**[]InfoResponseInventoryInner**](InfoResponseInventoryInner.md) |  | [optional] 
**CoinHistory** | Pointer to [**InfoResponseCoinHistory**](InfoResponseCoinHistory.md) |  | [optional] 

## Methods

### NewInfoResponse

`func NewInfoResponse() *InfoResponse`

NewInfoResponse instantiates a new InfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoResponseWithDefaults

`func NewInfoResponseWithDefaults() *InfoResponse`

NewInfoResponseWithDefaults instantiates a new InfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoins

`func (o *InfoResponse) GetCoins() int32`

GetCoins returns the Coins field if non-nil, zero value otherwise.

### GetCoinsOk

`func (o *InfoResponse) GetCoinsOk() (*int32, bool)`

GetCoinsOk returns a tuple with the Coins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoins

`func (o *InfoResponse) SetCoins(v int32)`

SetCoins sets Coins field to given value.

### HasCoins

`func (o *InfoResponse) HasCoins() bool`

HasCoins returns a boolean if a field has been set.

### GetInventory

`func (o *InfoResponse) GetInventory() []InfoResponseInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *InfoResponse) GetInventoryOk() (*[]InfoResponseInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *InfoResponse) SetInventory(v []InfoResponseInventoryInner)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *InfoResponse) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetCoinHistory

`func (o *InfoResponse) GetCoinHistory() InfoResponseCoinHistory`

GetCoinHistory returns the CoinHistory field if non-nil, zero value otherwise.

### GetCoinHistoryOk

`func (o *InfoResponse) GetCoinHistoryOk() (*InfoResponseCoinHistory, bool)`

GetCoinHistoryOk returns a tuple with the CoinHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinHistory

`func (o *InfoResponse) SetCoinHistory(v InfoResponseCoinHistory)`

SetCoinHistory sets CoinHistory field to given value.

### HasCoinHistory

`func (o *InfoResponse) HasCoinHistory() bool`

HasCoinHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


