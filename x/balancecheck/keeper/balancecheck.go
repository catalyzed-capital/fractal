package keeper

import (

	"encoding/binary"
	"encoding/json"
	//"fmt"
	"fractal/x/balancecheck/types"
	"net/http"
	"io/ioutil"
	"github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"

)

func (k Keeper) AppendUsdcBalance(ctx sdk.Context, usdcbalance types.Usdcbalance) uint64 {

	count := k.GetUsdcBalanceCount(ctx)
	usdcbalance = k.BalanceCheckUsdc(usdcbalance)
	usdcbalance.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsdcKey))
	appendedValue := k.cdc.MustMarshal(&usdcbalance)
	store.Set(k.GetUsdcIDBytes(usdcbalance.Id), appendedValue)
	k.SetUsdcbalanceCount(ctx, count+1)

	return count
	
}

func (k Keeper) GetUsdcBalanceCount(ctx sdk.Context) uint64 {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UsdcCountKey)
	bz := store.Get(byteKey)
	if bz == nil{
		return 0
	}

	return binary.BigEndian.Uint64(bz)

}

func (k Keeper) GetUsdcIDBytes(id uint64) []byte {

	bz :=make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz

}

func (k Keeper) SetUsdcbalanceCount(ctx sdk.Context, count uint64){
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UsdcCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

type WalletAddress struct{

	Data []struct{
		Address string `json:"address"`
		Currency string `json:"currency"`
		Chain string `json:"chain"`
	} `json:"data"`

}

func (k Keeper) BalanceCheckUsdc(usdcbalance types.Usdcbalance) types.Usdcbalance {

	url := "https://api-sandbox.circle.com/v1/businessAccount/wallets/addresses/deposit"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		usdcbalance.Chainaddress = "Request Error"
		usdcbalance.Unit = "Request Error"
		usdcbalance.Chain = "Request Error"
		return usdcbalance
	}

	bearerToken := "QVBJX0tFWTphNjIwZjNmMzk3OTRiNDcxZGZhOGRjZGQ3Y2EyZmY5ZjplOGZmOTExNGU3Mzg1MWM3OTZkN2E4MWU5YzBjN2Y3Ng=="
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+ bearerToken)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil{
		usdcbalance.Chainaddress = "Client Error"
		usdcbalance.Unit = "Client Error"
		usdcbalance.Chain = "Client Error"
		return usdcbalance
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		usdcbalance.Chainaddress = "Body Error"
		usdcbalance.Unit = "Body Error"
		usdcbalance.Chain = "Body Error"
		return usdcbalance
	}

	var walletAddr WalletAddress
	err = json.Unmarshal(body, &walletAddr)
	if err != nil{
		usdcbalance.Chainaddress = "Json Error"
		usdcbalance.Unit = "Json Error"
		usdcbalance.Chain = "Json Error"
		return usdcbalance
	}

	usdcbalance.Chainaddress = walletAddr.Data[0].Address
	usdcbalance.Unit = "CODE RAN"
	usdcbalance.Chain = walletAddr.Data[0].Chain

	return usdcbalance

}