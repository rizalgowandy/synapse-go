package entity

import "fmt"

type ErrResp struct {
	ErrDetailResp struct {
		Code string `json:"code"`
		En   string `json:"en"`
	} `json:"error"`
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Success   *bool  `json:"success,omitempty"`
}

func (e ErrResp) Valid() bool {
	if e.Success == nil {
		return false
	}
	return !*e.Success
}

func (e ErrResp) Error() string {
	if e.ErrDetailResp.En != "" {
		return fmt.Sprintf("synapse: %s", e.ErrDetailResp.En)
	}
	if e.ErrDetailResp.Code != "" {
		return fmt.Sprintf("synapse: %s", e.ErrDetailResp.Code)
	}
	return fmt.Sprintf("synapse: %s", e.ErrorCode)
}
