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
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCompactedValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"hasDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_CAUSE_VOLUNTARILY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resumeValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"cause\",\"type\":\"uint8\"}],\"name\":\"pauseValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_CAUSE_PUNISHMENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_LOCK_BLOCKS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
const ValidatorsBin = `0x6060604052341561000f57600080fd5b6109598061001e6000396000f3006060604052600436106100a35763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663181a8d6881146100af5780631c48c7ec146101155780633ccfd60b1461014857806347fdd5f81461015b5780635b14f183146101845780636e7f5bd3146101a357806376df2516146101b6578063be1d05c2146101db578063e1e158a5146101ee578063faeedb6114610213575b6100ad3334610241565b005b34156100ba57600080fd5b6100c2610349565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101015780820151838201526020016100e9565b505050509050019250505060405180910390f35b341561012057600080fd5b610134600160a060020a03600435166103a9565b604051901515815260200160405180910390f35b341561015357600080fd5b6100ad6103d3565b341561016657600080fd5b61016e6104ae565b60405160ff909116815260200160405180910390f35b341561018f57600080fd5b610134600160a060020a03600435166104b3565b34156101ae57600080fd5b6100ad6104d9565b34156101c157600080fd5b6100ad600160a060020a036004351660ff60243516610571565b34156101e657600080fd5b61016e6106b5565b34156101f957600080fd5b6102016106ba565b60405190815260200160405180910390f35b341561021e57600080fd5b6102266106c6565b60405165ffffffffffff909116815260200160405180910390f35b600160a060020a038216600090815260208190526040812080546bffffffffffffffffffffffff1981166bffffffffffffffffffffffff9182168501821617808355919291670de0b6b3a76400009116101561029c57600080fd5b6102a684836106cb565b82549091506000609860020a90910463ffffffff1611156102f0578154600180548392609860020a900463ffffffff169081106102df57fe5b600091825260209091200155610343565b6001805480820161030183826108d4565b506000918252602090912001819055600154825463ffffffff909116609860020a0276ffffffff00000000000000000000000000000000000000199091161782555b50505050565b6103516108fd565b600180548060200260200160405190810160405280929190818152602001828054801561039e57602002820191906000526020600020905b81548152600190910190602001808311610389575b505050505090505b90565b600160a060020a03166000908152602081905260408120546bffffffffffffffffffffffff161190565b6000806103df336103a9565b15156103ea57600080fd5b6103f3336104b3565b15156103fe57600080fd5b600160a060020a03331660009081526020819052604090208054909250609060020a900460ff1660011461043157600080fd5b815465ffffffffffff4381166c010000000000000000000000009092048116600a0116111561045f57600080fd5b5080546bffffffffffffffffffffffff1661047933610737565b600160a060020a03331681156108fc0282604051600060405180830381858888f1935050505015156104aa57600080fd5b5050565b600181565b600160a060020a0316600090815260208190526040812054609060020a900460ff161190565b60006104e4336104b3565b15156104ef57600080fd5b50600160a060020a03331660009081526020819052604090208054609060020a900460ff1660011461052057600080fd5b805472ffffffffffffff0000000000000000000000001916815561054433826106cb565b8154600180549091609860020a900463ffffffff1690811061056257fe5b60009182526020909120015550565b600061057c836104b3565b15801561058d575061058d836103a9565b151561059857600080fd5b600160ff831610156105a957600080fd5b33600160a060020a031683600160a060020a0316141515610600576105cd336103a9565b80156105df57506105dd336104b3565b155b15156105ea57600080fd5b60ff8216600114156105fb57600080fd5b610610565b60ff821660011461061057600080fd5b50600160a060020a0382166000908152602081905260409020805471ffffffffffff00000000000000000000000019166c010000000000000000000000004365ffffffffffff16021772ff0000000000000000000000000000000000001916609060020a60ff84160217815561068683826106cb565b8154600180549091609860020a900463ffffffff169081106106a457fe5b600091825260209091200155505050565b600281565b670de0b6b3a764000081565b600a81565b54600160a060020a03609060020a820460ff167c0100000000000000000000000000000000000000000000000000000000026401000000006bffffffffffffffffffffffff90931692909204740100000000000000000000000000000000000000000292909216161690565b600160a060020a0381166000908152602081905260408120805460019190839081906bffffffffffffffffffffffff1681901161077357600080fd5b83546000901161077f57fe5b82548454609860020a90910463ffffffff166000190192506001901180156107b2575083546000190163ffffffff831614155b15610868578354849060001981019081106107c957fe5b906000526020600020900154848363ffffffff168154811015156107e957fe5b600091825260209091200155835461081e90859063ffffffff851690811061080d57fe5b9060005260206000209001546103a6565b600160a060020a0381166000908152602087905260409020805476ffffffff000000000000000000000000000000000000001916609860020a6001860163ffffffff160217905590505b83548490600019810190811061087a57fe5b600091825260208220015583546108958560001983016108d4565b505050600160a060020a0390931660009081526020929092525060409020805476ffffffffffffffffffffffffffffffffffffffffffffff1916905550565b8154818355818115116108f8576000838152602090206108f891810190830161090f565b505050565b60206040519081016040526000815290565b6103a691905b808211156109295760008155600101610915565b50905600a165627a7a723058200b149ac9609c6d1db1f4d8621ae5da6916695eb2475b0be2e9fb5507f62a610c0029`

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

// HasDeposit is a free data retrieval call binding the contract method 0x1c48c7ec.
//
// Solidity: function hasDeposit(addr address) constant returns(bool)
func (_Validators *ValidatorsCaller) HasDeposit(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "hasDeposit", addr)
	return *ret0, err
}

// HasDeposit is a free data retrieval call binding the contract method 0x1c48c7ec.
//
// Solidity: function hasDeposit(addr address) constant returns(bool)
func (_Validators *ValidatorsSession) HasDeposit(addr common.Address) (bool, error) {
	return _Validators.Contract.HasDeposit(&_Validators.CallOpts, addr)
}

// HasDeposit is a free data retrieval call binding the contract method 0x1c48c7ec.
//
// Solidity: function hasDeposit(addr address) constant returns(bool)
func (_Validators *ValidatorsCallerSession) HasDeposit(addr common.Address) (bool, error) {
	return _Validators.Contract.HasDeposit(&_Validators.CallOpts, addr)
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(addr address) constant returns(bool)
func (_Validators *ValidatorsCaller) IsPaused(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isPaused", addr)
	return *ret0, err
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(addr address) constant returns(bool)
func (_Validators *ValidatorsSession) IsPaused(addr common.Address) (bool, error) {
	return _Validators.Contract.IsPaused(&_Validators.CallOpts, addr)
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(addr address) constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsPaused(addr common.Address) (bool, error) {
	return _Validators.Contract.IsPaused(&_Validators.CallOpts, addr)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x76df2516.
//
// Solidity: function pauseValidation(addr address, cause uint8) returns()
func (_Validators *ValidatorsTransactor) PauseValidation(opts *bind.TransactOpts, addr common.Address, cause uint8) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "pauseValidation", addr, cause)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x76df2516.
//
// Solidity: function pauseValidation(addr address, cause uint8) returns()
func (_Validators *ValidatorsSession) PauseValidation(addr common.Address, cause uint8) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, addr, cause)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x76df2516.
//
// Solidity: function pauseValidation(addr address, cause uint8) returns()
func (_Validators *ValidatorsTransactorSession) PauseValidation(addr common.Address, cause uint8) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, addr, cause)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0x6e7f5bd3.
//
// Solidity: function resumeValidation() returns()
func (_Validators *ValidatorsTransactor) ResumeValidation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "resumeValidation")
}

// ResumeValidation is a paid mutator transaction binding the contract method 0x6e7f5bd3.
//
// Solidity: function resumeValidation() returns()
func (_Validators *ValidatorsSession) ResumeValidation() (*types.Transaction, error) {
	return _Validators.Contract.ResumeValidation(&_Validators.TransactOpts)
}

// ResumeValidation is a paid mutator transaction binding the contract method 0x6e7f5bd3.
//
// Solidity: function resumeValidation() returns()
func (_Validators *ValidatorsTransactorSession) ResumeValidation() (*types.Transaction, error) {
	return _Validators.Contract.ResumeValidation(&_Validators.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Validators *ValidatorsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Validators *ValidatorsSession) Withdraw() (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Validators *ValidatorsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts)
}
