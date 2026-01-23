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

// StudyGroupMetaData contains all meta data concerning the StudyGroup contract.
var StudyGroupMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commitTime\",\"type\":\"uint256\"}],\"name\":\"CommitTracked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"ParticipantJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"ParticipantLeft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"penalizedParticipants\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penaltyAmount\",\"type\":\"uint256\"}],\"name\":\"StudyClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRefunded\",\"type\":\"uint256\"}],\"name\":\"StudyGroupTerminated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StudyStarted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"closeStudy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getAllCommitTimes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"participantList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"commitTimes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"getCommitTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getStudyDayInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"participantCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_studyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_penaltyAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_studyStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_studyEndTime\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"isParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStudyClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinStudy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveStudy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"startTodayStudy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"studyDays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"studyEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"studyName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"studyStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"terminateStudyGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commitTime\",\"type\":\"uint256\"}],\"name\":\"trackCommit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StudyGroupABI is the input ABI used to generate the binding from.
// Deprecated: Use StudyGroupMetaData.ABI instead.
var StudyGroupABI = StudyGroupMetaData.ABI

// StudyGroup is an auto generated Go binding around an Ethereum contract.
type StudyGroup struct {
	StudyGroupCaller     // Read-only binding to the contract
	StudyGroupTransactor // Write-only binding to the contract
	StudyGroupFilterer   // Log filterer for contract events
}

// StudyGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type StudyGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StudyGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StudyGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StudyGroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StudyGroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StudyGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StudyGroupSession struct {
	Contract     *StudyGroup       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StudyGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StudyGroupCallerSession struct {
	Contract *StudyGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StudyGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StudyGroupTransactorSession struct {
	Contract     *StudyGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StudyGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type StudyGroupRaw struct {
	Contract *StudyGroup // Generic contract binding to access the raw methods on
}

// StudyGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StudyGroupCallerRaw struct {
	Contract *StudyGroupCaller // Generic read-only contract binding to access the raw methods on
}

// StudyGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StudyGroupTransactorRaw struct {
	Contract *StudyGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStudyGroup creates a new instance of StudyGroup, bound to a specific deployed contract.
func NewStudyGroup(address common.Address, backend bind.ContractBackend) (*StudyGroup, error) {
	contract, err := bindStudyGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StudyGroup{StudyGroupCaller: StudyGroupCaller{contract: contract}, StudyGroupTransactor: StudyGroupTransactor{contract: contract}, StudyGroupFilterer: StudyGroupFilterer{contract: contract}}, nil
}

// NewStudyGroupCaller creates a new read-only instance of StudyGroup, bound to a specific deployed contract.
func NewStudyGroupCaller(address common.Address, caller bind.ContractCaller) (*StudyGroupCaller, error) {
	contract, err := bindStudyGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StudyGroupCaller{contract: contract}, nil
}

// NewStudyGroupTransactor creates a new write-only instance of StudyGroup, bound to a specific deployed contract.
func NewStudyGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*StudyGroupTransactor, error) {
	contract, err := bindStudyGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StudyGroupTransactor{contract: contract}, nil
}

// NewStudyGroupFilterer creates a new log filterer instance of StudyGroup, bound to a specific deployed contract.
func NewStudyGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*StudyGroupFilterer, error) {
	contract, err := bindStudyGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StudyGroupFilterer{contract: contract}, nil
}

// bindStudyGroup binds a generic wrapper to an already deployed contract.
func bindStudyGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StudyGroupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StudyGroup *StudyGroupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StudyGroup.Contract.StudyGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StudyGroup *StudyGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StudyGroup.Contract.StudyGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StudyGroup *StudyGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StudyGroup.Contract.StudyGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StudyGroup *StudyGroupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StudyGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StudyGroup *StudyGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StudyGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StudyGroup *StudyGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StudyGroup.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StudyGroup *StudyGroupCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StudyGroup *StudyGroupSession) Admin() (common.Address, error) {
	return _StudyGroup.Contract.Admin(&_StudyGroup.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StudyGroup *StudyGroupCallerSession) Admin() (common.Address, error) {
	return _StudyGroup.Contract.Admin(&_StudyGroup.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StudyGroup *StudyGroupCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StudyGroup *StudyGroupSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _StudyGroup.Contract.Balances(&_StudyGroup.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StudyGroup *StudyGroupCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _StudyGroup.Contract.Balances(&_StudyGroup.CallOpts, arg0)
}

// DepositAmount is a free data retrieval call binding the contract method 0x419759f5.
//
// Solidity: function depositAmount() view returns(uint256)
func (_StudyGroup *StudyGroupCaller) DepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "depositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositAmount is a free data retrieval call binding the contract method 0x419759f5.
//
// Solidity: function depositAmount() view returns(uint256)
func (_StudyGroup *StudyGroupSession) DepositAmount() (*big.Int, error) {
	return _StudyGroup.Contract.DepositAmount(&_StudyGroup.CallOpts)
}

// DepositAmount is a free data retrieval call binding the contract method 0x419759f5.
//
// Solidity: function depositAmount() view returns(uint256)
func (_StudyGroup *StudyGroupCallerSession) DepositAmount() (*big.Int, error) {
	return _StudyGroup.Contract.DepositAmount(&_StudyGroup.CallOpts)
}

// GetAllCommitTimes is a free data retrieval call binding the contract method 0x7dac3460.
//
// Solidity: function getAllCommitTimes(uint256 timestamp) view returns(address[] participantList, uint256[] commitTimes)
func (_StudyGroup *StudyGroupCaller) GetAllCommitTimes(opts *bind.CallOpts, timestamp *big.Int) (struct {
	ParticipantList []common.Address
	CommitTimes     []*big.Int
}, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "getAllCommitTimes", timestamp)

	outstruct := new(struct {
		ParticipantList []common.Address
		CommitTimes     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParticipantList = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CommitTimes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAllCommitTimes is a free data retrieval call binding the contract method 0x7dac3460.
//
// Solidity: function getAllCommitTimes(uint256 timestamp) view returns(address[] participantList, uint256[] commitTimes)
func (_StudyGroup *StudyGroupSession) GetAllCommitTimes(timestamp *big.Int) (struct {
	ParticipantList []common.Address
	CommitTimes     []*big.Int
}, error) {
	return _StudyGroup.Contract.GetAllCommitTimes(&_StudyGroup.CallOpts, timestamp)
}

// GetAllCommitTimes is a free data retrieval call binding the contract method 0x7dac3460.
//
// Solidity: function getAllCommitTimes(uint256 timestamp) view returns(address[] participantList, uint256[] commitTimes)
func (_StudyGroup *StudyGroupCallerSession) GetAllCommitTimes(timestamp *big.Int) (struct {
	ParticipantList []common.Address
	CommitTimes     []*big.Int
}, error) {
	return _StudyGroup.Contract.GetAllCommitTimes(&_StudyGroup.CallOpts, timestamp)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address participant) view returns(uint256)
func (_StudyGroup *StudyGroupCaller) GetBalance(opts *bind.CallOpts, participant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "getBalance", participant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address participant) view returns(uint256)
func (_StudyGroup *StudyGroupSession) GetBalance(participant common.Address) (*big.Int, error) {
	return _StudyGroup.Contract.GetBalance(&_StudyGroup.CallOpts, participant)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address participant) view returns(uint256)
func (_StudyGroup *StudyGroupCallerSession) GetBalance(participant common.Address) (*big.Int, error) {
	return _StudyGroup.Contract.GetBalance(&_StudyGroup.CallOpts, participant)
}

// GetCommitTime is a free data retrieval call binding the contract method 0x8f3708e8.
//
// Solidity: function getCommitTime(uint256 timestamp, address participant) view returns(uint256)
func (_StudyGroup *StudyGroupCaller) GetCommitTime(opts *bind.CallOpts, timestamp *big.Int, participant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "getCommitTime", timestamp, participant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitTime is a free data retrieval call binding the contract method 0x8f3708e8.
//
// Solidity: function getCommitTime(uint256 timestamp, address participant) view returns(uint256)
func (_StudyGroup *StudyGroupSession) GetCommitTime(timestamp *big.Int, participant common.Address) (*big.Int, error) {
	return _StudyGroup.Contract.GetCommitTime(&_StudyGroup.CallOpts, timestamp, participant)
}

// GetCommitTime is a free data retrieval call binding the contract method 0x8f3708e8.
//
// Solidity: function getCommitTime(uint256 timestamp, address participant) view returns(uint256)
func (_StudyGroup *StudyGroupCallerSession) GetCommitTime(timestamp *big.Int, participant common.Address) (*big.Int, error) {
	return _StudyGroup.Contract.GetCommitTime(&_StudyGroup.CallOpts, timestamp, participant)
}

// GetStudyDayInfo is a free data retrieval call binding the contract method 0x31a8911a.
//
// Solidity: function getStudyDayInfo(uint256 timestamp) view returns(uint256 participantCount, bool isClosed)
func (_StudyGroup *StudyGroupCaller) GetStudyDayInfo(opts *bind.CallOpts, timestamp *big.Int) (struct {
	ParticipantCount *big.Int
	IsClosed         bool
}, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "getStudyDayInfo", timestamp)

	outstruct := new(struct {
		ParticipantCount *big.Int
		IsClosed         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParticipantCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsClosed = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetStudyDayInfo is a free data retrieval call binding the contract method 0x31a8911a.
//
// Solidity: function getStudyDayInfo(uint256 timestamp) view returns(uint256 participantCount, bool isClosed)
func (_StudyGroup *StudyGroupSession) GetStudyDayInfo(timestamp *big.Int) (struct {
	ParticipantCount *big.Int
	IsClosed         bool
}, error) {
	return _StudyGroup.Contract.GetStudyDayInfo(&_StudyGroup.CallOpts, timestamp)
}

// GetStudyDayInfo is a free data retrieval call binding the contract method 0x31a8911a.
//
// Solidity: function getStudyDayInfo(uint256 timestamp) view returns(uint256 participantCount, bool isClosed)
func (_StudyGroup *StudyGroupCallerSession) GetStudyDayInfo(timestamp *big.Int) (struct {
	ParticipantCount *big.Int
	IsClosed         bool
}, error) {
	return _StudyGroup.Contract.GetStudyDayInfo(&_StudyGroup.CallOpts, timestamp)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address _participant) view returns(bool)
func (_StudyGroup *StudyGroupCaller) IsParticipant(opts *bind.CallOpts, _participant common.Address) (bool, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "isParticipant", _participant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address _participant) view returns(bool)
func (_StudyGroup *StudyGroupSession) IsParticipant(_participant common.Address) (bool, error) {
	return _StudyGroup.Contract.IsParticipant(&_StudyGroup.CallOpts, _participant)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address _participant) view returns(bool)
func (_StudyGroup *StudyGroupCallerSession) IsParticipant(_participant common.Address) (bool, error) {
	return _StudyGroup.Contract.IsParticipant(&_StudyGroup.CallOpts, _participant)
}

// IsStudyClosed is a free data retrieval call binding the contract method 0x926f73a3.
//
// Solidity: function isStudyClosed() view returns(bool)
func (_StudyGroup *StudyGroupCaller) IsStudyClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "isStudyClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStudyClosed is a free data retrieval call binding the contract method 0x926f73a3.
//
// Solidity: function isStudyClosed() view returns(bool)
func (_StudyGroup *StudyGroupSession) IsStudyClosed() (bool, error) {
	return _StudyGroup.Contract.IsStudyClosed(&_StudyGroup.CallOpts)
}

// IsStudyClosed is a free data retrieval call binding the contract method 0x926f73a3.
//
// Solidity: function isStudyClosed() view returns(bool)
func (_StudyGroup *StudyGroupCallerSession) IsStudyClosed() (bool, error) {
	return _StudyGroup.Contract.IsStudyClosed(&_StudyGroup.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_StudyGroup *StudyGroupCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_StudyGroup *StudyGroupSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _StudyGroup.Contract.Participants(&_StudyGroup.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_StudyGroup *StudyGroupCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _StudyGroup.Contract.Participants(&_StudyGroup.CallOpts, arg0)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_StudyGroup *StudyGroupCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_StudyGroup *StudyGroupSession) PaymentToken() (common.Address, error) {
	return _StudyGroup.Contract.PaymentToken(&_StudyGroup.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_StudyGroup *StudyGroupCallerSession) PaymentToken() (common.Address, error) {
	return _StudyGroup.Contract.PaymentToken(&_StudyGroup.CallOpts)
}

// PenaltyAmount is a free data retrieval call binding the contract method 0xf218f48c.
//
// Solidity: function penaltyAmount() view returns(uint256)
func (_StudyGroup *StudyGroupCaller) PenaltyAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "penaltyAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyAmount is a free data retrieval call binding the contract method 0xf218f48c.
//
// Solidity: function penaltyAmount() view returns(uint256)
func (_StudyGroup *StudyGroupSession) PenaltyAmount() (*big.Int, error) {
	return _StudyGroup.Contract.PenaltyAmount(&_StudyGroup.CallOpts)
}

// PenaltyAmount is a free data retrieval call binding the contract method 0xf218f48c.
//
// Solidity: function penaltyAmount() view returns(uint256)
func (_StudyGroup *StudyGroupCallerSession) PenaltyAmount() (*big.Int, error) {
	return _StudyGroup.Contract.PenaltyAmount(&_StudyGroup.CallOpts)
}

// StudyDays is a free data retrieval call binding the contract method 0x1a205a64.
//
// Solidity: function studyDays(uint256 ) view returns(bool isClosed)
func (_StudyGroup *StudyGroupCaller) StudyDays(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "studyDays", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StudyDays is a free data retrieval call binding the contract method 0x1a205a64.
//
// Solidity: function studyDays(uint256 ) view returns(bool isClosed)
func (_StudyGroup *StudyGroupSession) StudyDays(arg0 *big.Int) (bool, error) {
	return _StudyGroup.Contract.StudyDays(&_StudyGroup.CallOpts, arg0)
}

// StudyDays is a free data retrieval call binding the contract method 0x1a205a64.
//
// Solidity: function studyDays(uint256 ) view returns(bool isClosed)
func (_StudyGroup *StudyGroupCallerSession) StudyDays(arg0 *big.Int) (bool, error) {
	return _StudyGroup.Contract.StudyDays(&_StudyGroup.CallOpts, arg0)
}

// StudyEndTime is a free data retrieval call binding the contract method 0x7e9174ac.
//
// Solidity: function studyEndTime() view returns(uint256)
func (_StudyGroup *StudyGroupCaller) StudyEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "studyEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StudyEndTime is a free data retrieval call binding the contract method 0x7e9174ac.
//
// Solidity: function studyEndTime() view returns(uint256)
func (_StudyGroup *StudyGroupSession) StudyEndTime() (*big.Int, error) {
	return _StudyGroup.Contract.StudyEndTime(&_StudyGroup.CallOpts)
}

// StudyEndTime is a free data retrieval call binding the contract method 0x7e9174ac.
//
// Solidity: function studyEndTime() view returns(uint256)
func (_StudyGroup *StudyGroupCallerSession) StudyEndTime() (*big.Int, error) {
	return _StudyGroup.Contract.StudyEndTime(&_StudyGroup.CallOpts)
}

// StudyName is a free data retrieval call binding the contract method 0x30f7dd86.
//
// Solidity: function studyName() view returns(string)
func (_StudyGroup *StudyGroupCaller) StudyName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "studyName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StudyName is a free data retrieval call binding the contract method 0x30f7dd86.
//
// Solidity: function studyName() view returns(string)
func (_StudyGroup *StudyGroupSession) StudyName() (string, error) {
	return _StudyGroup.Contract.StudyName(&_StudyGroup.CallOpts)
}

// StudyName is a free data retrieval call binding the contract method 0x30f7dd86.
//
// Solidity: function studyName() view returns(string)
func (_StudyGroup *StudyGroupCallerSession) StudyName() (string, error) {
	return _StudyGroup.Contract.StudyName(&_StudyGroup.CallOpts)
}

// StudyStartTime is a free data retrieval call binding the contract method 0x5990d558.
//
// Solidity: function studyStartTime() view returns(uint256)
func (_StudyGroup *StudyGroupCaller) StudyStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudyGroup.contract.Call(opts, &out, "studyStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StudyStartTime is a free data retrieval call binding the contract method 0x5990d558.
//
// Solidity: function studyStartTime() view returns(uint256)
func (_StudyGroup *StudyGroupSession) StudyStartTime() (*big.Int, error) {
	return _StudyGroup.Contract.StudyStartTime(&_StudyGroup.CallOpts)
}

// StudyStartTime is a free data retrieval call binding the contract method 0x5990d558.
//
// Solidity: function studyStartTime() view returns(uint256)
func (_StudyGroup *StudyGroupCallerSession) StudyStartTime() (*big.Int, error) {
	return _StudyGroup.Contract.StudyStartTime(&_StudyGroup.CallOpts)
}

// CloseStudy is a paid mutator transaction binding the contract method 0xee43548d.
//
// Solidity: function closeStudy(uint256 timestamp) returns()
func (_StudyGroup *StudyGroupTransactor) CloseStudy(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _StudyGroup.contract.Transact(opts, "closeStudy", timestamp)
}

// CloseStudy is a paid mutator transaction binding the contract method 0xee43548d.
//
// Solidity: function closeStudy(uint256 timestamp) returns()
func (_StudyGroup *StudyGroupSession) CloseStudy(timestamp *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.CloseStudy(&_StudyGroup.TransactOpts, timestamp)
}

// CloseStudy is a paid mutator transaction binding the contract method 0xee43548d.
//
// Solidity: function closeStudy(uint256 timestamp) returns()
func (_StudyGroup *StudyGroupTransactorSession) CloseStudy(timestamp *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.CloseStudy(&_StudyGroup.TransactOpts, timestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0x71ec3f50.
//
// Solidity: function initialize(string _studyName, uint256 _depositAmount, uint256 _penaltyAmount, address _admin, uint256 _studyStartTime, uint256 _studyEndTime) returns()
func (_StudyGroup *StudyGroupTransactor) Initialize(opts *bind.TransactOpts, _studyName string, _depositAmount *big.Int, _penaltyAmount *big.Int, _admin common.Address, _studyStartTime *big.Int, _studyEndTime *big.Int) (*types.Transaction, error) {
	return _StudyGroup.contract.Transact(opts, "initialize", _studyName, _depositAmount, _penaltyAmount, _admin, _studyStartTime, _studyEndTime)
}

// Initialize is a paid mutator transaction binding the contract method 0x71ec3f50.
//
// Solidity: function initialize(string _studyName, uint256 _depositAmount, uint256 _penaltyAmount, address _admin, uint256 _studyStartTime, uint256 _studyEndTime) returns()
func (_StudyGroup *StudyGroupSession) Initialize(_studyName string, _depositAmount *big.Int, _penaltyAmount *big.Int, _admin common.Address, _studyStartTime *big.Int, _studyEndTime *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.Initialize(&_StudyGroup.TransactOpts, _studyName, _depositAmount, _penaltyAmount, _admin, _studyStartTime, _studyEndTime)
}

// Initialize is a paid mutator transaction binding the contract method 0x71ec3f50.
//
// Solidity: function initialize(string _studyName, uint256 _depositAmount, uint256 _penaltyAmount, address _admin, uint256 _studyStartTime, uint256 _studyEndTime) returns()
func (_StudyGroup *StudyGroupTransactorSession) Initialize(_studyName string, _depositAmount *big.Int, _penaltyAmount *big.Int, _admin common.Address, _studyStartTime *big.Int, _studyEndTime *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.Initialize(&_StudyGroup.TransactOpts, _studyName, _depositAmount, _penaltyAmount, _admin, _studyStartTime, _studyEndTime)
}

// JoinStudy is a paid mutator transaction binding the contract method 0xaa4e0204.
//
// Solidity: function joinStudy() returns()
func (_StudyGroup *StudyGroupTransactor) JoinStudy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StudyGroup.contract.Transact(opts, "joinStudy")
}

// JoinStudy is a paid mutator transaction binding the contract method 0xaa4e0204.
//
// Solidity: function joinStudy() returns()
func (_StudyGroup *StudyGroupSession) JoinStudy() (*types.Transaction, error) {
	return _StudyGroup.Contract.JoinStudy(&_StudyGroup.TransactOpts)
}

// JoinStudy is a paid mutator transaction binding the contract method 0xaa4e0204.
//
// Solidity: function joinStudy() returns()
func (_StudyGroup *StudyGroupTransactorSession) JoinStudy() (*types.Transaction, error) {
	return _StudyGroup.Contract.JoinStudy(&_StudyGroup.TransactOpts)
}

// LeaveStudy is a paid mutator transaction binding the contract method 0xc05a6471.
//
// Solidity: function leaveStudy() returns()
func (_StudyGroup *StudyGroupTransactor) LeaveStudy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StudyGroup.contract.Transact(opts, "leaveStudy")
}

// LeaveStudy is a paid mutator transaction binding the contract method 0xc05a6471.
//
// Solidity: function leaveStudy() returns()
func (_StudyGroup *StudyGroupSession) LeaveStudy() (*types.Transaction, error) {
	return _StudyGroup.Contract.LeaveStudy(&_StudyGroup.TransactOpts)
}

// LeaveStudy is a paid mutator transaction binding the contract method 0xc05a6471.
//
// Solidity: function leaveStudy() returns()
func (_StudyGroup *StudyGroupTransactorSession) LeaveStudy() (*types.Transaction, error) {
	return _StudyGroup.Contract.LeaveStudy(&_StudyGroup.TransactOpts)
}

// StartTodayStudy is a paid mutator transaction binding the contract method 0xf4aee10e.
//
// Solidity: function startTodayStudy(uint256 timestamp) returns()
func (_StudyGroup *StudyGroupTransactor) StartTodayStudy(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _StudyGroup.contract.Transact(opts, "startTodayStudy", timestamp)
}

// StartTodayStudy is a paid mutator transaction binding the contract method 0xf4aee10e.
//
// Solidity: function startTodayStudy(uint256 timestamp) returns()
func (_StudyGroup *StudyGroupSession) StartTodayStudy(timestamp *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.StartTodayStudy(&_StudyGroup.TransactOpts, timestamp)
}

// StartTodayStudy is a paid mutator transaction binding the contract method 0xf4aee10e.
//
// Solidity: function startTodayStudy(uint256 timestamp) returns()
func (_StudyGroup *StudyGroupTransactorSession) StartTodayStudy(timestamp *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.StartTodayStudy(&_StudyGroup.TransactOpts, timestamp)
}

// TerminateStudyGroup is a paid mutator transaction binding the contract method 0x32843098.
//
// Solidity: function terminateStudyGroup() returns()
func (_StudyGroup *StudyGroupTransactor) TerminateStudyGroup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StudyGroup.contract.Transact(opts, "terminateStudyGroup")
}

// TerminateStudyGroup is a paid mutator transaction binding the contract method 0x32843098.
//
// Solidity: function terminateStudyGroup() returns()
func (_StudyGroup *StudyGroupSession) TerminateStudyGroup() (*types.Transaction, error) {
	return _StudyGroup.Contract.TerminateStudyGroup(&_StudyGroup.TransactOpts)
}

// TerminateStudyGroup is a paid mutator transaction binding the contract method 0x32843098.
//
// Solidity: function terminateStudyGroup() returns()
func (_StudyGroup *StudyGroupTransactorSession) TerminateStudyGroup() (*types.Transaction, error) {
	return _StudyGroup.Contract.TerminateStudyGroup(&_StudyGroup.TransactOpts)
}

// TrackCommit is a paid mutator transaction binding the contract method 0xb6f0fd28.
//
// Solidity: function trackCommit(uint256 timestamp, address participant, uint256 commitTime) returns()
func (_StudyGroup *StudyGroupTransactor) TrackCommit(opts *bind.TransactOpts, timestamp *big.Int, participant common.Address, commitTime *big.Int) (*types.Transaction, error) {
	return _StudyGroup.contract.Transact(opts, "trackCommit", timestamp, participant, commitTime)
}

// TrackCommit is a paid mutator transaction binding the contract method 0xb6f0fd28.
//
// Solidity: function trackCommit(uint256 timestamp, address participant, uint256 commitTime) returns()
func (_StudyGroup *StudyGroupSession) TrackCommit(timestamp *big.Int, participant common.Address, commitTime *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.TrackCommit(&_StudyGroup.TransactOpts, timestamp, participant, commitTime)
}

// TrackCommit is a paid mutator transaction binding the contract method 0xb6f0fd28.
//
// Solidity: function trackCommit(uint256 timestamp, address participant, uint256 commitTime) returns()
func (_StudyGroup *StudyGroupTransactorSession) TrackCommit(timestamp *big.Int, participant common.Address, commitTime *big.Int) (*types.Transaction, error) {
	return _StudyGroup.Contract.TrackCommit(&_StudyGroup.TransactOpts, timestamp, participant, commitTime)
}

// StudyGroupCommitTrackedIterator is returned from FilterCommitTracked and is used to iterate over the raw logs and unpacked data for CommitTracked events raised by the StudyGroup contract.
type StudyGroupCommitTrackedIterator struct {
	Event *StudyGroupCommitTracked // Event containing the contract specifics and raw log

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
func (it *StudyGroupCommitTrackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudyGroupCommitTracked)
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
		it.Event = new(StudyGroupCommitTracked)
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
func (it *StudyGroupCommitTrackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudyGroupCommitTrackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudyGroupCommitTracked represents a CommitTracked event raised by the StudyGroup contract.
type StudyGroupCommitTracked struct {
	Timestamp   *big.Int
	Participant common.Address
	CommitTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommitTracked is a free log retrieval operation binding the contract event 0xc6fd0795982be08673ad2932226cedc4c07d4a2f2eedf06c59b9dc33ec27b39a.
//
// Solidity: event CommitTracked(uint256 indexed timestamp, address indexed participant, uint256 commitTime)
func (_StudyGroup *StudyGroupFilterer) FilterCommitTracked(opts *bind.FilterOpts, timestamp []*big.Int, participant []common.Address) (*StudyGroupCommitTrackedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _StudyGroup.contract.FilterLogs(opts, "CommitTracked", timestampRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &StudyGroupCommitTrackedIterator{contract: _StudyGroup.contract, event: "CommitTracked", logs: logs, sub: sub}, nil
}

// WatchCommitTracked is a free log subscription operation binding the contract event 0xc6fd0795982be08673ad2932226cedc4c07d4a2f2eedf06c59b9dc33ec27b39a.
//
// Solidity: event CommitTracked(uint256 indexed timestamp, address indexed participant, uint256 commitTime)
func (_StudyGroup *StudyGroupFilterer) WatchCommitTracked(opts *bind.WatchOpts, sink chan<- *StudyGroupCommitTracked, timestamp []*big.Int, participant []common.Address) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _StudyGroup.contract.WatchLogs(opts, "CommitTracked", timestampRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudyGroupCommitTracked)
				if err := _StudyGroup.contract.UnpackLog(event, "CommitTracked", log); err != nil {
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

// ParseCommitTracked is a log parse operation binding the contract event 0xc6fd0795982be08673ad2932226cedc4c07d4a2f2eedf06c59b9dc33ec27b39a.
//
// Solidity: event CommitTracked(uint256 indexed timestamp, address indexed participant, uint256 commitTime)
func (_StudyGroup *StudyGroupFilterer) ParseCommitTracked(log types.Log) (*StudyGroupCommitTracked, error) {
	event := new(StudyGroupCommitTracked)
	if err := _StudyGroup.contract.UnpackLog(event, "CommitTracked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudyGroupParticipantJoinedIterator is returned from FilterParticipantJoined and is used to iterate over the raw logs and unpacked data for ParticipantJoined events raised by the StudyGroup contract.
type StudyGroupParticipantJoinedIterator struct {
	Event *StudyGroupParticipantJoined // Event containing the contract specifics and raw log

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
func (it *StudyGroupParticipantJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudyGroupParticipantJoined)
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
		it.Event = new(StudyGroupParticipantJoined)
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
func (it *StudyGroupParticipantJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudyGroupParticipantJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudyGroupParticipantJoined represents a ParticipantJoined event raised by the StudyGroup contract.
type StudyGroupParticipantJoined struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantJoined is a free log retrieval operation binding the contract event 0x6cd71e167cab47cb2ccb763aada102691b585b848adb6395f1da4ac187b9214a.
//
// Solidity: event ParticipantJoined(address indexed participant)
func (_StudyGroup *StudyGroupFilterer) FilterParticipantJoined(opts *bind.FilterOpts, participant []common.Address) (*StudyGroupParticipantJoinedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _StudyGroup.contract.FilterLogs(opts, "ParticipantJoined", participantRule)
	if err != nil {
		return nil, err
	}
	return &StudyGroupParticipantJoinedIterator{contract: _StudyGroup.contract, event: "ParticipantJoined", logs: logs, sub: sub}, nil
}

// WatchParticipantJoined is a free log subscription operation binding the contract event 0x6cd71e167cab47cb2ccb763aada102691b585b848adb6395f1da4ac187b9214a.
//
// Solidity: event ParticipantJoined(address indexed participant)
func (_StudyGroup *StudyGroupFilterer) WatchParticipantJoined(opts *bind.WatchOpts, sink chan<- *StudyGroupParticipantJoined, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _StudyGroup.contract.WatchLogs(opts, "ParticipantJoined", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudyGroupParticipantJoined)
				if err := _StudyGroup.contract.UnpackLog(event, "ParticipantJoined", log); err != nil {
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

// ParseParticipantJoined is a log parse operation binding the contract event 0x6cd71e167cab47cb2ccb763aada102691b585b848adb6395f1da4ac187b9214a.
//
// Solidity: event ParticipantJoined(address indexed participant)
func (_StudyGroup *StudyGroupFilterer) ParseParticipantJoined(log types.Log) (*StudyGroupParticipantJoined, error) {
	event := new(StudyGroupParticipantJoined)
	if err := _StudyGroup.contract.UnpackLog(event, "ParticipantJoined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudyGroupParticipantLeftIterator is returned from FilterParticipantLeft and is used to iterate over the raw logs and unpacked data for ParticipantLeft events raised by the StudyGroup contract.
type StudyGroupParticipantLeftIterator struct {
	Event *StudyGroupParticipantLeft // Event containing the contract specifics and raw log

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
func (it *StudyGroupParticipantLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudyGroupParticipantLeft)
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
		it.Event = new(StudyGroupParticipantLeft)
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
func (it *StudyGroupParticipantLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudyGroupParticipantLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudyGroupParticipantLeft represents a ParticipantLeft event raised by the StudyGroup contract.
type StudyGroupParticipantLeft struct {
	Participant  common.Address
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterParticipantLeft is a free log retrieval operation binding the contract event 0xd9c208ba846167a97aa7c13f0494ae9de2cd214de16281410c3d97ab5da35eb8.
//
// Solidity: event ParticipantLeft(address indexed participant, uint256 refundAmount)
func (_StudyGroup *StudyGroupFilterer) FilterParticipantLeft(opts *bind.FilterOpts, participant []common.Address) (*StudyGroupParticipantLeftIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _StudyGroup.contract.FilterLogs(opts, "ParticipantLeft", participantRule)
	if err != nil {
		return nil, err
	}
	return &StudyGroupParticipantLeftIterator{contract: _StudyGroup.contract, event: "ParticipantLeft", logs: logs, sub: sub}, nil
}

// WatchParticipantLeft is a free log subscription operation binding the contract event 0xd9c208ba846167a97aa7c13f0494ae9de2cd214de16281410c3d97ab5da35eb8.
//
// Solidity: event ParticipantLeft(address indexed participant, uint256 refundAmount)
func (_StudyGroup *StudyGroupFilterer) WatchParticipantLeft(opts *bind.WatchOpts, sink chan<- *StudyGroupParticipantLeft, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _StudyGroup.contract.WatchLogs(opts, "ParticipantLeft", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudyGroupParticipantLeft)
				if err := _StudyGroup.contract.UnpackLog(event, "ParticipantLeft", log); err != nil {
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

// ParseParticipantLeft is a log parse operation binding the contract event 0xd9c208ba846167a97aa7c13f0494ae9de2cd214de16281410c3d97ab5da35eb8.
//
// Solidity: event ParticipantLeft(address indexed participant, uint256 refundAmount)
func (_StudyGroup *StudyGroupFilterer) ParseParticipantLeft(log types.Log) (*StudyGroupParticipantLeft, error) {
	event := new(StudyGroupParticipantLeft)
	if err := _StudyGroup.contract.UnpackLog(event, "ParticipantLeft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudyGroupStudyClosedIterator is returned from FilterStudyClosed and is used to iterate over the raw logs and unpacked data for StudyClosed events raised by the StudyGroup contract.
type StudyGroupStudyClosedIterator struct {
	Event *StudyGroupStudyClosed // Event containing the contract specifics and raw log

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
func (it *StudyGroupStudyClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudyGroupStudyClosed)
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
		it.Event = new(StudyGroupStudyClosed)
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
func (it *StudyGroupStudyClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudyGroupStudyClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudyGroupStudyClosed represents a StudyClosed event raised by the StudyGroup contract.
type StudyGroupStudyClosed struct {
	Timestamp             *big.Int
	PenalizedParticipants []common.Address
	PenaltyAmount         *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterStudyClosed is a free log retrieval operation binding the contract event 0xe4627dcf110aa70b73267922fb1d1c841476972398c9302d00b20a1d9853d8c0.
//
// Solidity: event StudyClosed(uint256 indexed timestamp, address[] penalizedParticipants, uint256 penaltyAmount)
func (_StudyGroup *StudyGroupFilterer) FilterStudyClosed(opts *bind.FilterOpts, timestamp []*big.Int) (*StudyGroupStudyClosedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _StudyGroup.contract.FilterLogs(opts, "StudyClosed", timestampRule)
	if err != nil {
		return nil, err
	}
	return &StudyGroupStudyClosedIterator{contract: _StudyGroup.contract, event: "StudyClosed", logs: logs, sub: sub}, nil
}

// WatchStudyClosed is a free log subscription operation binding the contract event 0xe4627dcf110aa70b73267922fb1d1c841476972398c9302d00b20a1d9853d8c0.
//
// Solidity: event StudyClosed(uint256 indexed timestamp, address[] penalizedParticipants, uint256 penaltyAmount)
func (_StudyGroup *StudyGroupFilterer) WatchStudyClosed(opts *bind.WatchOpts, sink chan<- *StudyGroupStudyClosed, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _StudyGroup.contract.WatchLogs(opts, "StudyClosed", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudyGroupStudyClosed)
				if err := _StudyGroup.contract.UnpackLog(event, "StudyClosed", log); err != nil {
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

// ParseStudyClosed is a log parse operation binding the contract event 0xe4627dcf110aa70b73267922fb1d1c841476972398c9302d00b20a1d9853d8c0.
//
// Solidity: event StudyClosed(uint256 indexed timestamp, address[] penalizedParticipants, uint256 penaltyAmount)
func (_StudyGroup *StudyGroupFilterer) ParseStudyClosed(log types.Log) (*StudyGroupStudyClosed, error) {
	event := new(StudyGroupStudyClosed)
	if err := _StudyGroup.contract.UnpackLog(event, "StudyClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudyGroupStudyGroupTerminatedIterator is returned from FilterStudyGroupTerminated and is used to iterate over the raw logs and unpacked data for StudyGroupTerminated events raised by the StudyGroup contract.
type StudyGroupStudyGroupTerminatedIterator struct {
	Event *StudyGroupStudyGroupTerminated // Event containing the contract specifics and raw log

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
func (it *StudyGroupStudyGroupTerminatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudyGroupStudyGroupTerminated)
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
		it.Event = new(StudyGroupStudyGroupTerminated)
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
func (it *StudyGroupStudyGroupTerminatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudyGroupStudyGroupTerminatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudyGroupStudyGroupTerminated represents a StudyGroupTerminated event raised by the StudyGroup contract.
type StudyGroupStudyGroupTerminated struct {
	Participants  []common.Address
	TotalRefunded *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStudyGroupTerminated is a free log retrieval operation binding the contract event 0xd808b1510f070636af55fdd79b2ab453621eec258030d7b99bc180067e28fb6c.
//
// Solidity: event StudyGroupTerminated(address[] participants, uint256 totalRefunded)
func (_StudyGroup *StudyGroupFilterer) FilterStudyGroupTerminated(opts *bind.FilterOpts) (*StudyGroupStudyGroupTerminatedIterator, error) {

	logs, sub, err := _StudyGroup.contract.FilterLogs(opts, "StudyGroupTerminated")
	if err != nil {
		return nil, err
	}
	return &StudyGroupStudyGroupTerminatedIterator{contract: _StudyGroup.contract, event: "StudyGroupTerminated", logs: logs, sub: sub}, nil
}

// WatchStudyGroupTerminated is a free log subscription operation binding the contract event 0xd808b1510f070636af55fdd79b2ab453621eec258030d7b99bc180067e28fb6c.
//
// Solidity: event StudyGroupTerminated(address[] participants, uint256 totalRefunded)
func (_StudyGroup *StudyGroupFilterer) WatchStudyGroupTerminated(opts *bind.WatchOpts, sink chan<- *StudyGroupStudyGroupTerminated) (event.Subscription, error) {

	logs, sub, err := _StudyGroup.contract.WatchLogs(opts, "StudyGroupTerminated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudyGroupStudyGroupTerminated)
				if err := _StudyGroup.contract.UnpackLog(event, "StudyGroupTerminated", log); err != nil {
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

// ParseStudyGroupTerminated is a log parse operation binding the contract event 0xd808b1510f070636af55fdd79b2ab453621eec258030d7b99bc180067e28fb6c.
//
// Solidity: event StudyGroupTerminated(address[] participants, uint256 totalRefunded)
func (_StudyGroup *StudyGroupFilterer) ParseStudyGroupTerminated(log types.Log) (*StudyGroupStudyGroupTerminated, error) {
	event := new(StudyGroupStudyGroupTerminated)
	if err := _StudyGroup.contract.UnpackLog(event, "StudyGroupTerminated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudyGroupStudyStartedIterator is returned from FilterStudyStarted and is used to iterate over the raw logs and unpacked data for StudyStarted events raised by the StudyGroup contract.
type StudyGroupStudyStartedIterator struct {
	Event *StudyGroupStudyStarted // Event containing the contract specifics and raw log

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
func (it *StudyGroupStudyStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudyGroupStudyStarted)
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
		it.Event = new(StudyGroupStudyStarted)
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
func (it *StudyGroupStudyStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudyGroupStudyStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudyGroupStudyStarted represents a StudyStarted event raised by the StudyGroup contract.
type StudyGroupStudyStarted struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStudyStarted is a free log retrieval operation binding the contract event 0xe4bad6f84ca0da462474e2232f201c0c7b83f546ed9687ebb9d959a14e9c928b.
//
// Solidity: event StudyStarted(uint256 indexed timestamp)
func (_StudyGroup *StudyGroupFilterer) FilterStudyStarted(opts *bind.FilterOpts, timestamp []*big.Int) (*StudyGroupStudyStartedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _StudyGroup.contract.FilterLogs(opts, "StudyStarted", timestampRule)
	if err != nil {
		return nil, err
	}
	return &StudyGroupStudyStartedIterator{contract: _StudyGroup.contract, event: "StudyStarted", logs: logs, sub: sub}, nil
}

// WatchStudyStarted is a free log subscription operation binding the contract event 0xe4bad6f84ca0da462474e2232f201c0c7b83f546ed9687ebb9d959a14e9c928b.
//
// Solidity: event StudyStarted(uint256 indexed timestamp)
func (_StudyGroup *StudyGroupFilterer) WatchStudyStarted(opts *bind.WatchOpts, sink chan<- *StudyGroupStudyStarted, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _StudyGroup.contract.WatchLogs(opts, "StudyStarted", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudyGroupStudyStarted)
				if err := _StudyGroup.contract.UnpackLog(event, "StudyStarted", log); err != nil {
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

// ParseStudyStarted is a log parse operation binding the contract event 0xe4bad6f84ca0da462474e2232f201c0c7b83f546ed9687ebb9d959a14e9c928b.
//
// Solidity: event StudyStarted(uint256 indexed timestamp)
func (_StudyGroup *StudyGroupFilterer) ParseStudyStarted(log types.Log) (*StudyGroupStudyStarted, error) {
	event := new(StudyGroupStudyStarted)
	if err := _StudyGroup.contract.UnpackLog(event, "StudyStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
