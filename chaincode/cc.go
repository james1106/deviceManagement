/*Device management*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"

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
	EmpID    string `json:"empId"`
}

// Employee to define employee structure
type Employee struct {
	ID     string  `json:"id"`
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
		fmt.Println("Error putting employees to state ledger")
		return nil, err
	}
	/* Write the assets with "assets" key state to the ledger*/
	assetbytes, err := json.Marshal(dummyAssets)
	//Writing
	err = stub.PutState(assetKey, assetbytes)
	if err != nil {
		fmt.Println("Error putting assets to state ledger")
		return nil, err
	}

	return nil, nil
}

//Assign function , expecting employee and asset
func (t *SimpleChaincode) assignAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running assign to employee")
	fmt.Printf("--------------------------------------ASSIGN ASSET METHOD----------------------------------------")

	var err error
	var selectedEmp Employee
	var employeeContainer []Employee
	var selectedAsset Asset
	var assetContainer []Asset
	var matchedEmpKey int
	//store selected asset and selected employee
	_ = json.Unmarshal([]byte(args[0]), &selectedEmp)
	_ = json.Unmarshal([]byte(args[1]), &selectedAsset)
	fmt.Printf("-------------------------------------- SELECTED EMPLOYEE----------------------------------------")
	fmt.Printf("%+v\n", selectedEmp)

	//Get employee and asset ist from state
	emplContainerbytes, err := stub.GetState("employees")
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	assetContainerbytes, err := stub.GetState("assets")
	if err != nil {
		return nil, errors.New("Failed to get state")
	}

	//convert back from bytes to perform operations
	_ = json.Unmarshal([]byte(emplContainerbytes), &employeeContainer)
	_ = json.Unmarshal([]byte(assetContainerbytes), &assetContainer)
	fmt.Printf("-------------------------------------- CURRENT employeeContainer----------------------------------------")
	fmt.Printf("%+v\n", employeeContainer)
	fmt.Printf("-------------------------------------- CURRENT ASET CONTAINER----------------------------------------")
	fmt.Printf("%+v\n", assetContainer)

	//assigning selected asset to selected employee
	for i := 0; i < len(employeeContainer); i++ {
		fmt.Printf("---------------------employee id---------------%s--------", employeeContainer[i].ID)
		fmt.Printf("---------------------selected employee id---------------%s--------", selectedEmp.ID)
		if employeeContainer[i].ID == selectedEmp.ID {
			selectedAsset.EmpID = employeeContainer[i].ID
			employeeContainer[i].Assets = append(employeeContainer[i].Assets, selectedAsset)
			fmt.Printf("-------------------------------------- AFTER UPDATING EMPLOYEE CONTAINER----------------------------------------")
			fmt.Printf("%+v\n", employeeContainer)
			matchedEmpKey = i
		}
	}
	//Update the assetContainer with emp key
	for i := 0; i < len(assetContainer); i++ {
		if assetContainer[i].SerialNo == selectedAsset.SerialNo {
			assetContainer[i].EmpID = employeeContainer[matchedEmpKey].ID
		}
	}
	fmt.Printf("-------------------------------------- UPDATED employeeContainer----------------------------------------")
	fmt.Printf("%+v\n", employeeContainer)
	fmt.Printf("-------------------------------------- UPDATED ASET CONTAINER----------------------------------------")
	fmt.Printf("%+v\n", assetContainer)

	/* Write the employee with "employees" key state to the ledger*/
	empbytes, err := json.Marshal(employeeContainer)
	//Writing
	err = stub.PutState("employees", empbytes)
	if err != nil {
		fmt.Println("Error putting employees to state ledger")
		return nil, err
	}

	/* Write the assets with "assets" key state to the ledger*/
	assetbytes, err := json.Marshal(assetContainer)
	//Writing
	err = stub.PutState("assets", assetbytes)
	if err != nil {
		fmt.Println("Error putting assets to state ledger")
		return nil, err
	}

	return nil, nil
}

//Expecting customer name and credit amount
func (t *SimpleChaincode) returnAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running return from emplyee")

	var err error
	var selectedEmp Employee
	var employeeContainer []Employee
	var selectedAsset Asset
	var assetContainer []Asset
	var matchedEmpAssetKey int
	var matchedEmpKey int
	//store selected asset and selected employee
	_ = json.Unmarshal([]byte(args[0]), &selectedEmp)
	_ = json.Unmarshal([]byte(args[1]), &selectedAsset)

	//Get employee and asset ist from state
	emplContainerbytes, err := stub.GetState("employees")
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	assetContainerbytes, err := stub.GetState("assets")
	if err != nil {
		return nil, errors.New("Failed to get state")
	}

	//convert back from bytes to perform operations
	_ = json.Unmarshal([]byte(emplContainerbytes), &employeeContainer)
	_ = json.Unmarshal([]byte(assetContainerbytes), &assetContainer)

	//finding the key of the selected asset to be removed from the selected employee
	for i := 0; i < len(employeeContainer); i++ {
		if employeeContainer[i].ID == selectedEmp.ID {
			matchedEmpKey = i
			for j := 0; j < len(employeeContainer[i].Assets); j++ {
				if selectedAsset.SerialNo == employeeContainer[i].Assets[j].SerialNo {
					matchedEmpAssetKey = j
					break
				}
			}

		}
	}
	/*REMOVE the asset from the employee by swapping it to the last element then popping out the last element from the array
	Delete without preserving order
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]*/
	employeeContainer[matchedEmpKey].Assets[matchedEmpAssetKey] = employeeContainer[matchedEmpKey].Assets[len(employeeContainer[matchedEmpKey].Assets)-1]
	employeeContainer[matchedEmpKey].Assets = employeeContainer[matchedEmpKey].Assets[:len(employeeContainer[matchedEmpKey].Assets)-1]

	//Also remove the emp id relation from the asset
	for i := 0; i < len(assetContainer); i++ {
		if assetContainer[i].SerialNo == selectedAsset.SerialNo {
			assetContainer[i].EmpID = ""
		}
	}

	/* Write the updated employee with "employees" key state to the ledger*/
	empbytes, err := json.Marshal(employeeContainer)
	//Writing
	err = stub.PutState("employees", empbytes)
	if err != nil {
		fmt.Println("Error putting employees to state ledger")
		return nil, err
	}

	/* Write the updated assets with "assets" key state to the ledger*/
	assetbytes, err := json.Marshal(assetContainer)
	//Writing
	err = stub.PutState("assets", assetbytes)
	if err != nil {
		fmt.Println("Error putting assets to state ledger")
		return nil, err
	}

	return nil, nil
}

//Add assets to system
func (t *SimpleChaincode) addAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running addAsset")

	var newAsset Asset
	var assets []Asset

	//New asset
	_ = json.Unmarshal([]byte(args[0]), &newAsset)
	fmt.Printf("--------------------------------------NEW ASSET----------------------------------------")
	fmt.Printf("%+v\n", newAsset)

	//getting current assets
	fmt.Printf("-------------------------------------CURRENTASSETS-------------------------------------")
	assetbytes, err := stub.GetState("assets")
	if err != nil {
		fmt.Printf("error getting current state of assets")
	}
	_ = json.Unmarshal([]byte(assetbytes), &assets)
	fmt.Printf("%+v\n", assets)

	//Updating assets
	fmt.Printf("--------------------------------------Updated ASSETS-----------------------------------")
	assets = append(assets, newAsset)
	fmt.Printf("%+v\n", assets)

	//Marshalling the string back to bytes to store in state
	updatedassetbytes, err := json.Marshal(assets)
	if err != nil {
		fmt.Println("Erro converting updated assets to bytes")
	}
	//Put it back to state assets
	err = stub.PutState("assets", updatedassetbytes)
	if err != nil {
		fmt.Println("Error updating the assets to ledger")
		return nil, err
	}

	return nil, nil

}

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage initialization , credit and delete of transactions.
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")

	// Handle different functions
	if function == "assignAsset" {
		// Transaction makes an assignment of assets
		fmt.Printf("Function is assignment of assets")
		return t.assignAsset(stub, args)
	} else if function == "returnAsset" {
		// Transaction makes a return of assets
		fmt.Printf("Function is debit")
		return t.returnAsset(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "addAsset" {
		//Add new assets to the system
		fmt.Printf("Function is  add asset")
		return t.addAsset(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

//Run method
func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")

	// Handle different functions
	if function == "assignAsset" {
		// Transaction makes an assignment of assets
		fmt.Printf("Function is assignment of assets")
		return t.assignAsset(stub, args)
	} else if function == "returnAsset" {
		// Transaction makes a return of assets
		fmt.Printf("Function is return of assets")
		return t.returnAsset(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "addAsset" {
		//Add new assets to the system
		fmt.Printf("Function is add asset")
		return t.addAsset(stub, args)
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

	return nil, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
