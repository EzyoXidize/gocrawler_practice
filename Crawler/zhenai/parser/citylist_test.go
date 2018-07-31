package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserCityList(t *testing.T) {
	contents,err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents)
	const resultSize  = 470
	expectedUrls := []string {
		"","","",
	}

	expectedCities := []string {
		"","","",
	}

	for i,url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d :%s ; but "+" was %s",
				i,url,result.Requests[i].Url)
		}
	}

	for i,city:= range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d :%s ; but "+" was %s",
				i,city,result.Items[i].(string))
		}
	}


	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d" + "requests; but had %d",resultSize,len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d" + "requests; but had %d",resultSize,len(result.Items))
	}
}
