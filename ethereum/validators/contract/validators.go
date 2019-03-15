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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a031916331790556101fd806100326000396000f3fe608060405234801561001057600080fd5b506004361061005d577c01000000000000000000000000000000000000000000000000000000006000350463715018a681146100625780638da5cb5b1461006c578063f2fde38b14610090575b600080fd5b61006a6100b6565b005b610074610122565b60408051600160a060020a039092168252519081900360200190f35b61006a600480360360208110156100a657600080fd5b5035600160a060020a0316610131565b600054600160a060020a031633146100cd57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461014857600080fd5b61015181610154565b50565b600160a060020a038116151561016957600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556fea165627a7a7230582012136b96509a538d25fd0b2f5da9e02339631e4fcf85501ab519c7f1cfe8c3de0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"getNodeAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"c_MIN_DEPOSIT_INCREMENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompactedValidators\",\"outputs\":[{\"name\":\"ValidatorsCompacted\",\"type\":\"bytes32[]\"},{\"name\":\"ValidatorsPubKeys\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"addDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"},{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"addInitialDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"isPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"},{\"name\":\"vFrom\",\"type\":\"bytes32\"},{\"name\":\"cause\",\"type\":\"uint8\"},{\"name\":\"punishValue\",\"type\":\"uint96\"}],\"name\":\"pauseValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"hasDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"resumeValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"getNodeReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveCompactedValidators\",\"outputs\":[{\"name\":\"ValidatorsCompacted\",\"type\":\"bytes32[]\"},{\"name\":\"ValidatorsPubKeys\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"c_MIN_DEPOSIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vPub\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint96\"},{\"name\":\"pauseBlockNumber\",\"type\":\"uint48\"},{\"name\":\"pauseCause\",\"type\":\"uint8\"},{\"name\":\"punishValue\",\"type\":\"uint96\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min_dep\",\"type\":\"uint256\"},{\"name\":\"min_dep_inc\",\"type\":\"uint256\"}],\"name\":\"setDepositBounds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"all\",\"type\":\"bool\"}],\"name\":\"enablePunishers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vPub\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint96\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vPub\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cause\",\"type\":\"uint8\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
const ValidatorsBin = `0x6080604052678ac7230489e8000060015568878678326eac9000006002556006805460ff1916905560008054600160a060020a03191633179055611b1e806100486000396000f3fe60806040526004361061013c576000357c0100000000000000000000000000000000000000000000000000000000900480638e19899e116100bd578063bfe0eb5211610081578063bfe0eb5214610431578063d5f20ff614610446578063d64c2ead146104c4578063f1f58d92146104f4578063f2fde38b146105205761013c565b80638e19899e1461037457806392c8a9661461039e5780639da0e3e3146103c8578063a36147d3146103f2578063acfb8f761461041c5761013c565b8063241b71bb11610104578063241b71bb146102af5780632d5f512a146102ed57806363338b1714610335578063715018a61461034a5780638da5cb5b1461035f5761013c565b80630b8604fc14610141578063155e1abc14610187578063181a8d68146101ae5780631c32d9821461025c57806320aef8d41461027b575b600080fd5b34801561014d57600080fd5b5061016b6004803603602081101561016457600080fd5b5035610553565b60408051600160a060020a039092168252519081900360200190f35b34801561019357600080fd5b5061019c61056e565b60408051918252519081900360200190f35b3480156101ba57600080fd5b506101c3610574565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156102075781810151838201526020016101ef565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561024657818101518382015260200161022e565b5050505090500194505050505060405180910390f35b6102796004803603602081101561027257600080fd5b5035610621565b005b6102796004803603606081101561029157600080fd5b50803590600160a060020a0360208201358116916040013516610630565b3480156102bb57600080fd5b506102d9600480360360208110156102d257600080fd5b50356108d5565b604080519115158252519081900360200190f35b3480156102f957600080fd5b506102796004803603608081101561031057600080fd5b50803590602081013590604081013560ff1690606001356001606060020a03166108f8565b34801561034157600080fd5b5061019c610cc5565b34801561035657600080fd5b50610279610d2f565b34801561036b57600080fd5b5061016b610d9b565b34801561038057600080fd5b506102796004803603602081101561039757600080fd5b5035610daa565b3480156103aa57600080fd5b506102d9600480360360208110156103c157600080fd5b5035610fe5565b3480156103d457600080fd5b50610279600480360360208110156103eb57600080fd5b5035611008565b3480156103fe57600080fd5b5061016b6004803603602081101561041557600080fd5b5035611326565b34801561042857600080fd5b506101c361137b565b34801561043d57600080fd5b5061019c61149b565b34801561045257600080fd5b506104706004803603602081101561046957600080fd5b50356114a1565b60408051600160a060020a0397881681526001606060020a03968716602082015265ffffffffffff9095168582015260ff9093166060850152931660808301529190921660a0830152519081900360c00190f35b3480156104d057600080fd5b50610279600480360360408110156104e757600080fd5b5080359060200135611509565b34801561050057600080fd5b506102796004803603602081101561051757600080fd5b5035151561152b565b34801561052c57600080fd5b506102796004803603602081101561054357600080fd5b5035600160a060020a0316611555565b600090815260036020526040902054600160a060020a031690565b60015481565b60608060058054806020026020016040519081016040528092919081815260200182805480156105c357602002820191906000526020600020905b8154815260200190600101908083116105af575b50505050509150600480548060200260200160405190810160405280929190818152602001828054801561061657602002820191906000526020600020905b815481526020019060010190808311610602575b505050505090509091565b61062d81600080610630565b50565b6001543410156106745760405160e560020a62461bcd0281526004018080602001828103825260258152602001806119656025913960400191505060405180910390fd5b600083815260036020526040902080546001606060020a0360a060020a8083048216340190911602600160a060020a03918216178083551615156107aa57600160a060020a0383161515610712576040805160e560020a62461bcd02815260206004820152601b60248201527f506c656173652073706563696679206e6f646520616464726573730000000000604482015290519081900360640190fd5b805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384811691909117825560018201805466010000000000007fff0000000000000000000000000000000000000000ffffffffffffffffffffff9091166b010000000000000000000000938616939093029290921766ff00000000000019169190911765ffffffffffff19164365ffffffffffff161790555b6001810154600067010000000000000090910463ffffffff1611156107d8576107d281611575565b5061087a565b60006107e3826115bd565b60048054600181810183557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9091018890556005805480830182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0019290925554908301805463ffffffff909216670100000000000000026affffffff0000000000000019909216919091179055505b60003411156108cf5780546040805186815260a060020a9092046001606060020a0316602083015280517fd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf07889281900390910190a15b50505050565b6000908152600360205260408120600101546601000000000000900460ff161190565b610901846108d5565b158015610912575061091284610fe5565b15156109525760405160e560020a62461bcd0281526004018080602001828103825260318152602001806119dc6031913960400191505060405180910390fd5b600160ff831610156109ae576040805160e560020a62461bcd02815260206004820152601860248201527f596f752073686f756c6420737065636966792063617573650000000000000000604482015290519081900360640190fd5b336109b884610553565b600160a060020a031614610a005760405160e560020a62461bcd028152600401808060200182810382526033815260200180611a8c6033913960400191505060405180910390fd5b60ff821660011480610a1c5750600054600160a060020a031633145b80610a29575060065460ff165b1515610a7f576040805160e560020a62461bcd02815260206004820152600e60248201527f57726f6e672070756e6973686572000000000000000000000000000000000000604482015290519081900360640190fd5b60ff82166002141580610aa05750655af3107a40006001606060020a038216105b1515610af6576040805160e560020a62461bcd02815260206004820152601360248201527f57726f6e6720626c6f636b73206e756d62657200000000000000000000000000604482015290519081900360640190fd5b838314610bb957610b0683610fe5565b8015610b185750610b16836108d5565b155b1515610b585760405160e560020a62461bcd02815260040180806020018281038252602881526020018061193d6028913960400191505060405180910390fd5b60ff821660011415610bb4576040805160e560020a62461bcd02815260206004820152601e60248201527f546869732070617573696e67206973206e6f7420766f6c756e74617279210000604482015290519081900360640190fd5b610c14565b60ff8216600114610c14576040805160e560020a62461bcd02815260206004820152601b60248201527f596f75206172652070617573696e6720766f6c756e746172696c790000000000604482015290519081900360640190fd5b600084815260036020526040902060018101805465ffffffffffff19164365ffffffffffff161766ff0000000000001916660100000000000060ff8616021790556002810180546bffffffffffffffffffffffff19166001606060020a038416179055610c8081611575565b506040805186815260ff8516602082015281517f2ffa21f275941345406452ac41a8bceea9dde4f76b233ac4fc9e53c31e399342929181900390910190a15050505050565b600554600090815b81811015610d2a576000600582815481101515610ce657fe5b60009182526020909120015490507cff0000000000000000000000000000000000000000000000000000000081161515610d21578360010193505b50600101610ccd565b505090565b600054600160a060020a03163314610d4657600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b610db381610fe5565b1515610e09576040805160e560020a62461bcd02815260206004820152601760248201527f596f752073686f756c642068617665206465706f736974000000000000000000604482015290519081900360640190fd5b610e12816108d5565b1515610e68576040805160e560020a62461bcd02815260206004820152601460248201527f596f752073686f756c6420626520706175736564000000000000000000000000604482015290519081900360640190fd5b60008181526003602052604090208054600160a060020a03163314610ed7576040805160e560020a62461bcd02815260206004820152601e60248201527f4f6e6c79206e6f646520697473656c662063616e207769746864726177210000604482015290519081900360640190fd5b6001818101546601000000000000900460ff1614610f295760405160e560020a62461bcd028152600401808060200182810382526032815260200180611a0d6032913960400191505060405180910390fd5b600181015465ffffffffffff43811691811661177001161115610f805760405160e560020a62461bcd02815260040180806020018281038252602b8152602001806119b1602b913960400191505060405180910390fd5b805460a060020a90046001606060020a03166000610f9d84611326565b9050610fa88461161f565b604051600160a060020a0382169083156108fc029084906000818181858888f19350505050158015610fde573d6000803e3d6000fd5b5050505050565b60009081526003602052604081205460a060020a90046001606060020a03161190565b611011816108d5565b1515611067576040805160e560020a62461bcd02815260206004820152601860248201527f54686973206e6f6465206973206e6f7420706175736564210000000000000000604482015290519081900360640190fd5b60008181526003602052604090208054600160a060020a031633146110d6576040805160e560020a62461bcd02815260206004820152601d60248201527f596f752063616e206f6e6c7920756e706175736520796f757273656c66000000604482015290519081900360640190fd5b6001818101546601000000000000900460ff16908114806110fa575060ff81166002145b80611108575060ff81166003145b15156111485760405160e560020a62461bcd028152600401808060200182810382526023815260200180611a3f6023913960400191505060405180910390fd5b60ff8116600214156111b55760028201546001830154436001606060020a0390921665ffffffffffff9091160111156111b55760405160e560020a62461bcd02815260040180806020018281038252602a815260200180611a62602a913960400191505060405180910390fd5b60ff81166003141561124e57600282015482546001606060020a0391821660a060020a909104909116101561121e5760405160e560020a62461bcd02815260040180806020018281038252602781526020018061198a6027913960400191505060405180910390fd5b60028201548254600160a060020a0381166001606060020a0392831660a060020a92839004841603909216021782555b600254825460a060020a90046001606060020a031610156112a35760405160e560020a62461bcd028152600401808060200182810382526034815260200180611abf6034913960400191505060405180910390fd5b60018201805466ffffffffffffff191690556002820180546bffffffffffffffffffffffff191690556112d582611575565b5081546040805185815260a060020a9092046001606060020a0316602083015280517fd816ca5f3fead11e08d6cf12bbf4d0cb3f96eec4513bf5e32712796d92cf07889281900390910190a1505050565b6000818152600360205260408120600101546b0100000000000000000000009004600160a060020a03168015156113635761136083610553565b90505b600160a060020a038116151561137557fe5b92915050565b6060806000611388610cc5565b9050806040519080825280602002602001820160405280156113b4578160200160208202803883390190505b509250806040519080825280602002602001820160405280156113e1578160200160208202803883390190505b506005549092506000805b8281101561149357600060058281548110151561140557fe5b60009182526020909120015490507cff000000000000000000000000000000000000000000000000000000008116151561148a5780878481518110151561144857fe5b60209081029091010152600480548390811061146057fe5b9060005260206000200154868481518110151561147957fe5b602090810290910101526001909201915b506001016113ec565b505050509091565b60025481565b600090815260036020526040902080546001820154600290920154600160a060020a03808316946001606060020a0360a060020a90940484169465ffffffffffff81169460ff66010000000000008304169416926b0100000000000000000000009091041690565b600054600160a060020a0316331461152057600080fd5b600291909155600155565b600054600160a060020a0316331461154257600080fd5b6006805460ff1916911515919091179055565b600054600160a060020a0316331461156c57600080fd5b61062d81611875565b6000611580826115bd565b600183015460058054909160001963ffffffff670100000000000000909204821601169081106115ac57fe5b600091825260209091200155919050565b80546001909101546601000000000000900460ff167c010000000000000000000000000000000000000000000000000000000002600160a060020a03821664010000000060a060020a938490046001606060020a031604909202919091171790565b60008181526003602081905260408220805491926004926005929160a060020a9091046001606060020a0316116116a0576040805160e560020a62461bcd02815260206004820152601560248201527f4465706f7369742073686f756c64206578697374210000000000000000000000604482015290519081900360640190fd5b82546000106116ab57fe5b600181810154845467010000000000000090910463ffffffff1660001901911080156116e2575083546000190163ffffffff821614155b156117c4578354849060001981019081106116f957fe5b9060005260206000200154848263ffffffff1681548110151561171857fe5b60009182526020909120015582548390600019810190811061173657fe5b9060005260206000200154838263ffffffff1681548110151561175557fe5b90600052602060002001819055506000848263ffffffff1681548110151561177957fe5b600091825260208083209190910154825287905260409020600190810180546affffffff00000000000000191667010000000000000092850163ffffffff1692909202919091179055505b8354849060001981019081106117d657fe5b600091825260208220015583546117f18560001983016118f2565b5082548390600019810190811061180457fe5b6000918252602082200155825461181f8460001983016118f2565b50505060009384525050602052604081209081556001810180547fff0000000000000000000000000000000000000000000000000000000000000016905560020180546bffffffffffffffffffffffff19169055565b600160a060020a038116151561188a57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b8154818355818111156119165760008381526020902061191691810190830161191b565b505050565b61193991905b808211156119355760008155600101611921565b5090565b9056fe596f752073686f756c64206e6f742062652070617573656420746f207061757365206f7468657273546f6f20736d616c6c2076616c756520746f2061646420746f20746865206465706f7369744465706f736974206973206e6f7420656e6f75676820746f20636f766572207468652066696e65596f752073686f756c64207761697420736f6d6520626c6f636b73206265666f72652077697468647261774e6f64652073686f756c64206e6f742062652070617573656420616e642073686f756c642068617665206465706f736974596f752073686f756c64206861766520766f6c756e746172696c7920706175736564206265666f7265207769746864726177596f752063616e206e6f7420756e70617573652066726f6d2074686973207374617465596f752073686f756c64207761697420736f6d6520626c6f636b73206265666f726520756e70617573654e6f64652073686f756c6420636f72726563746c792070617373206974732076616c696461746f72207075626c6963206b6579596f752063616e206e6f7420756e7061757365206265666f7265206465706f7369742065786365656473206d696e2076616c7565a165627a7a72305820834c3f1444b7d69f5280eb2d78f14f176c410c93fd2b7b102a9a64fd26f4ed850029`

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

// CMINDEPOSIT is a free data retrieval call binding the contract method 0xbfe0eb52.
//
// Solidity: function c_MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsCaller) CMINDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "c_MIN_DEPOSIT")
	return *ret0, err
}

// CMINDEPOSIT is a free data retrieval call binding the contract method 0xbfe0eb52.
//
// Solidity: function c_MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsSession) CMINDEPOSIT() (*big.Int, error) {
	return _Validators.Contract.CMINDEPOSIT(&_Validators.CallOpts)
}

// CMINDEPOSIT is a free data retrieval call binding the contract method 0xbfe0eb52.
//
// Solidity: function c_MIN_DEPOSIT() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CMINDEPOSIT() (*big.Int, error) {
	return _Validators.Contract.CMINDEPOSIT(&_Validators.CallOpts)
}

// CMINDEPOSITINCREMENT is a free data retrieval call binding the contract method 0x155e1abc.
//
// Solidity: function c_MIN_DEPOSIT_INCREMENT() constant returns(uint256)
func (_Validators *ValidatorsCaller) CMINDEPOSITINCREMENT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "c_MIN_DEPOSIT_INCREMENT")
	return *ret0, err
}

// CMINDEPOSITINCREMENT is a free data retrieval call binding the contract method 0x155e1abc.
//
// Solidity: function c_MIN_DEPOSIT_INCREMENT() constant returns(uint256)
func (_Validators *ValidatorsSession) CMINDEPOSITINCREMENT() (*big.Int, error) {
	return _Validators.Contract.CMINDEPOSITINCREMENT(&_Validators.CallOpts)
}

// CMINDEPOSITINCREMENT is a free data retrieval call binding the contract method 0x155e1abc.
//
// Solidity: function c_MIN_DEPOSIT_INCREMENT() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CMINDEPOSITINCREMENT() (*big.Int, error) {
	return _Validators.Contract.CMINDEPOSITINCREMENT(&_Validators.CallOpts)
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
// Solidity: function getValidator(vPub bytes32) constant returns(nodeAddr address, deposit uint96, pauseBlockNumber uint48, pauseCause uint8, punishValue uint96, receiver address)
func (_Validators *ValidatorsCaller) GetValidator(opts *bind.CallOpts, vPub [32]byte) (struct {
	NodeAddr         common.Address
	Deposit          *big.Int
	PauseBlockNumber *big.Int
	PauseCause       uint8
	PunishValue      *big.Int
	Receiver         common.Address
}, error) {
	ret := new(struct {
		NodeAddr         common.Address
		Deposit          *big.Int
		PauseBlockNumber *big.Int
		PauseCause       uint8
		PunishValue      *big.Int
		Receiver         common.Address
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "getValidator", vPub)
	return *ret, err
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(vPub bytes32) constant returns(nodeAddr address, deposit uint96, pauseBlockNumber uint48, pauseCause uint8, punishValue uint96, receiver address)
func (_Validators *ValidatorsSession) GetValidator(vPub [32]byte) (struct {
	NodeAddr         common.Address
	Deposit          *big.Int
	PauseBlockNumber *big.Int
	PauseCause       uint8
	PunishValue      *big.Int
	Receiver         common.Address
}, error) {
	return _Validators.Contract.GetValidator(&_Validators.CallOpts, vPub)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(vPub bytes32) constant returns(nodeAddr address, deposit uint96, pauseBlockNumber uint48, pauseCause uint8, punishValue uint96, receiver address)
func (_Validators *ValidatorsCallerSession) GetValidator(vPub [32]byte) (struct {
	NodeAddr         common.Address
	Deposit          *big.Int
	PauseBlockNumber *big.Int
	PauseCause       uint8
	PunishValue      *big.Int
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsSession) Owner() (common.Address, error) {
	return _Validators.Contract.Owner(&_Validators.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsCallerSession) Owner() (common.Address, error) {
	return _Validators.Contract.Owner(&_Validators.CallOpts)
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

// EnablePunishers is a paid mutator transaction binding the contract method 0xf1f58d92.
//
// Solidity: function enablePunishers(all bool) returns()
func (_Validators *ValidatorsTransactor) EnablePunishers(opts *bind.TransactOpts, all bool) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "enablePunishers", all)
}

// EnablePunishers is a paid mutator transaction binding the contract method 0xf1f58d92.
//
// Solidity: function enablePunishers(all bool) returns()
func (_Validators *ValidatorsSession) EnablePunishers(all bool) (*types.Transaction, error) {
	return _Validators.Contract.EnablePunishers(&_Validators.TransactOpts, all)
}

// EnablePunishers is a paid mutator transaction binding the contract method 0xf1f58d92.
//
// Solidity: function enablePunishers(all bool) returns()
func (_Validators *ValidatorsTransactorSession) EnablePunishers(all bool) (*types.Transaction, error) {
	return _Validators.Contract.EnablePunishers(&_Validators.TransactOpts, all)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x2d5f512a.
//
// Solidity: function pauseValidation(vPub bytes32, vFrom bytes32, cause uint8, punishValue uint96) returns()
func (_Validators *ValidatorsTransactor) PauseValidation(opts *bind.TransactOpts, vPub [32]byte, vFrom [32]byte, cause uint8, punishValue *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "pauseValidation", vPub, vFrom, cause, punishValue)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x2d5f512a.
//
// Solidity: function pauseValidation(vPub bytes32, vFrom bytes32, cause uint8, punishValue uint96) returns()
func (_Validators *ValidatorsSession) PauseValidation(vPub [32]byte, vFrom [32]byte, cause uint8, punishValue *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, vPub, vFrom, cause, punishValue)
}

// PauseValidation is a paid mutator transaction binding the contract method 0x2d5f512a.
//
// Solidity: function pauseValidation(vPub bytes32, vFrom bytes32, cause uint8, punishValue uint96) returns()
func (_Validators *ValidatorsTransactorSession) PauseValidation(vPub [32]byte, vFrom [32]byte, cause uint8, punishValue *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.PauseValidation(&_Validators.TransactOpts, vPub, vFrom, cause, punishValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validators.Contract.RenounceOwnership(&_Validators.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validators.Contract.RenounceOwnership(&_Validators.TransactOpts)
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

// SetDepositBounds is a paid mutator transaction binding the contract method 0xd64c2ead.
//
// Solidity: function setDepositBounds(min_dep uint256, min_dep_inc uint256) returns()
func (_Validators *ValidatorsTransactor) SetDepositBounds(opts *bind.TransactOpts, min_dep *big.Int, min_dep_inc *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setDepositBounds", min_dep, min_dep_inc)
}

// SetDepositBounds is a paid mutator transaction binding the contract method 0xd64c2ead.
//
// Solidity: function setDepositBounds(min_dep uint256, min_dep_inc uint256) returns()
func (_Validators *ValidatorsSession) SetDepositBounds(min_dep *big.Int, min_dep_inc *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetDepositBounds(&_Validators.TransactOpts, min_dep, min_dep_inc)
}

// SetDepositBounds is a paid mutator transaction binding the contract method 0xd64c2ead.
//
// Solidity: function setDepositBounds(min_dep uint256, min_dep_inc uint256) returns()
func (_Validators *ValidatorsTransactorSession) SetDepositBounds(min_dep *big.Int, min_dep_inc *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetDepositBounds(&_Validators.TransactOpts, min_dep, min_dep_inc)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Validators *ValidatorsTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Validators *ValidatorsSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TransferOwnership(&_Validators.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Validators *ValidatorsTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TransferOwnership(&_Validators.TransactOpts, _newOwner)
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

// ValidatorsOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Validators contract.
type ValidatorsOwnershipRenouncedIterator struct {
	Event *ValidatorsOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ValidatorsOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsOwnershipRenounced)
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
		it.Event = new(ValidatorsOwnershipRenounced)
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
func (it *ValidatorsOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsOwnershipRenounced represents a OwnershipRenounced event raised by the Validators contract.
type ValidatorsOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Validators *ValidatorsFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ValidatorsOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsOwnershipRenouncedIterator{contract: _Validators.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Validators *ValidatorsFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ValidatorsOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsOwnershipRenounced)
				if err := _Validators.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ValidatorsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Validators contract.
type ValidatorsOwnershipTransferredIterator struct {
	Event *ValidatorsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsOwnershipTransferred)
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
		it.Event = new(ValidatorsOwnershipTransferred)
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
func (it *ValidatorsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsOwnershipTransferred represents a OwnershipTransferred event raised by the Validators contract.
type ValidatorsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Validators *ValidatorsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsOwnershipTransferredIterator{contract: _Validators.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Validators *ValidatorsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsOwnershipTransferred)
				if err := _Validators.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
