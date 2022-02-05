package swell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	Id          string    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Active      bool      `json:"active,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Description string    `json:"description,omitempty"`
	Images      []Image   `json:"images,omitempty"`

	Sku             string `json:"sku,omitempty"`
	Type            string `json:"type,omitempty"`
	MetaDescription string `json:"meta_description,omitempty"`
	MetaKeywords    string `json:"meta_keywords,omitempty"`
	MetaTitle       string `json:"meta_title,omitempty"`

	ParentId string `json:"parent_id,omitempty"`
	Slug     string `json:"slug,omitempty"`
	Sort     int32  `json:"sort,omitempty"`
	Sorting  string `json:"sorting,omitempty"`
	TopId    string `json:"top_id,omitempty"`
}

type Image struct {
	// TODO add category image support
	Id string `json:"id"`
}

func (c *Client) GetCategory(id string) (Category, error) {
	result := Category{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/categories/%s", c.HostUrl, id), nil)
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

func (c *Client) getRequest(isUpdate bool, category Category) (*http.Request, error) {
	if isUpdate {
		fmt.Println("doing put")
		json, err := json.Marshal(category)
		if err != nil {
			panic(err)
		}
		fmt.Printf("json: %s", json)
		req, err := http.NewRequest("PUT", fmt.Sprintf("%s/categories/%s", c.HostUrl, category.Id), bytes.NewBuffer(json))
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		req.Header.Add("Content-Length", strconv.Itoa(len(json)))
		return req, err
	} else {
		fmt.Println("doing post")
		data := url.Values{}
		data.Set("name", category.Name)
		data.Set("descrption", category.Description)

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/categories", c.HostUrl), strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		return req, err
	}
}

func (c *Client) CreateCategory(category Category) (Category, error) {
	return c.MutateCategory(false, category)
}

func (c *Client) UpdateCategory(category Category) (Category, error) {
	return c.MutateCategory(true, category)
}

func (c *Client) MutateCategory(isUpdate bool, category Category) (Category, error) {

	//req, err := http.NewRequest("POST", fmt.Sprintf("%s/categories", c.HostUrl), strings.NewReader(data.Encode()))
	req, err := c.getRequest(isUpdate, category)
	if err != nil {
		log.Fatal(err)
	}

	result := Category{}
	errors := Error{}
	textBytes, err := c.doRequest(req)
	if err != nil {
		fmt.Printf("status code: %v\n", err)
		return result, err
	}
	err = json.Unmarshal(textBytes, &errors)
	if err != nil {
		return result, err
	}
	if len(errors.Errors) > 0 {

		// TODO clean this up and package it into an Error type
		fmt.Printf("errors: %v\n", errors.Errors)
		for k, v := range errors.Errors {
			fmt.Printf("error: %v: %v\n", k, v["message"])
		}

		return result, fmt.Errorf("API error2:", errors)
	}
	err = json.Unmarshal(textBytes, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
