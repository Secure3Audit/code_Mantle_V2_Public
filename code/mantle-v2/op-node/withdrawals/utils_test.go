package withdrawals

//import (
//	"encoding/json"
//	"math/big"
//	"os"
//	"path"
//	"testing"
//
//	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
//	"github.com/ethereum-optimism/optimism/op-node/testutils"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/stretchr/testify/require"
//)
//
//func TestParseMessagePassed(t *testing.T) {
//	tests := []struct {
//		name     string
//		file     string
//		expected *bindings.L2ToL1MessagePasserMessagePassed
//	}{
//		{
//			"withdrawal through bridge",
//			"bridge-withdrawal.json",
//			&bindings.L2ToL1MessagePasserMessagePassed{
//				Nonce:    big.NewInt(1),
//				Sender:   common.HexToAddress("0x4200000000000000000000000000000000000007"),
//				Target:   common.HexToAddress("0x6900000000000000000000000000000000000002"),
//				EthValue: big.NewInt(100000000000000000),
//				MntValue: new(big.Int),
//				GasLimit: big.NewInt(592897),
//				Data: hexutil.MustDecode(
//					"0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000000000000000000000000000000000000000090c0100000000000000000000000000000000000000000000000000000000000000a0359d9afe9cbaefaf52b443534aa33f1bb03cb679c4782b1b851a0ce8565f551400000000000000000000000000000000000000000000000000000000000001e4ff8daf150001000000000000000000000000000000000000000000000000000000000000000000000000000000000000420000000000000000000000000000000000001000000000000000000000000069000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000c41635f5fd000000000000000000000000ee3e7d56188ae7af8d5bab980908e3e91c0d7384000000000000000000000000ee3e7d56188ae7af8d5bab980908e3e91c0d7384000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000230780000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
//				),
//				WithdrawalHash: common.HexToHash("0xb1aa88c446aaccd03ab1af407c5ab28b6a3090822122959376916b88eea93850"),
//				Raw: types.Log{
//					Address: common.HexToAddress("0x4200000000000000000000000000000000000016"),
//					Topics: []common.Hash{
//						common.HexToHash("0x5da382596b838a63b4248e533d8e399b3b0f13ba6c6679f670489d44716cb173"),
//						common.HexToHash("0x0001000000000000000000000000000000000000000000000000000000000000"),
//						common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000007"),
//						common.HexToHash("0x0000000000000000000000006900000000000000000000000000000000000002"),
//					},
//					Data: hexutil.MustDecode(
//						"0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000000000000000000000000000000000000000090c0100000000000000000000000000000000000000000000000000000000000000a0359d9afe9cbaefaf52b443534aa33f1bb03cb679c4782b1b851a0ce8565f551400000000000000000000000000000000000000000000000000000000000001e4ff8daf150001000000000000000000000000000000000000000000000000000000000000000000000000000000000000420000000000000000000000000000000000001000000000000000000000000069000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000c41635f5fd000000000000000000000000ee3e7d56188ae7af8d5bab980908e3e91c0d7384000000000000000000000000ee3e7d56188ae7af8d5bab980908e3e91c0d7384000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000230780000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
//					),
//					BlockNumber: 0x80,
//					TxHash:      common.HexToHash("0xb1aa88c446aaccd03ab1af407c5ab28b6a3090822122959376916b88eea93850"),
//					TxIndex:     1,
//					BlockHash:   common.HexToHash("0x9d9308912c439ad196524cf8fc68d443095e649e0f6bb0144e231a9c3d8252c3"),
//					Index:       9,
//					Removed:     false,
//				},
//			},
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			f, err := os.Open(path.Join("testdata", test.file))
//			require.NoError(t, err)
//			dec := json.NewDecoder(f)
//			receipt := new(types.Receipt)
//			require.NoError(t, dec.Decode(receipt))
//			parsed, err := ParseMessagePassed(receipt)
//			require.NoError(t, err)
//			bytes, err := json.Marshal(parsed)
//			require.NoError(t, err)
//			t.Log(string(bytes))
//
//			// Have to do this weird thing to compare zero bigints.
//			// When they're deserialized from JSON, the internal byte
//			// array is an empty array whereas it is nil in the expectation.
//			//parsedNonce := parsed.Nonce
//			parsedETHValue := parsed.EthValue
//			parsedMNTValue := parsed.MntValue
//
//			//expNonce := test.expected.Nonce
//			expMNTValue := test.expected.MntValue
//			expETHValue := test.expected.EthValue
//
//			//testutils.RequireBigEqual(t, expNonce, parsedNonce)
//			testutils.RequireBigEqual(t, expMNTValue, parsedMNTValue)
//			testutils.RequireBigEqual(t, expETHValue, parsedETHValue)
//
//			parsed.Nonce = nil
//			parsed.EthValue = nil
//			parsed.MntValue = nil
//
//			test.expected.Nonce = nil
//			test.expected.EthValue = nil
//			test.expected.MntValue = nil
//
//			require.EqualValues(t, test.expected, parsed)
//		})
//	}
//}
