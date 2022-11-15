package thirdweb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventListener(t *testing.T) {
	nft := getNft()

	quantityEvent := 0
	_, err := nft.Events.AddEventListener("TokensMinted", func (event ContractEvent) {
		assert.Equal(t, event.EventName, "TokensMinted")
		assert.Equal(t, event.Data["to"], nft.helper.GetSignerAddress().String())
		quantityEvent++
	})
	assert.Nil(t, err)

	_, err = nft.Mint(context.Background(), &NFTMetadataInput{})
	assert.Nil(t, err)

	<- time.After(5 * time.Second)
	assert.Equal(t, 1, quantityEvent)
}

func TestGetEvents(t *testing.T) {
	nft := getNft()

	_, err := nft.Mint(context.Background(), &NFTMetadataInput{})
	assert.Nil(t, err)

	events, err := nft.Events.GetEvents(context.Background(), "TokensMinted", EventQueryOptions{})
	assert.Nil(t, err)

	assert.Equal(t, 1, len(events))
}