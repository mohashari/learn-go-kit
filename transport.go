package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

func makeCreateCustomerEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequets)
		msg, err := s.CreateCustomer(ctx, req.customer)
		return CreateCustomerResponse{Msg: msg, Err: err}, nil
	}
}

func makeGetCustomerByIdEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIdRequest)
		customerDetails, err := s.GetCustomerById(ctx, req.Id)
		if err != nil {
			return GetCustomerByIdResponse{
				Customer: customerDetails,
				Err:      "Id not found",
			}, nil
		}
		return GetCustomerByIdResponse{
			Customer: customerDetails,
			Err:      "",
		}, nil
	}
}

func makeGetAllCustomerEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customerDetails, err := s.GetAllCustomers(ctx)
		if err != nil {
			return GetAllCustomersResponse{
				Customer: customerDetails,
				Err:      "no data found",
			}, nil

		}
		return GetAllCustomersResponse{
			Customer: customerDetails,
			Err:      "",
		}, nil
	}
}

func makeDeleteCustomerEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCustomerRequest)
		msg, err := s.DeleteCustomer(ctx, req.CustomerId)
		if err != nil {
			return DeleteCustomerResponse{
				Msg: msg,
				Err: err,
			}, nil
		}
		return DeleteCustomerResponse{
			Msg: msg,
			Err: nil,
		}, nil
	}
}

func makeUppdateCustomerEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCustomerRequest)
		msg, err := s.UpdateCustomer(ctx, req.customer)
		return msg, err
	}
}

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateCustomerRequets
	fmt.Println("---------->>> into decodeCreateCustomerRequest decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetCustomerByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetCustomerByIdRequest
	fmt.Println("---------->>> into decodeGetCustomerByIdRequest decoding")
	vars := mux.Vars(r)
	req = GetCustomerByIdRequest{
		Id: vars["customerid"],
	}
	return req, nil
}

func decodeGetAllCustomersRequst(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("---------->>> into decodeGetAllCustomersRequst decoding")
	var req GetAllCustomersRequest
	return req, nil
}

func decodeDeleteCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("---------->>> into decodeDeleteCustomerRequest decoding")
	var req DeleteCustomerRequest
	vars := mux.Vars(r)
	req = DeleteCustomerRequest{
		CustomerId: vars["customerid"],
	}
	return req, nil
}

func decodeUpdateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("---------->>> into decodeUpdateCustomerRequest decoding")
	var req UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<-----------")
	return json.NewEncoder(w).Encode(response)
}

type (
	CreateCustomerRequets struct {
		customer Customer
	}
	CreateCustomerResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	GetCustomerByIdRequest struct {
		Id string `json:"customerid"`
	}
	GetCustomerByIdResponse struct {
		Customer interface{} `json:"customer,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	GetAllCustomersRequest struct {
	}
	GetAllCustomersResponse struct {
		Customer interface{} `json:"customer,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	DeleteCustomerRequest struct {
		CustomerId string `json:"customerid"`
	}
	DeleteCustomerResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
	UpdateCustomerRequest struct {
		customer Customer
	}
	UpdateCustomerResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
)
