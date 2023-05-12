/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetraders_openapi_go_client

import (
	"encoding/json"
)

// checks if the ShipModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipModule{}

// ShipModule A module can be installed in a ship and provides a set of capabilities such as storage space or quarters for crew. Module installations are permanent.
type ShipModule struct {
	Symbol string `json:"symbol"`
	Capacity *int32 `json:"capacity,omitempty"`
	Range *int32 `json:"range,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Requirements ShipRequirements `json:"requirements"`
}

// NewShipModule instantiates a new ShipModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipModule(symbol string, name string, requirements ShipRequirements) *ShipModule {
	this := ShipModule{}
	this.Symbol = symbol
	this.Name = name
	this.Requirements = requirements
	return &this
}

// NewShipModuleWithDefaults instantiates a new ShipModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipModuleWithDefaults() *ShipModule {
	this := ShipModule{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipModule) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipModule) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipModule) SetSymbol(v string) {
	o.Symbol = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *ShipModule) GetCapacity() int32 {
	if o == nil || IsNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipModule) GetCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *ShipModule) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *ShipModule) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ShipModule) GetRange() int32 {
	if o == nil || IsNil(o.Range) {
		var ret int32
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipModule) GetRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ShipModule) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given int32 and assigns it to the Range field.
func (o *ShipModule) SetRange(v int32) {
	o.Range = &v
}

// GetName returns the Name field value
func (o *ShipModule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipModule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipModule) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ShipModule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipModule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ShipModule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ShipModule) SetDescription(v string) {
	o.Description = &v
}

// GetRequirements returns the Requirements field value
func (o *ShipModule) GetRequirements() ShipRequirements {
	if o == nil {
		var ret ShipRequirements
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *ShipModule) GetRequirementsOk() (*ShipRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *ShipModule) SetRequirements(v ShipRequirements) {
	o.Requirements = v
}

func (o ShipModule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["requirements"] = o.Requirements
	return toSerialize, nil
}

type NullableShipModule struct {
	value *ShipModule
	isSet bool
}

func (v NullableShipModule) Get() *ShipModule {
	return v.value
}

func (v *NullableShipModule) Set(val *ShipModule) {
	v.value = val
	v.isSet = true
}

func (v NullableShipModule) IsSet() bool {
	return v.isSet
}

func (v *NullableShipModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipModule(val *ShipModule) *NullableShipModule {
	return &NullableShipModule{value: val, isSet: true}
}

func (v NullableShipModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


