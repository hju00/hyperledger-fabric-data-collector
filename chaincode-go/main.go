package main

import (
	"fmt"
	"smartcontractPool/smartcontractTemplate/temperature" // temperature 패키지를 import
	"smartcontractPool/smartcontractTemplate/vote"        // vote 패키지를 import

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	voteChaincode, err := contractapi.NewChaincode(new(vote.SmartContract))
	if err != nil {
		fmt.Printf("Error creating vote chaincode: %s", err.Error())
		return
	}

	temperatureChaincode, err := contractapi.NewChaincode(new(temperature.TemperatureSmartContract))
	if err != nil {
		fmt.Printf("Error creating temperature chaincode: %s", err.Error())
		return
	}

	fmt.Println("Starting chaincodes...")

	if err := voteChaincode.Start(); err != nil {
		fmt.Printf("Error starting vote chaincode: %s", err.Error())
	}

	if err := temperatureChaincode.Start(); err != nil {
		fmt.Printf("Error starting temperature chaincode: %s", err.Error())
	}
}
