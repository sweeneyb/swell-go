package swell

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestCheckCategoryParse(t *testing.T) {
	g := NewWithT(t)
	// text := getStaticJson()
	// results := Results{}
	// _ = json.Unmarshal(text, &results)
	// assert.Equal(t, results.Sku, "TV-MM-KBS-BL", "The two words should be the same.")
	subject := Category{}
	err := json.Unmarshal(getStaticCategoryJson(), &subject)
	if err != nil {
		t.Fatal(err)
	}
	g.Expect(subject.Id).To(Equal("60f199509111e70000000012"), "Ids should match.")
	g.Expect(subject.Active).To(BeTrue(), "standard price Active mismatch.")
	g.Expect(subject.Images[0].Id).To(Equal("5cab78ab2045865e3c8a3794"), "Image id mismatch")
	g.Expect(subject.DateCreated).To(Equal(time.Date(2021, 07, 16, 14, 36, 00, 84000000, time.UTC)))
}

func getStaticCategoryJson() []byte {
	text := `{
		"id": "60f199509111e70000000012",
		"name": "Example category",
		"active": true,
		"date_created": "2021-07-16T14:36:00.084Z",
		"date_updated": "2021-07-16T14:36:00.084Z",
		"description": "A long form description of the category",
		"images": [
		  {
			"id": "5cab78ab2045865e3c8a3794",
			"file": {
			  "id": "5cab78ab2045865e3c8a375",
			  "length": 66764,
			  "md5": "99194f53bfdea832553e7fa8ae8fd80f",
			  "content_type": "image/png",
			  "url": "http://cdn.swell.store/test/5ca24abb9c077817e5fe2b36/99194f53bfdea832553e7fa8ae8fd80f",
			  "width": 940,
			  "height": 600
			}
		  }
		],
		"meta_description": null,
		"meta_keywords": null,
		"meta_title": null,
		"parent_id": "60f199509111e70000000013",
		"slug": "example-category",
		"sort": 3,
		"sorting": "price_desc",
		"top_id": "60f199509111e70000000014"
	  }`
	return []byte(text)
}
