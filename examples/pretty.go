package examples

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

func MakePrettyJSON(data interface{}) error {
	bytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	result := pretty.Color(pretty.Pretty(bytes), nil)

	fmt.Println(string(result))

	return nil
}
