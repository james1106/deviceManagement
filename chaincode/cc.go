/*Device management*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example to simple Chaincode implementation
type SimpleChaincode struct {
}

// Asset to define asset stucture
type Asset struct {
	Make     string `json:"make"`
	Type     string `json:"type"`
	Model    string `json:"model"`
	SerialNo string `json:"serialNo"`
}

// Employee to define employee structure
type Employee struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Assets []Asset `json:"assets"`
}

//Init Method to initialize employees and assets
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")

	var err error
	var dummyEmp [3]Employee
	var dummyAssets [3]Asset
	//Error for wrong input
	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3 employee's data")
	}
	var employeeKey string
	var assetKey string
	employeeKey = "employees"
	assetKey = "assets"
	//Adding initial employees to Employee structure
	_ = json.Unmarshal([]byte(args[0]), &dummyEmp[0])
	_ = json.Unmarshal([]byte(args[1]), &dummyEmp[1])
	_ = json.Unmarshal([]byte(args[2]), &dummyEmp[2])

	//Addign initial assets to Asset structure
	_ = json.Unmarshal([]byte(args[3]), &dummyAssets[0])
	_ = json.Unmarshal([]byte(args[4]), &dummyAssets[1])
	_ = json.Unmarshal([]byte(args[5]), &dummyAssets[2])

	/* Write the employee with "employees" key state to the ledger*/
	empbytes, err := json.Marshal(dummyEmp)
	//Writing
	err = stub.PutState(employeeKey, empbytes)
	if err != nil {
		return nil, err
	}
	/* Write the assets with "assets" key state to the ledger*/
	assetbytes, err := json.Marshal(dummyAssets)
	//Writing
	err = stub.PutState(assetKey, assetbytes)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//Credit Function
//Expecting customer name and credit amount
func (t *SimpleChaincode) credit(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running credit")

	var custName string  // Customer Name
	var currentBal int   //Customer Balance  --Aval->currentBal
	var creditAmount int //Credit Amount
	var err error

	//Error for wrong input
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2- customer name and credit amount")
	}

	custName = args[0]

	// Get the current state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger

	custCurrentBalbytes, err := stub.GetState(custName)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if custCurrentBalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	currentBal, _ = strconv.Atoi(string(custCurrentBalbytes))

	// Credit Execution
	creditAmount, err = strconv.Atoi(args[1])
	currentBal = currentBal + creditAmount
	fmt.Printf("currentBal = %d\n", currentBal)

	// Write the state back to the ledger
	err = stub.PutState(custName, []byte(strconv.Itoa(currentBal)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//Expecting customer name and credit amount
func (t *SimpleChaincode) debit(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running credit")

	var custName string // Customer Name
	var currentBal int  //Customer Balance
	var debitAmount int
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	custName = args[0]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	custCurrentBalbytes, err := stub.GetState(custName)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if custCurrentBalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	currentBal, _ = strconv.Atoi(string(custCurrentBalbytes))

	// Debit Execution
	debitAmount, err = strconv.Atoi(args[1])
	currentBal = currentBal - debitAmount

	fmt.Printf("currentBal = %d\n", currentBal)

	// Write the state back to the ledger
	err = stub.PutState(custName, []byte(strconv.Itoa(currentBal)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Deletes a customer from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running deletion of customer")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 customer name")
	}

	cust := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(cust)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

//Updates the address
// func (t *SimpleChaincode) updateAddress(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
// 	fmt.Printf("Running updateAddress")

// 	var custAddress string
// 	var custAddressKey string //Customer address key to read write in ledger as key value of address
// 	var err error

// 	//Error for wrong input
// 	if len(args) != 2 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting 2- customer name and new address")
// 	}
// 	custAddress = args[1]
// 	custAddressKey = args[0] + "Add"

// 	fmt.Printf("new address :", custAddress)

// 	// Write the state back to the ledger with new address
// 	err = stub.PutState(custAddressKey, []byte(custAddress))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil, nil

// }

//Updates the address
func (t *SimpleChaincode) addAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running addAsset")

	// var newAsset Asset
	// var cuurentAssets []Asset
	// assetbytes, err := stub.GetState(args[0])
	// _ = json.Unmarshal([]byte(assetbytes), &cuurentAssets)

	//_ = json.Unmarshal([]byte(args[0]), &newAsset)
	// var custAddressKey string //Customer address key to read write in ledger as key value of address
	// var err error

	// //Error for wrong input
	// if len(args) != 2 {
	// 	return nil, errors.New("Incorrect number of arguments. Expecting 2- customer name and new address")
	// }
	// custAddress = args[1]
	// custAddressKey = args[0] + "Add"

	// fmt.Printf("new address :", custAddress)

	// // Write the state back to the ledger with new address
	// err = stub.PutState(custAddressKey, []byte(custAddress))
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil

}

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage initialization , credit and delete of transactions.
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")

	// Handle different functions
	if function == "credit" {
		// Transaction makes a credit to the customer
		fmt.Printf("Function is credit")
		return t.credit(stub, args)
	} else if function == "debit" {
		// Transaction makes a debit from the customer
		fmt.Printf("Function is debit")
		return t.debit(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an customer from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	} else if function == "updateAddress" {
		//Update Address
		fmt.Printf("Function is updated address")
		//return t.updateAddress(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")

	// Handle different functions
	if function == "credit" {
		// Transaction makes a credit to the customer
		fmt.Printf("Function is credit")
		return t.credit(stub, args)
	} else if function == "debit" {
		// Transaction makes a debit from the customer
		fmt.Printf("Function is debit")
		return t.debit(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	} else if function == "updateAddress" {
		//Update Address
		fmt.Printf("Function is updated address")
		//return t.updateAddress(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")

	if function != "query" {
		fmt.Printf("Function is query")
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	//get Employees
	if args[0] == "employees" {
		var e []Employee
		empbytes, err := stub.GetState(args[0])
		if err != nil {
			fmt.Println("error getting employees from state")
		}
		err = json.Unmarshal(empbytes, &e)
		return empbytes, nil
		// fmt.Println("--------------------------------------------------------------------------------------")
		// fmt.Println(e[0])
		// fmt.Println("--------------------------------------------------------------------------------------")
		// fmt.Println(e[0].Name)
		// fmt.Println("--------------------------------------------------------------------------------------")
		// fmt.Println(e[0].Assets[0].SerialNo)
		// fmt.Println("--------------------------------------------------------------------------------------")
		// fmt.Println(e[1].Assets[0].SerialNo)
		// fmt.Println("--------------------------------------------------------------------------------------")
		// fmt.Println(e[2].Assets[0].SerialNo)
		// fmt.Println("--------------------------------------------------------------------------------------")
		// fmt.Println(e[2].Assets[1].SerialNo)
	} else if args[0] == "assets" {
		//get Assets
		var a []Asset
		assetbytes, err := stub.GetState(args[0])
		if err != nil {
			fmt.Println("error getting assets from state")
		}
		err = json.Unmarshal(assetbytes, &a)
		return assetbytes, nil
	}

	// var custName string       // Entities
	// var custAddressKey string //Customer address key to read write in ledger as key value of Address
	// var resp []byte           //response result based on query key

	// if len(args) != 2 {
	// 	return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query and balance or address")
	// }

	// custName = args[0]
	// //Check query key for Balance and Address
	// if args[1] == "Balance" {
	// 	custAvailBalbytes, err := stub.GetState(custName)
	// 	if err != nil {
	// 		jsonResp := "{\"Error\":\"Failed to get state for " + custName + "\"}"
	// 		return nil, errors.New(jsonResp)
	// 	}
	// 	if custAvailBalbytes == nil {
	// 		jsonResp := "{\"Error\":\"Nil amount for " + custName + "\"}"
	// 		return nil, errors.New(jsonResp)
	// 	}
	// 	jsonResp := "{\"Name\":\"" + custName + "\",\"Amount\":\"" + string(custAvailBalbytes) + "\"}"
	// 	fmt.Printf("Query Response:%s\n", jsonResp)
	// 	resp = custAvailBalbytes
	// } else if args[1] == "Address" {
	// 	custAddressKey = args[0] + "Add"
	// 	custAddressbytes, err := stub.GetState(custAddressKey)
	// 	if err != nil {
	// 		jsonResp := "{\"Error\":\"Failed to get state for " + custAddressKey + "\"}"
	// 		return nil, errors.New(jsonResp)
	// 	}
	// 	if custAddressbytes == nil {
	// 		jsonResp := "{\"Error\":\"No address for " + custName + "\"}"
	// 		return nil, errors.New(jsonResp)
	// 	}
	// 	jsonResp := "{\"Name\":\"" + custName + "\",\"Address\":\"" + string(custAddressbytes) + "\"}"
	// 	fmt.Printf("Query Response:%s\n", jsonResp)
	// 	resp = custAddressbytes
	// }

	return nil, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
