package test

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"go-wallet/wallet"
	"testing"
)

// leather wallet
const pwd = "11447755pp@!"
const mne = "crucial tent equal movie volcano bleak erase letter vacuum rapid banner believe increase mushroom this memory price balcony scale witness expand green net abuse"
const addr = "bc1q3rzzzm7mny78tt0tt5s67zqpyykqah8hpemuu6"

// local wallet
const mne2 = "lab initial awful neglect balcony tell add please run chest say thing forward wait sphere absorb earn taste nerve blood because bleak allow unlock"
const pwd2 = "11447755pp@!"
const addr2 = "bc1q9e0mew6zhy67mph9ylcqnzpzuaq8z2ya4psgds"

func TestFromMne(t *testing.T) {

	fmt.Println("walletAddr:", addr)

	key, _ := wallet.GetPrivateFromMnemonic(mne, "")

	wif, a1, a2, a3, _ := wallet.GetBTCMainAddress(key, true)
	fmt.Println("wif:", wif)
	fmt.Println("ad1:", a1)
	fmt.Println("ad2:", a2)
	fmt.Println("ad3:", a3)
	fmt.Println("--------------------")
	//txscript.SegwitSigHashMidstate{}

	//for i := 0; i < 10; i++ {
	//	childKey, _ := wallet.GetPrivateFromMnemonicByChild(mne, "", uint32(i))
	//	_, _, a2, _, _ = wallet.GetBTCMainAddress(childKey, true)
	//	fmt.Println("addr:", a2)
	//	fmt.Println("--------------------")
	//}
}

func TestBtcGenerata(t *testing.T) {
	mnemonic, key, err := wallet.GenerataPrivate("11447755pp@!", 24)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mnemonic:", mnemonic)

	wif, a1, a2, a3, err := wallet.GetBTCTestNetAddress(key, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("addr:", wif)
	fmt.Println("addr:", a1)
	fmt.Println("addr:", a2)
	fmt.Println("addr:", a3)

}

func TestPrivate(t *testing.T) {
	key := "f46123d7ba1fdda75345a9c83a3dc883f8246b61577f3f020ff7d2a1d71b88a9"

	keyb, _ := hex.DecodeString(key)
	prv, _ := btcec.PrivKeyFromBytes(keyb)
	etcA := crypto.PubkeyToAddress(prv.ToECDSA().PublicKey)
	fmt.Println("eth:", etcA)

	// btc-okex : bc1phrtfjws5qphxe87kf3j466qlncthwqkhpnu2k7dmxa5a784atgfs86tzu4
	mnemonic := "liar fabric trade later deputy heart enroll hair link card mail home"
	pass := ""

	master, _ := wallet.GetMasterKey(mnemonic, pass)
	fmt.Println("master:", master)
	path := []uint32{wallet.PurposeBip44, wallet.ETC, wallet.Account, wallet.Change, wallet.ChildIndex}
	prv, _ = wallet.Driver(master, path)
	//fmt.Println("key-private:",key)
	//fmt.Println("key-pub:",key.PublicKey())
	etcb := crypto.PubkeyToAddress(prv.ToECDSA().PublicKey)
	fmt.Println("eth:", etcb)

	// btc
	masterBtc, _ := wallet.GetMasterKey(mnemonic, pass)
	path = []uint32{wallet.PurposeBip44, wallet.BTC, wallet.Account, wallet.ChildIndex}
	btcPrv, _ := wallet.Driver(masterBtc, path)
	wif, a1, a2, a3, _ := wallet.GetBTCMainAddress(btcPrv, true)
	fmt.Println("wif:", wif)
	fmt.Println("btc:", a1)
	fmt.Println("btc:", a2)
	fmt.Println("btc:", a3)

	// private :KysHxNxksxHtyDAFWYnXDdV6ADk2HsjBXRDyaAP5gS9d8DjkZabp
	key_t := "KysHxNxksxHtyDAFWYnXDdV6ADk2HsjBXRDyaAP5gS9d8DjkZabp"
	fmt.Println(hex.EncodeToString(base58.Decode(key_t)))
	keyb, err := hex.DecodeString(hex.EncodeToString(base58.Decode(key_t)))
	if err != nil {
		fmt.Println(err)
		return
	}
	pr, _ := btcec.PrivKeyFromBytes(keyb)
	fmt.Println(hex.EncodeToString(pr.Serialize()))
	wif, a1, a2, a3, _ = wallet.GetBTCMainAddress(pr, true)
	fmt.Println("wif:", wif)
	fmt.Println("btc:", a1)
	fmt.Println("btc:", a2)
	fmt.Println("btc:", a3)
}
