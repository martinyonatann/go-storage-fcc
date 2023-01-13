package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/martinyonatann/go-storage-fcc/api"
)

func main() {
	// connect to blockchain network
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		panic(err)
	}

	// private key of the deployer
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	// extract public key of the deployer from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// address of the deployer
	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	// chain id of the network
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	// Get Transaction Ops to make a valid Ethereum transaction
	auth, err := GetNextTransaction(client, fromAddr, privateKey, chainID)
	if err != nil {
		panic(err)
	}

	// deploy the contract
	addr, tx, storageApi, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Api contract deployed to %s\n", addr.Hex())
	fmt.Printf("Tx: %s\n", tx.Hash().Hex())

	// Get Favorite Number
	// Call Storage contract Retrieve function to get current favourite number
	favNumber, err := storageApi.Retrieve(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Favorite Number: %d\n", favNumber)

	// Set Favourite Number
	// Get Transaction Ops to make a valid Ethereum transaction
	auth, err = GetNextTransaction(client, fromAddr, privateKey, chainID)
	if err != nil {
		panic(err)
	}

	// Call Storage function to store fav number
	reply, err := storageApi.Store(auth, big.NewInt(20))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Reply: %s\n", reply.Hash().Hex())

	// Get Fav Number
	// P.S Retrieve is a Gas Free function hence no need to get Transaction Ops
	newFavNumber, err := storageApi.Retrieve(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Favorite Number: %d\n", newFavNumber)
}

// GetNextTransaction returns the next transaction in the pending transaction queue
// NOTE: this is not an optimized way
func GetNextTransaction(client *ethclient.Client, fromAddr common.Address, privateKey *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	// nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return nil, err
	}

	// sign the transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)             // in wei
	auth.GasLimit = uint64(3000000)        // in units
	auth.GasPrice = big.NewInt(1000000000) // in wei

	return auth, nil
}
