package mws

type Product struct {
	Client
}

/**
 * Get Competitive Pricing For ASIN
 * Gets competitive pricing and related information for a product identified by
 * the MarketplaceId and ASIN.
 *
 * @param mixed $request array of parameters for GetCompetitivePricingForASIN request or GetCompetitivePricingForASIN object itself
 * @see GetCompetitivePricingForASINRequest
 * @return GetCompetitivePricingForASINResponse
 *
 * @throws ProductsException
 */
func (api Product) GetCompetitivePricingForASIN() {
}

/**
 * Get Competitive Pricing For SKU
 * Gets competitive pricing and related information for a product identified by
 * the SellerId and SKU.
 *
 * @param mixed $request array of parameters for GetCompetitivePricingForSKU request or GetCompetitivePricingForSKU object itself
 * @see GetCompetitivePricingForSKURequest
 * @return GetCompetitivePricingForSKUResponse
 *
 * @throws ProductsException
 */
func (api Product) GetCompetitivePricingForSKU() {
}

/**
 * Get Lowest Offer Listings For ASIN
 * Gets some of the lowest prices based on the product identified by the given
 * MarketplaceId and ASIN.
 *
 * @param mixed $request array of parameters for GetLowestOfferListingsForASIN request or GetLowestOfferListingsForASIN object itself
 * @see GetLowestOfferListingsForASINRequest
 * @return GetLowestOfferListingsForASINResponse
 *
 * @throws ProductsException
 */
func (api Product) GetLowestOfferListingsForASIN() {
}

/**
 * Get Lowest Offer Listings For SKU
 * Gets some of the lowest prices based on the product identified by the given
 * SellerId and SKU.
 *
 * @param mixed $request array of parameters for GetLowestOfferListingsForSKU request or GetLowestOfferListingsForSKU object itself
 * @see GetLowestOfferListingsForSKURequest
 * @return GetLowestOfferListingsForSKUResponse
 *
 * @throws ProductsException
 */
func (api Product) GetLowestOfferListingsForSKU() {
}

/**
 * Get Lowest Priced Offers For ASIN
 * Retrieves the lowest priced offers based on the product identified by the given
 *     ASIN.
 *
 * @param mixed $request array of parameters for GetLowestPricedOffersForASIN request or GetLowestPricedOffersForASIN object itself
 * @see GetLowestPricedOffersForASINRequest
 * @return GetLowestPricedOffersForASINResponse
 *
 * @throws ProductsException
 */
func (api Product) GetLowestPricedOffersForASIN() {
}

/**
 * Get Lowest Priced Offers For SKU
 * Retrieves the lowest priced offers based on the product identified by the given
 *     SellerId and SKU.
 *
 * @param mixed $request array of parameters for GetLowestPricedOffersForSKU request or GetLowestPricedOffersForSKU object itself
 * @see GetLowestPricedOffersForSKURequest
 * @return GetLowestPricedOffersForSKUResponse
 *
 * @throws ProductsException
 */
func (api Product) GetLowestPricedOffersForSKU() {
}

/**
 * Get Matching Product
 * GetMatchingProduct will return the details (attributes) for the
 * given ASIN.
 *
 * @param mixed $request array of parameters for GetMatchingProduct request or GetMatchingProduct object itself
 * @see GetMatchingProductRequest
 * @return GetMatchingProductResponse
 *
 * @throws ProductsException
 */
func (api Product) GetMatchingProduct() {
}

/**
 * Get Matching Product For Id
 * GetMatchingProduct will return the details (attributes) for the
 * given Identifier list. Identifer type can be one of [SKU|ASIN|UPC|EAN|ISBN|GTIN|JAN]
 *
 * @param mixed $request array of parameters for GetMatchingProductForId request or GetMatchingProductForId object itself
 * @see GetMatchingProductForIdRequest
 * @return GetMatchingProductForIdResponse
 *
 * @throws ProductsException
 */
func (api Product) GetMatchingProductForId() {
}

/**
 * Get My Fees Estimate
 * Retrieves the fees estimate for the
 *         products identified by the given
 *         ASIN/SKU list.
 *
 * @param mixed $request array of parameters for GetMyFeesEstimate request or GetMyFeesEstimate object itself
 * @see GetMyFeesEstimateRequest
 * @return GetMyFeesEstimateResponse
 *
 * @throws ProductsException
 */
func (api Product) GetMyFeesEstimate() {
}

/**
 * Get My Price For ASIN
 * <!-- Wrong doc in current code -->
 *
 * @param mixed $request array of parameters for GetMyPriceForASIN request or GetMyPriceForASIN object itself
 * @see GetMyPriceForASINRequest
 * @return GetMyPriceForASINResponse
 *
 * @throws ProductsException
 */
func (api Product) GetMyPriceForASIN() {
}

/**
 * Get My Price For SKU
 * <!-- Wrong doc in current code -->
 *
 * @param mixed $request array of parameters for GetMyPriceForSKU request or GetMyPriceForSKU object itself
 * @see GetMyPriceForSKURequest
 * @return GetMyPriceForSKUResponse
 *
 * @throws ProductsException
 */
func (api Product) GetMyPriceForSKU() {
}

/**
 * Get Product Categories For ASIN
 * Gets categories information for a product identified by
 * the MarketplaceId and ASIN.
 *
 * @param mixed $request array of parameters for GetProductCategoriesForASIN request or GetProductCategoriesForASIN object itself
 * @see GetProductCategoriesForASINRequest
 * @return GetProductCategoriesForASINResponse
 *
 * @throws ProductsException
 */
func (api Product) GetProductCategoriesForASIN() {
}

/**
 * Get Product Categories For SKU
 * Gets categories information for a product identified by
 * the SellerId and SKU.
 *
 * @param mixed $request array of parameters for GetProductCategoriesForSKU request or GetProductCategoriesForSKU object itself
 * @see GetProductCategoriesForSKURequest
 * @return GetProductCategoriesForSKUResponse
 *
 * @throws ProductsException
 */
func (api Product) GetProductCategoriesForSKU() {
}

/**
 * Get Service Status
 * Returns the service status of a particular MWS API section. The operation
 * takes no input.
 * All API sections within the API are required to implement this operation.
 *
 * @param mixed $request array of parameters for GetServiceStatus request or GetServiceStatus object itself
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 *
 * @throws ProductsException
 */
func (api Product) GetServiceStatus() {
}

/**
 * List Matching Products
 * ListMatchingProducts can be used to
 * find products that match the given criteria.
 *
 * @param mixed $request array of parameters for ListMatchingProducts request or ListMatchingProducts object itself
 * @see ListMatchingProductsRequest
 * @return ListMatchingProductsResponse
 *
 * @throws ProductsException
 */
func (api Product) ListMatchingProducts() {
}
