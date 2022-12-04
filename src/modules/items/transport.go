package items

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeSaveEndpoint(service ItemService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		data, err := service.Save(req.Item)
		if err != nil {
			return createResponse{data, err.Error()}, nil
		}
		return createResponse{data, ""}, nil
	}
}

func decodeSaveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type createRequest struct {
	Item Item `json:"item"`
}

type createResponse struct {
	Data Item   `json:"data"`
	Err  string `json:"err,omitempty"`
}
