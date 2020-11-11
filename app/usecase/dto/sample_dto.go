package dto

type CreateSampleRequest struct {
	ID   string
	Name string `validate:"required"`
}

type OneSampleResponse struct {
	ID   string
	Name string
}
