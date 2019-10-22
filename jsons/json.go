// +build !jsoniter

package jsons

import (
	"encoding/json"
	"fmt"
)

func MyMarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	fmt.Println("use encoding/json package")
	return json.MarshalIndent(v, prefix, indent)
}
