// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"getNodeAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompactedValidators\",\"outputs\":[{\"name\":\"ValidatorsCompacted\",\"type\":\"bytes32[]\"},{\"name\":\"ValidatorsIndex\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"isPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_CAUSE_VOLUNTARILY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"},{\"name\":\"vFrom\",\"type\":\"bytes32\"},{\"name\":\"cause\",\"type\":\"uint8\"}],\"name\":\"pauseValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"},{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"addDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"hasDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"resumeValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_CAUSE_PUNISHMENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_LOCK_BLOCKS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
const ValidatorsBin = `0x6060604052341561000f57600080fd5b610b378061001e6000396000f3006060604052600436106100b95763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630b8604fc81146100be578063181a8d68146100f0578063241b71bb1461019c57806347fdd5f8146101c657806379354ef5146101ef5780637e332a46146102105780638e19899e1461022757806392c8a9661461023d5780639da0e3e314610253578063be1d05c214610269578063e1e158a51461027c578063faeedb61146102a1575b600080fd5b34156100c957600080fd5b6100d46004356102cf565b604051600160a060020a03909116815260200160405180910390f35b34156100fb57600080fd5b6101036102ea565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561014757808201518382015260200161012f565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561018657808201518382015260200161016e565b5050505090500194505050505060405180910390f35b34156101a757600080fd5b6101b26004356103a6565b604051901515815260200160405180910390f35b34156101d157600080fd5b6101d96103c9565b60405160ff909116815260200160405180910390f35b34156101fa57600080fd5b61020e60043560243560ff604435166103ce565b005b61020e600435600160a060020a0360243516610500565b341561023257600080fd5b61020e60043561069f565b341561024857600080fd5b6101b260043561078a565b341561025e57600080fd5b61020e6004356107b2565b341561027457600080fd5b6101d961085d565b341561028757600080fd5b61028f610862565b60405190815260200160405180910390f35b34156102ac57600080fd5b6102b461086e565b60405165ffffffffffff909116815260200160405180910390f35b600090815260208190526040902054600160a060020a031690565b6102f2610ab9565b6102fa610ab9565b600280548060200260200160405190810160405280929190818152602001828054801561034757602002820191906000526020600020905b81548152600190910190602001808311610332575b50505050509150600180548060200260200160405190810160405280929190818152602001828054801561039b57602002820191906000526020600020905b81548152600190910190602001808311610386575b505050505090509091565b6000908152602081905260408120600101546601000000000000900460ff161190565b600181565b60006103d9846103a6565b1580156103ea57506103ea8461078a565b15156103f557600080fd5b600160ff8316101561040657600080fd5b33600160a060020a0316610419846102cf565b600160a060020a03161461042c57600080fd5b83831461046f5761043c8361078a565b801561044e575061044c836103a6565b155b151561045957600080fd5b60ff82166001141561046a57600080fd5b61047f565b60ff821660011461047f57600080fd5b50600083815260208190526040902060018101805465ffffffffffff19164365ffffffffffff161766ff0000000000001916660100000000000060ff8516021790556104cb8482610873565b600182810154815467010000000000000090910463ffffffff169081106104ee57fe5b60009182526020909120015550505050565b600082815260208190526040812080546bffffffffffffffffffffffff60a060020a8083048216340182168102600160a060020a03909316929092178084559293928392670de0b6b3a76400009104909116101561055d57600080fd5b8254600160a060020a0316151561059557825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161783555b61059f8584610873565b6001840154909250600067010000000000000090910463ffffffff1611156106255760018360010160079054906101000a900463ffffffff160390508460018263ffffffff168154811015156105f157fe5b6000918252602090912001556002805483919063ffffffff841690811061061457fe5b600091825260209091200155610698565b600180548082016106368382610acb565b50600091825260209091200185905560028054600181016106578382610acb565b50600091825260209091200182905560018054908401805463ffffffff909216670100000000000000026affffffff00000000000000199092169190911790555b5050505050565b6000806106ab8361078a565b15156106b657600080fd5b6106bf836103a6565b15156106ca57600080fd5b6000838152602081905260409020805490925033600160a060020a039081169116146106f557600080fd5b6001828101546601000000000000900460ff161461071257600080fd5b600182015465ffffffffffff438116918116600a0116111561073357600080fd5b50805460a060020a90046bffffffffffffffffffffffff16610754836108ce565b600160a060020a03331681156108fc0282604051600060405180830381858888f19350505050151561078557600080fd5b505050565b60009081526020819052604081205460a060020a90046bffffffffffffffffffffffff161190565b60006107bd826103a6565b15156107c857600080fd5b506000818152602081905260409020805433600160a060020a039081169116146107f157600080fd5b6001818101546601000000000000900460ff161461080e57600080fd5b60018101805466ffffffffffffff1916905561082a8282610873565b600182810154815467010000000000000090910463ffffffff1690811061084d57fe5b6000918252602090912001555050565b600281565b670de0b6b3a764000081565b600a81565b60018101549054660100000000000090910460ff167c01000000000000000000000000000000000000000000000000000000000264010000000060a060020a928390046bffffffffffffffffffffffff160490910217919050565b600081815260208190526040812080546001916002918490819060a060020a90046bffffffffffffffffffffffff1681901161090957600080fd5b84546000901161091557fe5b60018360010160079054906101000a900463ffffffff160391506001858054905011801561094e575084546000190163ffffffff831614155b15610a325784548590600019810190811061096557fe5b906000526020600020900154858363ffffffff1681548110151561098557fe5b6000918252602090912001558354849060001981019081106109a357fe5b906000526020600020900154848363ffffffff168154811015156109c357fe5b6000918252602090912001558454859063ffffffff84169081106109e357fe5b600091825260208083209190910154808352908890526040909120600190810180546affffffff00000000000000191667010000000000000092860163ffffffff169290920291909117905590505b845485906000198101908110610a4457fe5b60009182526020822001558454610a5f866000198301610acb565b50835484906000198101908110610a7257fe5b60009182526020822001558354610a8d856000198301610acb565b50505060009485525050506020526040812090815560010180546affffffffffffffffffffff19169055565b60206040519081016040526000815290565b81548183558181151161078557600083815260209020610785918101908301610b0891905b80821115610b045760008155600101610af0565b5090565b905600a165627a7a723058207926caa61cc549b149fb31808fe197726d706abf09dc00bd724fc5298a91320e0029`

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

// DEPOSITLOCKBLOCKS is a free data retrieval call binding the contract method 0xfaeedb61.
//
// Solidity: function DEPOSIT_LOCK_BLOCKS() constant returns(uint48)
func (_Validators *ValidatorsCaller) DEPOSITLOCKBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "DEPOSIT_LOCK_BLOCKS")
	return *ret0, err
}

// DEPOSITLOCKBLOCKS is a free data retrieval call binding the contract method 0xfaeedb61.
//
// Solidity: function DEPOSIT_LOCK_BLOCKS() constant returns(uint48)
func (_Validators *ValidatorsSession) DEPOSITLOCKBLOCKS() (*big.Int, error) {
	return _Validators.Contract.DEPOSITLOCKBLOCKS(&_Validators.CallOpts)
}

// DEPOSITLOCKBLOCKS is a free data retrieval call binding the contract method 0xfaeedb61.
//
// Solidity: function DEPOSIT_LOCK_BLOCKS() constant returns(uint48)
func (_Validators *ValidatorsCallerSession) DEPOSITLOCKBLOCKS() (*big.Int, error) {
	return _Validators.Contract.DEPOSITLOCKBLOCKS(&_Validators.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsCaller) MINDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "MIN_DEPOSIT")
	return *ret0, err
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsSession) MINDEPOSIT() (*big.Int, error) {
	return _Validators.Contract.MINDEPOSIT(&_Validators.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MINDEPOSIT() (*big.Int, error) {
	return _Validators.Contract.MINDEPOSIT(&_Validators.CallOpts)
}

// PAUSECAUSEPUNISHMENT is a free data retrieval call binding the contract method 0xbe1d05c2.
//
// Solidity: function PAUSE_CAUSE_PUNISHMENT() constant returns(uint8)
func (_Validators *ValidatorsCaller) PAUSECAUSEPUNISHMENT(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "PAUSE_CAUSE_PUNISHMENT")
	return *ret0, err
}

// PAUSECAUSEPUNISHMENT is a free data retrieval call binding the contract method 0xbe1d05c2.
//
// Solidity: function PAUSE_CAUSE_PUNISHMENT() constant returns(uint8)
func (_Validators *ValidatorsSession) PAUSECAUSEPUNISHMENT() (uint8, error) {
	return _Validators.Contract.PAUSECAUSEPUNISHMENT(&_Validators.CallOpts)
}

// PAUSECAUSEPUNISHMENT is a free data retrieval call binding the contract method 0xbe1d05c2.
//
// Solidity: function PAUSE_CAUSE_PUNISHMENT() constant returns(uint8)
func (_Validators *ValidatorsCallerSession) PAUSECAUSEPUNISHMENT() (uint8, error) {
	return _Validators.Contract.PAUSECAUSEPUNISHMENT(&_Validators.CallOpts)
}

// PAUSECAUSEVOLUNTARILY is a free data retrieval call binding the contract method 0x47fdd5f8.
//
// Solidity: function PAUSE_CAUSE_VOLUNTARILY() constant returns(uint8)
func (_Validators *ValidatorsCaller) PAUSECAUSEVOLUNTARILY(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "PAUSE_CAUSE_VOLUNTARILY")
	return *ret0, err
}

// PAUSECAUSEVOLUNTARILY is a free data retrieval call binding the contract method 0x47fdd5f8.
//
// Solidity: function PAUSE_CAUSE_VOLUNTARILY() constant returns(uint8)
func (_Validators *ValidatorsSession) PAUSECAUSEVOLUNTARILY() (uint8, error) {
	return _Validators.Contract.PAUSECAUSEVOLUNTARILY(&_Validators.CallOpts)
}

// PAUSECAUSEVOLUNTARILY is a free data retrieval call binding the contract method 0x47fdd5f8.
//
// Solidity: function PAUSE_CAUSE_VOLUNTARILY() constant returns(uint8)
func (_Validators *ValidatorsCallerSession) PAUSECAUSEVOLUNTARILY() (uint8, error) {
	return _Validators.Contract.PAUSECAUSEVOLUNTARILY(&_Validators.CallOpts)
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsIndex bytes32[])
func (_Validators *ValidatorsCaller) GetCompactedValidators(opts *bind.CallOpts) (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsIndex     [][32]byte
}, error) {
	ret := new(struct {
		ValidatorsCompacted [][32]byte
		ValidatorsIndex     [][32]byte
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "getCompactedValidators")
	return *ret, err
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsIndex bytes32[])
func (_Validators *ValidatorsSession) GetCompactedValidators() (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsIndex     [][32]byte
}, error) {
	return _Validators.Contract.GetCompactedValidators(&_Validators.CallOpts)
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(ValidatorsCompacted bytes32[], ValidatorsIndex bytes32[])
func (_Validators *ValidatorsCallerSession) GetCompactedValidators() (struct {
	ValidatorsCompacted [][32]byte
	ValidatorsIndex     [][32]byte
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

// AddDeposit is a paid mutator transaction binding the contract method 0x7e332a46.
//
// Solidity: function addDeposit(vPub bytes32, nodeAddr address) returns()
func (_Validators *ValidatorsTransactor) AddDeposit(opts *bind.TransactOpts, vPub [32]byte, nodeAddr common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addDeposit", vPub, nodeAddr)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x7e332a46.
//
// Solidity: function addDeposit(vPub bytes32, nodeAddr address) returns()
func (_Validators *ValidatorsSession) AddDeposit(vPub [32]byte, nodeAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddDeposit(&_Validators.TransactOpts, vPub, nodeAddr)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x7e332a46.
//
// Solidity: function addDeposit(vPub bytes32, nodeAddr address) returns()
func (_Validators *ValidatorsTransactorSession) AddDeposit(vPub [32]byte, nodeAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddDeposit(&_Validators.TransactOpts, vPub, nodeAddr)
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
