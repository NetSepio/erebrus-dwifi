// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node

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
	PeaqDid   common.Address
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erebrusmanagerAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"did_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"validity\",\"type\":\"uint32\"}],\"name\":\"AddAttribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"NodeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ssid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"NodeOperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"did_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"RemoveAttribte\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"did_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"validity\",\"type\":\"uint32\"}],\"name\":\"UpdateAttribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"updatedStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"updatedRegion\",\"type\":\"string\"}],\"name\":\"VPNUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodename\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipaddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ispinfo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"VpnNodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ssid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"}],\"name\":\"WifiNodeOperatorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"VpnDeviceCheckpoints\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"did_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"validity_for\",\"type\":\"uint32\"}],\"name\":\"addAttribute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeID\",\"type\":\"uint256\"}],\"name\":\"deactivateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"peaqDid\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodename\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipaddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ispinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structErebrusRegistry.VPNNode\",\"name\":\"node\",\"type\":\"tuple\"}],\"name\":\"delegateRegisterVpnNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_peaqDid\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ssid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_pricePermin\",\"type\":\"uint256\"}],\"name\":\"delegateRegisterWifiNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"didToUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeID\",\"type\":\"uint256\"}],\"name\":\"getWifiDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"did_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"readAttribute\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"validity\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"created\",\"type\":\"uint256\"}],\"internalType\":\"structDID.Attribute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"peaqDid\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodename\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipaddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ispinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structErebrusRegistry.VPNNode\",\"name\":\"node\",\"type\":\"tuple\"}],\"name\":\"registerVpnNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_peaqDid\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ssid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_pricePermin\",\"type\":\"uint256\"}],\"name\":\"registerWifiNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"did_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"removeAttribute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"did_account\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"validity_for\",\"type\":\"uint32\"}],\"name\":\"updateAttribute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeID\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_region\",\"type\":\"string\"}],\"name\":\"updateVPNNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ssid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMin\",\"type\":\"uint256\"}],\"name\":\"updateWiFiNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"}],\"name\":\"vpnDeviceCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vpnNodeOperators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"peaqDid\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodename\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipaddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ispinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vpnTotalCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"}],\"name\":\"wifiDeviceCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wifiDeviceCheckpoints\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wifiNodeOperators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"peaqDid\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ssid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wifiTotalCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}],",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// VpnDeviceCheckpoints is a free data retrieval call binding the contract method 0xee5458b3.
//
// Solidity: function VpnDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Node *NodeCaller) VpnDeviceCheckpoints(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "VpnDeviceCheckpoints", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VpnDeviceCheckpoints is a free data retrieval call binding the contract method 0xee5458b3.
//
// Solidity: function VpnDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Node *NodeSession) VpnDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Node.Contract.VpnDeviceCheckpoints(&_Node.CallOpts, arg0, arg1)
}

// VpnDeviceCheckpoints is a free data retrieval call binding the contract method 0xee5458b3.
//
// Solidity: function VpnDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Node *NodeCallerSession) VpnDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Node.Contract.VpnDeviceCheckpoints(&_Node.CallOpts, arg0, arg1)
}

// DidToUser is a free data retrieval call binding the contract method 0xe4f4ff00.
//
// Solidity: function didToUser(address ) view returns(address)
func (_Node *NodeCaller) DidToUser(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "didToUser", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DidToUser is a free data retrieval call binding the contract method 0xe4f4ff00.
//
// Solidity: function didToUser(address ) view returns(address)
func (_Node *NodeSession) DidToUser(arg0 common.Address) (common.Address, error) {
	return _Node.Contract.DidToUser(&_Node.CallOpts, arg0)
}

// DidToUser is a free data retrieval call binding the contract method 0xe4f4ff00.
//
// Solidity: function didToUser(address ) view returns(address)
func (_Node *NodeCallerSession) DidToUser(arg0 common.Address) (common.Address, error) {
	return _Node.Contract.DidToUser(&_Node.CallOpts, arg0)
}

// GetWifiDetails is a free data retrieval call binding the contract method 0x387f68e1.
//
// Solidity: function getWifiDetails(uint256 nodeID) view returns(uint256 price, address owner)
func (_Node *NodeCaller) GetWifiDetails(opts *bind.CallOpts, nodeID *big.Int) (struct {
	Price *big.Int
	Owner common.Address
}, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getWifiDetails", nodeID)

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
func (_Node *NodeSession) GetWifiDetails(nodeID *big.Int) (struct {
	Price *big.Int
	Owner common.Address
}, error) {
	return _Node.Contract.GetWifiDetails(&_Node.CallOpts, nodeID)
}

// GetWifiDetails is a free data retrieval call binding the contract method 0x387f68e1.
//
// Solidity: function getWifiDetails(uint256 nodeID) view returns(uint256 price, address owner)
func (_Node *NodeCallerSession) GetWifiDetails(nodeID *big.Int) (struct {
	Price *big.Int
	Owner common.Address
}, error) {
	return _Node.Contract.GetWifiDetails(&_Node.CallOpts, nodeID)
}

// ReadAttribute is a free data retrieval call binding the contract method 0xb2028b7d.
//
// Solidity: function readAttribute(address did_account, bytes name) view returns((bytes,bytes,uint32,uint256))
func (_Node *NodeCaller) ReadAttribute(opts *bind.CallOpts, did_account common.Address, name []byte) (DIDAttribute, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "readAttribute", did_account, name)

	if err != nil {
		return *new(DIDAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new(DIDAttribute)).(*DIDAttribute)

	return out0, err

}

// ReadAttribute is a free data retrieval call binding the contract method 0xb2028b7d.
//
// Solidity: function readAttribute(address did_account, bytes name) view returns((bytes,bytes,uint32,uint256))
func (_Node *NodeSession) ReadAttribute(did_account common.Address, name []byte) (DIDAttribute, error) {
	return _Node.Contract.ReadAttribute(&_Node.CallOpts, did_account, name)
}

// ReadAttribute is a free data retrieval call binding the contract method 0xb2028b7d.
//
// Solidity: function readAttribute(address did_account, bytes name) view returns((bytes,bytes,uint32,uint256))
func (_Node *NodeCallerSession) ReadAttribute(did_account common.Address, name []byte) (DIDAttribute, error) {
	return _Node.Contract.ReadAttribute(&_Node.CallOpts, did_account, name)
}

// VpnNodeOperators is a free data retrieval call binding the contract method 0xd5399313.
//
// Solidity: function vpnNodeOperators(uint256 ) view returns(address user, address peaqDid, string nodename, string ipaddress, string ispinfo, string region, string location, uint8 status)
func (_Node *NodeCaller) VpnNodeOperators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User      common.Address
	PeaqDid   common.Address
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "vpnNodeOperators", arg0)

	outstruct := new(struct {
		User      common.Address
		PeaqDid   common.Address
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
	outstruct.PeaqDid = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
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
// Solidity: function vpnNodeOperators(uint256 ) view returns(address user, address peaqDid, string nodename, string ipaddress, string ispinfo, string region, string location, uint8 status)
func (_Node *NodeSession) VpnNodeOperators(arg0 *big.Int) (struct {
	User      common.Address
	PeaqDid   common.Address
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}, error) {
	return _Node.Contract.VpnNodeOperators(&_Node.CallOpts, arg0)
}

// VpnNodeOperators is a free data retrieval call binding the contract method 0xd5399313.
//
// Solidity: function vpnNodeOperators(uint256 ) view returns(address user, address peaqDid, string nodename, string ipaddress, string ispinfo, string region, string location, uint8 status)
func (_Node *NodeCallerSession) VpnNodeOperators(arg0 *big.Int) (struct {
	User      common.Address
	PeaqDid   common.Address
	Nodename  string
	Ipaddress string
	Ispinfo   string
	Region    string
	Location  string
	Status    uint8
}, error) {
	return _Node.Contract.VpnNodeOperators(&_Node.CallOpts, arg0)
}

// VpnTotalCheckpoints is a free data retrieval call binding the contract method 0xef5050fb.
//
// Solidity: function vpnTotalCheckpoints(uint256 ) view returns(uint256)
func (_Node *NodeCaller) VpnTotalCheckpoints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "vpnTotalCheckpoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VpnTotalCheckpoints is a free data retrieval call binding the contract method 0xef5050fb.
//
// Solidity: function vpnTotalCheckpoints(uint256 ) view returns(uint256)
func (_Node *NodeSession) VpnTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Node.Contract.VpnTotalCheckpoints(&_Node.CallOpts, arg0)
}

// VpnTotalCheckpoints is a free data retrieval call binding the contract method 0xef5050fb.
//
// Solidity: function vpnTotalCheckpoints(uint256 ) view returns(uint256)
func (_Node *NodeCallerSession) VpnTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Node.Contract.VpnTotalCheckpoints(&_Node.CallOpts, arg0)
}

// WifiDeviceCheckpoints is a free data retrieval call binding the contract method 0x4f663d31.
//
// Solidity: function wifiDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Node *NodeCaller) WifiDeviceCheckpoints(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "wifiDeviceCheckpoints", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WifiDeviceCheckpoints is a free data retrieval call binding the contract method 0x4f663d31.
//
// Solidity: function wifiDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Node *NodeSession) WifiDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Node.Contract.WifiDeviceCheckpoints(&_Node.CallOpts, arg0, arg1)
}

// WifiDeviceCheckpoints is a free data retrieval call binding the contract method 0x4f663d31.
//
// Solidity: function wifiDeviceCheckpoints(uint256 , uint256 ) view returns(string)
func (_Node *NodeCallerSession) WifiDeviceCheckpoints(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _Node.Contract.WifiDeviceCheckpoints(&_Node.CallOpts, arg0, arg1)
}

// WifiNodeOperators is a free data retrieval call binding the contract method 0x8f9c8ee2.
//
// Solidity: function wifiNodeOperators(uint256 ) view returns(address user, string deviceId, address peaqDid, string ssid, string location, uint256 pricePerMinute, bool isActive)
func (_Node *NodeCaller) WifiNodeOperators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User           common.Address
	DeviceId       string
	PeaqDid        common.Address
	Ssid           string
	Location       string
	PricePerMinute *big.Int
	IsActive       bool
}, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "wifiNodeOperators", arg0)

	outstruct := new(struct {
		User           common.Address
		DeviceId       string
		PeaqDid        common.Address
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
	outstruct.PeaqDid = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Ssid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Location = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.PricePerMinute = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// WifiNodeOperators is a free data retrieval call binding the contract method 0x8f9c8ee2.
//
// Solidity: function wifiNodeOperators(uint256 ) view returns(address user, string deviceId, address peaqDid, string ssid, string location, uint256 pricePerMinute, bool isActive)
func (_Node *NodeSession) WifiNodeOperators(arg0 *big.Int) (struct {
	User           common.Address
	DeviceId       string
	PeaqDid        common.Address
	Ssid           string
	Location       string
	PricePerMinute *big.Int
	IsActive       bool
}, error) {
	return _Node.Contract.WifiNodeOperators(&_Node.CallOpts, arg0)
}

// WifiNodeOperators is a free data retrieval call binding the contract method 0x8f9c8ee2.
//
// Solidity: function wifiNodeOperators(uint256 ) view returns(address user, string deviceId, address peaqDid, string ssid, string location, uint256 pricePerMinute, bool isActive)
func (_Node *NodeCallerSession) WifiNodeOperators(arg0 *big.Int) (struct {
	User           common.Address
	DeviceId       string
	PeaqDid        common.Address
	Ssid           string
	Location       string
	PricePerMinute *big.Int
	IsActive       bool
}, error) {
	return _Node.Contract.WifiNodeOperators(&_Node.CallOpts, arg0)
}

// WifiTotalCheckpoints is a free data retrieval call binding the contract method 0x0608a6b6.
//
// Solidity: function wifiTotalCheckpoints(uint256 ) view returns(uint256)
func (_Node *NodeCaller) WifiTotalCheckpoints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "wifiTotalCheckpoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WifiTotalCheckpoints is a free data retrieval call binding the contract method 0x0608a6b6.
//
// Solidity: function wifiTotalCheckpoints(uint256 ) view returns(uint256)
func (_Node *NodeSession) WifiTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Node.Contract.WifiTotalCheckpoints(&_Node.CallOpts, arg0)
}

// WifiTotalCheckpoints is a free data retrieval call binding the contract method 0x0608a6b6.
//
// Solidity: function wifiTotalCheckpoints(uint256 ) view returns(uint256)
func (_Node *NodeCallerSession) WifiTotalCheckpoints(arg0 *big.Int) (*big.Int, error) {
	return _Node.Contract.WifiTotalCheckpoints(&_Node.CallOpts, arg0)
}

// AddAttribute is a paid mutator transaction binding the contract method 0xcc4a70ca.
//
// Solidity: function addAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Node *NodeTransactor) AddAttribute(opts *bind.TransactOpts, did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "addAttribute", did_account, name, value, validity_for)
}

// AddAttribute is a paid mutator transaction binding the contract method 0xcc4a70ca.
//
// Solidity: function addAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Node *NodeSession) AddAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Node.Contract.AddAttribute(&_Node.TransactOpts, did_account, name, value, validity_for)
}

// AddAttribute is a paid mutator transaction binding the contract method 0xcc4a70ca.
//
// Solidity: function addAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Node *NodeTransactorSession) AddAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Node.Contract.AddAttribute(&_Node.TransactOpts, did_account, name, value, validity_for)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xe9d63830.
//
// Solidity: function deactivateNode(uint256 nodeID) returns()
func (_Node *NodeTransactor) DeactivateNode(opts *bind.TransactOpts, nodeID *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "deactivateNode", nodeID)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xe9d63830.
//
// Solidity: function deactivateNode(uint256 nodeID) returns()
func (_Node *NodeSession) DeactivateNode(nodeID *big.Int) (*types.Transaction, error) {
	return _Node.Contract.DeactivateNode(&_Node.TransactOpts, nodeID)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xe9d63830.
//
// Solidity: function deactivateNode(uint256 nodeID) returns()
func (_Node *NodeTransactorSession) DeactivateNode(nodeID *big.Int) (*types.Transaction, error) {
	return _Node.Contract.DeactivateNode(&_Node.TransactOpts, nodeID)
}

// DelegateRegisterVpnNode is a paid mutator transaction binding the contract method 0xf424a2df.
//
// Solidity: function delegateRegisterVpnNode((address,address,string,string,string,string,string,uint8) node) returns()
func (_Node *NodeTransactor) DelegateRegisterVpnNode(opts *bind.TransactOpts, node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "delegateRegisterVpnNode", node)
}

// DelegateRegisterVpnNode is a paid mutator transaction binding the contract method 0xf424a2df.
//
// Solidity: function delegateRegisterVpnNode((address,address,string,string,string,string,string,uint8) node) returns()
func (_Node *NodeSession) DelegateRegisterVpnNode(node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Node.Contract.DelegateRegisterVpnNode(&_Node.TransactOpts, node)
}

// DelegateRegisterVpnNode is a paid mutator transaction binding the contract method 0xf424a2df.
//
// Solidity: function delegateRegisterVpnNode((address,address,string,string,string,string,string,uint8) node) returns()
func (_Node *NodeTransactorSession) DelegateRegisterVpnNode(node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Node.Contract.DelegateRegisterVpnNode(&_Node.TransactOpts, node)
}

// DelegateRegisterWifiNode is a paid mutator transaction binding the contract method 0x130fb7ec.
//
// Solidity: function delegateRegisterWifiNode(address user, string _deviceId, address _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Node *NodeTransactor) DelegateRegisterWifiNode(opts *bind.TransactOpts, user common.Address, _deviceId string, _peaqDid common.Address, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "delegateRegisterWifiNode", user, _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// DelegateRegisterWifiNode is a paid mutator transaction binding the contract method 0x130fb7ec.
//
// Solidity: function delegateRegisterWifiNode(address user, string _deviceId, address _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Node *NodeSession) DelegateRegisterWifiNode(user common.Address, _deviceId string, _peaqDid common.Address, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Node.Contract.DelegateRegisterWifiNode(&_Node.TransactOpts, user, _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// DelegateRegisterWifiNode is a paid mutator transaction binding the contract method 0x130fb7ec.
//
// Solidity: function delegateRegisterWifiNode(address user, string _deviceId, address _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Node *NodeTransactorSession) DelegateRegisterWifiNode(user common.Address, _deviceId string, _peaqDid common.Address, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Node.Contract.DelegateRegisterWifiNode(&_Node.TransactOpts, user, _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// RegisterVpnNode is a paid mutator transaction binding the contract method 0x786b0bce.
//
// Solidity: function registerVpnNode((address,address,string,string,string,string,string,uint8) node) returns()
func (_Node *NodeTransactor) RegisterVpnNode(opts *bind.TransactOpts, node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "registerVpnNode", node)
}

// RegisterVpnNode is a paid mutator transaction binding the contract method 0x786b0bce.
//
// Solidity: function registerVpnNode((address,address,string,string,string,string,string,uint8) node) returns()
func (_Node *NodeSession) RegisterVpnNode(node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Node.Contract.RegisterVpnNode(&_Node.TransactOpts, node)
}

// RegisterVpnNode is a paid mutator transaction binding the contract method 0x786b0bce.
//
// Solidity: function registerVpnNode((address,address,string,string,string,string,string,uint8) node) returns()
func (_Node *NodeTransactorSession) RegisterVpnNode(node ErebrusRegistryVPNNode) (*types.Transaction, error) {
	return _Node.Contract.RegisterVpnNode(&_Node.TransactOpts, node)
}

// RegisterWifiNode is a paid mutator transaction binding the contract method 0xbc493cfd.
//
// Solidity: function registerWifiNode(string _deviceId, address _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Node *NodeTransactor) RegisterWifiNode(opts *bind.TransactOpts, _deviceId string, _peaqDid common.Address, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "registerWifiNode", _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// RegisterWifiNode is a paid mutator transaction binding the contract method 0xbc493cfd.
//
// Solidity: function registerWifiNode(string _deviceId, address _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Node *NodeSession) RegisterWifiNode(_deviceId string, _peaqDid common.Address, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Node.Contract.RegisterWifiNode(&_Node.TransactOpts, _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// RegisterWifiNode is a paid mutator transaction binding the contract method 0xbc493cfd.
//
// Solidity: function registerWifiNode(string _deviceId, address _peaqDid, string _ssid, string _location, uint256 _pricePermin) returns()
func (_Node *NodeTransactorSession) RegisterWifiNode(_deviceId string, _peaqDid common.Address, _ssid string, _location string, _pricePermin *big.Int) (*types.Transaction, error) {
	return _Node.Contract.RegisterWifiNode(&_Node.TransactOpts, _deviceId, _peaqDid, _ssid, _location, _pricePermin)
}

// RemoveAttribute is a paid mutator transaction binding the contract method 0xe8a81690.
//
// Solidity: function removeAttribute(address did_account, bytes name) returns(bool)
func (_Node *NodeTransactor) RemoveAttribute(opts *bind.TransactOpts, did_account common.Address, name []byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "removeAttribute", did_account, name)
}

// RemoveAttribute is a paid mutator transaction binding the contract method 0xe8a81690.
//
// Solidity: function removeAttribute(address did_account, bytes name) returns(bool)
func (_Node *NodeSession) RemoveAttribute(did_account common.Address, name []byte) (*types.Transaction, error) {
	return _Node.Contract.RemoveAttribute(&_Node.TransactOpts, did_account, name)
}

// RemoveAttribute is a paid mutator transaction binding the contract method 0xe8a81690.
//
// Solidity: function removeAttribute(address did_account, bytes name) returns(bool)
func (_Node *NodeTransactorSession) RemoveAttribute(did_account common.Address, name []byte) (*types.Transaction, error) {
	return _Node.Contract.RemoveAttribute(&_Node.TransactOpts, did_account, name)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x68b4b2c1.
//
// Solidity: function updateAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Node *NodeTransactor) UpdateAttribute(opts *bind.TransactOpts, did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "updateAttribute", did_account, name, value, validity_for)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x68b4b2c1.
//
// Solidity: function updateAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Node *NodeSession) UpdateAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Node.Contract.UpdateAttribute(&_Node.TransactOpts, did_account, name, value, validity_for)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x68b4b2c1.
//
// Solidity: function updateAttribute(address did_account, bytes name, bytes value, uint32 validity_for) returns(bool)
func (_Node *NodeTransactorSession) UpdateAttribute(did_account common.Address, name []byte, value []byte, validity_for uint32) (*types.Transaction, error) {
	return _Node.Contract.UpdateAttribute(&_Node.TransactOpts, did_account, name, value, validity_for)
}

// UpdateVPNNode is a paid mutator transaction binding the contract method 0xe48da8a4.
//
// Solidity: function updateVPNNode(uint256 nodeID, uint8 _status, string _region) returns()
func (_Node *NodeTransactor) UpdateVPNNode(opts *bind.TransactOpts, nodeID *big.Int, _status uint8, _region string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "updateVPNNode", nodeID, _status, _region)
}

// UpdateVPNNode is a paid mutator transaction binding the contract method 0xe48da8a4.
//
// Solidity: function updateVPNNode(uint256 nodeID, uint8 _status, string _region) returns()
func (_Node *NodeSession) UpdateVPNNode(nodeID *big.Int, _status uint8, _region string) (*types.Transaction, error) {
	return _Node.Contract.UpdateVPNNode(&_Node.TransactOpts, nodeID, _status, _region)
}

// UpdateVPNNode is a paid mutator transaction binding the contract method 0xe48da8a4.
//
// Solidity: function updateVPNNode(uint256 nodeID, uint8 _status, string _region) returns()
func (_Node *NodeTransactorSession) UpdateVPNNode(nodeID *big.Int, _status uint8, _region string) (*types.Transaction, error) {
	return _Node.Contract.UpdateVPNNode(&_Node.TransactOpts, nodeID, _status, _region)
}

// UpdateWiFiNode is a paid mutator transaction binding the contract method 0x8a7692dc.
//
// Solidity: function updateWiFiNode(uint256 nodeID, string ssid, string location, uint256 pricePerMin) returns()
func (_Node *NodeTransactor) UpdateWiFiNode(opts *bind.TransactOpts, nodeID *big.Int, ssid string, location string, pricePerMin *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "updateWiFiNode", nodeID, ssid, location, pricePerMin)
}

// UpdateWiFiNode is a paid mutator transaction binding the contract method 0x8a7692dc.
//
// Solidity: function updateWiFiNode(uint256 nodeID, string ssid, string location, uint256 pricePerMin) returns()
func (_Node *NodeSession) UpdateWiFiNode(nodeID *big.Int, ssid string, location string, pricePerMin *big.Int) (*types.Transaction, error) {
	return _Node.Contract.UpdateWiFiNode(&_Node.TransactOpts, nodeID, ssid, location, pricePerMin)
}

// UpdateWiFiNode is a paid mutator transaction binding the contract method 0x8a7692dc.
//
// Solidity: function updateWiFiNode(uint256 nodeID, string ssid, string location, uint256 pricePerMin) returns()
func (_Node *NodeTransactorSession) UpdateWiFiNode(nodeID *big.Int, ssid string, location string, pricePerMin *big.Int) (*types.Transaction, error) {
	return _Node.Contract.UpdateWiFiNode(&_Node.TransactOpts, nodeID, ssid, location, pricePerMin)
}

// VpnDeviceCheckpoint is a paid mutator transaction binding the contract method 0x1b9f0c95.
//
// Solidity: function vpnDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Node *NodeTransactor) VpnDeviceCheckpoint(opts *bind.TransactOpts, nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "vpnDeviceCheckpoint", nodeID, dataHash)
}

// VpnDeviceCheckpoint is a paid mutator transaction binding the contract method 0x1b9f0c95.
//
// Solidity: function vpnDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Node *NodeSession) VpnDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Node.Contract.VpnDeviceCheckpoint(&_Node.TransactOpts, nodeID, dataHash)
}

// VpnDeviceCheckpoint is a paid mutator transaction binding the contract method 0x1b9f0c95.
//
// Solidity: function vpnDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Node *NodeTransactorSession) VpnDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Node.Contract.VpnDeviceCheckpoint(&_Node.TransactOpts, nodeID, dataHash)
}

// WifiDeviceCheckpoint is a paid mutator transaction binding the contract method 0x63c3a7e7.
//
// Solidity: function wifiDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Node *NodeTransactor) WifiDeviceCheckpoint(opts *bind.TransactOpts, nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "wifiDeviceCheckpoint", nodeID, dataHash)
}

// WifiDeviceCheckpoint is a paid mutator transaction binding the contract method 0x63c3a7e7.
//
// Solidity: function wifiDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Node *NodeSession) WifiDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Node.Contract.WifiDeviceCheckpoint(&_Node.TransactOpts, nodeID, dataHash)
}

// WifiDeviceCheckpoint is a paid mutator transaction binding the contract method 0x63c3a7e7.
//
// Solidity: function wifiDeviceCheckpoint(uint256 nodeID, string dataHash) returns()
func (_Node *NodeTransactorSession) WifiDeviceCheckpoint(nodeID *big.Int, dataHash string) (*types.Transaction, error) {
	return _Node.Contract.WifiDeviceCheckpoint(&_Node.TransactOpts, nodeID, dataHash)
}

// NodeAddAttributeIterator is returned from FilterAddAttribute and is used to iterate over the raw logs and unpacked data for AddAttribute events raised by the Node contract.
type NodeAddAttributeIterator struct {
	Event *NodeAddAttribute // Event containing the contract specifics and raw log

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
func (it *NodeAddAttributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeAddAttribute)
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
		it.Event = new(NodeAddAttribute)
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
func (it *NodeAddAttributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeAddAttributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeAddAttribute represents a AddAttribute event raised by the Node contract.
type NodeAddAttribute struct {
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
func (_Node *NodeFilterer) FilterAddAttribute(opts *bind.FilterOpts) (*NodeAddAttributeIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "AddAttribute")
	if err != nil {
		return nil, err
	}
	return &NodeAddAttributeIterator{contract: _Node.contract, event: "AddAttribute", logs: logs, sub: sub}, nil
}

// WatchAddAttribute is a free log subscription operation binding the contract event 0x13aef52bc4a99da04591533072e304017e3fb76f43e7fadd25eb7f514c5ef6e5.
//
// Solidity: event AddAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Node *NodeFilterer) WatchAddAttribute(opts *bind.WatchOpts, sink chan<- *NodeAddAttribute) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "AddAttribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeAddAttribute)
				if err := _Node.contract.UnpackLog(event, "AddAttribute", log); err != nil {
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
func (_Node *NodeFilterer) ParseAddAttribute(log types.Log) (*NodeAddAttribute, error) {
	event := new(NodeAddAttribute)
	if err := _Node.contract.UnpackLog(event, "AddAttribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the Node contract.
type NodeNodeDeactivatedIterator struct {
	Event *NodeNodeDeactivated // Event containing the contract specifics and raw log

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
func (it *NodeNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeNodeDeactivated)
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
		it.Event = new(NodeNodeDeactivated)
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
func (it *NodeNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeNodeDeactivated represents a NodeDeactivated event raised by the Node contract.
type NodeNodeDeactivated struct {
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xd9957750e6343405c319eb99a4ec67fa11cfd66969318cbc71aa2d45fa53a349.
//
// Solidity: event NodeDeactivated(address indexed operatorAddress)
func (_Node *NodeFilterer) FilterNodeDeactivated(opts *bind.FilterOpts, operatorAddress []common.Address) (*NodeNodeDeactivatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "NodeDeactivated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeNodeDeactivatedIterator{contract: _Node.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xd9957750e6343405c319eb99a4ec67fa11cfd66969318cbc71aa2d45fa53a349.
//
// Solidity: event NodeDeactivated(address indexed operatorAddress)
func (_Node *NodeFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *NodeNodeDeactivated, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "NodeDeactivated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeNodeDeactivated)
				if err := _Node.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
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
func (_Node *NodeFilterer) ParseNodeDeactivated(log types.Log) (*NodeNodeDeactivated, error) {
	event := new(NodeNodeDeactivated)
	if err := _Node.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeNodeOperatorUpdatedIterator is returned from FilterNodeOperatorUpdated and is used to iterate over the raw logs and unpacked data for NodeOperatorUpdated events raised by the Node contract.
type NodeNodeOperatorUpdatedIterator struct {
	Event *NodeNodeOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *NodeNodeOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeNodeOperatorUpdated)
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
		it.Event = new(NodeNodeOperatorUpdated)
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
func (it *NodeNodeOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeNodeOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeNodeOperatorUpdated represents a NodeOperatorUpdated event raised by the Node contract.
type NodeNodeOperatorUpdated struct {
	OperatorAddress common.Address
	Ssid            string
	Location        string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorUpdated is a free log retrieval operation binding the contract event 0xf512e0b8ea262571de058d8b17a3c2d29913f35abcd4c173935ba6fa58b3833b.
//
// Solidity: event NodeOperatorUpdated(address indexed operatorAddress, string ssid, string location)
func (_Node *NodeFilterer) FilterNodeOperatorUpdated(opts *bind.FilterOpts, operatorAddress []common.Address) (*NodeNodeOperatorUpdatedIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "NodeOperatorUpdated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeNodeOperatorUpdatedIterator{contract: _Node.contract, event: "NodeOperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorUpdated is a free log subscription operation binding the contract event 0xf512e0b8ea262571de058d8b17a3c2d29913f35abcd4c173935ba6fa58b3833b.
//
// Solidity: event NodeOperatorUpdated(address indexed operatorAddress, string ssid, string location)
func (_Node *NodeFilterer) WatchNodeOperatorUpdated(opts *bind.WatchOpts, sink chan<- *NodeNodeOperatorUpdated, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "NodeOperatorUpdated", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeNodeOperatorUpdated)
				if err := _Node.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
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
func (_Node *NodeFilterer) ParseNodeOperatorUpdated(log types.Log) (*NodeNodeOperatorUpdated, error) {
	event := new(NodeNodeOperatorUpdated)
	if err := _Node.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRemoveAttribteIterator is returned from FilterRemoveAttribte and is used to iterate over the raw logs and unpacked data for RemoveAttribte events raised by the Node contract.
type NodeRemoveAttribteIterator struct {
	Event *NodeRemoveAttribte // Event containing the contract specifics and raw log

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
func (it *NodeRemoveAttribteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRemoveAttribte)
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
		it.Event = new(NodeRemoveAttribte)
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
func (it *NodeRemoveAttribteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRemoveAttribteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRemoveAttribte represents a RemoveAttribte event raised by the Node contract.
type NodeRemoveAttribte struct {
	DidAccount common.Address
	Name       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveAttribte is a free log retrieval operation binding the contract event 0x647d95b9099b095e5957558106bc51df77720f963539a0ca62ee9948e6554da1.
//
// Solidity: event RemoveAttribte(address did_account, bytes name)
func (_Node *NodeFilterer) FilterRemoveAttribte(opts *bind.FilterOpts) (*NodeRemoveAttribteIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "RemoveAttribte")
	if err != nil {
		return nil, err
	}
	return &NodeRemoveAttribteIterator{contract: _Node.contract, event: "RemoveAttribte", logs: logs, sub: sub}, nil
}

// WatchRemoveAttribte is a free log subscription operation binding the contract event 0x647d95b9099b095e5957558106bc51df77720f963539a0ca62ee9948e6554da1.
//
// Solidity: event RemoveAttribte(address did_account, bytes name)
func (_Node *NodeFilterer) WatchRemoveAttribte(opts *bind.WatchOpts, sink chan<- *NodeRemoveAttribte) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "RemoveAttribte")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRemoveAttribte)
				if err := _Node.contract.UnpackLog(event, "RemoveAttribte", log); err != nil {
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
func (_Node *NodeFilterer) ParseRemoveAttribte(log types.Log) (*NodeRemoveAttribte, error) {
	event := new(NodeRemoveAttribte)
	if err := _Node.contract.UnpackLog(event, "RemoveAttribte", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeUpdateAttributeIterator is returned from FilterUpdateAttribute and is used to iterate over the raw logs and unpacked data for UpdateAttribute events raised by the Node contract.
type NodeUpdateAttributeIterator struct {
	Event *NodeUpdateAttribute // Event containing the contract specifics and raw log

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
func (it *NodeUpdateAttributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeUpdateAttribute)
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
		it.Event = new(NodeUpdateAttribute)
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
func (it *NodeUpdateAttributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeUpdateAttributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeUpdateAttribute represents a UpdateAttribute event raised by the Node contract.
type NodeUpdateAttribute struct {
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
func (_Node *NodeFilterer) FilterUpdateAttribute(opts *bind.FilterOpts) (*NodeUpdateAttributeIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "UpdateAttribute")
	if err != nil {
		return nil, err
	}
	return &NodeUpdateAttributeIterator{contract: _Node.contract, event: "UpdateAttribute", logs: logs, sub: sub}, nil
}

// WatchUpdateAttribute is a free log subscription operation binding the contract event 0x018be4d5e2634aefdb51c4ddd186f33072cfb5f6baf685579c2e214995dc9e4f.
//
// Solidity: event UpdateAttribute(address sender, address did_account, bytes name, bytes value, uint32 validity)
func (_Node *NodeFilterer) WatchUpdateAttribute(opts *bind.WatchOpts, sink chan<- *NodeUpdateAttribute) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "UpdateAttribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeUpdateAttribute)
				if err := _Node.contract.UnpackLog(event, "UpdateAttribute", log); err != nil {
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
func (_Node *NodeFilterer) ParseUpdateAttribute(log types.Log) (*NodeUpdateAttribute, error) {
	event := new(NodeUpdateAttribute)
	if err := _Node.contract.UnpackLog(event, "UpdateAttribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeVPNUpdatedIterator is returned from FilterVPNUpdated and is used to iterate over the raw logs and unpacked data for VPNUpdated events raised by the Node contract.
type NodeVPNUpdatedIterator struct {
	Event *NodeVPNUpdated // Event containing the contract specifics and raw log

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
func (it *NodeVPNUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeVPNUpdated)
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
		it.Event = new(NodeVPNUpdated)
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
func (it *NodeVPNUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeVPNUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeVPNUpdated represents a VPNUpdated event raised by the Node contract.
type NodeVPNUpdated struct {
	NodeId        *big.Int
	UpdatedStatus uint8
	UpdatedRegion string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVPNUpdated is a free log retrieval operation binding the contract event 0x34349e39780aaa7fc3b8b465ca9a2ada7a8a0a941e269b03102ad3867ff34b66.
//
// Solidity: event VPNUpdated(uint256 nodeId, uint8 updatedStatus, string updatedRegion)
func (_Node *NodeFilterer) FilterVPNUpdated(opts *bind.FilterOpts) (*NodeVPNUpdatedIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "VPNUpdated")
	if err != nil {
		return nil, err
	}
	return &NodeVPNUpdatedIterator{contract: _Node.contract, event: "VPNUpdated", logs: logs, sub: sub}, nil
}

// WatchVPNUpdated is a free log subscription operation binding the contract event 0x34349e39780aaa7fc3b8b465ca9a2ada7a8a0a941e269b03102ad3867ff34b66.
//
// Solidity: event VPNUpdated(uint256 nodeId, uint8 updatedStatus, string updatedRegion)
func (_Node *NodeFilterer) WatchVPNUpdated(opts *bind.WatchOpts, sink chan<- *NodeVPNUpdated) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "VPNUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeVPNUpdated)
				if err := _Node.contract.UnpackLog(event, "VPNUpdated", log); err != nil {
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
func (_Node *NodeFilterer) ParseVPNUpdated(log types.Log) (*NodeVPNUpdated, error) {
	event := new(NodeVPNUpdated)
	if err := _Node.contract.UnpackLog(event, "VPNUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeVpnNodeRegisteredIterator is returned from FilterVpnNodeRegistered and is used to iterate over the raw logs and unpacked data for VpnNodeRegistered events raised by the Node contract.
type NodeVpnNodeRegisteredIterator struct {
	Event *NodeVpnNodeRegistered // Event containing the contract specifics and raw log

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
func (it *NodeVpnNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeVpnNodeRegistered)
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
		it.Event = new(NodeVpnNodeRegistered)
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
func (it *NodeVpnNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeVpnNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeVpnNodeRegistered represents a VpnNodeRegistered event raised by the Node contract.
type NodeVpnNodeRegistered struct {
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
func (_Node *NodeFilterer) FilterVpnNodeRegistered(opts *bind.FilterOpts) (*NodeVpnNodeRegisteredIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "VpnNodeRegistered")
	if err != nil {
		return nil, err
	}
	return &NodeVpnNodeRegisteredIterator{contract: _Node.contract, event: "VpnNodeRegistered", logs: logs, sub: sub}, nil
}

// WatchVpnNodeRegistered is a free log subscription operation binding the contract event 0x9e6fb352d55d4e4414d6ec189c261105e99d8baba96bd55c43a2013e2ee9453b.
//
// Solidity: event VpnNodeRegistered(uint256 nodeId, string nodename, string ipaddress, string ispinfo, string region, string location)
func (_Node *NodeFilterer) WatchVpnNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodeVpnNodeRegistered) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "VpnNodeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeVpnNodeRegistered)
				if err := _Node.contract.UnpackLog(event, "VpnNodeRegistered", log); err != nil {
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
func (_Node *NodeFilterer) ParseVpnNodeRegistered(log types.Log) (*NodeVpnNodeRegistered, error) {
	event := new(NodeVpnNodeRegistered)
	if err := _Node.contract.UnpackLog(event, "VpnNodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeWifiNodeOperatorRegisteredIterator is returned from FilterWifiNodeOperatorRegistered and is used to iterate over the raw logs and unpacked data for WifiNodeOperatorRegistered events raised by the Node contract.
type NodeWifiNodeOperatorRegisteredIterator struct {
	Event *NodeWifiNodeOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *NodeWifiNodeOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeWifiNodeOperatorRegistered)
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
		it.Event = new(NodeWifiNodeOperatorRegistered)
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
func (it *NodeWifiNodeOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeWifiNodeOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeWifiNodeOperatorRegistered represents a WifiNodeOperatorRegistered event raised by the Node contract.
type NodeWifiNodeOperatorRegistered struct {
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
func (_Node *NodeFilterer) FilterWifiNodeOperatorRegistered(opts *bind.FilterOpts, owner []common.Address) (*NodeWifiNodeOperatorRegisteredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "WifiNodeOperatorRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return &NodeWifiNodeOperatorRegisteredIterator{contract: _Node.contract, event: "WifiNodeOperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchWifiNodeOperatorRegistered is a free log subscription operation binding the contract event 0x72ab8f8888d362028a5ec40836cefca72088b92b1ef7afc7fb795556b02d6658.
//
// Solidity: event WifiNodeOperatorRegistered(uint256 nodeID, address indexed owner, string deviceId, string ssid, string location, uint256 pricePerMinute)
func (_Node *NodeFilterer) WatchWifiNodeOperatorRegistered(opts *bind.WatchOpts, sink chan<- *NodeWifiNodeOperatorRegistered, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "WifiNodeOperatorRegistered", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeWifiNodeOperatorRegistered)
				if err := _Node.contract.UnpackLog(event, "WifiNodeOperatorRegistered", log); err != nil {
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
func (_Node *NodeFilterer) ParseWifiNodeOperatorRegistered(log types.Log) (*NodeWifiNodeOperatorRegistered, error) {
	event := new(NodeWifiNodeOperatorRegistered)
	if err := _Node.contract.UnpackLog(event, "WifiNodeOperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
