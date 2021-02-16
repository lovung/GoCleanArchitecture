package presenter

import (
	"github.com/google/go-cmp/cmp"
)

// Response represents the response of every request
type Response struct {
	Meta   MetaResponse   `json:"meta,omitempty"`
	Data   interface{}    `json:"data,omitempty"`
	Errors ErrorResponses `json:"errors,omitempty"`
}

// IsEmpty check if the struct is empty or not
func (r Response) IsEmpty() bool {
	return cmp.Equal(r, Response{})
}

// MetaResponse represents meta-information
type MetaResponse struct {
	Code           int    `json:"code,omitempty"`
	Message        string `json:"message,omitempty"`
	Total          uint64 `json:"total,omitempty"`
	NextCursor     string `json:"next_cursor,omitempty"`
	PreviousCursor string `json:"previous_cursor,omitempty"`
}

// PagingRequest request
type PagingRequest struct {
	Size   uint64 `mapstructure:"size" json:"page[size]"`
	Number uint64 `mapstructure:"number" json:"page[number]"`
}

// IDResponse represent the common data response for many requests
type IDResponse struct {
	ID interface{} `json:"id,omitempty"`
}
