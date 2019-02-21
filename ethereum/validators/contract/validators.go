// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"getNodeAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompactedValidators\",\"outputs\":[{\"name\":\"ValidatorsCompacted\",\"type\":\"bytes32[]\"},{\"name\":\"ValidatorsPubKeys\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"addDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"},{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"addInitialDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"isPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"},{\"name\":\"vFrom\",\"type\":\"bytes32\"},{\"name\":\"cause\",\"type\":\"uint8\"}],\"name\":\"pauseValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"hasDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"resumeValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"getNodeReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveCompactedValidators\",\"outputs\":[{\"name\":\"ValidatorsCompacted\",\"type\":\"bytes32[]\"},{\"name\":\"ValidatorsPubKeys\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint96\"},{\"name\":\"pauseBlockNumber\",\"type\":\"uint48\"},{\"name\":\"pauseCause\",\"type\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vPub\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint96\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vPub\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cause\",\"type\":\"uint8\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"}]"

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
const ValidatorsBin = `0x608060405234801561001057600080fd5b5061168d806100206000396000f3fe6080604052600436106100df576000357c01000000000000000000000000000000000000000000000000000000009004806379354ef51161009c5780639da0e3e3116100765780639da0e3e31461031d578063a36147d314610347578063acfb8f7614610371578063d5f20ff614610386576100df565b806379354ef5146102905780638e19899e146102c957806392c8a966146102f3576100df565b80630b8604fc146100e4578063181a8d681461012a5780631c32d982146101d857806320aef8d4146101f7578063241b71bb1461022b57806363338b1714610269575b600080fd5b3480156100f057600080fd5b5061010e6004803603602081101561010757600080fd5b5035610402565b60408051600160a060020a039092168252519081900360200190f35b34801561013657600080fd5b5061013f61041d565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561018357818101518382015260200161016b565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156101c25781810151838201526020016101aa565b5050505090500194505050505060405180910390f35b6101f5600480360360208110156101ee57600080fd5b50356104ca565b005b6101f56004803603606081101561020d57600080fd5b50803590600160a060020a03602082013581169160400135166104d9565b34801561023757600080fd5b506102556004803603602081101561024e57600080fd5b503561078b565b604080519115158252519081900360200190f35b34801561027557600080fd5b5061027e6107ae565b60408051918252519081900360200190f35b34801561029c57600080fd5b506101f5600480360360608110156102b357600080fd5b508035906020810135906040013560ff16610818565b3480156102d557600080fd5b506101f5600480360360208110156102ec57600080fd5b5035610acc565b3480156102ff57600080fd5b506102556004803603602081101561031657600080fd5b5035610d0b565b34801561032957600080fd5b506101f56004803603602081101561034057600080fd5b5035610d33565b34801561035357600080fd5b5061010e6004803603602081101561036a57600080fd5b5035610fb9565b34801561037d57600080fd5b5061013f61100e565b34801561039257600080fd5b506103b0600480360360208110156103a957600080fd5b503561112e565b60408051600160a060020a0396871681526bffffffffffffffffffffffff909516602086015265ffffffffffff9093168484015260ff9091166060840152909216608082015290519081900360a00190f35b600090815260208190526040902054600160a060020a031690565b606080600280548060200260200160405190810160405280929190818152602001828054801561046c57602002820191906000526020600020905b815481526020019060010190808311610458575b5050505050915060018054806020026020016040519081016040528092919081815260200182805480156104bf57602002820191906000526020600020905b8154815260200190600101908083116104ab575b505050505090509091565b6104d6816000806104d9565b50565b678ac7230489e800003410156105235760405160e560020a62461bcd0281526004018080602001828103825260258152602001806114f46025913960400191505060405180910390fd5b600083815260208190526040902080546bffffffffffffffffffffffff60a060020a8083048216340190911602600160a060020a039182161780835516151561065e57600160a060020a03831615156105c6576040805160e560020a62461bcd02815260206004820152601b60248201527f506c656173652073706563696679206e6f646520616464726573730000000000604482015290519081900360640190fd5b805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384811691909117825560018201805466010000000000007fff0000000000000000000000000000000000000000ffffffffffffffffffffff9091166b010000000000000000000000938616939093029290921766ff00000000000019169190911765ffffffffffff19164365ffffffffffff161790555b6001810154600067010000000000000090910463ffffffff16111561068c5761068681611193565b5061072b565b6000610697826111db565b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018790556002805480830182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01919091558054908301805463ffffffff909216670100000000000000026affffffff0000000000000019909216919091179055505b60003411156107855780546040805186815260a060020a9092046bffffffffffffffffffffffff16602083015280517fd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf07889281900390910190a15b50505050565b6000908152602081905260408120600101546601000000000000900460ff161190565b600254600090815b818110156108135760006002828154811015156107cf57fe5b60009182526020909120015490507cff000000000000000000000000000000000000000000000000000000008116151561080a578360010193505b506001016107b6565b505090565b6108218361078b565b158015610832575061083283610d0b565b15156108725760405160e560020a62461bcd0281526004018080602001828103825260318152602001806115446031913960400191505060405180910390fd5b600160ff821610156108ce576040805160e560020a62461bcd02815260206004820152601860248201527f596f752073686f756c6420737065636966792063617573650000000000000000604482015290519081900360640190fd5b336108d883610402565b600160a060020a0316146109205760405160e560020a62461bcd0281526004018080602001828103825260338152602001806115d16033913960400191505060405180910390fd5b8282146109e35761093082610d0b565b801561094257506109408261078b565b155b15156109825760405160e560020a62461bcd0281526004018080602001828103825260288152602001806114cc6028913960400191505060405180910390fd5b60ff8116600114156109de576040805160e560020a62461bcd02815260206004820152601e60248201527f546869732070617573696e67206973206e6f7420766f6c756e74617279210000604482015290519081900360640190fd5b610a3e565b60ff8116600114610a3e576040805160e560020a62461bcd02815260206004820152601b60248201527f596f75206172652070617573696e672076616c756e746172696c790000000000604482015290519081900360640190fd5b600083815260208190526040902060018101805465ffffffffffff19164365ffffffffffff161766ff0000000000001916660100000000000060ff851602179055610a8881611193565b506040805185815260ff8416602082015281517f2ffa21f275941345406452ac41a8bceea9dde4f76b233ac4fc9e53c31e399342929181900390910190a150505050565b610ad581610d0b565b1515610b2b576040805160e560020a62461bcd02815260206004820152601760248201527f596f752073686f756c642068617665206465706f736974000000000000000000604482015290519081900360640190fd5b610b348161078b565b1515610b8a576040805160e560020a62461bcd02815260206004820152601460248201527f596f752073686f756c6420626520706175736564000000000000000000000000604482015290519081900360640190fd5b60008181526020819052604090208054600160a060020a03163314610bf9576040805160e560020a62461bcd02815260206004820152601e60248201527f4f6e6c79206e6f646520697473656c662063616e207769746864726177210000604482015290519081900360640190fd5b6001818101546601000000000000900460ff1614610c4b5760405160e560020a62461bcd0281526004018080602001828103825260328152602001806115756032913960400191505060405180910390fd5b600181015465ffffffffffff438116918116607801161115610ca15760405160e560020a62461bcd02815260040180806020018281038252602b815260200180611519602b913960400191505060405180910390fd5b805460a060020a90046bffffffffffffffffffffffff166000610cc384610fb9565b9050610cce84611242565b604051600160a060020a0382169083156108fc029084906000818181858888f19350505050158015610d04573d6000803e3d6000fd5b5050505050565b60009081526020819052604081205460a060020a90046bffffffffffffffffffffffff161190565b610d3c8161078b565b1515610d92576040805160e560020a62461bcd02815260206004820152601860248201527f54686973206e6f6465206973206e6f7420706175736564210000000000000000604482015290519081900360640190fd5b60008181526020819052604090208054600160a060020a03163314610e01576040805160e560020a62461bcd02815260206004820152601d60248201527f596f752063616e206f6e6c7920756e706175736520796f757273656c66000000604482015290519081900360640190fd5b6001818101546601000000000000900460ff161480610e31575060018101546601000000000000900460ff166002145b1515610e715760405160e560020a62461bcd02815260040180806020018281038252602a815260200180611638602a913960400191505060405180910390fd5b60018101546601000000000000900460ff166002141580610ea75750600181015465ffffffffffff4381169181166104b0011611155b1515610ee75760405160e560020a62461bcd02815260040180806020018281038252602a8152602001806115a7602a913960400191505060405180910390fd5b805468878678326eac90000060a060020a9091046bffffffffffffffffffffffff161015610f495760405160e560020a62461bcd0281526004018080602001828103825260348152602001806116046034913960400191505060405180910390fd5b60018101805466ffffffffffffff19169055610f6481611193565b5080546040805184815260a060020a9092046bffffffffffffffffffffffff16602083015280517fd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf07889281900390910190a15050565b6000818152602081905260408120600101546b0100000000000000000000009004600160a060020a0316801515610ff657610ff383610402565b90505b600160a060020a038116151561100857fe5b92915050565b606080600061101b6107ae565b905080604051908082528060200260200182016040528015611047578160200160208202803883390190505b50925080604051908082528060200260200182016040528015611074578160200160208202803883390190505b506002549092506000805b8281101561112657600060028281548110151561109857fe5b60009182526020909120015490507cff000000000000000000000000000000000000000000000000000000008116151561111d578087848151811015156110db57fe5b6020908102909101015260018054839081106110f357fe5b9060005260206000200154868481518110151561110c57fe5b602090810290910101526001909201915b5060010161107f565b505050509091565b60009081526020819052604090208054600190910154600160a060020a038083169360a060020a9093046bffffffffffffffffffffffff169265ffffffffffff8316926601000000000000810460ff16926b0100000000000000000000009091041690565b600061119e826111db565b600183015460028054909160001963ffffffff670100000000000000909204821601169081106111ca57fe5b600091825260209091200155919050565b80546001909101546601000000000000900460ff167c010000000000000000000000000000000000000000000000000000000002600160a060020a03821664010000000060a060020a938490046bffffffffffffffffffffffff1604909202919091171790565b6000818152602081905260408120805460019160029160a060020a90046bffffffffffffffffffffffff1684106112c3576040805160e560020a62461bcd02815260206004820152601560248201527f4465706f7369742073686f756c64206578697374210000000000000000000000604482015290519081900360640190fd5b82546000106112ce57fe5b600181810154845467010000000000000090910463ffffffff166000190191108015611305575083546000190163ffffffff821614155b156113e75783548490600019810190811061131c57fe5b9060005260206000200154848263ffffffff1681548110151561133b57fe5b60009182526020909120015582548390600019810190811061135957fe5b9060005260206000200154838263ffffffff1681548110151561137857fe5b90600052602060002001819055506000848263ffffffff1681548110151561139c57fe5b600091825260208083209190910154825287905260409020600190810180546affffffff00000000000000191667010000000000000092850163ffffffff1692909202919091179055505b8354849060001981019081106113f957fe5b60009182526020822001558354611414856000198301611481565b5082548390600019810190811061142757fe5b60009182526020822001558254611442846000198301611481565b505050600093845250506020526040812090815560010180547fff00000000000000000000000000000000000000000000000000000000000000169055565b8154818355818111156114a5576000838152602090206114a59181019083016114aa565b505050565b6114c891905b808211156114c457600081556001016114b0565b5090565b9056fe596f752073686f756c64206e6f742062652070617573656420746f207061757365206f7468657273546f6f20736d616c6c2076616c756520746f2061646420746f20746865206465706f736974596f752073686f756c64207761697420736f6d6520626c6f636b73206265666f72652077697468647261774e6f64652073686f756c64206e6f742062652070617573656420616e642073686f756c642068617665206465706f736974596f752073686f756c64206861766520766f6c756e746172696c7920706175736564206265666f7265207769746864726177596f752073686f756c64207761697420736f6d6520626c6f636b73206265666f726520756e70617573654e6f64652073686f756c6420636f72726563746c792070617373206974732076616c696461746f72207075626c6963206b6579596f752063616e206e6f7420756e7061757365206265666f7265206465706f7369742065786365656473206d696e2076616c7565596f752063616e206f6e6c7920756e706175736520766f6c756e7461727920706175736564206e6f6465a165627a7a7230582020715a821a9acb79d452b022c743d8d2e43d32261c49f59c3829c84c217094db0029`

// DeployValidators deploys a new Ethereum contract, binding an instance of Validators to it.
func DeployValidators(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validators, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// GetActiveCompactedValidators is a free data retrieval call binding the contract method 0xacfb8f76.
//
// Solidity: function getActiveCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsPubKeys bytes32[])
func (_Validators *ValidatorsCaller) GetActiveCompactedValidators(opts *bind.CallOpts) (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsPubKeys   [][32]byte
}, error) {
	ret := new(struct {
		ValidatorsCompacted [][32]byte
		ValidatorsPubKeys   [][32]byte
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "getActiveCompactedValidators")
	return *ret, err
}

// GetActiveCompactedValidators is a free data retrieval call binding the contract method 0xacfb8f76.
//
// Solidity: function getActiveCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsPubKeys bytes32[])
func (_Validators *ValidatorsSession) GetActiveCompactedValidators() (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsPubKeys   [][32]byte
}, error) {
	return _Validators.Contract.GetActiveCompactedValidators(&_Validators.CallOpts)
}

// GetActiveCompactedValidators is a free data retrieval call binding the contract method 0xacfb8f76.
//
// Solidity: function getActiveCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsPubKeys bytes32[])
func (_Validators *ValidatorsCallerSession) GetActiveCompactedValidators() (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsPubKeys   [][32]byte
}, error) {
	return _Validators.Contract.GetActiveCompactedValidators(&_Validators.CallOpts)
}

// GetActiveCount is a free data retrieval call binding the contract method 0x63338b17.
//
// Solidity: function getActiveCount() constant returns(count uint256)
func (_Validators *ValidatorsCaller) GetActiveCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getActiveCount")
	return *ret0, err
}

// GetActiveCount is a free data retrieval call binding the contract method 0x63338b17.
//
// Solidity: function getActiveCount() constant returns(count uint256)
func (_Validators *ValidatorsSession) GetActiveCount() (*big.Int, error) {
	return _Validators.Contract.GetActiveCount(&_Validators.CallOpts)
}

// GetActiveCount is a free data retrieval call binding the contract method 0x63338b17.
//
// Solidity: function getActiveCount() constant returns(count uint256)
func (_Validators *ValidatorsCallerSession) GetActiveCount() (*big.Int, error) {
	return _Validators.Contract.GetActiveCount(&_Validators.CallOpts)
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsPubKeys bytes32[])
func (_Validators *ValidatorsCaller) GetCompactedValidators(opts *bind.CallOpts) (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsPubKeys   [][32]byte
}, error) {
	ret := new(struct {
		ValidatorsCompacted [][32]byte
		ValidatorsPubKeys   [][32]byte
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "getCompactedValidators")
	return *ret, err
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsPubKeys bytes32[])
func (_Validators *ValidatorsSession) GetCompactedValidators() (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsPubKeys   [][32]byte
}, error) {
	return _Validators.Contract.GetCompactedValidators(&_Validators.CallOpts)
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsPubKeys bytes32[])
func (_Validators *ValidatorsCallerSession) GetCompactedValidators() (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsPubKeys   [][32]byte
}, error) {
	return _Validators.Contract.GetCompactedValidators(&_Validators.CallOpts)
}

// GetNodeAddr is a free data retrieval call binding the contract method 0x0b8604fc.
//
// Solidity: function getNodeAddr(vPub bytes32) constant returns(address)
func (_Validators *ValidatorsCaller) GetNodeAddr(opts *bind.CallOpts, vPub [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getNodeAddr", vPub)
	return *ret0, err
}

// GetNodeAddr is a free data retrieval call binding the contract method 0x0b8604fc.
//
// Solidity: function getNodeAddr(vPub bytes32) constant returns(address)
func (_Validators *ValidatorsSession) GetNodeAddr(vPub [32]byte) (common.Address, error) {
	return _Validators.Contract.GetNodeAddr(&_Validators.CallOpts, vPub)
}

// GetNodeAddr is a free data retrieval call binding the contract method 0x0b8604fc.
//
// Solidity: function getNodeAddr(vPub bytes32) constant returns(address)
func (_Validators *ValidatorsCallerSession) GetNodeAddr(vPub [32]byte) (common.Address, error) {
	return _Validators.Contract.GetNodeAddr(&_Validators.CallOpts, vPub)
}

// GetNodeReceiver is a free data retrieval call binding the contract method 0xa36147d3.
//
// Solidity: function getNodeReceiver(vPub bytes32) constant returns(address)
func (_Validators *ValidatorsCaller) GetNodeReceiver(opts *bind.CallOpts, vPub [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getNodeReceiver", vPub)
	return *ret0, err
}

// GetNodeReceiver is a free data retrieval call binding the contract method 0xa36147d3.
//
// Solidity: function getNodeReceiver(vPub bytes32) constant returns(address)
func (_Validators *ValidatorsSession) GetNodeReceiver(vPub [32]byte) (common.Address, error) {
	return _Validators.Contract.GetNodeReceiver(&_Validators.CallOpts, vPub)
}

// GetNodeReceiver is a free data retrieval call binding the contract method 0xa36147d3.
//
// Solidity: function getNodeReceiver(vPub bytes32) constant returns(address)
func (_Validators *ValidatorsCallerSession) GetNodeReceiver(vPub [32]byte) (common.Address, error) {
	return _Validators.Contract.GetNodeReceiver(&_Validators.CallOpts, vPub)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(vPub bytes32) constant returns(nodeAddr address, deposit uint96, pauseBlockNumber uint48, pauseCause uint8, receiver address)
func (_Validators *ValidatorsCaller) GetValidator(opts *bind.CallOpts, vPub [32]byte) (struct {
	NodeAddr         common.Address
	Deposit          *big.Int
	PauseBlockNumber *big.Int
	PauseCause       uint8
	Receiver         common.Address
}, error) {
	ret := new(struct {
		NodeAddr         common.Address
		Deposit          *big.Int
		PauseBlockNumber *big.Int
		PauseCause       uint8
		Receiver         common.Address
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "getValidator", vPub)
	return *ret, err
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(vPub bytes32) constant returns(nodeAddr address, deposit uint96, pauseBlockNumber uint48, pauseCause uint8, receiver address)
func (_Validators *ValidatorsSession) GetValidator(vPub [32]byte) (struct {
	NodeAddr         common.Address
	Deposit          *big.Int
	PauseBlockNumber *big.Int
	PauseCause       uint8
	Receiver         common.Address
}, error) {
	return _Validators.Contract.GetValidator(&_Validators.CallOpts, vPub)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(vPub bytes32) constant returns(nodeAddr address, deposit uint96, pauseBlockNumber uint48, pauseCause uint8, receiver address)
func (_Validators *ValidatorsCallerSession) GetValidator(vPub [32]byte) (struct {
	NodeAddr         common.Address
	Deposit          *big.Int
	PauseBlockNumber *big.Int
	PauseCause       uint8
	Receiver         common.Address
}, error) {
	return _Validators.Contract.GetValidator(&_Validators.CallOpts, vPub)
}

// HasDeposit is a free data retrieval call binding the contract method 0x92c8a966.
//
// Solidity: function hasDeposit(vPub bytes32) constant returns(bool)
func (_Validators *ValidatorsCaller) HasDeposit(opts *bind.CallOpts, vPub [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "hasDeposit", vPub)
	return *ret0, err
}

// HasDeposit is a free data retrieval call binding the contract method 0x92c8a966.
//
// Solidity: function hasDeposit(vPub bytes32) constant returns(bool)
func (_Validators *ValidatorsSession) HasDeposit(vPub [32]byte) (bool, error) {
	return _Validators.Contract.HasDeposit(&_Validators.CallOpts, vPub)
}

// HasDeposit is a free data retrieval call binding the contract method 0x92c8a966.
//
// Solidity: function hasDeposit(vPub bytes32) constant returns(bool)
func (_Validators *ValidatorsCallerSession) HasDeposit(vPub [32]byte) (bool, error) {
	return _Validators.Contract.HasDeposit(&_Validators.CallOpts, vPub)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(vPub bytes32) constant returns(bool)
func (_Validators *ValidatorsCaller) IsPaused(opts *bind.CallOpts, vPub [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isPaused", vPub)
	return *ret0, err
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(vPub bytes32) constant returns(bool)
func (_Validators *ValidatorsSession) IsPaused(vPub [32]byte) (bool, error) {
	return _Validators.Contract.IsPaused(&_Validators.CallOpts, vPub)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(vPub bytes32) constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsPaused(vPub [32]byte) (bool, error) {
	return _Validators.Contract.IsPaused(&_Validators.CallOpts, vPub)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x1c32d982.
//
// Solidity: function addDeposit(vPub bytes32) returns()
func (_Validators *ValidatorsTransactor) AddDeposit(opts *bind.TransactOpts, vPub [32]byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addDeposit", vPub)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x1c32d982.
//
// Solidity: function addDeposit(vPub bytes32) returns()
func (_Validators *ValidatorsSession) AddDeposit(vPub [32]byte) (*types.Transaction, error) {
	return _Validators.Contract.AddDeposit(&_Validators.TransactOpts, vPub)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x1c32d982.
//
// Solidity: function addDeposit(vPub bytes32) returns()
func (_Validators *ValidatorsTransactorSession) AddDeposit(vPub [32]byte) (*types.Transaction, error) {
	return _Validators.Contract.AddDeposit(&_Validators.TransactOpts, vPub)
}

// AddInitialDeposit is a paid mutator transaction binding the contract method 0x20aef8d4.
//
// Solidity: function addInitialDeposit(vPub bytes32, nodeAddr address, receiver address) returns()
func (_Validators *ValidatorsTransactor) AddInitialDeposit(opts *bind.TransactOpts, vPub [32]byte, nodeAddr common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addInitialDeposit", vPub, nodeAddr, receiver)
}

// AddInitialDeposit is a paid mutator transaction binding the contract method 0x20aef8d4.
//
// Solidity: function addInitialDeposit(vPub bytes32, nodeAddr address, receiver address) returns()
func (_Validators *ValidatorsSession) AddInitialDeposit(vPub [32]byte, nodeAddr common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddInitialDeposit(&_Validators.TransactOpts, vPub, nodeAddr, receiver)
}

// AddInitialDeposit is a paid mutator transaction binding the contract method 0x20aef8d4.
//
// Solidity: function addInitialDeposit(vPub bytes32, nodeAddr address, receiver address) returns()
func (_Validators *ValidatorsTransactorSession) AddInitialDeposit(vPub [32]byte, nodeAddr common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddInitialDeposit(&_Validators.TransactOpts, vPub, nodeAddr, receiver)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x79354ef5.
//
// Solidity: function pauseValidation(vPub bytes32, vFrom bytes32, cause uint8) returns()
func (_Validators *ValidatorsTransactor) PauseValidation(opts *bind.TransactOpts, vPub [32]byte, vFrom [32]byte, cause uint8) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "pauseValidation", vPub, vFrom, cause)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x79354ef5.
//
// Solidity: function pauseValidation(vPub bytes32, vFrom bytes32, cause uint8) returns()
func (_Validators *ValidatorsSession) PauseValidation(vPub [32]byte, vFrom [32]byte, cause uint8) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, vPub, vFrom, cause)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x79354ef5.
//
// Solidity: function pauseValidation(vPub bytes32, vFrom bytes32, cause uint8) returns()
func (_Validators *ValidatorsTransactorSession) PauseValidation(vPub [32]byte, vFrom [32]byte, cause uint8) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, vPub, vFrom, cause)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0x9da0e3e3.
//
// Solidity: function resumeValidation(vPub bytes32) returns()
func (_Validators *ValidatorsTransactor) ResumeValidation(opts *bind.TransactOpts, vPub [32]byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "resumeValidation", vPub)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0x9da0e3e3.
//
// Solidity: function resumeValidation(vPub bytes32) returns()
func (_Validators *ValidatorsSession) ResumeValidation(vPub [32]byte) (*types.Transaction, error) {
	return _Validators.Contract.ResumeValidation(&_Validators.TransactOpts, vPub)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0x9da0e3e3.
//
// Solidity: function resumeValidation(vPub bytes32) returns()
func (_Validators *ValidatorsTransactorSession) ResumeValidation(vPub [32]byte) (*types.Transaction, error) {
	return _Validators.Contract.ResumeValidation(&_Validators.TransactOpts, vPub)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(vPub bytes32) returns()
func (_Validators *ValidatorsTransactor) Withdraw(opts *bind.TransactOpts, vPub [32]byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdraw", vPub)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(vPub bytes32) returns()
func (_Validators *ValidatorsSession) Withdraw(vPub [32]byte) (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts, vPub)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(vPub bytes32) returns()
func (_Validators *ValidatorsTransactorSession) Withdraw(vPub [32]byte) (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts, vPub)
}

// ValidatorsValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Validators contract.
type ValidatorsValidatorRemovedIterator struct {
	Event *ValidatorsValidatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorRemoved represents a ValidatorRemoved event raised by the Validators contract.
type ValidatorsValidatorRemoved struct {
	VPub  [32]byte
	Cause uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0x2ffa21f275941345406452ac41a8bceea9dde4f76b233ac4fc9e53c31e399342.
//
// Solidity: event ValidatorRemoved(vPub bytes32, cause uint8)
func (_Validators *ValidatorsFilterer) FilterValidatorRemoved(opts *bind.FilterOpts) (*ValidatorsValidatorRemovedIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorRemovedIterator{contract: _Validators.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0x2ffa21f275941345406452ac41a8bceea9dde4f76b233ac4fc9e53c31e399342.
//
// Solidity: event ValidatorRemoved(vPub bytes32, cause uint8)
func (_Validators *ValidatorsFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorRemoved)
				if err := _Validators.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ValidatorsValidatorUpdatedIterator is returned from FilterValidatorUpdated and is used to iterate over the raw logs and unpacked data for ValidatorUpdated events raised by the Validators contract.
type ValidatorsValidatorUpdatedIterator struct {
	Event *ValidatorsValidatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorUpdated represents a ValidatorUpdated event raised by the Validators contract.
type ValidatorsValidatorUpdated struct {
	VPub    [32]byte
	Deposit *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorUpdated is a free log retrieval operation binding the contract event 0xd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf0788.
//
// Solidity: event ValidatorUpdated(vPub bytes32, deposit uint96)
func (_Validators *ValidatorsFilterer) FilterValidatorUpdated(opts *bind.FilterOpts) (*ValidatorsValidatorUpdatedIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorUpdatedIterator{contract: _Validators.contract, event: "ValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorUpdated is a free log subscription operation binding the contract event 0xd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf0788.
//
// Solidity: event ValidatorUpdated(vPub bytes32, deposit uint96)
func (_Validators *ValidatorsFilterer) WatchValidatorUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorUpdated) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
