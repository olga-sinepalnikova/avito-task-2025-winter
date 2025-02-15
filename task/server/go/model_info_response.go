/*
 * API Avito shop
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InfoResponse struct {

	// Количество доступных монет.
	Coins int32 `json:"coins,omitempty"`

	Inventory []InfoResponseInventoryInner `json:"inventory,omitempty"`

	CoinHistory InfoResponseCoinHistory `json:"coinHistory,omitempty"`
}
