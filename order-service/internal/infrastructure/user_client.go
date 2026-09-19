package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"ecommerce-microservice/order-service/internal/model"
	"ecommerce-microservice/order-service/internal/usecase"
)

type HTTPUserServiceClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHTTPUserServiceClient(baseURL string) *HTTPUserServiceClient {
	return &HTTPUserServiceClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (c *HTTPUserServiceClient) GetUserByID(
	ctx context.Context,
	userID int64,
) (usecase.User, error) {

	url := fmt.Sprintf(
		"%s/api/v1/users/%d",
		c.baseURL,
		userID,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return usecase.User{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return usecase.User{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return usecase.User{}, model.ErrUserNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return usecase.User{}, fmt.Errorf(
			"user service returned status %d",
			resp.StatusCode,
		)
	}

	var user usecase.User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return usecase.User{}, err
	}

	return user, nil
}
