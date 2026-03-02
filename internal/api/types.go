package api

type BodyCreateRequest struct {
	Url string `json:"url"`
}

type BodyCreateResponce struct {
	Data string `json:"data"`
}

type BodyGetRequest struct {
	Data string `json:"data"`
}

type BodyGetResponse struct {
	Url string `json:"url"`
}