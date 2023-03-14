/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Metadata The metadata of the Destination of which this Destination is an instance of. For example, Google Analytics or Amplitude.
type Metadata struct {
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

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata(
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
) *Metadata {
	this := Metadata{}
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

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetId returns the Id field value
func (o *Metadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Metadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Metadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Metadata) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Metadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Metadata) SetDescription(v string) {
	o.Description = v
}

// GetSlug returns the Slug field value
func (o *Metadata) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Metadata) SetSlug(v string) {
	o.Slug = v
}

// GetLogos returns the Logos field value
func (o *Metadata) GetLogos() Logos {
	if o == nil {
		var ret Logos
		return ret
	}

	return o.Logos
}

// GetLogosOk returns a tuple with the Logos field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetLogosOk() (*Logos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logos, true
}

// SetLogos sets field value
func (o *Metadata) SetLogos(v Logos) {
	o.Logos = v
}

// GetOptions returns the Options field value
func (o *Metadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		var ret []IntegrationOptionBeta
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetOptionsOk() ([]IntegrationOptionBeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *Metadata) SetOptions(v []IntegrationOptionBeta) {
	o.Options = v
}

// GetStatus returns the Status field value
func (o *Metadata) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Metadata) SetStatus(v string) {
	o.Status = v
}

// GetPreviousNames returns the PreviousNames field value
func (o *Metadata) GetPreviousNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PreviousNames
}

// GetPreviousNamesOk returns a tuple with the PreviousNames field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetPreviousNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousNames, true
}

// SetPreviousNames sets field value
func (o *Metadata) SetPreviousNames(v []string) {
	o.PreviousNames = v
}

// GetCategories returns the Categories field value
func (o *Metadata) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetCategoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *Metadata) SetCategories(v []string) {
	o.Categories = v
}

// GetWebsite returns the Website field value
func (o *Metadata) GetWebsite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Website
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Website, true
}

// SetWebsite sets field value
func (o *Metadata) SetWebsite(v string) {
	o.Website = v
}

// GetComponents returns the Components field value
func (o *Metadata) GetComponents() []DestinationMetadataComponentV1 {
	if o == nil {
		var ret []DestinationMetadataComponentV1
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetComponentsOk() ([]DestinationMetadataComponentV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *Metadata) SetComponents(v []DestinationMetadataComponentV1) {
	o.Components = v
}

// GetSupportedFeatures returns the SupportedFeatures field value
func (o *Metadata) GetSupportedFeatures() SupportedFeatures {
	if o == nil {
		var ret SupportedFeatures
		return ret
	}

	return o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetSupportedFeaturesOk() (*SupportedFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedFeatures, true
}

// SetSupportedFeatures sets field value
func (o *Metadata) SetSupportedFeatures(v SupportedFeatures) {
	o.SupportedFeatures = v
}

// GetSupportedMethods returns the SupportedMethods field value
func (o *Metadata) GetSupportedMethods() SupportedMethods {
	if o == nil {
		var ret SupportedMethods
		return ret
	}

	return o.SupportedMethods
}

// GetSupportedMethodsOk returns a tuple with the SupportedMethods field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetSupportedMethodsOk() (*SupportedMethods, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedMethods, true
}

// SetSupportedMethods sets field value
func (o *Metadata) SetSupportedMethods(v SupportedMethods) {
	o.SupportedMethods = v
}

// GetSupportedPlatforms returns the SupportedPlatforms field value
func (o *Metadata) GetSupportedPlatforms() SupportedPlatforms {
	if o == nil {
		var ret SupportedPlatforms
		return ret
	}

	return o.SupportedPlatforms
}

// GetSupportedPlatformsOk returns a tuple with the SupportedPlatforms field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetSupportedPlatformsOk() (*SupportedPlatforms, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedPlatforms, true
}

// SetSupportedPlatforms sets field value
func (o *Metadata) SetSupportedPlatforms(v SupportedPlatforms) {
	o.SupportedPlatforms = v
}

// GetActions returns the Actions field value
func (o *Metadata) GetActions() []DestinationMetadataActionV1 {
	if o == nil {
		var ret []DestinationMetadataActionV1
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetActionsOk() ([]DestinationMetadataActionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *Metadata) SetActions(v []DestinationMetadataActionV1) {
	o.Actions = v
}

// GetPresets returns the Presets field value
func (o *Metadata) GetPresets() []DestinationMetadataSubscriptionPresetV1 {
	if o == nil {
		var ret []DestinationMetadataSubscriptionPresetV1
		return ret
	}

	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetPresetsOk() ([]DestinationMetadataSubscriptionPresetV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Presets, true
}

// SetPresets sets field value
func (o *Metadata) SetPresets(v []DestinationMetadataSubscriptionPresetV1) {
	o.Presets = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *Metadata) GetContacts() []Contact {
	if o == nil || o.Contacts == nil {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetContactsOk() ([]Contact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *Metadata) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *Metadata) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetPartnerOwned returns the PartnerOwned field value if set, zero value otherwise.
func (o *Metadata) GetPartnerOwned() bool {
	if o == nil || o.PartnerOwned == nil {
		var ret bool
		return ret
	}
	return *o.PartnerOwned
}

// GetPartnerOwnedOk returns a tuple with the PartnerOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetPartnerOwnedOk() (*bool, bool) {
	if o == nil || o.PartnerOwned == nil {
		return nil, false
	}
	return o.PartnerOwned, true
}

// HasPartnerOwned returns a boolean if a field has been set.
func (o *Metadata) HasPartnerOwned() bool {
	if o != nil && o.PartnerOwned != nil {
		return true
	}

	return false
}

// SetPartnerOwned gets a reference to the given bool and assigns it to the PartnerOwned field.
func (o *Metadata) SetPartnerOwned(v bool) {
	o.PartnerOwned = &v
}

// GetSupportedRegions returns the SupportedRegions field value if set, zero value otherwise.
func (o *Metadata) GetSupportedRegions() []string {
	if o == nil || o.SupportedRegions == nil {
		var ret []string
		return ret
	}
	return o.SupportedRegions
}

// GetSupportedRegionsOk returns a tuple with the SupportedRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetSupportedRegionsOk() ([]string, bool) {
	if o == nil || o.SupportedRegions == nil {
		return nil, false
	}
	return o.SupportedRegions, true
}

// HasSupportedRegions returns a boolean if a field has been set.
func (o *Metadata) HasSupportedRegions() bool {
	if o != nil && o.SupportedRegions != nil {
		return true
	}

	return false
}

// SetSupportedRegions gets a reference to the given []string and assigns it to the SupportedRegions field.
func (o *Metadata) SetSupportedRegions(v []string) {
	o.SupportedRegions = v
}

// GetRegionEndpoints returns the RegionEndpoints field value if set, zero value otherwise.
func (o *Metadata) GetRegionEndpoints() []string {
	if o == nil || o.RegionEndpoints == nil {
		var ret []string
		return ret
	}
	return o.RegionEndpoints
}

// GetRegionEndpointsOk returns a tuple with the RegionEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetRegionEndpointsOk() ([]string, bool) {
	if o == nil || o.RegionEndpoints == nil {
		return nil, false
	}
	return o.RegionEndpoints, true
}

// HasRegionEndpoints returns a boolean if a field has been set.
func (o *Metadata) HasRegionEndpoints() bool {
	if o != nil && o.RegionEndpoints != nil {
		return true
	}

	return false
}

// SetRegionEndpoints gets a reference to the given []string and assigns it to the RegionEndpoints field.
func (o *Metadata) SetRegionEndpoints(v []string) {
	o.RegionEndpoints = v
}

func (o Metadata) MarshalJSON() ([]byte, error) {
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

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
