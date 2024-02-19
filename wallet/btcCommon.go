package wallet

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

// GenerataPrivate 随机生成助记词私钥 bip32， bip44
func GenerataPrivate(password string, mneNumber int) (string, *btcec.PrivateKey, error) {
	mne, bipKey, _ := NewRandMaterKey(password, mneNumber)
	btcKey, err := Driver(bipKey, BtcMainPath)
	if err != nil {
		return "", nil, err
	}
	return mne, btcKey, nil
}

func GetPrivateFromMnemonicByChild(mne, password string, child uint32) (*btcec.PrivateKey, error) {
	bip32key, err := GetMasterKey(mne, password)
	if err != nil {
		return nil, err
	}
	// todo bug
	childKey, _ := bip32key.NewChildKey(child)
	btcKey, err := Driver(childKey, BtcMainPath)
	if err != nil {
		return nil, err
	}
	return btcKey, nil
}

func GetPrivateFromMnemonic(mne, password string) (*btcec.PrivateKey, error) {
	bip32key, err := GetMasterKey(mne, password)
	if err != nil {
		return nil, err
	}
	btcKey, err := Driver(bip32key, BtcMainPath)
	if err != nil {
		return nil, err
	}
	return btcKey, nil
}

// GetBTCMainAddress btc main addreess
//
//	btcec.PrivateKey
//	compressed	by serializing the public key compressed rather than uncompressed.
func GetBTCMainAddress(prvKey *btcec.PrivateKey, compress bool) (wif, p2pkhAddress, p2wkhAddress, scriptAddress string, err error) {
	net := &chaincfg.MainNetParams
	return GetBtcAddress(prvKey, net, compress)
}

func GetBTCTestNetAddress(prvKey *btcec.PrivateKey, compress bool) (wif, p2pkhAddress, p2wkhAddress, scriptAddress string, err error) {
	net := &chaincfg.TestNet3Params
	return GetBtcAddress(prvKey, net, compress)
}

func GetBtcAddress(prvKey *btcec.PrivateKey, net *chaincfg.Params, compress bool) (wif, p2pkhAddress, p2wkhAddress, scriptAddress string, err error) {
	// generate the wif(wallet import format) string
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, compress)
	if err != nil {
		return
	}
	wif = btcwif.String()

	// generate a normal p2pkh address
	serializedPubKey := btcwif.SerializePubKey()
	addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, net)
	if err != nil {
		return
	}
	p2pkhAddress = addressPubKey.EncodeAddress()

	// generate a normal p2wkh address from the pubkey hash
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, net)
	//addressWitnessPubKeyHash, err := btcutil.NewAddressTaproot(witnessProg, net)
	if err != nil {
		return
	}
	p2wkhAddress = addressWitnessPubKeyHash.EncodeAddress()

	// generate an address which is
	// backwards compatible to Bitcoin nodes running 0.6.0 onwards, but
	// allows us to take advantage of segwit's scripting improvments,
	// and malleability fixes.
	serializedScript, err := txscript.PayToAddrScript(addressWitnessPubKeyHash)
	if err != nil {
		return
	}
	addressScriptHash, err := btcutil.NewAddressScriptHash(serializedScript, net)
	if err != nil {
		return
	}
	scriptAddress = addressScriptHash.EncodeAddress()

	return
}
