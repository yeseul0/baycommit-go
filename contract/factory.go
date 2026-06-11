// Code generated from Factory.sol ABI - DO NOT EDIT manually.
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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: `[
		{
			"anonymous": false,
			"inputs": [
				{"indexed": true, "internalType": "address", "name": "proxyAddress", "type": "address"}
			],
			"name": "ProxyCreated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{"indexed": false, "internalType": "string", "name": "message", "type": "string"},
				{"indexed": false, "internalType": "bool",   "name": "success", "type": "bool"}
			],
			"name": "DebugLog",
			"type": "event"
		},
		{
			"inputs": [
				{"internalType": "string",  "name": "_studyName",      "type": "string"},
				{"internalType": "uint256", "name": "_depositAmount",   "type": "uint256"},
				{"internalType": "uint256", "name": "_penaltyAmount",   "type": "uint256"},
				{"internalType": "address", "name": "_studyAdmin",      "type": "address"},
				{"internalType": "uint256", "name": "_studyStartTime",  "type": "uint256"},
				{"internalType": "uint256", "name": "_studyEndTime",    "type": "uint256"}
			],
			"name": "createProxy",
			"outputs": [{"internalType": "address", "name": "", "type": "address"}],
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`,
}

// FactoryABI is the input ABI used to generate the binding from.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller
	FactoryTransactor
	FactoryFilterer
}

type FactoryCaller struct {
	contract *bind.BoundContract
}

type FactoryTransactor struct {
	contract *bind.BoundContract
}

type FactoryFilterer struct {
	contract *bind.BoundContract
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, *parsed, backend, backend, backend)
	return &Factory{
		FactoryCaller:     FactoryCaller{contract: contract},
		FactoryTransactor: FactoryTransactor{contract: contract},
		FactoryFilterer:   FactoryFilterer{contract: contract},
	}, nil
}

// CreateProxy invokes the createProxy Solidity function.
func (_Factory *FactoryTransactor) CreateProxy(
	opts *bind.TransactOpts,
	studyName string,
	depositAmount *big.Int,
	penaltyAmount *big.Int,
	studyAdmin common.Address,
	studyStartTime *big.Int,
	studyEndTime *big.Int,
) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createProxy",
		studyName, depositAmount, penaltyAmount, studyAdmin, studyStartTime, studyEndTime)
}

// FactoryProxyCreated represents a ProxyCreated event raised by the Factory contract.
type FactoryProxyCreated struct {
	ProxyAddress common.Address
	Raw          types.Log
}

// ParseProxyCreated parses the ProxyCreated event from a transaction receipt log.
func (_Factory *FactoryFilterer) ParseProxyCreated(log types.Log) (*FactoryProxyCreated, error) {
	event := new(FactoryProxyCreated)
	if err := _Factory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
