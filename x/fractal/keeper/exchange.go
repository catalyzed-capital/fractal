package keeper

import (
	"encoding/binary"

	"fractal/x/fractal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetExchangeCount get the total number of exchange
func (k Keeper) GetExchangeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ExchangeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetExchangeCount set the total number of exchange
func (k Keeper) SetExchangeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ExchangeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendExchange appends a exchange in the store with a new id and update the count
func (k Keeper) AppendExchange(
	ctx sdk.Context,
	exchange types.Exchange,
) uint64 {
	// Create the exchange
	count := k.GetExchangeCount(ctx)

	// Set the ID of the appended value
	exchange.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeKey))
	appendedValue := k.cdc.MustMarshal(&exchange)
	store.Set(GetExchangeIDBytes(exchange.Id), appendedValue)

	// Update exchange count
	k.SetExchangeCount(ctx, count+1)

	return count
}

// SetExchange set a specific exchange in the store
func (k Keeper) SetExchange(ctx sdk.Context, exchange types.Exchange) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeKey))
	b := k.cdc.MustMarshal(&exchange)
	store.Set(GetExchangeIDBytes(exchange.Id), b)
}

// GetExchange returns a exchange from its id
func (k Keeper) GetExchange(ctx sdk.Context, id uint64) (val types.Exchange, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeKey))
	b := store.Get(GetExchangeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExchange removes a exchange from the store
func (k Keeper) RemoveExchange(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeKey))
	store.Delete(GetExchangeIDBytes(id))
}

// GetAllExchange returns all exchange
func (k Keeper) GetAllExchange(ctx sdk.Context) (list []types.Exchange) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Exchange
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetExchangeIDBytes returns the byte representation of the ID
func GetExchangeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetExchangeIDFromBytes returns ID in uint64 format from a byte array
func GetExchangeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
