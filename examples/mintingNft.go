package examples

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/pkg/nftlabs"
	"log"
	"math/big"
)

func main() {
	nftContractAddress := ""
	chainRpcUrl := "https://rpc-mumbai.maticvigil.com" // change this

	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	sdk, err := nftlabs.NewSdk(client, &nftlabs.SdkOptions{PrivateKey: "// TODO"})
	if err != nil {
		log.Fatal(err)
	}

	nftModule, err := sdk.GetNftModule(nftContractAddress)
	if err != nil {
		log.Fatal(err)
	}

	if result, err := nftModule.Mint(nftlabs.MintNftMetadata{
		Name:                 "",
		Description:          "",
		Image:                "",
		ExternalUrl:          "",
		SellerFeeBasisPoints: big.NewInt(1000), // 10%
		FeeRecipient:         "",
		BackgroundColor:      "",
	}); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Minted new nft with ID %d", result.Id)
	}
}
