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
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCompactedValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vAddr\",\"type\":\"address\"}],\"name\":\"hasDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vAddr\",\"type\":\"address\"}],\"name\":\"getNodeAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vAddr\",\"type\":\"address\"},{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"addDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_CAUSE_VOLUNTARILY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vAddr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vAddr\",\"type\":\"address\"}],\"name\":\"isPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vAddr\",\"type\":\"address\"},{\"name\":\"vFrom\",\"type\":\"address\"},{\"name\":\"cause\",\"type\":\"uint8\"}],\"name\":\"pauseValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_CAUSE_PUNISHMENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vAddr\",\"type\":\"address\"}],\"name\":\"resumeValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_LOCK_BLOCKS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
const ValidatorsBin = `0x6060604052341561000f57600080fd5b610a778061001e6000396000f3006060604052600436106100b95763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663181a8d6881146100be5780631c48c7ec1461012457806320aca18c146101575780632f1c563e1461019257806347fdd5f8146101ae57806351cff8d9146101d75780635b14f183146101f65780636393ac7b14610215578063be1d05c214610240578063d67d0a3a14610253578063e1e158a514610272578063faeedb6114610297575b600080fd5b34156100c957600080fd5b6100d16102c5565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101105780820151838201526020016100f8565b505050509050019250505060405180910390f35b341561012f57600080fd5b610143600160a060020a0360043516610325565b604051901515815260200160405180910390f35b341561016257600080fd5b610176600160a060020a0360043516610356565b604051600160a060020a03909116815260200160405180910390f35b6101ac600160a060020a0360043581169060243516610374565b005b34156101b957600080fd5b6101c16104c2565b60405160ff909116815260200160405180910390f35b34156101e257600080fd5b6101ac600160a060020a03600435166104c7565b341561020157600080fd5b610143600160a060020a03600435166105b4565b341561022057600080fd5b6101ac600160a060020a036004358116906024351660ff604435166105e0565b341561024b57600080fd5b6101c1610728565b341561025e57600080fd5b6101ac600160a060020a036004351661072d565b341561027d57600080fd5b6102856107dc565b60405190815260200160405180910390f35b34156102a257600080fd5b6102aa6107e8565b60405165ffffffffffff909116815260200160405180910390f35b6102cd6109fc565b600180548060200260200160405190810160405280929190818152602001828054801561031a57602002820191906000526020600020905b81548152600190910190602001808311610305575b505050505090505b90565b600160a060020a031660009081526020819052604081205460a060020a90046bffffffffffffffffffffffff161190565b600160a060020a039081166000908152602081905260409020541690565b600160a060020a03808316600090815260208190526040812080546bffffffffffffffffffffffff60a060020a808304821634018216810292909516919091178083559193670de0b6b3a764000092041610156103d057600080fd5b8154600160a060020a0316151561040857815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161782555b61041284836107ed565b6001830154909150600067010000000000000090910463ffffffff16111561046a576001828101548154839291670100000000000000900463ffffffff1690811061045957fe5b6000918252602090912001556104bc565b6001805480820161047b8382610a0e565b50600091825260209091200181905560018054908301805463ffffffff909216670100000000000000026affffffff00000000000000199092169190911790555b50505050565b600181565b6000806104d383610325565b15156104de57600080fd5b6104e7836105b4565b15156104f257600080fd5b600160a060020a038381166000908152602081905260409020805490935033821691161461051f57600080fd5b6001828101546601000000000000900460ff161461053c57600080fd5b600182015465ffffffffffff438116918116600a0116111561055d57600080fd5b50805460a060020a90046bffffffffffffffffffffffff1661057e83610856565b600160a060020a03331681156108fc0282604051600060405180830381858888f1935050505015156105af57600080fd5b505050565b600160a060020a03166000908152602081905260408120600101546601000000000000900460ff161190565b60006105eb846105b4565b1580156105fc57506105fc84610325565b151561060757600080fd5b600160ff8316101561061857600080fd5b33600160a060020a031661062b84610356565b600160a060020a03161461063e57600080fd5b600160a060020a038481169084161461068d5761065a83610325565b801561066c575061066a836105b4565b155b151561067757600080fd5b60ff82166001141561068857600080fd5b61069d565b60ff821660011461069d57600080fd5b50600160a060020a038316600090815260208190526040902060018101805465ffffffffffff19164365ffffffffffff161766ff0000000000001916660100000000000060ff8516021790556106f384826107ed565b600182810154815467010000000000000090910463ffffffff1690811061071657fe5b60009182526020909120015550505050565b600281565b6000610738826105b4565b151561074357600080fd5b50600160a060020a0381811660009081526020819052604090208054909133811691161461077057600080fd5b6001818101546601000000000000900460ff161461078d57600080fd5b60018101805466ffffffffffffff191690556107a933826107ed565b600182810154815467010000000000000090910463ffffffff169081106107cc57fe5b6000918252602090912001555050565b670de0b6b3a764000081565b600a81565b60018101549054600160a060020a03660100000000000090920460ff167c01000000000000000000000000000000000000000000000000000000000264010000000060a060020a928390046bffffffffffffffffffffffff160490910292909216919091161690565b600160a060020a03811660009081526020819052604081208054600191908390819060a060020a90046bffffffffffffffffffffffff1681901161089957600080fd5b8354600090116108a557fe5b60018360010160079054906101000a900463ffffffff16039150600184805490501180156108de575083546000190163ffffffff831614155b15610996578354849060001981019081106108f557fe5b906000526020600020900154848363ffffffff1681548110151561091557fe5b600091825260209091200155835461094a90859063ffffffff851690811061093957fe5b906000526020600020900154610322565b600160a060020a0381166000908152602087905260409020600190810180546affffffff00000000000000191667010000000000000092860163ffffffff169290920291909117905590505b8354849060001981019081106109a857fe5b600091825260208220015583546109c3856000198301610a0e565b505050600160a060020a039093166000908152602092909252506040812090815560010180546affffffffffffffffffffff1916905550565b60206040519081016040526000815290565b8154818355818115116105af576000838152602090206105af91810190830161032291905b80821115610a475760008155600101610a33565b50905600a165627a7a72305820be64d50614e74597ce3400b7e342c6c6150ba61fea0bd83f497c23f1131d12950029`

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
	return address, tx, &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}}, nil
}

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
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
	contract, err := bindValidators(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// DEPOSIT_LOCK_BLOCKS is a free data retrieval call binding the contract method 0xfaeedb61.
//
// Solidity: function DEPOSIT_LOCK_BLOCKS() constant returns(uint48)
func (_Validators *ValidatorsCaller) DEPOSIT_LOCK_BLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "DEPOSIT_LOCK_BLOCKS")
	return *ret0, err
}

// DEPOSIT_LOCK_BLOCKS is a free data retrieval call binding the contract method 0xfaeedb61.
//
// Solidity: function DEPOSIT_LOCK_BLOCKS() constant returns(uint48)
func (_Validators *ValidatorsSession) DEPOSIT_LOCK_BLOCKS() (*big.Int, error) {
	return _Validators.Contract.DEPOSIT_LOCK_BLOCKS(&_Validators.CallOpts)
}

// DEPOSIT_LOCK_BLOCKS is a free data retrieval call binding the contract method 0xfaeedb61.
//
// Solidity: function DEPOSIT_LOCK_BLOCKS() constant returns(uint48)
func (_Validators *ValidatorsCallerSession) DEPOSIT_LOCK_BLOCKS() (*big.Int, error) {
	return _Validators.Contract.DEPOSIT_LOCK_BLOCKS(&_Validators.CallOpts)
}

// MIN_DEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsCaller) MIN_DEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "MIN_DEPOSIT")
	return *ret0, err
}

// MIN_DEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsSession) MIN_DEPOSIT() (*big.Int, error) {
	return _Validators.Contract.MIN_DEPOSIT(&_Validators.CallOpts)
}

// MIN_DEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MIN_DEPOSIT() (*big.Int, error) {
	return _Validators.Contract.MIN_DEPOSIT(&_Validators.CallOpts)
}

// PAUSE_CAUSE_PUNISHMENT is a free data retrieval call binding the contract method 0xbe1d05c2.
//
// Solidity: function PAUSE_CAUSE_PUNISHMENT() constant returns(uint8)
func (_Validators *ValidatorsCaller) PAUSE_CAUSE_PUNISHMENT(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "PAUSE_CAUSE_PUNISHMENT")
	return *ret0, err
}

// PAUSE_CAUSE_PUNISHMENT is a free data retrieval call binding the contract method 0xbe1d05c2.
//
// Solidity: function PAUSE_CAUSE_PUNISHMENT() constant returns(uint8)
func (_Validators *ValidatorsSession) PAUSE_CAUSE_PUNISHMENT() (uint8, error) {
	return _Validators.Contract.PAUSE_CAUSE_PUNISHMENT(&_Validators.CallOpts)
}

// PAUSE_CAUSE_PUNISHMENT is a free data retrieval call binding the contract method 0xbe1d05c2.
//
// Solidity: function PAUSE_CAUSE_PUNISHMENT() constant returns(uint8)
func (_Validators *ValidatorsCallerSession) PAUSE_CAUSE_PUNISHMENT() (uint8, error) {
	return _Validators.Contract.PAUSE_CAUSE_PUNISHMENT(&_Validators.CallOpts)
}

// PAUSE_CAUSE_VOLUNTARILY is a free data retrieval call binding the contract method 0x47fdd5f8.
//
// Solidity: function PAUSE_CAUSE_VOLUNTARILY() constant returns(uint8)
func (_Validators *ValidatorsCaller) PAUSE_CAUSE_VOLUNTARILY(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "PAUSE_CAUSE_VOLUNTARILY")
	return *ret0, err
}

// PAUSE_CAUSE_VOLUNTARILY is a free data retrieval call binding the contract method 0x47fdd5f8.
//
// Solidity: function PAUSE_CAUSE_VOLUNTARILY() constant returns(uint8)
func (_Validators *ValidatorsSession) PAUSE_CAUSE_VOLUNTARILY() (uint8, error) {
	return _Validators.Contract.PAUSE_CAUSE_VOLUNTARILY(&_Validators.CallOpts)
}

// PAUSE_CAUSE_VOLUNTARILY is a free data retrieval call binding the contract method 0x47fdd5f8.
//
// Solidity: function PAUSE_CAUSE_VOLUNTARILY() constant returns(uint8)
func (_Validators *ValidatorsCallerSession) PAUSE_CAUSE_VOLUNTARILY() (uint8, error) {
	return _Validators.Contract.PAUSE_CAUSE_VOLUNTARILY(&_Validators.CallOpts)
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(bytes32[])
func (_Validators *ValidatorsCaller) GetCompactedValidators(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getCompactedValidators")
	return *ret0, err
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(bytes32[])
func (_Validators *ValidatorsSession) GetCompactedValidators() ([][32]byte, error) {
	return _Validators.Contract.GetCompactedValidators(&_Validators.CallOpts)
}

// GetCompactedValidators is a free data retrieval call binding the contract method 0x181a8d68.
//
// Solidity: function getCompactedValidators() constant returns(bytes32[])
func (_Validators *ValidatorsCallerSession) GetCompactedValidators() ([][32]byte, error) {
	return _Validators.Contract.GetCompactedValidators(&_Validators.CallOpts)
}

// GetNodeAddr is a free data retrieval call binding the contract method 0x20aca18c.
//
// Solidity: function getNodeAddr(vAddr address) constant returns(address)
func (_Validators *ValidatorsCaller) GetNodeAddr(opts *bind.CallOpts, vAddr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getNodeAddr", vAddr)
	return *ret0, err
}

// GetNodeAddr is a free data retrieval call binding the contract method 0x20aca18c.
//
// Solidity: function getNodeAddr(vAddr address) constant returns(address)
func (_Validators *ValidatorsSession) GetNodeAddr(vAddr common.Address) (common.Address, error) {
	return _Validators.Contract.GetNodeAddr(&_Validators.CallOpts, vAddr)
}

// GetNodeAddr is a free data retrieval call binding the contract method 0x20aca18c.
//
// Solidity: function getNodeAddr(vAddr address) constant returns(address)
func (_Validators *ValidatorsCallerSession) GetNodeAddr(vAddr common.Address) (common.Address, error) {
	return _Validators.Contract.GetNodeAddr(&_Validators.CallOpts, vAddr)
}

// HasDeposit is a free data retrieval call binding the contract method 0x1c48c7ec.
//
// Solidity: function hasDeposit(vAddr address) constant returns(bool)
func (_Validators *ValidatorsCaller) HasDeposit(opts *bind.CallOpts, vAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "hasDeposit", vAddr)
	return *ret0, err
}

// HasDeposit is a free data retrieval call binding the contract method 0x1c48c7ec.
//
// Solidity: function hasDeposit(vAddr address) constant returns(bool)
func (_Validators *ValidatorsSession) HasDeposit(vAddr common.Address) (bool, error) {
	return _Validators.Contract.HasDeposit(&_Validators.CallOpts, vAddr)
}

// HasDeposit is a free data retrieval call binding the contract method 0x1c48c7ec.
//
// Solidity: function hasDeposit(vAddr address) constant returns(bool)
func (_Validators *ValidatorsCallerSession) HasDeposit(vAddr common.Address) (bool, error) {
	return _Validators.Contract.HasDeposit(&_Validators.CallOpts, vAddr)
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(vAddr address) constant returns(bool)
func (_Validators *ValidatorsCaller) IsPaused(opts *bind.CallOpts, vAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isPaused", vAddr)
	return *ret0, err
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(vAddr address) constant returns(bool)
func (_Validators *ValidatorsSession) IsPaused(vAddr common.Address) (bool, error) {
	return _Validators.Contract.IsPaused(&_Validators.CallOpts, vAddr)
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(vAddr address) constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsPaused(vAddr common.Address) (bool, error) {
	return _Validators.Contract.IsPaused(&_Validators.CallOpts, vAddr)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x2f1c563e.
//
// Solidity: function addDeposit(vAddr address, nodeAddr address) returns()
func (_Validators *ValidatorsTransactor) AddDeposit(opts *bind.TransactOpts, vAddr common.Address, nodeAddr common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addDeposit", vAddr, nodeAddr)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x2f1c563e.
//
// Solidity: function addDeposit(vAddr address, nodeAddr address) returns()
func (_Validators *ValidatorsSession) AddDeposit(vAddr common.Address, nodeAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddDeposit(&_Validators.TransactOpts, vAddr, nodeAddr)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x2f1c563e.
//
// Solidity: function addDeposit(vAddr address, nodeAddr address) returns()
func (_Validators *ValidatorsTransactorSession) AddDeposit(vAddr common.Address, nodeAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddDeposit(&_Validators.TransactOpts, vAddr, nodeAddr)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x6393ac7b.
//
// Solidity: function pauseValidation(vAddr address, vFrom address, cause uint8) returns()
func (_Validators *ValidatorsTransactor) PauseValidation(opts *bind.TransactOpts, vAddr common.Address, vFrom common.Address, cause uint8) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "pauseValidation", vAddr, vFrom, cause)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x6393ac7b.
//
// Solidity: function pauseValidation(vAddr address, vFrom address, cause uint8) returns()
func (_Validators *ValidatorsSession) PauseValidation(vAddr common.Address, vFrom common.Address, cause uint8) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, vAddr, vFrom, cause)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x6393ac7b.
//
// Solidity: function pauseValidation(vAddr address, vFrom address, cause uint8) returns()
func (_Validators *ValidatorsTransactorSession) PauseValidation(vAddr common.Address, vFrom common.Address, cause uint8) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, vAddr, vFrom, cause)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0xd67d0a3a.
//
// Solidity: function resumeValidation(vAddr address) returns()
func (_Validators *ValidatorsTransactor) ResumeValidation(opts *bind.TransactOpts, vAddr common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "resumeValidation", vAddr)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0xd67d0a3a.
//
// Solidity: function resumeValidation(vAddr address) returns()
func (_Validators *ValidatorsSession) ResumeValidation(vAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ResumeValidation(&_Validators.TransactOpts, vAddr)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0xd67d0a3a.
//
// Solidity: function resumeValidation(vAddr address) returns()
func (_Validators *ValidatorsTransactorSession) ResumeValidation(vAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ResumeValidation(&_Validators.TransactOpts, vAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(vAddr address) returns()
func (_Validators *ValidatorsTransactor) Withdraw(opts *bind.TransactOpts, vAddr common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdraw", vAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(vAddr address) returns()
func (_Validators *ValidatorsSession) Withdraw(vAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts, vAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(vAddr address) returns()
func (_Validators *ValidatorsTransactorSession) Withdraw(vAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts, vAddr)
}
