package replacer_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/streamwest-1629/go_object_replacer/replacer"
)

func TestReplacer(t *testing.T) {

	type Person struct {
		Name string `replacer:"name!"`
		Age  uint   `replacer:"age!"`
	}

	type RootObj struct {
		FamilyName string   `replacer:"family_name!"`
		Checked    bool     `replacer:"checked"`
		Patriarch  Person   `replacer:"<-"`
		Partner    *Person  `replacer:"partner"`
		Children   []Person `replacer:"children"`
	}

	conv, err := replacer.MakeConverter(RootObj{})
	if err != nil {
		t.Fatal(err.Error())
	}

	testMap := map[string]interface{}{
		"family_name": "Adams",
		"name":        "Smith",
		"checked":     true,
		"age":         52,
		"partner": map[string]interface{}{
			"name": "Julia",
			"age":  34,
		},
		"children": []map[string]interface{}{
			{
				"name": "Amy",
				"age":  15,
			},
			{
				"name": "Franz",
				"age":  13,
			},
		},
	}

	dest := &RootObj{}
	if err := conv.Convert(testMap, dest); err != nil {
		t.Fatal(err.Error())
	}

	b, _ := json.Marshal(dest)
	log.Println(string(b))
}
