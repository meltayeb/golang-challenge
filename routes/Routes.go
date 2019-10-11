package routes

import (
    "github.com/go-task/handlers"
    "net/http"
)

type Route struct {
Name        string
Method      string
Pattern     string
HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

    Route{
        "TransactionIndex",
        "GET",
        "/api/payment/transaction",
        handlers.TransactionIndex,
    },

}