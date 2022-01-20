package swell

import (
	"fmt"
	"net/http"
	"testing"
)

func TestTextSize(t *testing.T) {
	text, _ := getStaticJson()
	if len(text) < 100 {
		t.Fatal("short texts")
	}
}

func TestProductParse(t *testing.T) {
	bytes, _ := getStaticJson()
	results, err := parseBytes(bytes)
	if err != nil {
		t.Fatal(err)
	}
	if( len (results) != 3) {
		t.Fatal("wrong result size")
	}
	fmt.Printf("%v", len(results))
}

func getStaticJson() ([]byte, error) {
	text := `{
		"count": 3,
		"results": [
		  {
			"name": "45 Minute Coaching Session",
			"sku": null,
			"type": "digital",
			"active": true,
			"images": null,
			"purchase_options": {
			  "standard": {
				"active": true,
				"price": 45,
				"sale": false,
				"sale_price": null,
				"prices": []
			  }
			},
			"variable": false,
			"description": "This is a free-form sessions where we talk through the specifics of your project. &nbsp;Pairing is available, and any resulting IP belongs to you.",
			"tags": [],
			"meta_title": null,
			"meta_description": null,
			"slug": "45-minute-coaching-session",
			"attributes": {},
			"delivery": null,
			"virtual": true,
			"bundle": null,
			"price": 45,
			"stock_tracking": false,
			"options": [],
			"currency": "USD",
			"date_created": "2022-01-09T18:15:44.301Z",
			"stock_status": null,
			"date_updated": "2022-01-09T18:52:14.225Z",
			"prices": [],
			"sale": false,
			"sale_price": null,
			"stock_purchasable": false,
			"popularity": 1,
			"id": "61db26508e5fd0013d42bdcb"
		  },
		  {
			"name": "Moccamaster KBS Black",
			"sku": "TV-MM-KBS-BL",
			"active": true,
			"images": [
			  {
				"file": {
				  "id": "5f6fd37b33b5a2371f471d9c",
				  "date_uploaded": "2020-09-26T23:49:15.628Z",
				  "length": 533204,
				  "md5": "4d4ad535acaa7c46ac6aebdbcba80801",
				  "filename": null,
				  "content_type": "image/jpeg",
				  "metadata": null,
				  "url": "https://cdn.schema.io/test-theme/5f6fd37b33b5a2371f471d9c/4d4ad535acaa7c46ac6aebdbcba80801",
				  "width": 1500,
				  "height": 1500
				},
				"id": "5e5198f272f5542e005fa520"
			  }
			],
			"bundle": false,
			"sale": false,
			"price": 298,
			"sale_price": null,
			"cost": null,
			"prices": [],
			"variable": false,
			"description": "Loved by coffee aficionados and design snobs alike, the Moccamaster has made a name for itself around the world as a high-quality filter brewer that&#39;s also extremely easy to use. Handmade in The Netherlands by Dutch company <a href=\"https://technivorm.com/\">Technivorm</a> since 1964, the Moccamaster heats water to a precise temperature before dripping it through a conical paper filter above a receptacle. An integrated heating element in the base maintains the brewed coffee at optimal drinking temperature for 40 minutes, then switches itself off.<br><br>The KBS features an angled glass pot and brews 8 cups of coffee in six minutes. All Moccamasters are certified by the Specialty Coffee Association (SCA) and come with a 5-year manufacturer&#39;s warranty.&nbsp;",
			"tags": [],
			"meta_title": null,
			"meta_description": "Handmade in The Netherlands, the KBS filter coffee brewer is loved by coffee aficionados and design snobs alike. It brews 8 cups of coffee in six minutes and features an angled glass pot. ",
			"slug": "moccamaster-kbs-black",
			"attributes": {
			  "brand": "Technivorm"
			},
			"type": "standard",
			"stock_tracking": true,
			"options": [],
			"content": {},
			"currency": "USD",
			"delivery": "shipment",
			"tax_class": "standard",
			"date_created": "2020-02-22T21:10:38.240Z",
			"stock_status": "out_of_stock",
			"category_index": {
			  "sort": {
				"5e31db0a08eb153204cd4a1d": 0
			  },
			  "id": [
				"5e31db0a08eb153204cd4a1d"
			  ]
			},
			"date_updated": "2022-01-09T18:33:45.169Z",
			"stock_purchasable": true,
			"popularity": 3,
			"id": "5e5198ce72f5542e005fa4c8"
		  },
		  {
			"name": "AeroPress Paper Filters (350 pack)",
			"sku": "AERO-FLT-350",
			"active": true,
			"images": [
			  {
				"file": {
				  "id": "5f6fd3521abceb3936ee2418",
				  "date_uploaded": "2020-09-26T23:48:34.333Z",
				  "length": 244084,
				  "md5": "c29a6e7c2fddde366149ccf829985e27",
				  "filename": null,
				  "content_type": "image/jpeg",
				  "metadata": null,
				  "url": "https://cdn.schema.io/test-theme/5f6fd3521abceb3936ee2418/c29a6e7c2fddde366149ccf829985e27",
				  "width": 1500,
				  "height": 1042
				},
				"id": "5e31eab3db8c386823a75eeb"
			  }
			],
			"bundle": false,
			"sale": false,
			"price": 4.95,
			"sale_price": null,
			"cost": null,
			"prices": [],
			"variable": false,
			"description": "<p>These paper AeroPress Micro Filters help you brew coffee with a smooth flavour and clean body by preventing micro grounds from entering your cup. Paper Micro Filters make cleaning up a breeze after brewing a mug of Aeropress coffee. Simply press the plunger until the filter and used grounds are forced into the garbage.<br><br>This pack comes with 350 single-use filters of 6.4cm diameter and are designed to be used with the Aeropress or other brewers that use circular filters.<br><br><br></p>",
			"tags": [],
			"meta_title": null,
			"meta_description": "These paper AeroPress Micro Filters help you brew coffee with a smooth flavour and clean body by disabling micro grounds to enter your final mug.",
			"slug": "aero-press-paper-filters-350-pack",
			"attributes": {
			  "brand": "Aerobie"
			},
			"type": "standard",
			"stock_tracking": true,
			"options": [],
			"content": {},
			"currency": "USD",
			"delivery": "shipment",
			"tax_class": "standard",
			"date_created": "2020-01-29T20:09:31.118Z",
			"stock_status": "out_of_stock",
			"category_index": {
			  "sort": {
				"5e31dbe2ae1309046a52f343": 1
			  },
			  "id": [
				"5e31dbe2ae1309046a52f343"
			  ]
			},
			"date_updated": "2022-01-09T19:13:55.373Z",
			"stock_purchasable": true,
			"stock_level": -1,
			"popularity": 1,
			"id": "5e31e67be53f9a59d89600f1"
		  }
		],
		"page": 1
	  }`
	return []byte(text), nil
}
