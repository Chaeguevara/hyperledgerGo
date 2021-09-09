package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

//Init is called during chaincode instantitation to initialize any data
// 버전이 변경되지 않는다면 비어있는 init을 사용해라(From description)
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// get the args from the transaction proposal
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("잘못된 아규먼트. 키와 값을 기대함")
	}

	//최초상태 만들기
	//Set up any variable or assets here by calling stub.PutState()

	// key와 value를 장부에 저장
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("자산을 만드는데 실패함: %s", args[0]))
	}
	return shim.Success(nil)

}

//체인코드 부르기
//부르기?적용하기?(Invoke)는 체인코드의 트랜잭션당 호출됨
//각 tx는 Init으로 생긴 자산에 대해 get 또는 set실행
//set을 통해 새로운 자산 생성할수도(key-value)
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	//ChaincodeStubInterface에서 arg추출해야함
	//필요 arg는 chaincode app func의 이름임
	fn, args := stub.GetFunctionAndParameters()

	//set 인지 get인지 판별 -> shim.Success 또는 shim.Error
	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	//결과 return
	return shim.Success([]byte(result))
}

//Imple chaincode app
//chaincodestubinterface.putstte,getstate이용
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

// key에 해당하는 자산의 value값을 가져옴
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

//shim Start를 실행시키는 main
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
