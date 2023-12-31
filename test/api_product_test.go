/*
GroupBy Retail

Testing ProductAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gbretailapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/groupby/gb-retailapi-go-client"
)

func Test_gbretailapi_ProductAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductAPIService GetByProductIds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.GetByProductIds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
