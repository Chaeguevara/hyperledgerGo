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

	//체인코드 부르기
	//부르기?적용하기?(Invoke)는 체인코드의 트랜잭션당 호출됨
	//각 tx는
}
