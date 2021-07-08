// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_event

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

// ContractEventABI is the input ABI used to generate the binding from.
const ContractEventABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"AddOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"CancelOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"FinishOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"typa\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"GameOver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"typa\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SysFeeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"endRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintToSelf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"startGame\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"typa\",\"type\":\"uint256\"}],\"name\":\"vote1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"typa\",\"type\":\"uint256\"}],\"name\":\"vote2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ContractEvent is an auto generated Go binding around an Ethereum contract.
type ContractEvent struct {
	ContractEventCaller     // Read-only binding to the contract
	ContractEventTransactor // Write-only binding to the contract
	ContractEventFilterer   // Log filterer for contract events
}

// ContractEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractEventSession struct {
	Contract     *ContractEvent    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractEventCallerSession struct {
	Contract *ContractEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ContractEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractEventTransactorSession struct {
	Contract     *ContractEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractEventRaw struct {
	Contract *ContractEvent // Generic contract binding to access the raw methods on
}

// ContractEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractEventCallerRaw struct {
	Contract *ContractEventCaller // Generic read-only contract binding to access the raw methods on
}

// ContractEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractEventTransactorRaw struct {
	Contract *ContractEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractEvent creates a new instance of ContractEvent, bound to a specific deployed contract.
func NewContractEvent(address common.Address, backend bind.ContractBackend) (*ContractEvent, error) {
	contract, err := bindContractEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractEvent{ContractEventCaller: ContractEventCaller{contract: contract}, ContractEventTransactor: ContractEventTransactor{contract: contract}, ContractEventFilterer: ContractEventFilterer{contract: contract}}, nil
}

// NewContractEventCaller creates a new read-only instance of ContractEvent, bound to a specific deployed contract.
func NewContractEventCaller(address common.Address, caller bind.ContractCaller) (*ContractEventCaller, error) {
	contract, err := bindContractEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractEventCaller{contract: contract}, nil
}

// NewContractEventTransactor creates a new write-only instance of ContractEvent, bound to a specific deployed contract.
func NewContractEventTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractEventTransactor, error) {
	contract, err := bindContractEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractEventTransactor{contract: contract}, nil
}

// NewContractEventFilterer creates a new log filterer instance of ContractEvent, bound to a specific deployed contract.
func NewContractEventFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractEventFilterer, error) {
	contract, err := bindContractEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractEventFilterer{contract: contract}, nil
}

// bindContractEvent binds a generic wrapper to an already deployed contract.
func bindContractEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractEventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractEvent *ContractEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractEvent.Contract.ContractEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractEvent *ContractEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEvent.Contract.ContractEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractEvent *ContractEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractEvent.Contract.ContractEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractEvent *ContractEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractEvent *ContractEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractEvent *ContractEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractEvent.Contract.contract.Transact(opts, method, params...)
}

// SysFeeAccount is a free data retrieval call binding the contract method 0xfe7419b6.
//
// Solidity: function SysFeeAccount() view returns(address)
func (_ContractEvent *ContractEventCaller) SysFeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "SysFeeAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SysFeeAccount is a free data retrieval call binding the contract method 0xfe7419b6.
//
// Solidity: function SysFeeAccount() view returns(address)
func (_ContractEvent *ContractEventSession) SysFeeAccount() (common.Address, error) {
	return _ContractEvent.Contract.SysFeeAccount(&_ContractEvent.CallOpts)
}

// SysFeeAccount is a free data retrieval call binding the contract method 0xfe7419b6.
//
// Solidity: function SysFeeAccount() view returns(address)
func (_ContractEvent *ContractEventCallerSession) SysFeeAccount() (common.Address, error) {
	return _ContractEvent.Contract.SysFeeAccount(&_ContractEvent.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ContractEvent *ContractEventCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ContractEvent *ContractEventSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ContractEvent.Contract.BalanceOf(&_ContractEvent.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ContractEvent *ContractEventCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ContractEvent.Contract.BalanceOf(&_ContractEvent.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ContractEvent *ContractEventCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ContractEvent *ContractEventSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ContractEvent.Contract.GetApproved(&_ContractEvent.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ContractEvent *ContractEventCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ContractEvent.Contract.GetApproved(&_ContractEvent.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ContractEvent *ContractEventCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ContractEvent *ContractEventSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ContractEvent.Contract.IsApprovedForAll(&_ContractEvent.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ContractEvent *ContractEventCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ContractEvent.Contract.IsApprovedForAll(&_ContractEvent.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractEvent *ContractEventCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractEvent *ContractEventSession) Name() (string, error) {
	return _ContractEvent.Contract.Name(&_ContractEvent.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractEvent *ContractEventCallerSession) Name() (string, error) {
	return _ContractEvent.Contract.Name(&_ContractEvent.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, uint256 price, uint256 tokenId, uint256 createTime)
func (_ContractEvent *ContractEventCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller     common.Address
	Price      *big.Int
	TokenId    *big.Int
	CreateTime *big.Int
}, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Seller     common.Address
		Price      *big.Int
		TokenId    *big.Int
		CreateTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CreateTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, uint256 price, uint256 tokenId, uint256 createTime)
func (_ContractEvent *ContractEventSession) Orders(arg0 *big.Int) (struct {
	Seller     common.Address
	Price      *big.Int
	TokenId    *big.Int
	CreateTime *big.Int
}, error) {
	return _ContractEvent.Contract.Orders(&_ContractEvent.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, uint256 price, uint256 tokenId, uint256 createTime)
func (_ContractEvent *ContractEventCallerSession) Orders(arg0 *big.Int) (struct {
	Seller     common.Address
	Price      *big.Int
	TokenId    *big.Int
	CreateTime *big.Int
}, error) {
	return _ContractEvent.Contract.Orders(&_ContractEvent.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ContractEvent *ContractEventCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ContractEvent *ContractEventSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ContractEvent.Contract.OwnerOf(&_ContractEvent.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ContractEvent *ContractEventCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ContractEvent.Contract.OwnerOf(&_ContractEvent.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractEvent *ContractEventCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractEvent *ContractEventSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractEvent.Contract.SupportsInterface(&_ContractEvent.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractEvent *ContractEventCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractEvent.Contract.SupportsInterface(&_ContractEvent.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractEvent *ContractEventCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractEvent *ContractEventSession) Symbol() (string, error) {
	return _ContractEvent.Contract.Symbol(&_ContractEvent.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractEvent *ContractEventCallerSession) Symbol() (string, error) {
	return _ContractEvent.Contract.Symbol(&_ContractEvent.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ContractEvent *ContractEventCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ContractEvent *ContractEventSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ContractEvent.Contract.TokenURI(&_ContractEvent.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ContractEvent *ContractEventCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ContractEvent.Contract.TokenURI(&_ContractEvent.CallOpts, tokenId)
}

// VoteAmount is a free data retrieval call binding the contract method 0xef58c784.
//
// Solidity: function voteAmount() view returns(uint256)
func (_ContractEvent *ContractEventCaller) VoteAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractEvent.contract.Call(opts, &out, "voteAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteAmount is a free data retrieval call binding the contract method 0xef58c784.
//
// Solidity: function voteAmount() view returns(uint256)
func (_ContractEvent *ContractEventSession) VoteAmount() (*big.Int, error) {
	return _ContractEvent.Contract.VoteAmount(&_ContractEvent.CallOpts)
}

// VoteAmount is a free data retrieval call binding the contract method 0xef58c784.
//
// Solidity: function voteAmount() view returns(uint256)
func (_ContractEvent *ContractEventCallerSession) VoteAmount() (*big.Int, error) {
	return _ContractEvent.Contract.VoteAmount(&_ContractEvent.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Approve(&_ContractEvent.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Approve(&_ContractEvent.TransactOpts, to, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns(bool)
func (_ContractEvent *ContractEventTransactor) Buy(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "buy", tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns(bool)
func (_ContractEvent *ContractEventSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Buy(&_ContractEvent.TransactOpts, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns(bool)
func (_ContractEvent *ContractEventTransactorSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Buy(&_ContractEvent.TransactOpts, tokenId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactor) Cancel(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "cancel", tokenId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 tokenId) returns()
func (_ContractEvent *ContractEventSession) Cancel(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Cancel(&_ContractEvent.TransactOpts, tokenId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactorSession) Cancel(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Cancel(&_ContractEvent.TransactOpts, tokenId)
}

// EndRound is a paid mutator transaction binding the contract method 0x5a1e6ca1.
//
// Solidity: function endRound(uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactor) EndRound(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "endRound", tokenId)
}

// EndRound is a paid mutator transaction binding the contract method 0x5a1e6ca1.
//
// Solidity: function endRound(uint256 tokenId) returns()
func (_ContractEvent *ContractEventSession) EndRound(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.EndRound(&_ContractEvent.TransactOpts, tokenId)
}

// EndRound is a paid mutator transaction binding the contract method 0x5a1e6ca1.
//
// Solidity: function endRound(uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactorSession) EndRound(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.EndRound(&_ContractEvent.TransactOpts, tokenId)
}

// MintToSelf is a paid mutator transaction binding the contract method 0x5201248e.
//
// Solidity: function mintToSelf(string tokenURI) payable returns(uint256)
func (_ContractEvent *ContractEventTransactor) MintToSelf(opts *bind.TransactOpts, tokenURI string) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "mintToSelf", tokenURI)
}

// MintToSelf is a paid mutator transaction binding the contract method 0x5201248e.
//
// Solidity: function mintToSelf(string tokenURI) payable returns(uint256)
func (_ContractEvent *ContractEventSession) MintToSelf(tokenURI string) (*types.Transaction, error) {
	return _ContractEvent.Contract.MintToSelf(&_ContractEvent.TransactOpts, tokenURI)
}

// MintToSelf is a paid mutator transaction binding the contract method 0x5201248e.
//
// Solidity: function mintToSelf(string tokenURI) payable returns(uint256)
func (_ContractEvent *ContractEventTransactorSession) MintToSelf(tokenURI string) (*types.Transaction, error) {
	return _ContractEvent.Contract.MintToSelf(&_ContractEvent.TransactOpts, tokenURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.SafeTransferFrom(&_ContractEvent.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.SafeTransferFrom(&_ContractEvent.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ContractEvent *ContractEventTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ContractEvent *ContractEventSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractEvent.Contract.SafeTransferFrom0(&_ContractEvent.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ContractEvent *ContractEventTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ContractEvent.Contract.SafeTransferFrom0(&_ContractEvent.TransactOpts, from, to, tokenId, _data)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 tokenId, uint256 price) returns(bool)
func (_ContractEvent *ContractEventTransactor) Sell(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "sell", tokenId, price)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 tokenId, uint256 price) returns(bool)
func (_ContractEvent *ContractEventSession) Sell(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Sell(&_ContractEvent.TransactOpts, tokenId, price)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 tokenId, uint256 price) returns(bool)
func (_ContractEvent *ContractEventTransactorSession) Sell(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Sell(&_ContractEvent.TransactOpts, tokenId, price)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractEvent *ContractEventTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractEvent *ContractEventSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractEvent.Contract.SetApprovalForAll(&_ContractEvent.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractEvent *ContractEventTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractEvent.Contract.SetApprovalForAll(&_ContractEvent.TransactOpts, operator, approved)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 tokenId) returns(bool)
func (_ContractEvent *ContractEventTransactor) StartGame(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "startGame", tokenId)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 tokenId) returns(bool)
func (_ContractEvent *ContractEventSession) StartGame(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.StartGame(&_ContractEvent.TransactOpts, tokenId)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 tokenId) returns(bool)
func (_ContractEvent *ContractEventTransactorSession) StartGame(tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.StartGame(&_ContractEvent.TransactOpts, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.TransferFrom(&_ContractEvent.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ContractEvent *ContractEventTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.TransferFrom(&_ContractEvent.TransactOpts, from, to, tokenId)
}

// Vote1 is a paid mutator transaction binding the contract method 0x1f38a52f.
//
// Solidity: function vote1(uint256 tokenId, uint256 typa) payable returns(bool success)
func (_ContractEvent *ContractEventTransactor) Vote1(opts *bind.TransactOpts, tokenId *big.Int, typa *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "vote1", tokenId, typa)
}

// Vote1 is a paid mutator transaction binding the contract method 0x1f38a52f.
//
// Solidity: function vote1(uint256 tokenId, uint256 typa) payable returns(bool success)
func (_ContractEvent *ContractEventSession) Vote1(tokenId *big.Int, typa *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Vote1(&_ContractEvent.TransactOpts, tokenId, typa)
}

// Vote1 is a paid mutator transaction binding the contract method 0x1f38a52f.
//
// Solidity: function vote1(uint256 tokenId, uint256 typa) payable returns(bool success)
func (_ContractEvent *ContractEventTransactorSession) Vote1(tokenId *big.Int, typa *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Vote1(&_ContractEvent.TransactOpts, tokenId, typa)
}

// Vote2 is a paid mutator transaction binding the contract method 0x8b8e626a.
//
// Solidity: function vote2(uint256 tokenId, uint256 typa) returns(bool success)
func (_ContractEvent *ContractEventTransactor) Vote2(opts *bind.TransactOpts, tokenId *big.Int, typa *big.Int) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "vote2", tokenId, typa)
}

// Vote2 is a paid mutator transaction binding the contract method 0x8b8e626a.
//
// Solidity: function vote2(uint256 tokenId, uint256 typa) returns(bool success)
func (_ContractEvent *ContractEventSession) Vote2(tokenId *big.Int, typa *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Vote2(&_ContractEvent.TransactOpts, tokenId, typa)
}

// Vote2 is a paid mutator transaction binding the contract method 0x8b8e626a.
//
// Solidity: function vote2(uint256 tokenId, uint256 typa) returns(bool success)
func (_ContractEvent *ContractEventTransactorSession) Vote2(tokenId *big.Int, typa *big.Int) (*types.Transaction, error) {
	return _ContractEvent.Contract.Vote2(&_ContractEvent.TransactOpts, tokenId, typa)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractEvent *ContractEventTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEvent.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractEvent *ContractEventSession) Withdraw() (*types.Transaction, error) {
	return _ContractEvent.Contract.Withdraw(&_ContractEvent.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractEvent *ContractEventTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ContractEvent.Contract.Withdraw(&_ContractEvent.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractEvent *ContractEventTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEvent.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractEvent *ContractEventSession) Receive() (*types.Transaction, error) {
	return _ContractEvent.Contract.Receive(&_ContractEvent.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractEvent *ContractEventTransactorSession) Receive() (*types.Transaction, error) {
	return _ContractEvent.Contract.Receive(&_ContractEvent.TransactOpts)
}

// ContractEventAddOrderIterator is returned from FilterAddOrder and is used to iterate over the raw logs and unpacked data for AddOrder events raised by the ContractEvent contract.
type ContractEventAddOrderIterator struct {
	Event *ContractEventAddOrder // Event containing the contract specifics and raw log

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
func (it *ContractEventAddOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventAddOrder)
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
		it.Event = new(ContractEventAddOrder)
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
func (it *ContractEventAddOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventAddOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventAddOrder represents a AddOrder event raised by the ContractEvent contract.
type ContractEventAddOrder struct {
	Seller  common.Address
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddOrder is a free log retrieval operation binding the contract event 0xc9ebdee0f95a19e039d4896abd76fcd8601916a080e6f74e848d597cf0de647a.
//
// Solidity: event AddOrder(address indexed seller, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) FilterAddOrder(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int, time []*big.Int) (*ContractEventAddOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "AddOrder", sellerRule, tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventAddOrderIterator{contract: _ContractEvent.contract, event: "AddOrder", logs: logs, sub: sub}, nil
}

// WatchAddOrder is a free log subscription operation binding the contract event 0xc9ebdee0f95a19e039d4896abd76fcd8601916a080e6f74e848d597cf0de647a.
//
// Solidity: event AddOrder(address indexed seller, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) WatchAddOrder(opts *bind.WatchOpts, sink chan<- *ContractEventAddOrder, seller []common.Address, tokenId []*big.Int, time []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "AddOrder", sellerRule, tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventAddOrder)
				if err := _ContractEvent.contract.UnpackLog(event, "AddOrder", log); err != nil {
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

// ParseAddOrder is a log parse operation binding the contract event 0xc9ebdee0f95a19e039d4896abd76fcd8601916a080e6f74e848d597cf0de647a.
//
// Solidity: event AddOrder(address indexed seller, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) ParseAddOrder(log types.Log) (*ContractEventAddOrder, error) {
	event := new(ContractEventAddOrder)
	if err := _ContractEvent.contract.UnpackLog(event, "AddOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ContractEvent contract.
type ContractEventApprovalIterator struct {
	Event *ContractEventApproval // Event containing the contract specifics and raw log

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
func (it *ContractEventApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventApproval)
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
		it.Event = new(ContractEventApproval)
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
func (it *ContractEventApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventApproval represents a Approval event raised by the ContractEvent contract.
type ContractEventApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ContractEvent *ContractEventFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractEventApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventApprovalIterator{contract: _ContractEvent.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ContractEvent *ContractEventFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractEventApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventApproval)
				if err := _ContractEvent.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ContractEvent *ContractEventFilterer) ParseApproval(log types.Log) (*ContractEventApproval, error) {
	event := new(ContractEventApproval)
	if err := _ContractEvent.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ContractEvent contract.
type ContractEventApprovalForAllIterator struct {
	Event *ContractEventApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractEventApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventApprovalForAll)
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
		it.Event = new(ContractEventApprovalForAll)
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
func (it *ContractEventApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventApprovalForAll represents a ApprovalForAll event raised by the ContractEvent contract.
type ContractEventApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ContractEvent *ContractEventFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractEventApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventApprovalForAllIterator{contract: _ContractEvent.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ContractEvent *ContractEventFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractEventApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventApprovalForAll)
				if err := _ContractEvent.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ContractEvent *ContractEventFilterer) ParseApprovalForAll(log types.Log) (*ContractEventApprovalForAll, error) {
	event := new(ContractEventApprovalForAll)
	if err := _ContractEvent.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventCancelOrderIterator is returned from FilterCancelOrder and is used to iterate over the raw logs and unpacked data for CancelOrder events raised by the ContractEvent contract.
type ContractEventCancelOrderIterator struct {
	Event *ContractEventCancelOrder // Event containing the contract specifics and raw log

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
func (it *ContractEventCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventCancelOrder)
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
		it.Event = new(ContractEventCancelOrder)
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
func (it *ContractEventCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventCancelOrder represents a CancelOrder event raised by the ContractEvent contract.
type ContractEventCancelOrder struct {
	Seller  common.Address
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelOrder is a free log retrieval operation binding the contract event 0x1e0e8a7901ee871b939f18e3937f188af21bf9885762ec9c84917abc9fbaf1fa.
//
// Solidity: event CancelOrder(address indexed seller, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) FilterCancelOrder(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int, time []*big.Int) (*ContractEventCancelOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "CancelOrder", sellerRule, tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventCancelOrderIterator{contract: _ContractEvent.contract, event: "CancelOrder", logs: logs, sub: sub}, nil
}

// WatchCancelOrder is a free log subscription operation binding the contract event 0x1e0e8a7901ee871b939f18e3937f188af21bf9885762ec9c84917abc9fbaf1fa.
//
// Solidity: event CancelOrder(address indexed seller, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) WatchCancelOrder(opts *bind.WatchOpts, sink chan<- *ContractEventCancelOrder, seller []common.Address, tokenId []*big.Int, time []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "CancelOrder", sellerRule, tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventCancelOrder)
				if err := _ContractEvent.contract.UnpackLog(event, "CancelOrder", log); err != nil {
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

// ParseCancelOrder is a log parse operation binding the contract event 0x1e0e8a7901ee871b939f18e3937f188af21bf9885762ec9c84917abc9fbaf1fa.
//
// Solidity: event CancelOrder(address indexed seller, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) ParseCancelOrder(log types.Log) (*ContractEventCancelOrder, error) {
	event := new(ContractEventCancelOrder)
	if err := _ContractEvent.contract.UnpackLog(event, "CancelOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventFinishOrderIterator is returned from FilterFinishOrder and is used to iterate over the raw logs and unpacked data for FinishOrder events raised by the ContractEvent contract.
type ContractEventFinishOrderIterator struct {
	Event *ContractEventFinishOrder // Event containing the contract specifics and raw log

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
func (it *ContractEventFinishOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventFinishOrder)
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
		it.Event = new(ContractEventFinishOrder)
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
func (it *ContractEventFinishOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventFinishOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventFinishOrder represents a FinishOrder event raised by the ContractEvent contract.
type ContractEventFinishOrder struct {
	Buyer   common.Address
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinishOrder is a free log retrieval operation binding the contract event 0x382dce283af4282f2ee4126bd1d59671390fcab2477542ed6bdd175676daa2f8.
//
// Solidity: event FinishOrder(address indexed buyer, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) FilterFinishOrder(opts *bind.FilterOpts, buyer []common.Address, tokenId []*big.Int, time []*big.Int) (*ContractEventFinishOrderIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "FinishOrder", buyerRule, tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventFinishOrderIterator{contract: _ContractEvent.contract, event: "FinishOrder", logs: logs, sub: sub}, nil
}

// WatchFinishOrder is a free log subscription operation binding the contract event 0x382dce283af4282f2ee4126bd1d59671390fcab2477542ed6bdd175676daa2f8.
//
// Solidity: event FinishOrder(address indexed buyer, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) WatchFinishOrder(opts *bind.WatchOpts, sink chan<- *ContractEventFinishOrder, buyer []common.Address, tokenId []*big.Int, time []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeRule []interface{}
	for _, timeItem := range time {
		timeRule = append(timeRule, timeItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "FinishOrder", buyerRule, tokenIdRule, timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventFinishOrder)
				if err := _ContractEvent.contract.UnpackLog(event, "FinishOrder", log); err != nil {
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

// ParseFinishOrder is a log parse operation binding the contract event 0x382dce283af4282f2ee4126bd1d59671390fcab2477542ed6bdd175676daa2f8.
//
// Solidity: event FinishOrder(address indexed buyer, uint256 indexed tokenId, uint256 indexed time)
func (_ContractEvent *ContractEventFilterer) ParseFinishOrder(log types.Log) (*ContractEventFinishOrder, error) {
	event := new(ContractEventFinishOrder)
	if err := _ContractEvent.contract.UnpackLog(event, "FinishOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventGameOverIterator is returned from FilterGameOver and is used to iterate over the raw logs and unpacked data for GameOver events raised by the ContractEvent contract.
type ContractEventGameOverIterator struct {
	Event *ContractEventGameOver // Event containing the contract specifics and raw log

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
func (it *ContractEventGameOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventGameOver)
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
		it.Event = new(ContractEventGameOver)
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
func (it *ContractEventGameOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventGameOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventGameOver represents a GameOver event raised by the ContractEvent contract.
type ContractEventGameOver struct {
	TokenId *big.Int
	Typa    *big.Int
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGameOver is a free log retrieval operation binding the contract event 0x2f4c1c334141992a9407d0a9c3d93283e5c5a1045746eece29bd1793dd859619.
//
// Solidity: event GameOver(uint256 indexed tokenId, uint256 indexed typa, uint256 indexed reward)
func (_ContractEvent *ContractEventFilterer) FilterGameOver(opts *bind.FilterOpts, tokenId []*big.Int, typa []*big.Int, reward []*big.Int) (*ContractEventGameOverIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var typaRule []interface{}
	for _, typaItem := range typa {
		typaRule = append(typaRule, typaItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "GameOver", tokenIdRule, typaRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventGameOverIterator{contract: _ContractEvent.contract, event: "GameOver", logs: logs, sub: sub}, nil
}

// WatchGameOver is a free log subscription operation binding the contract event 0x2f4c1c334141992a9407d0a9c3d93283e5c5a1045746eece29bd1793dd859619.
//
// Solidity: event GameOver(uint256 indexed tokenId, uint256 indexed typa, uint256 indexed reward)
func (_ContractEvent *ContractEventFilterer) WatchGameOver(opts *bind.WatchOpts, sink chan<- *ContractEventGameOver, tokenId []*big.Int, typa []*big.Int, reward []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var typaRule []interface{}
	for _, typaItem := range typa {
		typaRule = append(typaRule, typaItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "GameOver", tokenIdRule, typaRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventGameOver)
				if err := _ContractEvent.contract.UnpackLog(event, "GameOver", log); err != nil {
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

// ParseGameOver is a log parse operation binding the contract event 0x2f4c1c334141992a9407d0a9c3d93283e5c5a1045746eece29bd1793dd859619.
//
// Solidity: event GameOver(uint256 indexed tokenId, uint256 indexed typa, uint256 indexed reward)
func (_ContractEvent *ContractEventFilterer) ParseGameOver(log types.Log) (*ContractEventGameOver, error) {
	event := new(ContractEventGameOver)
	if err := _ContractEvent.contract.UnpackLog(event, "GameOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ContractEvent contract.
type ContractEventTransferIterator struct {
	Event *ContractEventTransfer // Event containing the contract specifics and raw log

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
func (it *ContractEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventTransfer)
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
		it.Event = new(ContractEventTransfer)
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
func (it *ContractEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventTransfer represents a Transfer event raised by the ContractEvent contract.
type ContractEventTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ContractEvent *ContractEventFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractEventTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventTransferIterator{contract: _ContractEvent.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ContractEvent *ContractEventFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractEventTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventTransfer)
				if err := _ContractEvent.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ContractEvent *ContractEventFilterer) ParseTransfer(log types.Log) (*ContractEventTransfer, error) {
	event := new(ContractEventTransfer)
	if err := _ContractEvent.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the ContractEvent contract.
type ContractEventVoteIterator struct {
	Event *ContractEventVote // Event containing the contract specifics and raw log

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
func (it *ContractEventVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventVote)
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
		it.Event = new(ContractEventVote)
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
func (it *ContractEventVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventVote represents a Vote event raised by the ContractEvent contract.
type ContractEventVote struct {
	Voter   common.Address
	TokenId *big.Int
	Typa    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0xafd3f234c1f8e944129b26b206d98e5752ad3336a4059938b4a3e990e9588530.
//
// Solidity: event Vote(address indexed voter, uint256 indexed tokenId, uint256 indexed typa)
func (_ContractEvent *ContractEventFilterer) FilterVote(opts *bind.FilterOpts, voter []common.Address, tokenId []*big.Int, typa []*big.Int) (*ContractEventVoteIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var typaRule []interface{}
	for _, typaItem := range typa {
		typaRule = append(typaRule, typaItem)
	}

	logs, sub, err := _ContractEvent.contract.FilterLogs(opts, "Vote", voterRule, tokenIdRule, typaRule)
	if err != nil {
		return nil, err
	}
	return &ContractEventVoteIterator{contract: _ContractEvent.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0xafd3f234c1f8e944129b26b206d98e5752ad3336a4059938b4a3e990e9588530.
//
// Solidity: event Vote(address indexed voter, uint256 indexed tokenId, uint256 indexed typa)
func (_ContractEvent *ContractEventFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *ContractEventVote, voter []common.Address, tokenId []*big.Int, typa []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var typaRule []interface{}
	for _, typaItem := range typa {
		typaRule = append(typaRule, typaItem)
	}

	logs, sub, err := _ContractEvent.contract.WatchLogs(opts, "Vote", voterRule, tokenIdRule, typaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventVote)
				if err := _ContractEvent.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0xafd3f234c1f8e944129b26b206d98e5752ad3336a4059938b4a3e990e9588530.
//
// Solidity: event Vote(address indexed voter, uint256 indexed tokenId, uint256 indexed typa)
func (_ContractEvent *ContractEventFilterer) ParseVote(log types.Log) (*ContractEventVote, error) {
	event := new(ContractEventVote)
	if err := _ContractEvent.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
