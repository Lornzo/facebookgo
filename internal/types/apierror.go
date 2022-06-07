package types

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	Message      string `json:"message"`
	Type         string `json:"type"`
	Code         int64  `json:"code"`
	ErrorSubCode int64  `json:"error_subcode"`
	FBTraceID    string `json:"fbtrace_id"`
}

func (ae APIError) Error() error {

	var (
		jsonString string
		err        error
	)

	if ae.Code == 0 {
		return nil
	}

	if jsonString, err = ae.ToJson(); err != nil {
		return err
	}

	return fmt.Errorf(jsonString)
}

func (ae APIError) ToJson() (string, error) {

	var (
		bytes []byte
		err   error
	)

	if bytes, err = json.Marshal(ae); err != nil {
		return "", err
	}

	return string(bytes), nil
}
