/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs see our [API home page](https://help.sumologic.com/APIs). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/Manage/Security/Access-Keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-301-Error-Moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-403-Error-This-operation-is-not-allowed-for-your-account-type) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Entitlements struct for Entitlements
type Entitlements struct {
	// Details of the contract type. `AnnualBucket` are contracts that buy and consume ingest on yearly basis. `Credits` are contracts that buy a single unit called credits for all our features. `DailyAverage` are contracts that buy and consume ingest on a monthly basis.
	ContractType string `json:"contractType"`
	// Text denoting the type of entitlement. - `continuous` for Continuous Analytics, - `frequent` for Frequent Analytics, - `storage` for Total Storage, - `metrics` for Metrics.
	EntitlementType string `json:"entitlementType"`
	// The label of an entitlement is the plan name displayed on the accounts page in our user interface.
	Label string `json:"label"`
	Capacity Capacity `json:"capacity"`
	// Contains the capacities that were part of the contract.
	Capacities *[]Capacity `json:"capacities,omitempty"`
}

// NewEntitlements instantiates a new Entitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlements(contractType string, entitlementType string, label string, capacity Capacity) *Entitlements {
	this := Entitlements{}
	this.ContractType = contractType
	this.EntitlementType = entitlementType
	this.Label = label
	this.Capacity = capacity
	return &this
}

// NewEntitlementsWithDefaults instantiates a new Entitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsWithDefaults() *Entitlements {
	this := Entitlements{}
	return &this
}

// GetContractType returns the ContractType field value
func (o *Entitlements) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *Entitlements) GetContractTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *Entitlements) SetContractType(v string) {
	o.ContractType = v
}

// GetEntitlementType returns the EntitlementType field value
func (o *Entitlements) GetEntitlementType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementType
}

// GetEntitlementTypeOk returns a tuple with the EntitlementType field value
// and a boolean to check if the value has been set.
func (o *Entitlements) GetEntitlementTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntitlementType, true
}

// SetEntitlementType sets field value
func (o *Entitlements) SetEntitlementType(v string) {
	o.EntitlementType = v
}

// GetLabel returns the Label field value
func (o *Entitlements) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Entitlements) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Entitlements) SetLabel(v string) {
	o.Label = v
}

// GetCapacity returns the Capacity field value
func (o *Entitlements) GetCapacity() Capacity {
	if o == nil {
		var ret Capacity
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *Entitlements) GetCapacityOk() (*Capacity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *Entitlements) SetCapacity(v Capacity) {
	o.Capacity = v
}

// GetCapacities returns the Capacities field value if set, zero value otherwise.
func (o *Entitlements) GetCapacities() []Capacity {
	if o == nil || o.Capacities == nil {
		var ret []Capacity
		return ret
	}
	return *o.Capacities
}

// GetCapacitiesOk returns a tuple with the Capacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetCapacitiesOk() (*[]Capacity, bool) {
	if o == nil || o.Capacities == nil {
		return nil, false
	}
	return o.Capacities, true
}

// HasCapacities returns a boolean if a field has been set.
func (o *Entitlements) HasCapacities() bool {
	if o != nil && o.Capacities != nil {
		return true
	}

	return false
}

// SetCapacities gets a reference to the given []Capacity and assigns it to the Capacities field.
func (o *Entitlements) SetCapacities(v []Capacity) {
	o.Capacities = &v
}

func (o Entitlements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractType"] = o.ContractType
	}
	if true {
		toSerialize["entitlementType"] = o.EntitlementType
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["capacity"] = o.Capacity
	}
	if o.Capacities != nil {
		toSerialize["capacities"] = o.Capacities
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlements struct {
	value *Entitlements
	isSet bool
}

func (v NullableEntitlements) Get() *Entitlements {
	return v.value
}

func (v *NullableEntitlements) Set(val *Entitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlements(val *Entitlements) *NullableEntitlements {
	return &NullableEntitlements{value: val, isSet: true}
}

func (v NullableEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


