/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DestinationMetadata The catalog item matched by id.
type DestinationMetadata struct {
	// The id of the Destination metadata.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The user-friendly name of the Destination.  Config API note: equal to `displayName`.
	Name string `json:"name"`
	// The description of the Destination.
	Description string `json:"description"`
	// The slug used to identify the Destination in the Segment app.
	Slug  string `json:"slug"`
	Logos Logos  `json:"logos"`
	// Options configured for the Destination.
	Options []IntegrationOptionBeta `json:"options"`
	// Support status of the Destination.
	Status string `json:"status"`
	// A list of names previously used by the Destination.
	PreviousNames []string `json:"previousNames"`
	// A list of categories with which the Destination is associated.
	Categories []string `json:"categories"`
	// A website URL for this Destination.
	Website string `json:"website"`
	// A list of components this Destination provides.
	Components         []DestinationMetadataComponentV1 `json:"components"`
	SupportedFeatures  SupportedFeatures                `json:"supportedFeatures"`
	SupportedMethods   SupportedMethods                 `json:"supportedMethods"`
	SupportedPlatforms SupportedPlatforms               `json:"supportedPlatforms"`
	// Actions available for the Destination.
	Actions []DestinationMetadataActionV1 `json:"actions"`
	// Predefined Destination subscriptions that can optionally be applied when connecting a new instance of the Destination.
	Presets []DestinationMetadataSubscriptionPresetV1 `json:"presets"`
	// Contact info for Integration Owners.
	Contacts []Contact `json:"contacts,omitempty"`
	// Partner Owned flag.
	PartnerOwned *bool `json:"partnerOwned,omitempty"`
	// A list of supported regions for this Destination.
	SupportedRegions []string `json:"supportedRegions,omitempty"`
	// The list of regional endpoints for this Destination.
	RegionEndpoints []string `json:"regionEndpoints,omitempty"`
}

// NewDestinationMetadata instantiates a new DestinationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadata(
	id string,
	name string,
	description string,
	slug string,
	logos Logos,
	options []IntegrationOptionBeta,
	status string,
	previousNames []string,
	categories []string,
	website string,
	components []DestinationMetadataComponentV1,
	supportedFeatures SupportedFeatures,
	supportedMethods SupportedMethods,
	supportedPlatforms SupportedPlatforms,
	actions []DestinationMetadataActionV1,
	presets []DestinationMetadataSubscriptionPresetV1,
) *DestinationMetadata {
	this := DestinationMetadata{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Slug = slug
	this.Logos = logos
	this.Options = options
	this.Status = status
	this.PreviousNames = previousNames
	this.Categories = categories
	this.Website = website
	this.Components = components
	this.SupportedFeatures = supportedFeatures
	this.SupportedMethods = supportedMethods
	this.SupportedPlatforms = supportedPlatforms
	this.Actions = actions
	this.Presets = presets
	return &this
}

// NewDestinationMetadataWithDefaults instantiates a new DestinationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataWithDefaults() *DestinationMetadata {
	this := DestinationMetadata{}
	return &this
}

// GetId returns the Id field value
func (o *DestinationMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DestinationMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DestinationMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DestinationMetadata) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *DestinationMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DestinationMetadata) SetDescription(v string) {
	o.Description = v
}

// GetSlug returns the Slug field value
func (o *DestinationMetadata) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DestinationMetadata) SetSlug(v string) {
	o.Slug = v
}

// GetLogos returns the Logos field value
func (o *DestinationMetadata) GetLogos() Logos {
	if o == nil {
		var ret Logos
		return ret
	}

	return o.Logos
}

// GetLogosOk returns a tuple with the Logos field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetLogosOk() (*Logos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logos, true
}

// SetLogos sets field value
func (o *DestinationMetadata) SetLogos(v Logos) {
	o.Logos = v
}

// GetOptions returns the Options field value
func (o *DestinationMetadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		var ret []IntegrationOptionBeta
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetOptionsOk() ([]IntegrationOptionBeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *DestinationMetadata) SetOptions(v []IntegrationOptionBeta) {
	o.Options = v
}

// GetStatus returns the Status field value
func (o *DestinationMetadata) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DestinationMetadata) SetStatus(v string) {
	o.Status = v
}

// GetPreviousNames returns the PreviousNames field value
func (o *DestinationMetadata) GetPreviousNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PreviousNames
}

// GetPreviousNamesOk returns a tuple with the PreviousNames field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetPreviousNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousNames, true
}

// SetPreviousNames sets field value
func (o *DestinationMetadata) SetPreviousNames(v []string) {
	o.PreviousNames = v
}

// GetCategories returns the Categories field value
func (o *DestinationMetadata) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetCategoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *DestinationMetadata) SetCategories(v []string) {
	o.Categories = v
}

// GetWebsite returns the Website field value
func (o *DestinationMetadata) GetWebsite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Website
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Website, true
}

// SetWebsite sets field value
func (o *DestinationMetadata) SetWebsite(v string) {
	o.Website = v
}

// GetComponents returns the Components field value
func (o *DestinationMetadata) GetComponents() []DestinationMetadataComponentV1 {
	if o == nil {
		var ret []DestinationMetadataComponentV1
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetComponentsOk() ([]DestinationMetadataComponentV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *DestinationMetadata) SetComponents(v []DestinationMetadataComponentV1) {
	o.Components = v
}

// GetSupportedFeatures returns the SupportedFeatures field value
func (o *DestinationMetadata) GetSupportedFeatures() SupportedFeatures {
	if o == nil {
		var ret SupportedFeatures
		return ret
	}

	return o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetSupportedFeaturesOk() (*SupportedFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedFeatures, true
}

// SetSupportedFeatures sets field value
func (o *DestinationMetadata) SetSupportedFeatures(v SupportedFeatures) {
	o.SupportedFeatures = v
}

// GetSupportedMethods returns the SupportedMethods field value
func (o *DestinationMetadata) GetSupportedMethods() SupportedMethods {
	if o == nil {
		var ret SupportedMethods
		return ret
	}

	return o.SupportedMethods
}

// GetSupportedMethodsOk returns a tuple with the SupportedMethods field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetSupportedMethodsOk() (*SupportedMethods, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedMethods, true
}

// SetSupportedMethods sets field value
func (o *DestinationMetadata) SetSupportedMethods(v SupportedMethods) {
	o.SupportedMethods = v
}

// GetSupportedPlatforms returns the SupportedPlatforms field value
func (o *DestinationMetadata) GetSupportedPlatforms() SupportedPlatforms {
	if o == nil {
		var ret SupportedPlatforms
		return ret
	}

	return o.SupportedPlatforms
}

// GetSupportedPlatformsOk returns a tuple with the SupportedPlatforms field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetSupportedPlatformsOk() (*SupportedPlatforms, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedPlatforms, true
}

// SetSupportedPlatforms sets field value
func (o *DestinationMetadata) SetSupportedPlatforms(v SupportedPlatforms) {
	o.SupportedPlatforms = v
}

// GetActions returns the Actions field value
func (o *DestinationMetadata) GetActions() []DestinationMetadataActionV1 {
	if o == nil {
		var ret []DestinationMetadataActionV1
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetActionsOk() ([]DestinationMetadataActionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *DestinationMetadata) SetActions(v []DestinationMetadataActionV1) {
	o.Actions = v
}

// GetPresets returns the Presets field value
func (o *DestinationMetadata) GetPresets() []DestinationMetadataSubscriptionPresetV1 {
	if o == nil {
		var ret []DestinationMetadataSubscriptionPresetV1
		return ret
	}

	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetPresetsOk() ([]DestinationMetadataSubscriptionPresetV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Presets, true
}

// SetPresets sets field value
func (o *DestinationMetadata) SetPresets(v []DestinationMetadataSubscriptionPresetV1) {
	o.Presets = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *DestinationMetadata) GetContacts() []Contact {
	if o == nil || o.Contacts == nil {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetContactsOk() ([]Contact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *DestinationMetadata) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *DestinationMetadata) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetPartnerOwned returns the PartnerOwned field value if set, zero value otherwise.
func (o *DestinationMetadata) GetPartnerOwned() bool {
	if o == nil || o.PartnerOwned == nil {
		var ret bool
		return ret
	}
	return *o.PartnerOwned
}

// GetPartnerOwnedOk returns a tuple with the PartnerOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetPartnerOwnedOk() (*bool, bool) {
	if o == nil || o.PartnerOwned == nil {
		return nil, false
	}
	return o.PartnerOwned, true
}

// HasPartnerOwned returns a boolean if a field has been set.
func (o *DestinationMetadata) HasPartnerOwned() bool {
	if o != nil && o.PartnerOwned != nil {
		return true
	}

	return false
}

// SetPartnerOwned gets a reference to the given bool and assigns it to the PartnerOwned field.
func (o *DestinationMetadata) SetPartnerOwned(v bool) {
	o.PartnerOwned = &v
}

// GetSupportedRegions returns the SupportedRegions field value if set, zero value otherwise.
func (o *DestinationMetadata) GetSupportedRegions() []string {
	if o == nil || o.SupportedRegions == nil {
		var ret []string
		return ret
	}
	return o.SupportedRegions
}

// GetSupportedRegionsOk returns a tuple with the SupportedRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetSupportedRegionsOk() ([]string, bool) {
	if o == nil || o.SupportedRegions == nil {
		return nil, false
	}
	return o.SupportedRegions, true
}

// HasSupportedRegions returns a boolean if a field has been set.
func (o *DestinationMetadata) HasSupportedRegions() bool {
	if o != nil && o.SupportedRegions != nil {
		return true
	}

	return false
}

// SetSupportedRegions gets a reference to the given []string and assigns it to the SupportedRegions field.
func (o *DestinationMetadata) SetSupportedRegions(v []string) {
	o.SupportedRegions = v
}

// GetRegionEndpoints returns the RegionEndpoints field value if set, zero value otherwise.
func (o *DestinationMetadata) GetRegionEndpoints() []string {
	if o == nil || o.RegionEndpoints == nil {
		var ret []string
		return ret
	}
	return o.RegionEndpoints
}

// GetRegionEndpointsOk returns a tuple with the RegionEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadata) GetRegionEndpointsOk() ([]string, bool) {
	if o == nil || o.RegionEndpoints == nil {
		return nil, false
	}
	return o.RegionEndpoints, true
}

// HasRegionEndpoints returns a boolean if a field has been set.
func (o *DestinationMetadata) HasRegionEndpoints() bool {
	if o != nil && o.RegionEndpoints != nil {
		return true
	}

	return false
}

// SetRegionEndpoints gets a reference to the given []string and assigns it to the RegionEndpoints field.
func (o *DestinationMetadata) SetRegionEndpoints(v []string) {
	o.RegionEndpoints = v
}

func (o DestinationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["logos"] = o.Logos
	}
	if true {
		toSerialize["options"] = o.Options
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["previousNames"] = o.PreviousNames
	}
	if true {
		toSerialize["categories"] = o.Categories
	}
	if true {
		toSerialize["website"] = o.Website
	}
	if true {
		toSerialize["components"] = o.Components
	}
	if true {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if true {
		toSerialize["supportedMethods"] = o.SupportedMethods
	}
	if true {
		toSerialize["supportedPlatforms"] = o.SupportedPlatforms
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if true {
		toSerialize["presets"] = o.Presets
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if o.PartnerOwned != nil {
		toSerialize["partnerOwned"] = o.PartnerOwned
	}
	if o.SupportedRegions != nil {
		toSerialize["supportedRegions"] = o.SupportedRegions
	}
	if o.RegionEndpoints != nil {
		toSerialize["regionEndpoints"] = o.RegionEndpoints
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationMetadata struct {
	value *DestinationMetadata
	isSet bool
}

func (v NullableDestinationMetadata) Get() *DestinationMetadata {
	return v.value
}

func (v *NullableDestinationMetadata) Set(val *DestinationMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadata(val *DestinationMetadata) *NullableDestinationMetadata {
	return &NullableDestinationMetadata{value: val, isSet: true}
}

func (v NullableDestinationMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
