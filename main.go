package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := GetDBconn()
	r := mux.NewRouter()
	var svc AccountService
	svc = accountservice{}
	{
		repository, err := NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
		}
		svc = NewService(repository, logger)
	}

	CreateAccountHandler := httptransport.NewServer(
		makeCreateCustomerEndpoint(svc),
		decodeCreateCustomerRequest,
		encodeResponse,
	)
	GetByCustomerIdHandler := httptransport.NewServer(
		makeGetCustomerByIdEndpoint(svc),
		decodeGetCustomerByIdRequest,
		encodeResponse,
	)
	GetAllCustomersHandler := httptransport.NewServer(
		makeGetAllCustomerEndpoint(svc),
		decodeGetAllCustomersRequst,
		encodeResponse,
	)
	DeleteCustomerHandler := httptransport.NewServer(
		makeDeleteCustomerEndpoint(svc),
		decodeDeleteCustomerRequest,
		encodeResponse,
	)

	UpdateCustomerHandler := httptransport.NewServer(
		makeUppdateCustomerEndpoint(svc),
		decodeUpdateCustomerRequest,
		encodeResponse,
	)

	http.Handle("/", r)
	http.Handle("/account", CreateAccountHandler)
	r.Handle("/account", UpdateCustomerHandler).Methods("PUT")
	r.Handle("/account/{customerid}", GetByCustomerIdHandler).Methods("GET")
	r.Handle("/account", GetAllCustomersHandler).Methods("GET")
	r.Handle("/account/{customerid}", DeleteCustomerHandler).Methods("DELETE")

	logger.Log("msg", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
