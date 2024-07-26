package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AccessLogContract struct {
    contractapi.Contract
}

type Log struct {
    UserID        string `json:"user_id"`
    AccessMethod  string `json:"access_method"`
    Timestamp     string `json:"timestamp"`
}

func (c *AccessLogContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
    return nil
}

func (c *AccessLogContract) CreateLog(ctx contractapi.TransactionContextInterface, userID string, accessMethod string, timestamp string) error {
    log := Log{
        UserID:       userID,
        AccessMethod: accessMethod,
        Timestamp:    timestamp,
    }

    logAsBytes, err := json.Marshal(log)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(userID, logAsBytes)
}

func (c *AccessLogContract) QueryLog(ctx contractapi.TransactionContextInterface, userID string) (*Log, error) {
    logAsBytes, err := ctx.GetStub().GetState(userID)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if logAsBytes == nil {
        return nil, fmt.Errorf("the log %s does not exist", userID)
    }

    var log Log
    err = json.Unmarshal(logAsBytes, &log)
    if err != nil {
        return nil, err
    }

    return &log, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&AccessLogContract{})
    if err != nil {
        fmt.Printf("Error creating access log chaincode: %s", err.Error())
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting access log chaincode: %s", err.Error())
    }
}
