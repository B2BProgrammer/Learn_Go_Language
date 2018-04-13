package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"strconv"
)

func main(){
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Error in starting the simple chaincode",err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	fmt.Println("MedChain is Starting up")
	_,args := stub.GetFunctionAndParameters()
	var Aval int
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect arguments")
	}

	Aval,err := strconv.Atoi(arg[0])
	if err != nil {
		return shim.Error("Expecting a numeric string argument to Init()")
	}

	err = stub.PutState("hl3_ui", []byte("1.0.0"))
	if err != nil {
		return shim.Error(err.Error())
	}

	// Writing the data to the ledger
	err = stub.PutState("SelfTest",[]byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Println("- ready for action")
		return shim.Success(nil)
	}
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response  {

}
