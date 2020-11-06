package parser

import (
	"io/ioutil"
	"learn-go/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "断念肉嘟嘟")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 "+
			"elementl; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:       "断念肉嘟嘟",
		Gender:     "男",
		Age:        27,
		Height:     7,
		Weight:     57,
		Income:     "8001-10000元",
		Marriage:   "离异",
		Education:  "硕士",
		Occupation: "财务",
		Hokou:      "大连市",
		Xinzuo:     "天秤座",
		House:      "有房",
		Car:        "有车",
	}

	if profile != expected {
		t.Errorf("expected %v;but was %v", expected, profile)
	}
}
