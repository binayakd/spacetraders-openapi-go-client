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

// checks if the JumpShip200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JumpShip200ResponseData{}

// JumpShip200ResponseData struct for JumpShip200ResponseData
type JumpShip200ResponseData struct {
	Cooldown Cooldown `json:"cooldown"`
	Nav *ShipNav `json:"nav,omitempty"`
}

// NewJumpShip200ResponseData instantiates a new JumpShip200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJumpShip200ResponseData(cooldown Cooldown) *JumpShip200ResponseData {
	this := JumpShip200ResponseData{}
	this.Cooldown = cooldown
	return &this
}

// NewJumpShip200ResponseDataWithDefaults instantiates a new JumpShip200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJumpShip200ResponseDataWithDefaults() *JumpShip200ResponseData {
	this := JumpShip200ResponseData{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *JumpShip200ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *JumpShip200ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *JumpShip200ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetNav returns the Nav field value if set, zero value otherwise.
func (o *JumpShip200ResponseData) GetNav() ShipNav {
	if o == nil || IsNil(o.Nav) {
		var ret ShipNav
		return ret
	}
	return *o.Nav
}

// GetNavOk returns a tuple with the Nav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JumpShip200ResponseData) GetNavOk() (*ShipNav, bool) {
	if o == nil || IsNil(o.Nav) {
		return nil, false
	}
	return o.Nav, true
}

// HasNav returns a boolean if a field has been set.
func (o *JumpShip200ResponseData) HasNav() bool {
	if o != nil && !IsNil(o.Nav) {
		return true
	}

	return false
}

// SetNav gets a reference to the given ShipNav and assigns it to the Nav field.
func (o *JumpShip200ResponseData) SetNav(v ShipNav) {
	o.Nav = &v
}

func (o JumpShip200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JumpShip200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	if !IsNil(o.Nav) {
		toSerialize["nav"] = o.Nav
	}
	return toSerialize, nil
}

type NullableJumpShip200ResponseData struct {
	value *JumpShip200ResponseData
	isSet bool
}

func (v NullableJumpShip200ResponseData) Get() *JumpShip200ResponseData {
	return v.value
}

func (v *NullableJumpShip200ResponseData) Set(val *JumpShip200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableJumpShip200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableJumpShip200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJumpShip200ResponseData(val *JumpShip200ResponseData) *NullableJumpShip200ResponseData {
	return &NullableJumpShip200ResponseData{value: val, isSet: true}
}

func (v NullableJumpShip200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJumpShip200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


