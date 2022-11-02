// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    "net/http/httptest"
    "net/http"
)

func Test_Echo(t *testing.T) {
    fakeServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("content-type", "application/json")
			_, _ = w.Write([]byte(`{
                    "data": {
                    	"message": "Hello Go-SDK!"
                    }
                }`))
		}),
		)
		defer fakeServer.Close()

    configuration := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "Segment (OpenAPI-Generator test)",
		Debug:         false,
		Servers: ServerConfigurations{
			{
				URL:         fakeServer.URL,
				Description: "Local",
			},
		},
		OperationServers: map[string]ServerConfigurations{},
	}

    apiClient := NewAPIClient(configuration)

    t.Run("Test Echo", func(t *testing.T) {
        resp, httpRes, err := apiClient.TestingApi.Echo(context.Background()).Message("Hello Go-SDK!").Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        assert.Equal(t, resp.Data.Message, "Hello Go-SDK!")
    })
}