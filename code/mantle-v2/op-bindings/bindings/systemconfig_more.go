// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1012_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1011_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1013_storage\"},{\"astId\":1010,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"baseFee\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c80638da5cb5b116100d8578063cc731b021161008c578063f45e65d811610066578063f45e65d814610443578063f68016b71461044c578063ffa1ad741461046057600080fd5b8063cc731b02146102f3578063e81b2c6d14610427578063f2fde38b1461043057600080fd5b8063b40a817c116100bd578063b40a817c146102ba578063c71973f6146102cd578063c9b26f61146102e057600080fd5b80638da5cb5b14610289578063935f029e146102a757600080fd5b80634add321d1161012f57806354fd4d501161011457806354fd4d50146102635780636ef25c3a14610278578063715018a61461028157600080fd5b80634add321d1461021b5780634f16540b1461023c57600080fd5b806318d139181161016057806318d13918146101ad5780631fd19ee1146101c0578063468606981461020857600080fd5b806304c7a26f1461017c5780630c18c16214610191575b600080fd5b61018f61018a3660046114d9565b610468565b005b61019a60655481565b6040519081526020015b60405180910390f35b61018f6101bb366004611555565b61070a565b7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08545b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a4565b61018f610216366004611577565b6107ce565b6102236107fe565b60405167ffffffffffffffff90911681526020016101a4565b61019a7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b61026b610829565b6040516101a4919061160a565b61019a606a5481565b61018f6108cc565b60335473ffffffffffffffffffffffffffffffffffffffff166101e3565b61018f6102b536600461161d565b6108e0565b61018f6102c836600461163f565b610979565b61018f6102db36600461165a565b610a5f565b61018f6102ee366004611577565b610a73565b6103b76040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516101a49190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b61019a60675481565b61018f61043e366004611555565b610aa3565b61019a60665481565b6068546102239067ffffffffffffffff1681565b61019a600081565b600054610100900460ff16158080156104885750600054600160ff909116105b806104a25750303b1580156104a2575060005460ff166001145b610533576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561059157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610599610b73565b6105a289610aa3565b606588905560668790556067869055606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8716179055606a8490557f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0883905561061782610c12565b61061f6107fe565b67ffffffffffffffff168567ffffffffffffffff16101561069c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f7700604482015260640161052a565b80156106ff57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b610712611086565b61073a817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516107c2919061160a565b60405180910390a35050565b6107d6611086565b606a819055604080516020808201849052825180830390910181529082019091526004610791565b6069546000906108249063ffffffff6a01000000000000000000008204811691166116a5565b905090565b60606108547f0000000000000000000000000000000000000000000000000000000000000000611107565b61087d7f0000000000000000000000000000000000000000000000000000000000000000611107565b6108a67f0000000000000000000000000000000000000000000000000000000000000000611107565b6040516020016108b8939291906116d1565b604051602081830303815290604052905090565b6108d4611086565b6108de6000611244565b565b6108e8611086565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161096c919061160a565b60405180910390a3505050565b610981611086565b6109896107fe565b67ffffffffffffffff168167ffffffffffffffff161015610a06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f7700604482015260640161052a565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610791565b610a67611086565b610a7081610c12565b50565b610a7b611086565b6067819055604080516020808201849052825180830390910181529082019091526000610791565b610aab611086565b73ffffffffffffffffffffffffffffffffffffffff8116610b4e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161052a565b610a7081611244565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b600054610100900460ff16610c0a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161052a565b6108de6112bb565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff161115610cc2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d617820626173650000000000000000000000606482015260840161052a565b6001816040015160ff1611610d59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e20310000000000000000000000000000000000606482015260840161052a565b6068546080820151825167ffffffffffffffff90921691610d7a9190611747565b63ffffffff161115610de8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f7700604482015260640161052a565b6000816020015160ff1611610e7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f7420626520300000000000000000000000000000000000606482015260840161052a565b8051602082015163ffffffff82169160ff90911690610e9f908290611795565b610ea991906117b8565b63ffffffff1614610f3c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d6974000000000000000000606482015260840161052a565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b60335473ffffffffffffffffffffffffffffffffffffffff1633146108de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161052a565b60608160000361114a57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611174578061115e816117e4565b915061116d9050600a8361181c565b915061114e565b60008167ffffffffffffffff81111561118f5761118f61139c565b6040519080825280601f01601f1916602001820160405280156111b9576020820181803683370190505b5090505b841561123c576111ce600183611830565b91506111db600a86611847565b6111e690603061185b565b60f81b8183815181106111fb576111fb611873565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611235600a8661181c565b94506111bd565b949350505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611352576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161052a565b6108de33611244565b803573ffffffffffffffffffffffffffffffffffffffff8116811461137f57600080fd5b919050565b803567ffffffffffffffff8116811461137f57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803563ffffffff8116811461137f57600080fd5b803560ff8116811461137f57600080fd5b80356fffffffffffffffffffffffffffffffff8116811461137f57600080fd5b600060c0828403121561142257600080fd5b60405160c0810181811067ffffffffffffffff8211171561146c577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290508061147b836113cb565b8152611489602084016113df565b602082015261149a604084016113df565b60408201526114ab606084016113cb565b60608201526114bc608084016113cb565b60808201526114cd60a084016113f0565b60a08201525092915050565b6000806000806000806000806101a0898b0312156114f657600080fd5b6114ff8961135b565b975060208901359650604089013595506060890135945061152260808a01611384565b935060a0890135925061153760c08a0161135b565b91506115468a60e08b01611410565b90509295985092959890939650565b60006020828403121561156757600080fd5b6115708261135b565b9392505050565b60006020828403121561158957600080fd5b5035919050565b60005b838110156115ab578181015183820152602001611593565b838111156115ba576000848401525b50505050565b600081518084526115d8816020860160208601611590565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061157060208301846115c0565b6000806040838503121561163057600080fd5b50508035926020909101359150565b60006020828403121561165157600080fd5b61157082611384565b600060c0828403121561166c57600080fd5b6115708383611410565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff8083168185168083038211156116c8576116c8611676565b01949350505050565b600084516116e3818460208901611590565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161171f816001850160208a01611590565b6001920191820152835161173a816002840160208801611590565b0160020195945050505050565b600063ffffffff8083168185168083038211156116c8576116c8611676565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600063ffffffff808416806117ac576117ac611766565b92169190910492915050565b600063ffffffff808316818516818304811182151516156117db576117db611676565b02949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361181557611815611676565b5060010190565b60008261182b5761182b611766565b500490565b60008282101561184257611842611676565b500390565b60008261185657611856611766565b500690565b6000821982111561186e5761186e611676565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}