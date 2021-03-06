package main

import (
	"dataMarket/context"
	"dataMarket/contracts"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	announcementContract := new(contracts.AnnouncementContract)
	announcementContract.Name = "AnnouncementContract"
	announcementContract.TransactionContextHandler = new(context.TransactionContext)
	announcementContract.BeforeTransaction = context.SearchIdentitiesHandler

	identificationContract := new(contracts.IdentificationContract)
	identificationContract.Name = "IdentificationContract"
	identificationContract.TransactionContextHandler = new(context.TransactionContext)
	identificationContract.BeforeTransaction = context.SearchIdentitiesHandler

	queryContract := new(contracts.QueryContract)
	queryContract.Name = "QueryContract"
	queryContract.TransactionContextHandler = new(context.TransactionContext)
	queryContract.BeforeTransaction = context.SearchIdentitiesHandler

	categoryContract := new(contracts.CategoryContract)
	categoryContract.Name = "CategoryContract"

	chaincode, err := contractapi.NewChaincode(announcementContract, identificationContract, queryContract, categoryContract)

	if err != nil {
		panic(err.Error())
	}

	if err := chaincode.Start(); err != nil {
		panic(err.Error())
	}
}
