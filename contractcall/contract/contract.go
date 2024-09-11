// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// DIDAttribute is an auto generated low-level Go binding around an user-defined struct.
type DIDAttribute struct {
	Name     []byte
	Value    []byte
	Validity uint32
	Created  *big.Int
}

// ErebrusRegistryVPNNode is an auto generated low-level Go binding around an user-defined struct.
type ErebrusRegistryVPNNode struct {
	User      common.Address
	PeaqDid   string
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"erebrusmanagerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"AddAttribute\",\"type\":\"event\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"did_account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"name\":\"NodeDeactivated\",\"type\":\"event\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"NodeOperatorUpdated\",\"type\":\"event\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ssid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"location\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"name\":\"RemoveAttribte\",\"type\":\"event\",\"inputs\":[{\"name\":\"did_account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"name\":\"UpdateAttribute\",\"type\":\"event\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"did_account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"name\":\"VPNUpdated\",\"type\":\"event\",\"inputs\":[{\"name\":\"nodeId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"updatedStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"updatedRegion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"name\":\"VpnNodeRegistered\",\"type\":\"event\",\"inputs\":[{\"name\":\"nodeId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nodename\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ipaddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ispinfo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"region\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"location\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"name\":\"WifiNodeOperatorRegistered\",\"type\":\"event\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deviceId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ssid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"location\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"pricePerMinute\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"VpnDeviceCheckpoints\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"name\":\"addAttribute\",\"type\":\"function\",\"inputs\":[{\"name\":\"did_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validity_for\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"deactivateNode\",\"type\":\"function\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"didToUser\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"getWifiDetails\",\"type\":\"function\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"readAttribute\",\"type\":\"function\",\"inputs\":[{\"name\":\"did_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"created\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"internalType\":\"structDID.Attribute\"}],\"stateMutability\":\"view\"},{\"name\":\"registerVpnNode\",\"type\":\"function\",\"inputs\":[{\"name\":\"node\",\"type\":\"tuple\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"peaqDid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nodename\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ipaddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ispinfo\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"region\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"location\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"internalType\":\"structErebrusRegistry.VPNNode\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"registerWifiNode\",\"type\":\"function\",\"inputs\":[{\"name\":\"_deviceId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_peaqDid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_ssid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_location\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_pricePermin\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"removeAttribute\",\"type\":\"function\",\"inputs\":[{\"name\":\"did_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"updateAttribute\",\"type\":\"function\",\"inputs\":[{\"name\":\"did_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validity_for\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"updateVPNNode\",\"type\":\"function\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_status\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_region\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"updateWiFiNode\",\"type\":\"function\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ssid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"location\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"pricePerMin\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"vpnDeviceCheckpoint\",\"type\":\"function\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataHash\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"vpnNodeOperators\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"peaqDid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nodename\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ipaddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ispinfo\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"region\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"location\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"name\":\"vpnTotalCheckpoints\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"wifiDeviceCheckpoint\",\"type\":\"function\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataHash\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"wifiDeviceCheckpoints\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"name\":\"wifiNodeOperators\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deviceId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"peaqDid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ssid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"location\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"pricePerMinute\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"wifiTotalCheckpoints\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// VpnDeviceCheckpoints is a free data retrieval call binding the contract method 0xee5458b3.
//
// Solidity: function VpnDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Contract *ContractCaller) VpnDeviceCheckpoints(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "VpnDeviceCheckpoints", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VpnDeviceCheckpoints is a free data retrieval call binding the contract method 0xee5458b3.
//
// Solidity: function VpnDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Contract *ContractSession) VpnDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.VpnDeviceCheckpoints(&_Contract.CallOpts, arg0, arg1)
}

// VpnDeviceCheckpoints is a free data retrieval call binding the contract method 0xee5458b3.
//
// Solidity: function VpnDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Contract *ContractCallerSession) VpnDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.VpnDeviceCheckpoints(&_Contract.CallOpts, arg0, arg1)
}

// DidToUser is a free data retrieval call binding the contract method 0xe4f4ff00.
//
// Solidity: function didToUser(address ) view returns(address)
func (_Contract *ContractCaller) DidToUser(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "didToUser", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DidToUser is a free data retrieval call binding the contract method 0xe4f4ff00.
//
// Solidity: function didToUser(address ) view returns(address)
func (_Contract *ContractSession) DidToUser(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.DidToUser(&_Contract.CallOpts, arg0)
}

// DidToUser is a free data retrieval call binding the contract method 0xe4f4ff00.
//
// Solidity: function didToUser(address ) view returns(address)
func (_Contract *ContractCallerSession) DidToUser(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.DidToUser(&_Contract.CallOpts, arg0)
}

// GetWifiDetails is a free data retrieval call binding the contract method 0x387f68e1.
//
// Solidity: function getWifiDetails(uint256 nodeID) view returns(uint256 price, address owner)
func (_Contract *ContractCaller) GetWifiDetails(opts *bind.CallOpts, nodeID *big.Int) (struct {
	Price *big.Int
	Owner common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWifiDetails", nodeID)

	outstruct := new(struct {
		Price *big.Int
		Owner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetWifiDetails is a free data retrieval call binding the contract method 0x387f68e1.
//
// Solidity: function getWifiDetails(uint256 nodeID) view returns(uint256 price, address owner)
func (_Contract *ContractSession) GetWifiDetails(nodeID *big.Int) (struct {
	Price *big.Int
	Owner common.Address
}, error) {
	return _Contract.Contract.GetWifiDetails(&_Contract.CallOpts, nodeID)
}

// GetWifiDetails is a free data retrieval call binding the contract method 0x387f68e1.
//
// Solidity: function getWifiDetails(uint256 nodeID) view returns(uint256 price, address owner)
func (_Contract *ContractCallerSession) GetWifiDetails(nodeID *big.Int) (struct {
	Price *big.Int
	Owner common.Address
}, error) {
	return _Contract.Contract.GetWifiDetails(&_Contract.CallOpts, nodeID)
}

// ReadAttribute is a free data retrieval call binding the contract method 0xb2028b7d.
//
// Solidity: function readAttribute(address did_account, bytes name) view returns((bytes,bytes,uint32,uint256))
func (_Contract *ContractCaller) ReadAttribute(opts *bind.CallOpts, did_account common.Address, name []byte) (DIDAttribute, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "readAttribute", did_account, name)

	if err != nil {
		return *new(DIDAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new(DIDAttribute)).(*DIDAttribute)

	return out0, err

}

// ReadAttribute is a free data retrieval call binding the contract method 0xb2028b7d.
//
// Solidity: function readAttribute(address did_account, bytes name) view returns((bytes,bytes,uint32,uint256))
func (_Contract *ContractSession) ReadAttribute(did_account common.Address, name []byte) (DIDAttribute, error) {
	return _Contract.Contract.ReadAttribute(&_Contract.CallOpts, did_account, name)
}

// ReadAttribute is a free data retrieval call binding the contract method 0xb2028b7d.
//
// Solidity: function readAttribute(address did_account, bytes name) view returns((bytes,bytes,uint32,uint256))
func (_Contract *ContractCallerSession) ReadAttribute(did_account common.Address, name []byte) (DIDAttribute, error) {
	return _Contract.Contract.ReadAttribute(&_Contract.CallOpts, did_account, name)
}

// VpnNodeOperators is a free data retrieval call binding the contract method 0xd5399313.
//
// Solidity: function vpnNodeOperators(uint256 ) view returns(address user, string peaqDid, string nodename, string ipaddress, string ispinfo, string region, string location, uint8 status)
func (_Contract *ContractCaller) VpnNodeOperators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User      common.Address
	PeaqDid   string
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vpnNodeOperators", arg0)

	outstruct := new(struct {
		User      common.Address
		PeaqDid   string
		Nodename  string
		Ipaddress string
		Ispinfo   string
		Region    string
		Location  string
		Status    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PeaqDid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Nodename = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Ipaddress = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Ispinfo = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Region = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Location = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// VpnNodeOperators is a free data retrieval call binding the contract method 0xd5399313.
//
// Solidity: function vpnNodeOperators(uint256 ) view returns(address user, string peaqDid, string nodename, string ipaddress, string ispinfo, string region, string location, uint8 status)
func (_Contract *ContractSession) VpnNodeOperators(arg0 *big.Int) (struct {
	User      common.Address
	PeaqDid   string
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}, error) {
	return _Contract.Contract.VpnNodeOperators(&_Contract.CallOpts, arg0)
}

// VpnNodeOperators is a free data retrieval call binding the contract method 0xd5399313.
//
// Solidity: function vpnNodeOperators(uint256 ) view returns(address user, string peaqDid, string nodename, string ipaddress, string ispinfo, string region, string location, uint8 status)
func (_Contract *ContractCallerSession) VpnNodeOperators(arg0 *big.Int) (struct {
	User      common.Address
	PeaqDid   string
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}, error) {
	return _Contract.Contract.VpnNodeOperators(&_Contract.CallOpts, arg0)
}

// VpnTotalCheckpoints is a free data retrieval call binding the contract method 0xef5050fb.
//
// Solidity: function vpnTotalCheckpoints(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) VpnTotalCheckpoints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vpnTotalCheckpoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VpnTotalCheckpoints is a free data retrieval call binding the contract method 0xef5050fb.
//
// Solidity: function vpnTotalCheckpoints(uint256 ) view returns(uint256)
func (_Contract *ContractSession) VpnTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.VpnTotalCheckpoints(&_Contract.CallOpts, arg0)
}

// VpnTotalCheckpoints is a free data retrieval call binding the contract method 0xef5050fb.
//
// Solidity: function vpnTotalCheckpoints(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) VpnTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.VpnTotalCheckpoints(&_Contract.CallOpts, arg0)
}

// WifiDeviceCheckpoints is a free data retrieval call binding the contract method 0x4f663d31.
//
// Solidity: function wifiDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Contract *ContractCaller) WifiDeviceCheckpoints(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "wifiDeviceCheckpoints", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WifiDeviceCheckpoints is a free data retrieval call binding the contract method 0x4f663d31.
//
// Solidity: function wifiDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Contract *ContractSession) WifiDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.WifiDeviceCheckpoints(&_Contract.CallOpts, arg0, arg1)
}

// WifiDeviceCheckpoints is a free data retrieval call binding the contract method 0x4f663d31.
//
// Solidity: function wifiDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Contract *ContractCallerSession) WifiDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Contract.Contract.WifiDeviceCheckpoints(&_Contract.CallOpts, arg0, arg1)
}

// WifiNodeOperators is a free data retrieval call binding the contract method 0x8f9c8ee2.
//
// Solidity: function wifiNodeOperators(uint256 ) view returns(address user, string deviceId, string peaqDid, string ssid, string location, uint256 pricePerMinute, bool isActive)
func (_Contract *ContractCaller) WifiNodeOperators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User           common.Address
	DeviceId       string
	PeaqDid        string
	Ssid           string
	Location       string
	PricePerMinute *big.Int
	IsActive       bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "wifiNodeOperators", arg0)

	outstruct := new(struct {
		User           common.Address
		DeviceId       string
		PeaqDid        string
		Ssid           string
		Location       string
		PricePerMinute *big.Int
		IsActive       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DeviceId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.PeaqDid = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Ssid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Location = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.PricePerMinute = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// WifiNodeOperators is a free data retrieval call binding the contract method 0x8f9c8ee2.
//
// Solidity: function wifiNodeOperators(uint256 ) view returns(address user, string deviceId, string peaqDid, string ssid, string location, uint256 pricePerMinute, bool isActive)
func (_Contract *ContractSession) WifiNodeOperators(arg0 *big.Int) (struct {
	User           common.Address
	DeviceId       string
	PeaqDid        string
	Ssid           string
	Location       string
	PricePerMinute *big.Int
	IsActive       bool
}, error) {
	return _Contract.Contract.WifiNodeOperators(&_Contract.CallOpts, arg0)
}

// WifiNodeOperators is a free data retrieval call binding the contract method 0x8f9c8ee2.
//
// Solidity: function wifiNodeOperators(uint256 ) view returns(address user, string deviceId, string peaqDid, string ssid, string location, uint256 pricePerMinute, bool isActive)
func (_Contract *ContractCallerSession) WifiNodeOperators(arg0 *big.Int) (struct {
	User           common.Address
	DeviceId       string
	PeaqDid        string
	Ssid           string
	Location       string
	PricePerMinute *big.Int
	IsActive       bool
}, error) {
	return _Contract.Contract.WifiNodeOperators(&_Contract.CallOpts, arg0)
}

// WifiTotalCheckpoints is a free data retrieval call binding the contract method 0x0608a6b6.
//
// Solidity: function wifiTotalCheckpoints(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) WifiTotalCheckpoints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "wifiTotalCheckpoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WifiTotalCheckpoints is a free data retrieval call binding the contract method 0x0608a6b6.
//
// Solidity: function wifiTotalCheckpoints(uint256 ) view returns(uint256)
func (_Contract *ContractSession) WifiTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.WifiTotalCheckpoints(&_Contract.CallOpts, arg0)
}

// WifiTotalCheckpoints is a free data retrieval call binding the contract method 0x0608a6b6.
//
// Solidity: function wifiTotalCheckpoints(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) WifiTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.WifiTotalCheckpoints(&_Contract.CallOpts, arg0)
}

// AddAttribute is a paid mutator transaction binding the contract method 0xcc4a70ca.
//
// Solidity: function addAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Contract *ContractTransactor) AddAttribute(opts *bind.TransactOpts, did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addAttribute", did_account, name, value, validity_for)
}

// AddAttribute is a paid mutator transaction binding the contract method 0xcc4a70ca.
//
// Solidity: function addAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Contract *ContractSession) AddAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Contract.Contract.AddAttribute(&_Contract.TransactOpts, did_account, name, value, validity_for)
}

// AddAttribute is a paid mutator transaction binding the contract method 0xcc4a70ca.
//
// Solidity: function addAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Contract *ContractTransactorSession) AddAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Contract.Contract.AddAttribute(&_Contract.TransactOpts, did_account, name, value, validity_for)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xe9d63830.
//
// Solidity: function deactivateNode(uint256 nodeID) returns()
func (_Contract *ContractTransactor) DeactivateNode(opts *bind.TransactOpts, nodeID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateNode", nodeID)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xe9d63830.
//
// Solidity: function deactivateNode(uint256 nodeID) returns()
func (_Contract *ContractSession) DeactivateNode(nodeID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateNode(&_Contract.TransactOpts, nodeID)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xe9d63830.
//
// Solidity: function deactivateNode(uint256 nodeID) returns()
func (_Contract *ContractTransactorSession) DeactivateNode(nodeID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateNode(&_Contract.TransactOpts, nodeID)
}

// RegisterVpnNode is a paid mutator transaction binding the contract method 0xd04b7ae4.
//
// Solidity: function registerVpnNode((address,string,string,string,string,string,string,uint8) node) returns()
func (_Contract *ContractTransactor) RegisterVpnNode(opts *bind.TransactOpts, node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerVpnNode", node)
}

// RegisterVpnNode is a paid mutator transaction binding the contract method 0xd04b7ae4.
//
// Solidity: function registerVpnNode((address,string,string,string,string,string,string,uint8) node) returns()
func (_Contract *ContractSession) RegisterVpnNode(node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Contract.Contract.RegisterVpnNode(&_Contract.TransactOpts, node)
}

// RegisterVpnNode is a paid mutator transaction binding the contract method 0xd04b7ae4.
//
// Solidity: function registerVpnNode((address,string,string,string,string,string,string,uint8) node) returns()
func (_Contract *ContractTransactorSession) RegisterVpnNode(node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Contract.Contract.RegisterVpnNode(&_Contract.TransactOpts, node)
}

// RegisterWifiNode is a paid mutator transaction binding the contract method 0x04795f58.
//
// Solidity: function registerWifiNode(string _deviceId, string _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Contract *ContractTransactor) RegisterWifiNode(opts *bind.TransactOpts, _deviceId string, _peaqDid string, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerWifiNode", _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// RegisterWifiNode is a paid mutator transaction binding the contract method 0x04795f58.
//
// Solidity: function registerWifiNode(string _deviceId, string _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Contract *ContractSession) RegisterWifiNode(_deviceId string, _peaqDid string, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterWifiNode(&_Contract.TransactOpts, _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// RegisterWifiNode is a paid mutator transaction binding the contract method 0x04795f58.
//
// Solidity: function registerWifiNode(string _deviceId, string _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Contract *ContractTransactorSession) RegisterWifiNode(_deviceId string, _peaqDid string, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterWifiNode(&_Contract.TransactOpts, _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// RemoveAttribute is a paid mutator transaction binding the contract method 0xe8a81690.
//
// Solidity: function removeAttribute(address did_account, bytes name) returns(bool)
func (_Contract *ContractTransactor) RemoveAttribute(opts *bind.TransactOpts, did_account common.Address, name []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeAttribute", did_account, name)
}

// RemoveAttribute is a paid mutator transaction binding the contract method 0xe8a81690.
//
// Solidity: function removeAttribute(address did_account, bytes name) returns(bool)
func (_Contract *ContractSession) RemoveAttribute(did_account common.Address, name []byte) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAttribute(&_Contract.TransactOpts, did_account, name)
}

// RemoveAttribute is a paid mutator transaction binding the contract method 0xe8a81690.
//
// Solidity: function removeAttribute(address did_account, bytes name) returns(bool)
func (_Contract *ContractTransactorSession) RemoveAttribute(did_account common.Address, name []byte) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAttribute(&_Contract.TransactOpts, did_account, name)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x68b4b2c1.
//
// Solidity: function updateAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Contract *ContractTransactor) UpdateAttribute(opts *bind.TransactOpts, did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAttribute", did_account, name, value, validity_for)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x68b4b2c1.
//
// Solidity: function updateAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Contract *ContractSession) UpdateAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttribute(&_Contract.TransactOpts, did_account, name, value, validity_for)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x68b4b2c1.
//
// Solidity: function updateAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Contract *ContractTransactorSession) UpdateAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttribute(&_Contract.TransactOpts, did_account, name, value, validity_for)
}

// UpdateVPNNode is a paid mutator transaction binding the contract method 0xe48da8a4.
//
// Solidity: function updateVPNNode(uint256 nodeID, uint8 _status, string _region) returns()
func (_Contract *ContractTransactor) UpdateVPNNode(opts *bind.TransactOpts, nodeID *big.Int, _status uint8, _region string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateVPNNode", nodeID, _status, _region)
}

// UpdateVPNNode is a paid mutator transaction binding the contract method 0xe48da8a4.
//
// Solidity: function updateVPNNode(uint256 nodeID, uint8 _status, string _region) returns()
func (_Contract *ContractSession) UpdateVPNNode(nodeID *big.Int, _status uint8, _region string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateVPNNode(&_Contract.TransactOpts, nodeID, _status, _region)
}

// UpdateVPNNode is a paid mutator transaction binding the contract method 0xe48da8a4.
//
// Solidity: function updateVPNNode(uint256 nodeID, uint8 _status, string _region) returns()
func (_Contract *ContractTransactorSession) UpdateVPNNode(nodeID *big.Int, _status uint8, _region string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateVPNNode(&_Contract.TransactOpts, nodeID, _status, _region)
}

// UpdateWiFiNode is a paid mutator transaction binding the contract method 0x8a7692dc.
//
// Solidity: function updateWiFiNode(uint256 nodeID, string ssid, string location, uint256 pricePerMin) returns()
func (_Contract *ContractTransactor) UpdateWiFiNode(opts *bind.TransactOpts, nodeID *big.Int, ssid string, location string, pricePerMin *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateWiFiNode", nodeID, ssid, location, pricePerMin)
}

// UpdateWiFiNode is a paid mutator transaction binding the contract method 0x8a7692dc.
//
// Solidity: function updateWiFiNode(uint256 nodeID, string ssid, string location, uint256 pricePerMin) returns()
func (_Contract *ContractSession) UpdateWiFiNode(nodeID *big.Int, ssid string, location string, pricePerMin *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateWiFiNode(&_Contract.TransactOpts, nodeID, ssid, location, pricePerMin)
}

// UpdateWiFiNode is a paid mutator transaction binding the contract method 0x8a7692dc.
//
// Solidity: function updateWiFiNode(uint256 nodeID, string ssid, string location, uint256 pricePerMin) returns()
func (_Contract *ContractTransactorSession) UpdateWiFiNode(nodeID *big.Int, ssid string, location string, pricePerMin *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateWiFiNode(&_Contract.TransactOpts, nodeID, ssid, location, pricePerMin)
}

// VpnDeviceCheckpoint is a paid mutator transaction binding the contract method 0x1b9f0c95.
//
// Solidity: function vpnDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Contract *ContractTransactor) VpnDeviceCheckpoint(opts *bind.TransactOpts, nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "vpnDeviceCheckpoint", nodeID, dataHash)
}

// VpnDeviceCheckpoint is a paid mutator transaction binding the contract method 0x1b9f0c95.
//
// Solidity: function vpnDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Contract *ContractSession) VpnDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Contract.Contract.VpnDeviceCheckpoint(&_Contract.TransactOpts, nodeID, dataHash)
}

// VpnDeviceCheckpoint is a paid mutator transaction binding the contract method 0x1b9f0c95.
//
// Solidity: function vpnDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Contract *ContractTransactorSession) VpnDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Contract.Contract.VpnDeviceCheckpoint(&_Contract.TransactOpts, nodeID, dataHash)
}

// WifiDeviceCheckpoint is a paid mutator transaction binding the contract method 0x63c3a7e7.
//
// Solidity: function wifiDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Contract *ContractTransactor) WifiDeviceCheckpoint(opts *bind.TransactOpts, nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "wifiDeviceCheckpoint", nodeID, dataHash)
}

// WifiDeviceCheckpoint is a paid mutator transaction binding the contract method 0x63c3a7e7.
//
// Solidity: function wifiDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Contract *ContractSession) WifiDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Contract.Contract.WifiDeviceCheckpoint(&_Contract.TransactOpts, nodeID, dataHash)
}

// WifiDeviceCheckpoint is a paid mutator transaction binding the contract method 0x63c3a7e7.
//
// Solidity: function wifiDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Contract *ContractTransactorSession) WifiDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Contract.Contract.WifiDeviceCheckpoint(&_Contract.TransactOpts, nodeID, dataHash)
}

// ContractAddAttributeIterator is returned from FilterAddAttribute and is used to iterate over the raw logs and unpacked data for AddAttribute events raised by the Contract contract.
type ContractAddAttributeIterator struct {
	Event *ContractAddAttribute // Event containing the contract specifics and raw log

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
func (it *ContractAddAttributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddAttribute)
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
		it.Event = new(ContractAddAttribute)
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
func (it *ContractAddAttributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddAttributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddAttribute represents a AddAttribute event raised by the Contract contract.
type ContractAddAttribute struct {
	Sender     common.Address
	DidAccount common.Address
	Name       []byte
	Value      []byte
	Validity   uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddAttribute is a free log retrieval operation binding the contract event 0x13aef52bc4a99da04591533072e304017e3fb76f43e7fadd25eb7f514c5ef6e5.
//
// Solidity: event AddAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Contract *ContractFilterer) FilterAddAttribute(opts *bind.FilterOpts) (*ContractAddAttributeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddAttribute")
	if err != nil {
		return nil, err
	}
	return &ContractAddAttributeIterator{contract: _Contract.contract, event: "AddAttribute", logs: logs, sub: sub}, nil
}

// WatchAddAttribute is a free log subscription operation binding the contract event 0x13aef52bc4a99da04591533072e304017e3fb76f43e7fadd25eb7f514c5ef6e5.
//
// Solidity: event AddAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Contract *ContractFilterer) WatchAddAttribute(opts *bind.WatchOpts, sink chan<- *ContractAddAttribute) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddAttribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddAttribute)
				if err := _Contract.contract.UnpackLog(event, "AddAttribute", log); err != nil {
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

// ParseAddAttribute is a log parse operation binding the contract event 0x13aef52bc4a99da04591533072e304017e3fb76f43e7fadd25eb7f514c5ef6e5.
//
// Solidity: event AddAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Contract *ContractFilterer) ParseAddAttribute(log types.Log) (*ContractAddAttribute, error) {
	event := new(ContractAddAttribute)
	if err := _Contract.contract.UnpackLog(event, "AddAttribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the Contract contract.
type ContractNodeDeactivatedIterator struct {
	Event *ContractNodeDeactivated // Event containing the contract specifics and raw log

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
func (it *ContractNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeDeactivated)
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
		it.Event = new(ContractNodeDeactivated)
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
func (it *ContractNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeDeactivated represents a NodeDeactivated event raised by the Contract contract.
type ContractNodeDeactivated struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xd9957750e6343405c319eb99a4ec67fa11cfd66969318cbc71aa2d45fa53a349.
//
// Solidity: event NodeDeactivated(address indexed operatorAddress)
func (_Contract *ContractFilterer) FilterNodeDeactivated(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractNodeDeactivatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeDeactivated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractNodeDeactivatedIterator{contract: _Contract.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xd9957750e6343405c319eb99a4ec67fa11cfd66969318cbc71aa2d45fa53a349.
//
// Solidity: event NodeDeactivated(address indexed operatorAddress)
func (_Contract *ContractFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *ContractNodeDeactivated, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeDeactivated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeDeactivated)
				if err := _Contract.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
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

// ParseNodeDeactivated is a log parse operation binding the contract event 0xd9957750e6343405c319eb99a4ec67fa11cfd66969318cbc71aa2d45fa53a349.
//
// Solidity: event NodeDeactivated(address indexed operatorAddress)
func (_Contract *ContractFilterer) ParseNodeDeactivated(log types.Log) (*ContractNodeDeactivated, error) {
	event := new(ContractNodeDeactivated)
	if err := _Contract.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeOperatorUpdatedIterator is returned from FilterNodeOperatorUpdated and is used to iterate over the raw logs and unpacked data for NodeOperatorUpdated events raised by the Contract contract.
type ContractNodeOperatorUpdatedIterator struct {
	Event *ContractNodeOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractNodeOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeOperatorUpdated)
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
		it.Event = new(ContractNodeOperatorUpdated)
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
func (it *ContractNodeOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeOperatorUpdated represents a NodeOperatorUpdated event raised by the Contract contract.
type ContractNodeOperatorUpdated struct {
	OperatorAddress common.Address
	Ssid            string
	Location        string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorUpdated is a free log retrieval operation binding the contract event 0xf512e0b8ea262571de058d8b17a3c2d29913f35abcd4c173935ba6fa58b3833b.
//
// Solidity: event NodeOperatorUpdated(address indexed operatorAddress, string ssid, string location)
func (_Contract *ContractFilterer) FilterNodeOperatorUpdated(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractNodeOperatorUpdatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeOperatorUpdated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractNodeOperatorUpdatedIterator{contract: _Contract.contract, event: "NodeOperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorUpdated is a free log subscription operation binding the contract event 0xf512e0b8ea262571de058d8b17a3c2d29913f35abcd4c173935ba6fa58b3833b.
//
// Solidity: event NodeOperatorUpdated(address indexed operatorAddress, string ssid, string location)
func (_Contract *ContractFilterer) WatchNodeOperatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractNodeOperatorUpdated, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeOperatorUpdated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeOperatorUpdated)
				if err := _Contract.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
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

// ParseNodeOperatorUpdated is a log parse operation binding the contract event 0xf512e0b8ea262571de058d8b17a3c2d29913f35abcd4c173935ba6fa58b3833b.
//
// Solidity: event NodeOperatorUpdated(address indexed operatorAddress, string ssid, string location)
func (_Contract *ContractFilterer) ParseNodeOperatorUpdated(log types.Log) (*ContractNodeOperatorUpdated, error) {
	event := new(ContractNodeOperatorUpdated)
	if err := _Contract.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRemoveAttribteIterator is returned from FilterRemoveAttribte and is used to iterate over the raw logs and unpacked data for RemoveAttribte events raised by the Contract contract.
type ContractRemoveAttribteIterator struct {
	Event *ContractRemoveAttribte // Event containing the contract specifics and raw log

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
func (it *ContractRemoveAttribteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRemoveAttribte)
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
		it.Event = new(ContractRemoveAttribte)
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
func (it *ContractRemoveAttribteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRemoveAttribteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRemoveAttribte represents a RemoveAttribte event raised by the Contract contract.
type ContractRemoveAttribte struct {
	DidAccount common.Address
	Name       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveAttribte is a free log retrieval operation binding the contract event 0x647d95b9099b095e5957558106bc51df77720f963539a0ca62ee9948e6554da1.
//
// Solidity: event RemoveAttribte(address did_account, bytes name)
func (_Contract *ContractFilterer) FilterRemoveAttribte(opts *bind.FilterOpts) (*ContractRemoveAttribteIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RemoveAttribte")
	if err != nil {
		return nil, err
	}
	return &ContractRemoveAttribteIterator{contract: _Contract.contract, event: "RemoveAttribte", logs: logs, sub: sub}, nil
}

// WatchRemoveAttribte is a free log subscription operation binding the contract event 0x647d95b9099b095e5957558106bc51df77720f963539a0ca62ee9948e6554da1.
//
// Solidity: event RemoveAttribte(address did_account, bytes name)
func (_Contract *ContractFilterer) WatchRemoveAttribte(opts *bind.WatchOpts, sink chan<- *ContractRemoveAttribte) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RemoveAttribte")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRemoveAttribte)
				if err := _Contract.contract.UnpackLog(event, "RemoveAttribte", log); err != nil {
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

// ParseRemoveAttribte is a log parse operation binding the contract event 0x647d95b9099b095e5957558106bc51df77720f963539a0ca62ee9948e6554da1.
//
// Solidity: event RemoveAttribte(address did_account, bytes name)
func (_Contract *ContractFilterer) ParseRemoveAttribte(log types.Log) (*ContractRemoveAttribte, error) {
	event := new(ContractRemoveAttribte)
	if err := _Contract.contract.UnpackLog(event, "RemoveAttribte", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateAttributeIterator is returned from FilterUpdateAttribute and is used to iterate over the raw logs and unpacked data for UpdateAttribute events raised by the Contract contract.
type ContractUpdateAttributeIterator struct {
	Event *ContractUpdateAttribute // Event containing the contract specifics and raw log

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
func (it *ContractUpdateAttributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateAttribute)
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
		it.Event = new(ContractUpdateAttribute)
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
func (it *ContractUpdateAttributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateAttributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateAttribute represents a UpdateAttribute event raised by the Contract contract.
type ContractUpdateAttribute struct {
	Sender     common.Address
	DidAccount common.Address
	Name       []byte
	Value      []byte
	Validity   uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAttribute is a free log retrieval operation binding the contract event 0x018be4d5e2634aefdb51c4ddd186f33072cfb5f6baf685579c2e214995dc9e4f.
//
// Solidity: event UpdateAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Contract *ContractFilterer) FilterUpdateAttribute(opts *bind.FilterOpts) (*ContractUpdateAttributeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateAttribute")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateAttributeIterator{contract: _Contract.contract, event: "UpdateAttribute", logs: logs, sub: sub}, nil
}

// WatchUpdateAttribute is a free log subscription operation binding the contract event 0x018be4d5e2634aefdb51c4ddd186f33072cfb5f6baf685579c2e214995dc9e4f.
//
// Solidity: event UpdateAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Contract *ContractFilterer) WatchUpdateAttribute(opts *bind.WatchOpts, sink chan<- *ContractUpdateAttribute) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateAttribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateAttribute)
				if err := _Contract.contract.UnpackLog(event, "UpdateAttribute", log); err != nil {
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

// ParseUpdateAttribute is a log parse operation binding the contract event 0x018be4d5e2634aefdb51c4ddd186f33072cfb5f6baf685579c2e214995dc9e4f.
//
// Solidity: event UpdateAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Contract *ContractFilterer) ParseUpdateAttribute(log types.Log) (*ContractUpdateAttribute, error) {
	event := new(ContractUpdateAttribute)
	if err := _Contract.contract.UnpackLog(event, "UpdateAttribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVPNUpdatedIterator is returned from FilterVPNUpdated and is used to iterate over the raw logs and unpacked data for VPNUpdated events raised by the Contract contract.
type ContractVPNUpdatedIterator struct {
	Event *ContractVPNUpdated // Event containing the contract specifics and raw log

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
func (it *ContractVPNUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVPNUpdated)
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
		it.Event = new(ContractVPNUpdated)
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
func (it *ContractVPNUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVPNUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVPNUpdated represents a VPNUpdated event raised by the Contract contract.
type ContractVPNUpdated struct {
	NodeId        *big.Int
	UpdatedStatus uint8
	UpdatedRegion string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVPNUpdated is a free log retrieval operation binding the contract event 0x34349e39780aaa7fc3b8b465ca9a2ada7a8a0a941e269b03102ad3867ff34b66.
//
// Solidity: event VPNUpdated(uint256 nodeId, uint8 updatedStatus, string updatedRegion)
func (_Contract *ContractFilterer) FilterVPNUpdated(opts *bind.FilterOpts) (*ContractVPNUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VPNUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractVPNUpdatedIterator{contract: _Contract.contract, event: "VPNUpdated", logs: logs, sub: sub}, nil
}

// WatchVPNUpdated is a free log subscription operation binding the contract event 0x34349e39780aaa7fc3b8b465ca9a2ada7a8a0a941e269b03102ad3867ff34b66.
//
// Solidity: event VPNUpdated(uint256 nodeId, uint8 updatedStatus, string updatedRegion)
func (_Contract *ContractFilterer) WatchVPNUpdated(opts *bind.WatchOpts, sink chan<- *ContractVPNUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VPNUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVPNUpdated)
				if err := _Contract.contract.UnpackLog(event, "VPNUpdated", log); err != nil {
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

// ParseVPNUpdated is a log parse operation binding the contract event 0x34349e39780aaa7fc3b8b465ca9a2ada7a8a0a941e269b03102ad3867ff34b66.
//
// Solidity: event VPNUpdated(uint256 nodeId, uint8 updatedStatus, string updatedRegion)
func (_Contract *ContractFilterer) ParseVPNUpdated(log types.Log) (*ContractVPNUpdated, error) {
	event := new(ContractVPNUpdated)
	if err := _Contract.contract.UnpackLog(event, "VPNUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVpnNodeRegisteredIterator is returned from FilterVpnNodeRegistered and is used to iterate over the raw logs and unpacked data for VpnNodeRegistered events raised by the Contract contract.
type ContractVpnNodeRegisteredIterator struct {
	Event *ContractVpnNodeRegistered // Event containing the contract specifics and raw log

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
func (it *ContractVpnNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVpnNodeRegistered)
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
		it.Event = new(ContractVpnNodeRegistered)
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
func (it *ContractVpnNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVpnNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVpnNodeRegistered represents a VpnNodeRegistered event raised by the Contract contract.
type ContractVpnNodeRegistered struct {
	NodeId    *big.Int
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVpnNodeRegistered is a free log retrieval operation binding the contract event 0x9e6fb352d55d4e4414d6ec189c261105e99d8baba96bd55c43a2013e2ee9453b.
//
// Solidity: event VpnNodeRegistered(uint256 nodeId, string nodename, string ipaddress, string ispinfo, string region, string location)
func (_Contract *ContractFilterer) FilterVpnNodeRegistered(opts *bind.FilterOpts) (*ContractVpnNodeRegisteredIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VpnNodeRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractVpnNodeRegisteredIterator{contract: _Contract.contract, event: "VpnNodeRegistered", logs: logs, sub: sub}, nil
}

// WatchVpnNodeRegistered is a free log subscription operation binding the contract event 0x9e6fb352d55d4e4414d6ec189c261105e99d8baba96bd55c43a2013e2ee9453b.
//
// Solidity: event VpnNodeRegistered(uint256 nodeId, string nodename, string ipaddress, string ispinfo, string region, string location)
func (_Contract *ContractFilterer) WatchVpnNodeRegistered(opts *bind.WatchOpts, sink chan<- *ContractVpnNodeRegistered) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VpnNodeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVpnNodeRegistered)
				if err := _Contract.contract.UnpackLog(event, "VpnNodeRegistered", log); err != nil {
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

// ParseVpnNodeRegistered is a log parse operation binding the contract event 0x9e6fb352d55d4e4414d6ec189c261105e99d8baba96bd55c43a2013e2ee9453b.
//
// Solidity: event VpnNodeRegistered(uint256 nodeId, string nodename, string ipaddress, string ispinfo, string region, string location)
func (_Contract *ContractFilterer) ParseVpnNodeRegistered(log types.Log) (*ContractVpnNodeRegistered, error) {
	event := new(ContractVpnNodeRegistered)
	if err := _Contract.contract.UnpackLog(event, "VpnNodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWifiNodeOperatorRegisteredIterator is returned from FilterWifiNodeOperatorRegistered and is used to iterate over the raw logs and unpacked data for WifiNodeOperatorRegistered events raised by the Contract contract.
type ContractWifiNodeOperatorRegisteredIterator struct {
	Event *ContractWifiNodeOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractWifiNodeOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWifiNodeOperatorRegistered)
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
		it.Event = new(ContractWifiNodeOperatorRegistered)
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
func (it *ContractWifiNodeOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWifiNodeOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWifiNodeOperatorRegistered represents a WifiNodeOperatorRegistered event raised by the Contract contract.
type ContractWifiNodeOperatorRegistered struct {
	NodeID         *big.Int
	Owner          common.Address
	DeviceId       string
	Ssid           string
	Location       string
	PricePerMinute *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWifiNodeOperatorRegistered is a free log retrieval operation binding the contract event 0x72ab8f8888d362028a5ec40836cefca72088b92b1ef7afc7fb795556b02d6658.
//
// Solidity: event WifiNodeOperatorRegistered(uint256 nodeID, address indexed owner, string deviceId, string ssid, string location, uint256 pricePerMinute)
func (_Contract *ContractFilterer) FilterWifiNodeOperatorRegistered(opts *bind.FilterOpts, owner []common.Address) (*ContractWifiNodeOperatorRegisteredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "WifiNodeOperatorRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractWifiNodeOperatorRegisteredIterator{contract: _Contract.contract, event: "WifiNodeOperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchWifiNodeOperatorRegistered is a free log subscription operation binding the contract event 0x72ab8f8888d362028a5ec40836cefca72088b92b1ef7afc7fb795556b02d6658.
//
// Solidity: event WifiNodeOperatorRegistered(uint256 nodeID, address indexed owner, string deviceId, string ssid, string location, uint256 pricePerMinute)
func (_Contract *ContractFilterer) WatchWifiNodeOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractWifiNodeOperatorRegistered, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "WifiNodeOperatorRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWifiNodeOperatorRegistered)
				if err := _Contract.contract.UnpackLog(event, "WifiNodeOperatorRegistered", log); err != nil {
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

// ParseWifiNodeOperatorRegistered is a log parse operation binding the contract event 0x72ab8f8888d362028a5ec40836cefca72088b92b1ef7afc7fb795556b02d6658.
//
// Solidity: event WifiNodeOperatorRegistered(uint256 nodeID, address indexed owner, string deviceId, string ssid, string location, uint256 pricePerMinute)
func (_Contract *ContractFilterer) ParseWifiNodeOperatorRegistered(log types.Log) (*ContractWifiNodeOperatorRegistered, error) {
	event := new(ContractWifiNodeOperatorRegistered)
	if err := _Contract.contract.UnpackLog(event, "WifiNodeOperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
