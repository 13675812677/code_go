package modelExtend

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"coldchain_go/model"
)
//OPERATE_GENERATE_START

//query_Generate_Start
	func query(stub shim.ChaincodeStubInterface) (error, *model.ResourceBox) {
		var resourceBox *model.ResourceBox
		
		//query_User_Start
		//query_User_End
		
	return nil, resourceBox
}
//query_Generate_End

//OPERATE_GENERATE_END