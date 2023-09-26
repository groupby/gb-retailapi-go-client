# ProductDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Relative path to product in Google Retail system. | [optional] 
**Id** | Pointer to **string** | Product id in Google Retail system. | [optional] 
**Type** | Pointer to **string** | Product type. Possible values: PRIMARY, VARIANT. If the product has variant list and the request specifies the variantIds, requested variants will be the first in the response. | [optional] 
**PrimaryProductId** | Pointer to **string** | Product ID that is primary in relation to the current one | [optional] 
**CollectionMemberIds** | Pointer to **[]string** | The of the collection members when product type is COLLECTION | [optional] 
**Gtin** | Pointer to **string** | Global Trade Item Number can be used by a company to uniquely identify all of its trade items.GTIN defines trade items as products or services that are priced, ordered or invoiced at any point in the supply chain. | [optional] 
**Categories** | Pointer to **[]string** | Product categories (array). | [optional] 
**Title** | Pointer to **string** | Product title. | [optional] 
**Brands** | Pointer to **[]string** | Product brands. | [optional] 
**Description** | Pointer to **string** | Product description. | [optional] 
**LanguageCode** | Pointer to **string** | Language of the title/description and other string attributes. Use language tags defined by [BCP 47][https://www.rfc-editor.org/rfc/bcp/bcp47.txt]. For product search this field is in use. It defaults to &#39;en-US&#39; if unset. | [optional] 
**Attributes** | Pointer to [**map[string]ProductCustomAttribute**](ProductCustomAttribute.md) | Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer.  Max entries count: 200. Length limit of 128 characters. | [optional] 
**Tags** | Pointer to **[]string** | Product tags (array). | [optional] 
**PriceInfo** | Pointer to [**ProductDtoPriceInfo**](ProductDtoPriceInfo.md) |  | [optional] 
**Rating** | Pointer to [**ProductDtoRating**](ProductDtoRating.md) |  | [optional] 
**AvailableTime** | Pointer to [**ProductDtoAvailableTime**](ProductDtoAvailableTime.md) |  | [optional] 
**Availability** | Pointer to **string** | The online availability of the product. Default to IN_STOCK | [optional] 
**AvailableQuantity** | Pointer to **int32** | The available quantity of the item. | [optional] 
**FulfillmentInfos** | Pointer to [**[]FulfillmentInfo**](FulfillmentInfo.md) | Fulfillment information, such as the store IDs for in-store pickup or region IDs for different shipping methods. | [optional] 
**Uri** | Pointer to **string** | Link to the appropriate product. | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) | Product Image. | [optional] 
**Audience** | Pointer to [**ProductDtoAudience**](ProductDtoAudience.md) |  | [optional] 
**ColorInfo** | Pointer to [**ProductDtoColorInfo**](ProductDtoColorInfo.md) |  | [optional] 
**Sizes** | Pointer to **[]string** | Product sizes (array). | [optional] 
**Materials** | Pointer to **[]string** | The material of the product. For example, &#39;leather&#39;, &#39;wooden&#39;. A maximum of 20 values are allowed. Length limit of 128 characters | [optional] 
**Patterns** | Pointer to **[]string** | The pattern or graphic print of the product. For example, &#39;striped&#39;, &#39;polka dot&#39;, &#39;paisley&#39;. A maximum of 20 values are allowed per product. Length limit of 128 characters. | [optional] 
**Conditions** | Pointer to **[]string** | The condition of the product. Strongly encouraged to use the standardvalues: &#39;new&#39;, &#39;refurbished&#39;, &#39;used&#39;. A maximum of 5 values are allowed per product. Length limit of 128 characters. | [optional] 
**PublishTime** | Pointer to [**ProductDtoPublishTime**](ProductDtoPublishTime.md) |  | [optional] 
**RetrievableFields** | Pointer to [**ProductDtoRetrievableFields**](ProductDtoRetrievableFields.md) |  | [optional] 
**Promotions** | Pointer to [**[]Promotion**](Promotion.md) | The promotions applied to the product. A maximum of 10 values are allowed per product. | [optional] 
**Variants** | Pointer to [**[]ProductDto**](ProductDto.md) | If the product has variant list and the request specifies the variantIds, requested variants will be the first in the response. | [optional] 

## Methods

### NewProductDto

`func NewProductDto() *ProductDto`

NewProductDto instantiates a new ProductDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDtoWithDefaults

`func NewProductDtoWithDefaults() *ProductDto`

NewProductDtoWithDefaults instantiates a new ProductDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ProductDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ProductDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrimaryProductId

`func (o *ProductDto) GetPrimaryProductId() string`

GetPrimaryProductId returns the PrimaryProductId field if non-nil, zero value otherwise.

### GetPrimaryProductIdOk

`func (o *ProductDto) GetPrimaryProductIdOk() (*string, bool)`

GetPrimaryProductIdOk returns a tuple with the PrimaryProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryProductId

`func (o *ProductDto) SetPrimaryProductId(v string)`

SetPrimaryProductId sets PrimaryProductId field to given value.

### HasPrimaryProductId

`func (o *ProductDto) HasPrimaryProductId() bool`

HasPrimaryProductId returns a boolean if a field has been set.

### GetCollectionMemberIds

`func (o *ProductDto) GetCollectionMemberIds() []string`

GetCollectionMemberIds returns the CollectionMemberIds field if non-nil, zero value otherwise.

### GetCollectionMemberIdsOk

`func (o *ProductDto) GetCollectionMemberIdsOk() (*[]string, bool)`

GetCollectionMemberIdsOk returns a tuple with the CollectionMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMemberIds

`func (o *ProductDto) SetCollectionMemberIds(v []string)`

SetCollectionMemberIds sets CollectionMemberIds field to given value.

### HasCollectionMemberIds

`func (o *ProductDto) HasCollectionMemberIds() bool`

HasCollectionMemberIds returns a boolean if a field has been set.

### GetGtin

`func (o *ProductDto) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductDto) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductDto) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductDto) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetCategories

`func (o *ProductDto) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ProductDto) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ProductDto) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ProductDto) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTitle

`func (o *ProductDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProductDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProductDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProductDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBrands

`func (o *ProductDto) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *ProductDto) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *ProductDto) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *ProductDto) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetDescription

`func (o *ProductDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLanguageCode

`func (o *ProductDto) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *ProductDto) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *ProductDto) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *ProductDto) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetAttributes

`func (o *ProductDto) GetAttributes() map[string]ProductCustomAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProductDto) GetAttributesOk() (*map[string]ProductCustomAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProductDto) SetAttributes(v map[string]ProductCustomAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProductDto) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *ProductDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPriceInfo

`func (o *ProductDto) GetPriceInfo() ProductDtoPriceInfo`

GetPriceInfo returns the PriceInfo field if non-nil, zero value otherwise.

### GetPriceInfoOk

`func (o *ProductDto) GetPriceInfoOk() (*ProductDtoPriceInfo, bool)`

GetPriceInfoOk returns a tuple with the PriceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInfo

`func (o *ProductDto) SetPriceInfo(v ProductDtoPriceInfo)`

SetPriceInfo sets PriceInfo field to given value.

### HasPriceInfo

`func (o *ProductDto) HasPriceInfo() bool`

HasPriceInfo returns a boolean if a field has been set.

### GetRating

`func (o *ProductDto) GetRating() ProductDtoRating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ProductDto) GetRatingOk() (*ProductDtoRating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ProductDto) SetRating(v ProductDtoRating)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ProductDto) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetAvailableTime

`func (o *ProductDto) GetAvailableTime() ProductDtoAvailableTime`

GetAvailableTime returns the AvailableTime field if non-nil, zero value otherwise.

### GetAvailableTimeOk

`func (o *ProductDto) GetAvailableTimeOk() (*ProductDtoAvailableTime, bool)`

GetAvailableTimeOk returns a tuple with the AvailableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTime

`func (o *ProductDto) SetAvailableTime(v ProductDtoAvailableTime)`

SetAvailableTime sets AvailableTime field to given value.

### HasAvailableTime

`func (o *ProductDto) HasAvailableTime() bool`

HasAvailableTime returns a boolean if a field has been set.

### GetAvailability

`func (o *ProductDto) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ProductDto) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ProductDto) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ProductDto) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetAvailableQuantity

`func (o *ProductDto) GetAvailableQuantity() int32`

GetAvailableQuantity returns the AvailableQuantity field if non-nil, zero value otherwise.

### GetAvailableQuantityOk

`func (o *ProductDto) GetAvailableQuantityOk() (*int32, bool)`

GetAvailableQuantityOk returns a tuple with the AvailableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableQuantity

`func (o *ProductDto) SetAvailableQuantity(v int32)`

SetAvailableQuantity sets AvailableQuantity field to given value.

### HasAvailableQuantity

`func (o *ProductDto) HasAvailableQuantity() bool`

HasAvailableQuantity returns a boolean if a field has been set.

### GetFulfillmentInfos

`func (o *ProductDto) GetFulfillmentInfos() []FulfillmentInfo`

GetFulfillmentInfos returns the FulfillmentInfos field if non-nil, zero value otherwise.

### GetFulfillmentInfosOk

`func (o *ProductDto) GetFulfillmentInfosOk() (*[]FulfillmentInfo, bool)`

GetFulfillmentInfosOk returns a tuple with the FulfillmentInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentInfos

`func (o *ProductDto) SetFulfillmentInfos(v []FulfillmentInfo)`

SetFulfillmentInfos sets FulfillmentInfos field to given value.

### HasFulfillmentInfos

`func (o *ProductDto) HasFulfillmentInfos() bool`

HasFulfillmentInfos returns a boolean if a field has been set.

### GetUri

`func (o *ProductDto) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ProductDto) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ProductDto) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ProductDto) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetImages

`func (o *ProductDto) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ProductDto) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ProductDto) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *ProductDto) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetAudience

`func (o *ProductDto) GetAudience() ProductDtoAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ProductDto) GetAudienceOk() (*ProductDtoAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ProductDto) SetAudience(v ProductDtoAudience)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *ProductDto) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetColorInfo

`func (o *ProductDto) GetColorInfo() ProductDtoColorInfo`

GetColorInfo returns the ColorInfo field if non-nil, zero value otherwise.

### GetColorInfoOk

`func (o *ProductDto) GetColorInfoOk() (*ProductDtoColorInfo, bool)`

GetColorInfoOk returns a tuple with the ColorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorInfo

`func (o *ProductDto) SetColorInfo(v ProductDtoColorInfo)`

SetColorInfo sets ColorInfo field to given value.

### HasColorInfo

`func (o *ProductDto) HasColorInfo() bool`

HasColorInfo returns a boolean if a field has been set.

### GetSizes

`func (o *ProductDto) GetSizes() []string`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *ProductDto) GetSizesOk() (*[]string, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *ProductDto) SetSizes(v []string)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *ProductDto) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetMaterials

`func (o *ProductDto) GetMaterials() []string`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *ProductDto) GetMaterialsOk() (*[]string, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *ProductDto) SetMaterials(v []string)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *ProductDto) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### GetPatterns

`func (o *ProductDto) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *ProductDto) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *ProductDto) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *ProductDto) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetConditions

`func (o *ProductDto) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ProductDto) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ProductDto) SetConditions(v []string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ProductDto) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPublishTime

`func (o *ProductDto) GetPublishTime() ProductDtoPublishTime`

GetPublishTime returns the PublishTime field if non-nil, zero value otherwise.

### GetPublishTimeOk

`func (o *ProductDto) GetPublishTimeOk() (*ProductDtoPublishTime, bool)`

GetPublishTimeOk returns a tuple with the PublishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTime

`func (o *ProductDto) SetPublishTime(v ProductDtoPublishTime)`

SetPublishTime sets PublishTime field to given value.

### HasPublishTime

`func (o *ProductDto) HasPublishTime() bool`

HasPublishTime returns a boolean if a field has been set.

### GetRetrievableFields

`func (o *ProductDto) GetRetrievableFields() ProductDtoRetrievableFields`

GetRetrievableFields returns the RetrievableFields field if non-nil, zero value otherwise.

### GetRetrievableFieldsOk

`func (o *ProductDto) GetRetrievableFieldsOk() (*ProductDtoRetrievableFields, bool)`

GetRetrievableFieldsOk returns a tuple with the RetrievableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievableFields

`func (o *ProductDto) SetRetrievableFields(v ProductDtoRetrievableFields)`

SetRetrievableFields sets RetrievableFields field to given value.

### HasRetrievableFields

`func (o *ProductDto) HasRetrievableFields() bool`

HasRetrievableFields returns a boolean if a field has been set.

### GetPromotions

`func (o *ProductDto) GetPromotions() []Promotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *ProductDto) GetPromotionsOk() (*[]Promotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *ProductDto) SetPromotions(v []Promotion)`

SetPromotions sets Promotions field to given value.

### HasPromotions

`func (o *ProductDto) HasPromotions() bool`

HasPromotions returns a boolean if a field has been set.

### GetVariants

`func (o *ProductDto) GetVariants() []ProductDto`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ProductDto) GetVariantsOk() (*[]ProductDto, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ProductDto) SetVariants(v []ProductDto)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ProductDto) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


