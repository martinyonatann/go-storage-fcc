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
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_favoriteNumber\",\"type\":\"uint256\"}],\"name\":\"addPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToFavoriteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"people\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"favoriteNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_favoriteNumber\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061057c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632e64cec11461005c5780636057361d146100735780636f760f41146100885780638bab8dd51461009b5780639e7a13ad146100c6575b600080fd5b6000545b6040519081526020015b60405180910390f35b610086610081366004610245565b600055565b005b610086610096366004610301565b6100e7565b6100606100a9366004610346565b805160208183018101805160018252928201919093012091525481565b6100d96100d4366004610245565b61018d565b60405161006a9291906103a7565b6040805180820190915281815260208101838152600280546001810182556000829052835191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101918255915190917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf0190610165908261046a565b50505080600183604051610179919061052a565b908152604051908190036020019020555050565b6002818154811061019d57600080fd5b600091825260209091206002909102018054600182018054919350906101c2906103e1565b80601f01602080910402602001604051908101604052809291908181526020018280546101ee906103e1565b801561023b5780601f106102105761010080835404028352916020019161023b565b820191906000526020600020905b81548152906001019060200180831161021e57829003601f168201915b5050505050905082565b60006020828403121561025757600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261028557600080fd5b813567ffffffffffffffff808211156102a0576102a061025e565b604051601f8301601f19908116603f011681019082821181831017156102c8576102c861025e565b816040528381528660208588010111156102e157600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561031457600080fd5b823567ffffffffffffffff81111561032b57600080fd5b61033785828601610274565b95602094909401359450505050565b60006020828403121561035857600080fd5b813567ffffffffffffffff81111561036f57600080fd5b61037b84828501610274565b949350505050565b60005b8381101561039e578181015183820152602001610386565b50506000910152565b82815260406020820152600082518060408401526103cc816060850160208701610383565b601f01601f1916919091016060019392505050565b600181811c908216806103f557607f821691505b60208210810361041557634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561046557600081815260208120601f850160051c810160208610156104425750805b601f850160051c820191505b818110156104615782815560010161044e565b5050505b505050565b815167ffffffffffffffff8111156104845761048461025e565b6104988161049284546103e1565b8461041b565b602080601f8311600181146104cd57600084156104b55750858301515b600019600386901b1c1916600185901b178555610461565b600085815260208120601f198616915b828110156104fc578886015182559484019460019091019084016104dd565b508582101561051a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000825161053c818460208701610383565b919091019291505056fea2646970667358221220d049abcceac819d033c44cf6dfb810aff0f541d589fd4d918262625b4c3b1a9064736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

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
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_Api *ApiCaller) NameToFavoriteNumber(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "nameToFavoriteNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_Api *ApiSession) NameToFavoriteNumber(arg0 string) (*big.Int, error) {
	return _Api.Contract.NameToFavoriteNumber(&_Api.CallOpts, arg0)
}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_Api *ApiCallerSession) NameToFavoriteNumber(arg0 string) (*big.Int, error) {
	return _Api.Contract.NameToFavoriteNumber(&_Api.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_Api *ApiCaller) People(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "people", arg0)

	outstruct := new(struct {
		FavoriteNumber *big.Int
		Name           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FavoriteNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_Api *ApiSession) People(arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	return _Api.Contract.People(&_Api.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_Api *ApiCallerSession) People(arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	return _Api.Contract.People(&_Api.CallOpts, arg0)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Api *ApiCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Api *ApiSession) Retrieve() (*big.Int, error) {
	return _Api.Contract.Retrieve(&_Api.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Api *ApiCallerSession) Retrieve() (*big.Int, error) {
	return _Api.Contract.Retrieve(&_Api.CallOpts)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_Api *ApiTransactor) AddPerson(opts *bind.TransactOpts, _name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addPerson", _name, _favoriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_Api *ApiSession) AddPerson(_name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddPerson(&_Api.TransactOpts, _name, _favoriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_Api *ApiTransactorSession) AddPerson(_name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddPerson(&_Api.TransactOpts, _name, _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Api *ApiTransactor) Store(opts *bind.TransactOpts, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "store", _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Api *ApiSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Store(&_Api.TransactOpts, _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Api *ApiTransactorSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Store(&_Api.TransactOpts, _favoriteNumber)
}
