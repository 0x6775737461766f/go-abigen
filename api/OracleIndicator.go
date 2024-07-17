// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OracleIndicatorDataFeed is an auto generated low-level Go binding around an user-defined struct.
type OracleIndicatorDataFeed struct {
	Value      *big.Int
	Updatedat  *big.Int
	Decimal    uint8
	Confidence uint8
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"READ_ONLY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDate\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"updatedat\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"confidence\",\"type\":\"uint8\"}],\"internalType\":\"structOracleIndicator.DataFeed\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getInterval\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLast\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"updatedat\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"confidence\",\"type\":\"uint8\"}],\"internalType\":\"structOracleIndicator.DataFeed\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indicators\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"updatedat\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"confidence\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_updatedat\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_confidence\",\"type\":\"uint8\"}],\"name\":\"saveIndicator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Api *ApiCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Api *ApiSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Api.Contract.DEFAULTADMINROLE(&_Api.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Api.Contract.DEFAULTADMINROLE(&_Api.CallOpts)
}

// READONLY is a free data retrieval call binding the contract method 0x15eecf21.
//
// Solidity: function READ_ONLY() view returns(bytes32)
func (_Api *ApiCaller) READONLY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "READ_ONLY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// READONLY is a free data retrieval call binding the contract method 0x15eecf21.
//
// Solidity: function READ_ONLY() view returns(bytes32)
func (_Api *ApiSession) READONLY() ([32]byte, error) {
	return _Api.Contract.READONLY(&_Api.CallOpts)
}

// READONLY is a free data retrieval call binding the contract method 0x15eecf21.
//
// Solidity: function READ_ONLY() view returns(bytes32)
func (_Api *ApiCallerSession) READONLY() ([32]byte, error) {
	return _Api.Contract.READONLY(&_Api.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_Api *ApiCaller) Decimal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "decimal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_Api *ApiSession) Decimal() (uint8, error) {
	return _Api.Contract.Decimal(&_Api.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_Api *ApiCallerSession) Decimal() (uint8, error) {
	return _Api.Contract.Decimal(&_Api.CallOpts)
}

// GetDate is a free data retrieval call binding the contract method 0x2b57298b.
//
// Solidity: function getDate(uint256 _timestamp) view returns((int256,uint256,uint8,uint8))
func (_Api *ApiCaller) GetDate(opts *bind.CallOpts, _timestamp *big.Int) (OracleIndicatorDataFeed, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getDate", _timestamp)

	if err != nil {
		return *new(OracleIndicatorDataFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleIndicatorDataFeed)).(*OracleIndicatorDataFeed)

	return out0, err

}

// GetDate is a free data retrieval call binding the contract method 0x2b57298b.
//
// Solidity: function getDate(uint256 _timestamp) view returns((int256,uint256,uint8,uint8))
func (_Api *ApiSession) GetDate(_timestamp *big.Int) (OracleIndicatorDataFeed, error) {
	return _Api.Contract.GetDate(&_Api.CallOpts, _timestamp)
}

// GetDate is a free data retrieval call binding the contract method 0x2b57298b.
//
// Solidity: function getDate(uint256 _timestamp) view returns((int256,uint256,uint8,uint8))
func (_Api *ApiCallerSession) GetDate(_timestamp *big.Int) (OracleIndicatorDataFeed, error) {
	return _Api.Contract.GetDate(&_Api.CallOpts, _timestamp)
}

// GetInterval is a free data retrieval call binding the contract method 0x3488ecb3.
//
// Solidity: function getInterval(uint256 _start, uint256 _end) view returns(int256)
func (_Api *ApiCaller) GetInterval(opts *bind.CallOpts, _start *big.Int, _end *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getInterval", _start, _end)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterval is a free data retrieval call binding the contract method 0x3488ecb3.
//
// Solidity: function getInterval(uint256 _start, uint256 _end) view returns(int256)
func (_Api *ApiSession) GetInterval(_start *big.Int, _end *big.Int) (*big.Int, error) {
	return _Api.Contract.GetInterval(&_Api.CallOpts, _start, _end)
}

// GetInterval is a free data retrieval call binding the contract method 0x3488ecb3.
//
// Solidity: function getInterval(uint256 _start, uint256 _end) view returns(int256)
func (_Api *ApiCallerSession) GetInterval(_start *big.Int, _end *big.Int) (*big.Int, error) {
	return _Api.Contract.GetInterval(&_Api.CallOpts, _start, _end)
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() view returns((int256,uint256,uint8,uint8))
func (_Api *ApiCaller) GetLast(opts *bind.CallOpts) (OracleIndicatorDataFeed, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getLast")

	if err != nil {
		return *new(OracleIndicatorDataFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleIndicatorDataFeed)).(*OracleIndicatorDataFeed)

	return out0, err

}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() view returns((int256,uint256,uint8,uint8))
func (_Api *ApiSession) GetLast() (OracleIndicatorDataFeed, error) {
	return _Api.Contract.GetLast(&_Api.CallOpts)
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() view returns((int256,uint256,uint8,uint8))
func (_Api *ApiCallerSession) GetLast() (OracleIndicatorDataFeed, error) {
	return _Api.Contract.GetLast(&_Api.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Api *ApiCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Api *ApiSession) GetName() (string, error) {
	return _Api.Contract.GetName(&_Api.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Api *ApiCallerSession) GetName() (string, error) {
	return _Api.Contract.GetName(&_Api.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Api *ApiCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Api *ApiSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Api.Contract.GetRoleAdmin(&_Api.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Api *ApiCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Api.Contract.GetRoleAdmin(&_Api.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Api *ApiCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Api *ApiSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Api.Contract.HasRole(&_Api.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Api *ApiCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Api.Contract.HasRole(&_Api.CallOpts, role, account)
}

// Indicators is a free data retrieval call binding the contract method 0x92c871d2.
//
// Solidity: function indicators(uint256 ) view returns(int256 value, uint256 updatedat, uint8 decimal, uint8 confidence)
func (_Api *ApiCaller) Indicators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value      *big.Int
	Updatedat  *big.Int
	Decimal    uint8
	Confidence uint8
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "indicators", arg0)

	outstruct := new(struct {
		Value      *big.Int
		Updatedat  *big.Int
		Decimal    uint8
		Confidence uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Updatedat = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Decimal = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Confidence = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Indicators is a free data retrieval call binding the contract method 0x92c871d2.
//
// Solidity: function indicators(uint256 ) view returns(int256 value, uint256 updatedat, uint8 decimal, uint8 confidence)
func (_Api *ApiSession) Indicators(arg0 *big.Int) (struct {
	Value      *big.Int
	Updatedat  *big.Int
	Decimal    uint8
	Confidence uint8
}, error) {
	return _Api.Contract.Indicators(&_Api.CallOpts, arg0)
}

// Indicators is a free data retrieval call binding the contract method 0x92c871d2.
//
// Solidity: function indicators(uint256 ) view returns(int256 value, uint256 updatedat, uint8 decimal, uint8 confidence)
func (_Api *ApiCallerSession) Indicators(arg0 *big.Int) (struct {
	Value      *big.Int
	Updatedat  *big.Int
	Decimal    uint8
	Confidence uint8
}, error) {
	return _Api.Contract.Indicators(&_Api.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Api.Contract.SupportsInterface(&_Api.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Api.Contract.SupportsInterface(&_Api.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Api *ApiTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Api *ApiSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.GrantRole(&_Api.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Api *ApiTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.GrantRole(&_Api.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Api *ApiTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Api *ApiSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Api.Contract.RenounceRole(&_Api.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Api *ApiTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Api.Contract.RenounceRole(&_Api.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Api *ApiTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Api *ApiSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.RevokeRole(&_Api.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Api *ApiTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.RevokeRole(&_Api.TransactOpts, role, account)
}

// SaveIndicator is a paid mutator transaction binding the contract method 0x77c6e440.
//
// Solidity: function saveIndicator(uint256 _timestamp, int256 _value, uint256 _updatedat, uint8 _confidence) returns()
func (_Api *ApiTransactor) SaveIndicator(opts *bind.TransactOpts, _timestamp *big.Int, _value *big.Int, _updatedat *big.Int, _confidence uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "saveIndicator", _timestamp, _value, _updatedat, _confidence)
}

// SaveIndicator is a paid mutator transaction binding the contract method 0x77c6e440.
//
// Solidity: function saveIndicator(uint256 _timestamp, int256 _value, uint256 _updatedat, uint8 _confidence) returns()
func (_Api *ApiSession) SaveIndicator(_timestamp *big.Int, _value *big.Int, _updatedat *big.Int, _confidence uint8) (*types.Transaction, error) {
	return _Api.Contract.SaveIndicator(&_Api.TransactOpts, _timestamp, _value, _updatedat, _confidence)
}

// SaveIndicator is a paid mutator transaction binding the contract method 0x77c6e440.
//
// Solidity: function saveIndicator(uint256 _timestamp, int256 _value, uint256 _updatedat, uint8 _confidence) returns()
func (_Api *ApiTransactorSession) SaveIndicator(_timestamp *big.Int, _value *big.Int, _updatedat *big.Int, _confidence uint8) (*types.Transaction, error) {
	return _Api.Contract.SaveIndicator(&_Api.TransactOpts, _timestamp, _value, _updatedat, _confidence)
}

// ApiRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Api contract.
type ApiRoleAdminChangedIterator struct {
	Event *ApiRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ApiRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRoleAdminChanged)
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
		it.Event = new(ApiRoleAdminChanged)
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
func (it *ApiRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRoleAdminChanged represents a RoleAdminChanged event raised by the Api contract.
type ApiRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Api *ApiFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ApiRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ApiRoleAdminChangedIterator{contract: _Api.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Api *ApiFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ApiRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRoleAdminChanged)
				if err := _Api.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Api *ApiFilterer) ParseRoleAdminChanged(log types.Log) (*ApiRoleAdminChanged, error) {
	event := new(ApiRoleAdminChanged)
	if err := _Api.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Api contract.
type ApiRoleGrantedIterator struct {
	Event *ApiRoleGranted // Event containing the contract specifics and raw log

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
func (it *ApiRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRoleGranted)
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
		it.Event = new(ApiRoleGranted)
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
func (it *ApiRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRoleGranted represents a RoleGranted event raised by the Api contract.
type ApiRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApiRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiRoleGrantedIterator{contract: _Api.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ApiRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRoleGranted)
				if err := _Api.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) ParseRoleGranted(log types.Log) (*ApiRoleGranted, error) {
	event := new(ApiRoleGranted)
	if err := _Api.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Api contract.
type ApiRoleRevokedIterator struct {
	Event *ApiRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ApiRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRoleRevoked)
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
		it.Event = new(ApiRoleRevoked)
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
func (it *ApiRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRoleRevoked represents a RoleRevoked event raised by the Api contract.
type ApiRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApiRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiRoleRevokedIterator{contract: _Api.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ApiRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRoleRevoked)
				if err := _Api.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) ParseRoleRevoked(log types.Log) (*ApiRoleRevoked, error) {
	event := new(ApiRoleRevoked)
	if err := _Api.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
