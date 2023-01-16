package controllers

import (
	"fmt"
	"net/http"

	"github.com/NonsoAmadi10/bitcoin-wallet/utils"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/labstack/echo/v4"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	MasterPubKey   string `json:"master_pub_key"`
	MasterPrivKey  string `json:"master_priv_key"`
	DerivedPrivKey string `json:"derived_priv_key"`
	DerivedPubKey  string `json:"derived_pub_key"`
	Mnemonic       string `json:"mnemonic"`
}

type P2SHAddress struct {
	MasterPrivKey string `json:"master_priv_key" validate:"required"`
}

type P2SHAddressResponse struct {
	Address      string `json:"p2sh_address" validate:"required"`
	RedeemScript string `json:"redeem_script" validate:"required"`
}

func CreateWallet(c echo.Context) error {

	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, utils.GetEnv("SECRET_PHRASE"))

	masterKey, _ := bip32.NewMasterKey(seed)

	publicKey := masterKey.PublicKey()

	// Generate a Bip32 key for the Testnet
	testnetKey, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)

	// Generate a private key using the Bip32 key
	derivedPrivateKey := testnetKey.Key

	// Generate a public key from the private key
	derivedPublicKey := testnetKey.PublicKey()

	result := Wallet{
		MasterPrivKey:  masterKey.B58Serialize(),
		MasterPubKey:   publicKey.PublicKey().B58Serialize(),
		DerivedPrivKey: string(derivedPrivateKey),
		DerivedPubKey:  derivedPublicKey.String(),
		Mnemonic:       mnemonic,
	}

	response := &utils.Response{
		Data:    result,
		Message: "Wallet Generated! Please keep this data a secret as you would need it to initiate a transaction",
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

func GenerateP2SHAddresses(c echo.Context) (err error) {

	reqBody := new(P2SHAddress)
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(reqBody); err != nil {
		return err
	}

	masterKey, err := hdkeychain.NewKeyFromString(reqBody.MasterPrivKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey, err := masterKey.ECPubKey()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Derive the testnet address using the public key
	testnetParams := &chaincfg.TestNet3Params
	//mainnetParams := &chaincfg.MainNetParams
	p2shAddr, err := btcutil.NewAddressScriptHash(publicKey.SerializeUncompressed(), testnetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Generate the p2sh redeem script
	// Get the public key hash
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	// Create the redeem script
	redeemScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(pubKeyHash).Script()
	if err != nil {
		fmt.Println(err)
		return
	}

	response := &P2SHAddressResponse{
		Address:      p2shAddr.EncodeAddress(),
		RedeemScript: string(redeemScript),
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}
