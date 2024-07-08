package main

import (
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleChaincode struct {
    contractapi.Contract
}

func (cc *SimpleChaincode) Init(ctx contractapi.TransactionContextInterface) error {
    fmt.Println("Chaincode initialized")
    return nil
}

func (cc *SimpleChaincode) Invoke(ctx contractapi.TransactionContextInterface) error {
    fmt.Println("Invoke called")
    return nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&SimpleChaincode{})
    if err != nil {
        fmt.Printf("Error creating chaincode: %s", err.Error())
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting chaincode: %s", err.Error())
    }
}
