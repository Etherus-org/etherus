package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"

	ethUtils "github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core"
)

// defaultGenesisBlob is the JSON representation of the default
// genesis file in $GOPATH/src/github.com/ya-enot/etherus/setup/genesis.json
// nolint=lll

// Get from remix storage with:
// JSON.stringify(Object.values(<storageObject>).reduce((acc, v) => {acc[v.key] = v.value; return acc}, {}), null, '   ')
var defaultGenesisBlob = []byte(`
{
    "config": {
        "chainId": 32019,
        "homesteadBlock": 0,
        "byzantiumBlock": 0,
        "eip150Block": 0,
        "eip155Block": 0,
        "eip158Block": 0
    },
    "nonce": "0xdeadbeefdeadbeef",
    "timestamp": "0x00",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "difficulty": "0x40",
    "gasLimit": "0x8000000",
    "alloc": {
		"0x0000000000000000000000000000000000000001": {	"balance": "1" },
		"0x0000000000000000000000000000000000000002": {	"balance": "1" },
		"0x0000000000000000000000000000000000000003": {	"balance": "1" },
		"0x0000000000000000000000000000000000000004": {	"balance": "1" },
		"0x0000000000000000000000000000000000000005": {	"balance": "1" },
		"0x0000000000000000000000000000000000000006": {	"balance": "1" },
		"0x0000000000000000000000000000000000000007": {	"balance": "1" },
		"0x0000000000000000000000000000000000000008": {	"balance": "1" },
		"0x0000000000000000000000000000000000000fff": {
			"balance": "2500000000000000000000",
			"code": "0x60806040526004361061011f5760003560e01c80638e19899e116100a0578063bfe0eb5211610064578063bfe0eb5214610414578063d5f20ff614610429578063d64c2ead146104a7578063f1f58d92146104d7578063f2fde38b146105035761011f565b80638e19899e1461035757806392c8a966146103815780639da0e3e3146103ab578063a36147d3146103d5578063acfb8f76146103ff5761011f565b8063241b71bb116100e7578063241b71bb146102925780632d5f512a146102d057806363338b1714610318578063715018a61461032d5780638da5cb5b146103425761011f565b80630b8604fc14610124578063155e1abc1461016a578063181a8d68146101915780631c32d9821461023f57806320aef8d41461025e575b600080fd5b34801561013057600080fd5b5061014e6004803603602081101561014757600080fd5b5035610536565b604080516001600160a01b039092168252519081900360200190f35b34801561017657600080fd5b5061017f610551565b60408051918252519081900360200190f35b34801561019d57600080fd5b506101a6610557565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156101ea5781810151838201526020016101d2565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610229578181015183820152602001610211565b5050505090500194505050505060405180910390f35b61025c6004803603602081101561025557600080fd5b5035610604565b005b61025c6004803603606081101561027457600080fd5b508035906001600160a01b0360208201358116916040013516610613565b34801561029e57600080fd5b506102bc600480360360208110156102b557600080fd5b503561088d565b604080519115158252519081900360200190f35b3480156102dc57600080fd5b5061025c600480360360808110156102f357600080fd5b50803590602081013590604081013560ff1690606001356001600160601b03166108ae565b34801561032457600080fd5b5061017f610c6b565b34801561033957600080fd5b5061025c610cbb565b34801561034e57600080fd5b5061014e610d1a565b34801561036357600080fd5b5061025c6004803603602081101561037a57600080fd5b5035610d29565b34801561038d57600080fd5b506102bc600480360360208110156103a457600080fd5b5035610f5d565b3480156103b757600080fd5b5061025c600480360360208110156103ce57600080fd5b5035610f81565b3480156103e157600080fd5b5061014e600480360360208110156103f857600080fd5b5035611293565b34801561040b57600080fd5b506101a66112dc565b34801561042057600080fd5b5061017f6113e1565b34801561043557600080fd5b506104536004803603602081101561044c57600080fd5b50356113e7565b604080516001600160a01b0397881681526001600160601b03968716602082015265ffffffffffff9095168582015260ff9093166060850152931660808301529190921660a0830152519081900360c00190f35b3480156104b357600080fd5b5061025c600480360360408110156104ca57600080fd5b5080359060200135611444565b3480156104e357600080fd5b5061025c600480360360208110156104fa57600080fd5b50351515611466565b34801561050f57600080fd5b5061025c6004803603602081101561052657600080fd5b50356001600160a01b0316611490565b6000908152600360205260409020546001600160a01b031690565b60015481565b60608060058054806020026020016040519081016040528092919081815260200182805480156105a657602002820191906000526020600020905b815481526020019060010190808311610592575b5050505050915060048054806020026020016040519081016040528092919081815260200182805480156105f957602002820191906000526020600020905b8154815260200190600101908083116105e5575b505050505090509091565b61061081600080610613565b50565b60015434101561065757604051600160e51b62461bcd0281526004018080602001828103825260258152602001806118396025913960400191505060405180910390fd5b600083815260036020526040902080546001600160601b03600160a01b80830482163401909116026001600160a01b039182161780835516610771576001600160a01b0383166106f15760408051600160e51b62461bcd02815260206004820152601b60248201527f506c656173652073706563696679206e6f646520616464726573730000000000604482015290519081900360640190fd5b80546001600160a01b0319166001600160a01b03848116919091178255600182018054600160301b7fff0000000000000000000000000000000000000000ffffffffffffffffffffff909116600160581b938616939093029290921766ff00000000000019169190911765ffffffffffff19164365ffffffffffff161790555b6001810154600160381b900463ffffffff161561079757610791816114b0565b50610835565b60006107a2826114f6565b60048054600181810183557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9091018890556005805480830182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0019290925554908301805463ffffffff909216600160381b026affffffff0000000000000019909216919091179055505b341561088757805460408051868152600160a01b9092046001600160601b0316602083015280517fd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf07889281900390910190a15b50505050565b600090815260036020526040902060010154600160301b900460ff16151590565b6108b78461088d565b1580156108c857506108c884610f5d565b61090657604051600160e51b62461bcd0281526004018080602001828103825260318152602001806118b06031913960400191505060405180910390fd5b600160ff831610156109625760408051600160e51b62461bcd02815260206004820152601860248201527f596f752073686f756c6420737065636966792063617573650000000000000000604482015290519081900360640190fd5b3361096c84610536565b6001600160a01b0316146109b457604051600160e51b62461bcd0281526004018080602001828103825260338152602001806119606033913960400191505060405180910390fd5b60ff8216600114806109d057506000546001600160a01b031633145b806109dd575060065460ff165b610a315760408051600160e51b62461bcd02815260206004820152600e60248201527f57726f6e672070756e6973686572000000000000000000000000000000000000604482015290519081900360640190fd5b60ff82166002141580610a525750655af3107a40006001600160601b038216105b610aa65760408051600160e51b62461bcd02815260206004820152601360248201527f57726f6e6720626c6f636b73206e756d62657200000000000000000000000000604482015290519081900360640190fd5b828414610b6757610ab683610f5d565b8015610ac85750610ac68361088d565b155b610b0657604051600160e51b62461bcd0281526004018080602001828103825260288152602001806118116028913960400191505060405180910390fd5b60ff821660011415610b625760408051600160e51b62461bcd02815260206004820152601e60248201527f546869732070617573696e67206973206e6f7420766f6c756e74617279210000604482015290519081900360640190fd5b610bc2565b60ff8216600114610bc25760408051600160e51b62461bcd02815260206004820152601b60248201527f596f75206172652070617573696e6720766f6c756e746172696c790000000000604482015290519081900360640190fd5b600084815260036020526040902060018101805465ffffffffffff19164365ffffffffffff161766ff0000000000001916600160301b60ff8616021790556002810180546001600160601b0319166001600160601b038416179055610c26816114b0565b506040805186815260ff8516602082015281517f2ffa21f275941345406452ac41a8bceea9dde4f76b233ac4fc9e53c31e399342929181900390910190a15050505050565b600554600090815b81811015610cb657600060058281548110610c8a57fe5b6000918252602090912001549050600160e01b60ff028116610cad578360010193505b50600101610c73565b505090565b6000546001600160a01b03163314610cd257600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031681565b610d3281610f5d565b610d865760408051600160e51b62461bcd02815260206004820152601760248201527f596f752073686f756c642068617665206465706f736974000000000000000000604482015290519081900360640190fd5b610d8f8161088d565b610de35760408051600160e51b62461bcd02815260206004820152601460248201527f596f752073686f756c6420626520706175736564000000000000000000000000604482015290519081900360640190fd5b600081815260036020526040902080546001600160a01b03163314610e525760408051600160e51b62461bcd02815260206004820152601e60248201527f4f6e6c79206e6f646520697473656c662063616e207769746864726177210000604482015290519081900360640190fd5b600181810154600160301b900460ff1614610ea157604051600160e51b62461bcd0281526004018080602001828103825260328152602001806118e16032913960400191505060405180910390fd5b600181015465ffffffffffff43811691811661177001161115610ef857604051600160e51b62461bcd02815260040180806020018281038252602b815260200180611885602b913960400191505060405180910390fd5b8054600160a01b90046001600160601b03166000610f1584611293565b9050610f2084611534565b6040516001600160a01b0382169083156108fc029084906000818181858888f19350505050158015610f56573d6000803e3d6000fd5b5050505050565b600090815260036020526040902054600160a01b90046001600160601b0316151590565b610f8a8161088d565b610fde5760408051600160e51b62461bcd02815260206004820152601860248201527f54686973206e6f6465206973206e6f7420706175736564210000000000000000604482015290519081900360640190fd5b600081815260036020526040902080546001600160a01b0316331461104d5760408051600160e51b62461bcd02815260206004820152601d60248201527f596f752063616e206f6e6c7920756e706175736520796f757273656c66000000604482015290519081900360640190fd5b600181810154600160301b900460ff169081148061106e575060ff81166002145b8061107c575060ff81166003145b6110ba57604051600160e51b62461bcd0281526004018080602001828103825260238152602001806119136023913960400191505060405180910390fd5b60ff8116600214156111275760028201546001830154436001600160601b0390921665ffffffffffff90911601111561112757604051600160e51b62461bcd02815260040180806020018281038252602a815260200180611936602a913960400191505060405180910390fd5b60ff8116600314156111c057600282015482546001600160601b03918216600160a01b909104909116101561119057604051600160e51b62461bcd02815260040180806020018281038252602781526020018061185e6027913960400191505060405180910390fd5b600282015482546001600160a01b0381166001600160601b03928316600160a01b92839004841603909216021782555b6002548254600160a01b90046001600160601b0316101561121557604051600160e51b62461bcd0281526004018080602001828103825260348152602001806119936034913960400191505060405180910390fd5b60018201805466ffffffffffffff191690556002820180546001600160601b0319169055611242826114b0565b50815460408051858152600160a01b9092046001600160601b0316602083015280517fd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf07889281900390910190a1505050565b600081815260036020526040812060010154600160581b90046001600160a01b0316806112c6576112c383610536565b90505b6001600160a01b0381166112d657fe5b92915050565b60608060006112e9610c6b565b905080604051908082528060200260200182016040528015611315578160200160208202803883390190505b50925080604051908082528060200260200182016040528015611342578160200160208202803883390190505b506005549092506000805b828110156113d95760006005828154811061136457fe5b6000918252602090912001549050600160e01b60ff0281166113d0578087848151811061138d57fe5b602002602001018181525050600482815481106113a657fe5b90600052602060002001548684815181106113bd57fe5b6020026020010181815250508260010192505b5060010161134d565b505050509091565b60025481565b6000908152600360205260409020805460018201546002909201546001600160a01b03808316946001600160601b03600160a01b90940484169465ffffffffffff81169460ff600160301b830416941692600160581b9091041690565b6000546001600160a01b0316331461145b57600080fd5b600291909155600155565b6000546001600160a01b0316331461147d57600080fd5b6006805460ff1916911515919091179055565b6000546001600160a01b031633146114a757600080fd5b61061081611758565b60006114bb826114f6565b600560018460010160079054906101000a900463ffffffff160363ffffffff16815481106114e557fe5b600091825260209091200155919050565b8054600190910154600160301b900460ff1660e01b6001600160a01b038216600160a01b9092046001600160601b031660201c60a01b919091171790565b60008181526003602081905260409091208054600491600591600160a01b90046001600160601b03166115b15760408051600160e51b62461bcd02815260206004820152601560248201527f4465706f7369742073686f756c64206578697374210000000000000000000000604482015290519081900360640190fd5b82546115b957fe5b6001818101548454600160381b90910463ffffffff1660001901911080156115ec575083546000190163ffffffff821614155b156116c45783548490600019810190811061160357fe5b9060005260206000200154848263ffffffff168154811061162057fe5b60009182526020909120015582548390600019810190811061163e57fe5b9060005260206000200154838263ffffffff168154811061165b57fe5b90600052602060002001819055506000848263ffffffff168154811061167d57fe5b600091825260208083209190910154825287905260409020600190810180546affffffff000000000000001916600160381b92850163ffffffff1692909202919091179055505b8354849060001981019081106116d657fe5b600091825260208220015583546116f18560001983016117c6565b5082548390600019810190811061170457fe5b6000918252602082200155825461171f8460001983016117c6565b50505060009384525050602052604081209081556001810180546001600160f81b031916905560020180546001600160601b0319169055565b6001600160a01b03811661176b57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b8154818355818111156117ea576000838152602090206117ea9181019083016117ef565b505050565b61180d91905b8082111561180957600081556001016117f5565b5090565b9056fe596f752073686f756c64206e6f742062652070617573656420746f207061757365206f7468657273546f6f20736d616c6c2076616c756520746f2061646420746f20746865206465706f7369744465706f736974206973206e6f7420656e6f75676820746f20636f766572207468652066696e65596f752073686f756c64207761697420736f6d6520626c6f636b73206265666f72652077697468647261774e6f64652073686f756c64206e6f742062652070617573656420616e642073686f756c642068617665206465706f736974596f752073686f756c64206861766520766f6c756e746172696c7920706175736564206265666f7265207769746864726177596f752063616e206e6f7420756e70617573652066726f6d2074686973207374617465596f752073686f756c64207761697420736f6d6520626c6f636b73206265666f726520756e70617573654e6f64652073686f756c6420636f72726563746c792070617373206974732076616c696461746f72207075626c6963206b6579596f752063616e206e6f7420756e7061757365206265666f7265206465706f7369742065786365656473206d696e2076616c7565a165627a7a723058205151995dc42595fea74462595a0324057b4fa29139441b84f3c4f337432a4eb80029",
			"storage": {
				"0x036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0": "0x00000000000000878678326ed8750a7c127d3d5ed315c8286fa54316b18b639e",
				"0x0000000000000000000000000000000000000000000000000000000000000000": "0x000000000000000000000000d8750a7c127d3d5ed315c8286fa54316b18b639e",
				"0x0000000000000000000000000000000000000000000000000000000000000002": "0x0000000000000000000000000000000000000000000000878678326eac900000",
				"0x0000000000000000000000000000000000000000000000000000000000000001": "0x0000000000000000000000000000000000000000000000008ac7230489e80000",
				"0x8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b": "0xcc930f7483d6f42d73821587a44061ae0a2653ae9d843cc3e5503770613c4573",
				"0x0000000000000000000000000000000000000000000000000000000000000005": "0x0000000000000000000000000000000000000000000000000000000000000001",
				"0xda29c1b0fb44c9518c3bade9c33ad246e0d2feb9f061b6ec561d975ab2e89047": "0x0000000000000000000000000000000000000000000000000100000000118c31",
				"0xda29c1b0fb44c9518c3bade9c33ad246e0d2feb9f061b6ec561d975ab2e89046": "0x000000878678326eac900000d8750a7c127d3d5ed315c8286fa54316b18b639e",
				"0x0000000000000000000000000000000000000000000000000000000000000004": "0x0000000000000000000000000000000000000000000000000000000000000001"
			 }
		},
        "0x7eff122b94897ea5b0e2a9abf47b86337fafebdc": { "balance": "10000000000000000000000000000000000" },
		"0x89713c7d8c15ab70f392eb2674fcd51ec8e8f83b": { "balance": "10000000000000000000000000000000000" },
		"0xd8750A7c127D3D5ed315C8286fa54316B18B639e": { "balance": "247499999999999999999992" },
		"0xD610125550e9322AA3d8B0460eF5F79289b8DF55": { "balance": "250000000000000000000000" },
		"0x39bCd189bEE867B97E81255Ee0c122A796511993": { "balance": "250000000000000000000000" },
		"0x8DB79362a436e66DD09C551FB8534d84Ca63a89D": { "balance": "250000000000000000000000" },
		"0xB4Bc620461a4D5304Ff035015954A2C9Fc2D5759": { "balance": "250000000000000000000000" },
		"0x754B8c833047B61927EF3A7bbD0eD77E63881281": { "balance": "250000000000000000000000" },
		"0x4C7Cd51B8A79726C551eD13abe37ECD30017C431": { "balance": "250000000000000000000000" },
		"0xC5005a46EE842C38c529b2dFd6701c65Af4FF6fE": { "balance": "250000000000000000000000" },
		"0x7fb3955e556cC242E473a1f0b479053cA4639672": { "balance": "250000000000000000000000" },
		"0x64Bf814B8782B4F32df5d765cc1f35E0298aA957": { "balance": "250000000000000000000000" },
		"0x183333Bf79827743F5003633B1C1aE71f9D54185": { "balance": "250000000000000000000000" },
		"0xa6034Fb0A091c9e4f8C0147d0A15D7197C2404A7": { "balance": "250000000000000000000000" }
    }
}`)

/* For state conversion
function toKV(o){
    let out = {}
    for(let v of Object.values(o)){
        out[v.key] = v.value
    }
    return JSON.stringify(out, null, '    ')
}*/

var blankGenesis = new(core.Genesis)

var errBlankGenesis = errors.New("could not parse a valid/non-blank Genesis")

// ParseGenesisOrDefault tries to read the content from provided
// genesisPath. If the path is empty or doesn't exist, it will
// use defaultGenesisBytes as the fallback genesis source. Otherwise,
// it will open that path and if it encounters an error that doesn't
// satisfy os.IsNotExist, it returns that error.
func ParseGenesisOrDefault(genesisPath string) (*core.Genesis, error) {
	var genesisBlob = defaultGenesisBlob[:]
	if len(genesisPath) > 0 {
		blob, err := ioutil.ReadFile(genesisPath)
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}
		if len(blob) >= 2 { // Expecting atleast "{}"
			genesisBlob = blob
		}
	}

	genesis := new(core.Genesis)
	if err := json.Unmarshal(genesisBlob, genesis); err != nil {
		ethUtils.Fatalf("Error parsing genesis json: %v", err)
		return nil, err
	}

	if reflect.DeepEqual(blankGenesis, genesis) {
		return nil, errBlankGenesis
	}

	return genesis, nil
}
