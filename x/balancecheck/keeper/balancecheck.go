package keeper

import (

	//"encoding/binary"
	"encoding/json"
	"fmt"
	"fractal/x/balancecheck/types"
	"net/http"
	//"github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"

)

func (k Keeper) AppendUsdcBalance(ctx sdk.Context, usdcbalance types.Usdcbalance) map[string]interface{} {

	url := "https://api-sandbox.circle.com/v1/businessAccount/balances"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error on request")
		
	}

	bearerToken := "QVBJX0tFWTphNjIwZjNmMzk3OTRiNDcxZGZhOGRjZGQ3Y2EyZmY5ZjplOGZmOTExNGU3Mzg1MWM3OTZkN2E4MWU5YzBjN2Y3Ng=="
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+ bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println("Error on sending bearer token")
	
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("json decoder error")

	}

	fmt.Println(data)

	return data
	
}