package wallet

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

const bash uint32 = 0x80000000

const PurposeBip44 = bash + 44
const PurposeBIP49 = bash + 49

const BTC = bash + 0
const BTCTest = bash + 1
const ETC = bash + 60
const FIL = bash + 461

const Account = bash + 0
const Change uint32 = 0
const ChildIndex uint32 = 0

var BtcMainPath = []uint32{PurposeBip44, BTC, Account, Change, ChildIndex}
var OneBtcTestPath = []uint32{PurposeBip44, BTCTest, Account, Change, ChildIndex}

func NewRandMaterKey(password string, mneNumber int) (string, *bip32.Key, error) {
	if mneNumber != 12 && mneNumber != 24 {
		return "", nil, fmt.Errorf("Mnemonic number is err")
	}
	entropy, err := bip39.NewEntropy(mneNumber * 128 / 12)
	if err != nil {
		return "", nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", nil, err
	}
	seed := bip39.NewSeed(mnemonic, password)
	masterKey, err := bip32.NewMasterKey(seed)
	return mnemonic, masterKey, err
}

func GetMasterKey(mnemonic string, password string) (*bip32.Key, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)

	if err != nil {
		return nil, err
	}
	mastKey, err := bip32.NewMasterKey(seed)
	return mastKey, err
}

func Driver(materKey *bip32.Key, path []uint32) (*btcec.PrivateKey, error) {
	key := materKey
	for _, v := range path {
		k, err := key.NewChildKey(v)
		if err != nil {
			return nil, err
		}
		key = k
	}
	prv, _ := btcec.PrivKeyFromBytes(key.Key)
	return prv, nil
}
