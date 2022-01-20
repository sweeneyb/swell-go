package swell

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Results struct {
	Count   int       `json:"count"`
	Results []Product `json:"results"`
}
type Product struct {
	Name string `json:"name"`
	Sku  string `json:"sku"`
}

func (c *Client) GetProducts() ([]Product, error) {
	ctx := context.Background()
	return c.GetProductsWithContext(ctx)
}
func (c *Client) GetProductsWithContext(ctx context.Context) ([]Product, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products", c.HostUrl), nil)
	if err != nil {
		log.Fatal(err)
	}

	textBytes, err := c.doRequest(req)
	if err != nil {
		fmt.Printf("status code: %v\n", err)
		return nil, err
	}

	return parseBytes(textBytes)

}

func parseBytes(textBytes []byte) ([]Product, error) {
	results := Results{}
	err := json.Unmarshal(textBytes, &results)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return results.Results, nil
}
