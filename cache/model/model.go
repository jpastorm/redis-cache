package model

import "encoding/json"

type Cache struct {
	Id   uint
	Name string
	Data json.RawMessage
}
