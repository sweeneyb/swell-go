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
	Id     string `json:"id"`
	Name   string `json:"name"`
	Sku    string `json:"sku"`
	Type   string `json:"type"`
	Active bool   `json:"active`
	// TODO images
	Purchase_Options map[string]Purchase_Option `json:"purchase_options"`
	Variable         bool                       `json:"variable"`
	Description      string                     `json:"description"`
	Slug             string                     `json:"slug"`
	Currency         string                     `json:"currency"`
}

type Purchase_Option struct {
	Active     bool      `json:"active"`
	Price      float64   `json:"price"`
	Sale       bool      `json:"sale"`
	Sale_Price float64   `json:"sale_price"`
	Prices     []float64 `json:"prices"`
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

	results := Results{}
	err = json.Unmarshal(textBytes, &results)
	if err != nil {
		return nil, err
	}
	return results.Results, nil
}

func (c *Client) GetProduct(id string) (Product, error) {
	result := Product{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products/%s", c.HostUrl, id), nil)
	if err != nil {
		log.Fatal(err)
	}

	textBytes, err := c.doRequest(req)
	if err != nil {
		fmt.Printf("status code: %v\n", err)
		return result, err
	}

	err = json.Unmarshal(textBytes, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
