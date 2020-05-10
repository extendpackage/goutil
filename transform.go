package goutil

import (
	"encoding/json"
)

//TransfromObject transform `from` to `to
func TransfromObject(from interface{}, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, to)
}
