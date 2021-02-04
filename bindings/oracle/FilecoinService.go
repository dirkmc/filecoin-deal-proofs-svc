// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FilecoinServiceABI is the input ABI used to generate the binding from.
const FilecoinServiceABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dataCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pieceCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"provider\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signedEpoch\",\"type\":\"uint256\"}],\"name\":\"StoredCid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"updatedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dataCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pieceCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"provider\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"submitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dataCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pieceCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"provider\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FilecoinService is an auto generated Go binding around an Ethereum contract.
type FilecoinService struct {
	FilecoinServiceCaller     // Read-only binding to the contract
	FilecoinServiceTransactor // Write-only binding to the contract
	FilecoinServiceFilterer   // Log filterer for contract events
}

// FilecoinServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type FilecoinServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilecoinServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FilecoinServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilecoinServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FilecoinServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilecoinServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FilecoinServiceSession struct {
	Contract     *FilecoinService  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FilecoinServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FilecoinServiceCallerSession struct {
	Contract *FilecoinServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FilecoinServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FilecoinServiceTransactorSession struct {
	Contract     *FilecoinServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FilecoinServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type FilecoinServiceRaw struct {
	Contract *FilecoinService // Generic contract binding to access the raw methods on
}

// FilecoinServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FilecoinServiceCallerRaw struct {
	Contract *FilecoinServiceCaller // Generic read-only contract binding to access the raw methods on
}

// FilecoinServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FilecoinServiceTransactorRaw struct {
	Contract *FilecoinServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFilecoinService creates a new instance of FilecoinService, bound to a specific deployed contract.
func NewFilecoinService(address common.Address, backend bind.ContractBackend) (*FilecoinService, error) {
	contract, err := bindFilecoinService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FilecoinService{FilecoinServiceCaller: FilecoinServiceCaller{contract: contract}, FilecoinServiceTransactor: FilecoinServiceTransactor{contract: contract}, FilecoinServiceFilterer: FilecoinServiceFilterer{contract: contract}}, nil
}

// NewFilecoinServiceCaller creates a new read-only instance of FilecoinService, bound to a specific deployed contract.
func NewFilecoinServiceCaller(address common.Address, caller bind.ContractCaller) (*FilecoinServiceCaller, error) {
	contract, err := bindFilecoinService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FilecoinServiceCaller{contract: contract}, nil
}

// NewFilecoinServiceTransactor creates a new write-only instance of FilecoinService, bound to a specific deployed contract.
func NewFilecoinServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*FilecoinServiceTransactor, error) {
	contract, err := bindFilecoinService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FilecoinServiceTransactor{contract: contract}, nil
}

// NewFilecoinServiceFilterer creates a new log filterer instance of FilecoinService, bound to a specific deployed contract.
func NewFilecoinServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*FilecoinServiceFilterer, error) {
	contract, err := bindFilecoinService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FilecoinServiceFilterer{contract: contract}, nil
}

// bindFilecoinService binds a generic wrapper to an already deployed contract.
func bindFilecoinService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FilecoinServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilecoinService *FilecoinServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilecoinService.Contract.FilecoinServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilecoinService *FilecoinServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilecoinService.Contract.FilecoinServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilecoinService *FilecoinServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilecoinService.Contract.FilecoinServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilecoinService *FilecoinServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilecoinService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilecoinService *FilecoinServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilecoinService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilecoinService *FilecoinServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilecoinService.Contract.contract.Transact(opts, method, params...)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_FilecoinService *FilecoinServiceCaller) Managers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FilecoinService.contract.Call(opts, &out, "managers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_FilecoinService *FilecoinServiceSession) Managers(arg0 common.Address) (bool, error) {
	return _FilecoinService.Contract.Managers(&_FilecoinService.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_FilecoinService *FilecoinServiceCallerSession) Managers(arg0 common.Address) (bool, error) {
	return _FilecoinService.Contract.Managers(&_FilecoinService.CallOpts, arg0)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FilecoinService *FilecoinServiceCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FilecoinService.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FilecoinService *FilecoinServiceSession) MerkleRoot() ([32]byte, error) {
	return _FilecoinService.Contract.MerkleRoot(&_FilecoinService.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FilecoinService *FilecoinServiceCallerSession) MerkleRoot() ([32]byte, error) {
	return _FilecoinService.Contract.MerkleRoot(&_FilecoinService.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilecoinService *FilecoinServiceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FilecoinService.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilecoinService *FilecoinServiceSession) Owner() (common.Address, error) {
	return _FilecoinService.Contract.Owner(&_FilecoinService.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilecoinService *FilecoinServiceCallerSession) Owner() (common.Address, error) {
	return _FilecoinService.Contract.Owner(&_FilecoinService.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 updatedAtTimestamp, bytes32 merkleRoot, uint256 epoch)
func (_FilecoinService *FilecoinServiceCaller) State(opts *bind.CallOpts) (struct {
	UpdatedAtTimestamp *big.Int
	MerkleRoot         [32]byte
	Epoch              *big.Int
}, error) {
	var out []interface{}
	err := _FilecoinService.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		UpdatedAtTimestamp *big.Int
		MerkleRoot         [32]byte
		Epoch              *big.Int
	})

	outstruct.UpdatedAtTimestamp = out[0].(*big.Int)
	outstruct.MerkleRoot = out[1].([32]byte)
	outstruct.Epoch = out[2].(*big.Int)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 updatedAtTimestamp, bytes32 merkleRoot, uint256 epoch)
func (_FilecoinService *FilecoinServiceSession) State() (struct {
	UpdatedAtTimestamp *big.Int
	MerkleRoot         [32]byte
	Epoch              *big.Int
}, error) {
	return _FilecoinService.Contract.State(&_FilecoinService.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 updatedAtTimestamp, bytes32 merkleRoot, uint256 epoch)
func (_FilecoinService *FilecoinServiceCallerSession) State() (struct {
	UpdatedAtTimestamp *big.Int
	MerkleRoot         [32]byte
	Epoch              *big.Int
}, error) {
	return _FilecoinService.Contract.State(&_FilecoinService.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8d834449.
//
// Solidity: function verifyProof(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch, bytes32[] merkleProof) view returns(bool)
func (_FilecoinService *FilecoinServiceCaller) VerifyProof(opts *bind.CallOpts, dataCid string, pieceCid string, dealId *big.Int, provider string, startEpoch *big.Int, endEpoch *big.Int, signedEpoch *big.Int, merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _FilecoinService.contract.Call(opts, &out, "verifyProof", dataCid, pieceCid, dealId, provider, startEpoch, endEpoch, signedEpoch, merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x8d834449.
//
// Solidity: function verifyProof(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch, bytes32[] merkleProof) view returns(bool)
func (_FilecoinService *FilecoinServiceSession) VerifyProof(dataCid string, pieceCid string, dealId *big.Int, provider string, startEpoch *big.Int, endEpoch *big.Int, signedEpoch *big.Int, merkleProof [][32]byte) (bool, error) {
	return _FilecoinService.Contract.VerifyProof(&_FilecoinService.CallOpts, dataCid, pieceCid, dealId, provider, startEpoch, endEpoch, signedEpoch, merkleProof)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8d834449.
//
// Solidity: function verifyProof(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch, bytes32[] merkleProof) view returns(bool)
func (_FilecoinService *FilecoinServiceCallerSession) VerifyProof(dataCid string, pieceCid string, dealId *big.Int, provider string, startEpoch *big.Int, endEpoch *big.Int, signedEpoch *big.Int, merkleProof [][32]byte) (bool, error) {
	return _FilecoinService.Contract.VerifyProof(&_FilecoinService.CallOpts, dataCid, pieceCid, dealId, provider, startEpoch, endEpoch, signedEpoch, merkleProof)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_FilecoinService *FilecoinServiceTransactor) AddManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _FilecoinService.contract.Transact(opts, "addManager", _manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_FilecoinService *FilecoinServiceSession) AddManager(_manager common.Address) (*types.Transaction, error) {
	return _FilecoinService.Contract.AddManager(&_FilecoinService.TransactOpts, _manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_FilecoinService *FilecoinServiceTransactorSession) AddManager(_manager common.Address) (*types.Transaction, error) {
	return _FilecoinService.Contract.AddManager(&_FilecoinService.TransactOpts, _manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address _manager) returns()
func (_FilecoinService *FilecoinServiceTransactor) RemoveManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _FilecoinService.contract.Transact(opts, "removeManager", _manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address _manager) returns()
func (_FilecoinService *FilecoinServiceSession) RemoveManager(_manager common.Address) (*types.Transaction, error) {
	return _FilecoinService.Contract.RemoveManager(&_FilecoinService.TransactOpts, _manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address _manager) returns()
func (_FilecoinService *FilecoinServiceTransactorSession) RemoveManager(_manager common.Address) (*types.Transaction, error) {
	return _FilecoinService.Contract.RemoveManager(&_FilecoinService.TransactOpts, _manager)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x01930e36.
//
// Solidity: function submitProof(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch, bytes32[] merkleProof) returns()
func (_FilecoinService *FilecoinServiceTransactor) SubmitProof(opts *bind.TransactOpts, dataCid string, pieceCid string, dealId *big.Int, provider string, startEpoch *big.Int, endEpoch *big.Int, signedEpoch *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _FilecoinService.contract.Transact(opts, "submitProof", dataCid, pieceCid, dealId, provider, startEpoch, endEpoch, signedEpoch, merkleProof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x01930e36.
//
// Solidity: function submitProof(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch, bytes32[] merkleProof) returns()
func (_FilecoinService *FilecoinServiceSession) SubmitProof(dataCid string, pieceCid string, dealId *big.Int, provider string, startEpoch *big.Int, endEpoch *big.Int, signedEpoch *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _FilecoinService.Contract.SubmitProof(&_FilecoinService.TransactOpts, dataCid, pieceCid, dealId, provider, startEpoch, endEpoch, signedEpoch, merkleProof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x01930e36.
//
// Solidity: function submitProof(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch, bytes32[] merkleProof) returns()
func (_FilecoinService *FilecoinServiceTransactorSession) SubmitProof(dataCid string, pieceCid string, dealId *big.Int, provider string, startEpoch *big.Int, endEpoch *big.Int, signedEpoch *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _FilecoinService.Contract.SubmitProof(&_FilecoinService.TransactOpts, dataCid, pieceCid, dealId, provider, startEpoch, endEpoch, signedEpoch, merkleProof)
}

// UpdateState is a paid mutator transaction binding the contract method 0x4c139b55.
//
// Solidity: function updateState(bytes32 _merkleRoot, uint256 _epoch) returns()
func (_FilecoinService *FilecoinServiceTransactor) UpdateState(opts *bind.TransactOpts, _merkleRoot [32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _FilecoinService.contract.Transact(opts, "updateState", _merkleRoot, _epoch)
}

// UpdateState is a paid mutator transaction binding the contract method 0x4c139b55.
//
// Solidity: function updateState(bytes32 _merkleRoot, uint256 _epoch) returns()
func (_FilecoinService *FilecoinServiceSession) UpdateState(_merkleRoot [32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _FilecoinService.Contract.UpdateState(&_FilecoinService.TransactOpts, _merkleRoot, _epoch)
}

// UpdateState is a paid mutator transaction binding the contract method 0x4c139b55.
//
// Solidity: function updateState(bytes32 _merkleRoot, uint256 _epoch) returns()
func (_FilecoinService *FilecoinServiceTransactorSession) UpdateState(_merkleRoot [32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _FilecoinService.Contract.UpdateState(&_FilecoinService.TransactOpts, _merkleRoot, _epoch)
}

// FilecoinServiceStoredCidIterator is returned from FilterStoredCid and is used to iterate over the raw logs and unpacked data for StoredCid events raised by the FilecoinService contract.
type FilecoinServiceStoredCidIterator struct {
	Event *FilecoinServiceStoredCid // Event containing the contract specifics and raw log

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
func (it *FilecoinServiceStoredCidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilecoinServiceStoredCid)
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
		it.Event = new(FilecoinServiceStoredCid)
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
func (it *FilecoinServiceStoredCidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilecoinServiceStoredCidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilecoinServiceStoredCid represents a StoredCid event raised by the FilecoinService contract.
type FilecoinServiceStoredCid struct {
	DataCid     string
	PieceCid    string
	DealId      *big.Int
	Provider    string
	StartEpoch  *big.Int
	EndEpoch    *big.Int
	SignedEpoch *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStoredCid is a free log retrieval operation binding the contract event 0xbd510a4e8d7985f3ff7bd53f296135db76010e4f86e77264c9896d3aa3d055e8.
//
// Solidity: event StoredCid(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch)
func (_FilecoinService *FilecoinServiceFilterer) FilterStoredCid(opts *bind.FilterOpts) (*FilecoinServiceStoredCidIterator, error) {

	logs, sub, err := _FilecoinService.contract.FilterLogs(opts, "StoredCid")
	if err != nil {
		return nil, err
	}
	return &FilecoinServiceStoredCidIterator{contract: _FilecoinService.contract, event: "StoredCid", logs: logs, sub: sub}, nil
}

// WatchStoredCid is a free log subscription operation binding the contract event 0xbd510a4e8d7985f3ff7bd53f296135db76010e4f86e77264c9896d3aa3d055e8.
//
// Solidity: event StoredCid(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch)
func (_FilecoinService *FilecoinServiceFilterer) WatchStoredCid(opts *bind.WatchOpts, sink chan<- *FilecoinServiceStoredCid) (event.Subscription, error) {

	logs, sub, err := _FilecoinService.contract.WatchLogs(opts, "StoredCid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilecoinServiceStoredCid)
				if err := _FilecoinService.contract.UnpackLog(event, "StoredCid", log); err != nil {
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

// ParseStoredCid is a log parse operation binding the contract event 0xbd510a4e8d7985f3ff7bd53f296135db76010e4f86e77264c9896d3aa3d055e8.
//
// Solidity: event StoredCid(string dataCid, string pieceCid, uint256 dealId, string provider, uint256 startEpoch, uint256 endEpoch, uint256 signedEpoch)
func (_FilecoinService *FilecoinServiceFilterer) ParseStoredCid(log types.Log) (*FilecoinServiceStoredCid, error) {
	event := new(FilecoinServiceStoredCid)
	if err := _FilecoinService.contract.UnpackLog(event, "StoredCid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
