// +build jsoniter

package jsons

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func MyMarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	fmt.Println("use jsoniter package")
	return json.MarshalIndent(v, prefix, indent)
}
