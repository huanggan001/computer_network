package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"testing"
)

func Test_de(t *testing.T) {
	client, _ := consulapi.NewClient(&consulapi.Config{Address: "127.0.0.1:8500"})
	_, err := client.Catalog().Deregister(&consulapi.CatalogDeregistration{
		Node:      "service10",
		Address:   "10.60.82.38",
		ServiceID: "10",
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
}
