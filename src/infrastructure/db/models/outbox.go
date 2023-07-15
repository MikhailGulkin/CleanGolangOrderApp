package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSON json.RawMessage

type Outbox struct {
	Base
	Exchange  string
	Route     string
	Payload   JSON `gorm:"type:json"`
	Processed bool
}

func GetJSONPayload(b []byte) (JSON, error) {
	result := json.RawMessage{}
	err := json.Unmarshal(b, &result)
	if err != nil {
		return JSON{}, err
	}
	return JSON(result), err
}
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
