package thirdweb

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mitchellh/mapstructure"
	"github.com/thirdweb-dev/go-sdk/abi"

	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
)

// The contract deployer lets you deploy new contracts to the blockchain using
// just the thirdweb SDK. You can access the contract deployer interface as follows:
//
//	import (
//		"github.com/thirdweb-dev/go-sdk/thirdweb"
//	)
//
//	privateKey = "..."
//
//	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	// Now you can deploy a contract
//	address, err := sdk.Deployer.DeployNFTCollection(
//		&thirdweb.DeployNFTCollectionMetadata{
//			Name: "Go NFT",
//		}
//	})
type ContractDeployer struct {
	*ProviderHandler
	factory *abi.TWFactory
	helper  *contractHelper
	storage storage
}

func newContractDeployer(provider *ethclient.Client, privateKey string, storage storage) (*ContractDeployer, error) {
	handler, err := NewProviderHandler(provider, privateKey)
	if err != nil {
		return nil, err
	}

	chainId, err := provider.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	factoryAddress, err := getContractAddressByChainId(ChainID(chainId.Int64()), "TWFactory")
	if err != nil {
		return nil, err
	}

	helper, err := newContractHelper(common.HexToAddress(factoryAddress), provider, privateKey)
	if err != nil {
		return nil, err
	}

	factory, err := abi.NewTWFactory(common.HexToAddress(factoryAddress), provider)
	if err != nil {
		return nil, err
	}

	contractDeployer := &ContractDeployer{
		handler,
		factory,
		helper,
		storage,
	}

	return contractDeployer, nil
}

// Deploy a new NFT Collection contract.
//
// metadata: the contract metadata
//
// returns: the address of the deployed contract
//
// Example
//
//		address, err := sdk.Deployer.DeployNFTCollection(
//	     context.Background(),
//			&thirdweb.DeployNFTCollectionMetadata{
//				Name: "Go NFT",
//			}
//		})
func (deployer *ContractDeployer) DeployNFTCollection(ctx context.Context, metadata *DeployNFTCollectionMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract(ctx, "nft-collection", metadata)
}

// Deploy a new Edition contract.
//
// metadata: the contract metadata
//
// returns: the address of the deployed contract
//
// Example
//
//		address, err := sdk.Deployer.DeployEdition(
//	     context.Background(),
//			&thirdweb.DeployEditionMetadata{
//				Name: "Go Edition",
//			}
//		})
func (deployer *ContractDeployer) DeployEdition(ctx context.Context, metadata *DeployEditionMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract(ctx, "edition", metadata)
}

// Deploy a new Token contract.
//
// metadata: the contract metadata
//
// returns: the address of the deployed contract
//
// Example
//
//		address, err := sdk.Deployer.DeployToken(
//	     context.Background(),
//			&thirdweb.DeployTokenMetadata{
//				Name: "Go Token",
//			}
//		})
func (deployer *ContractDeployer) DeployToken(ctx context.Context, metadata *DeployTokenMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract(ctx, "token", metadata)
}

// Deploy a new NFT Drop contract.
//
// metadata: the contract metadata
//
// returns: the address of the deployed contract
//
// Example
//
//		address, err := sdk.Deployer.DeployNFTDrop(
//	     context.Background(),
//			&thirdweb.DeployNFTDropMetadata{
//				Name: "Go NFT Drop",
//			}
//		})
func (deployer *ContractDeployer) DeployNFTDrop(ctx context.Context, metadata *DeployNFTDropMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract(ctx, "nft-drop", metadata)
}

// Deploy a new Edition Drop contract.
//
// metadata: the contract metadata
//
// returns: the address of the deployed contract
//
// Example
//
//		address, err := sdk.Deployer.DeployEditionDrop(
//	     context.Background(),
//			&thirdweb.DeployEditionDropMetadata{
//				Name: "Go Edition Drop",
//			}
//		})
func (deployer *ContractDeployer) DeployEditionDrop(ctx context.Context, metadata *DeployEditionDropMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract(ctx, "edition-drop", metadata)
}

// Deploy a new Multiwrap contract.
//
// metadata: the contract metadata
//
// returns: the address of the deployed contract
//
// Example
//
//		address, err := sdk.Deployer.DeployMultiwrap(
//	     context.Background()
//			&thirdweb.DeployMultiwrapMetadata{
//				Name: "Go Multiwrap",
//			}
//		})
func (deployer *ContractDeployer) DeployMultiwrap(ctx context.Context, metadata *DeployMultiwrapMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract(ctx, "multiwrap", metadata)
}

// Deploy a new Marketplace contract.
//
// metadata: the contract metadata
//
// returns: the address of the deployed contract
//
// Example
//
//		address, err := sdk.Deployer.DeployMarketplace(
//	     context.Background()
//			&thirdweb.DeployMarketplaceMetadata{
//				Name: "Go Marketplace",
//			}
//		})
func (deployer *ContractDeployer) DeployMarketplace(ctx context.Context, metadata *DeployMarketplaceMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract(ctx, "marketplace", metadata)
}

func (deployer *ContractDeployer) deployContract(ctx context.Context, contractType string, metadata interface{}) (string, error) {
	metadataToUpload := map[string]interface{}{}
	err := mapstructure.Decode(metadata, &metadataToUpload)
	if err != nil {
		return "", err
	}

	// Set merkle default to {} for drop contracts
	if _, ok := metadataToUpload["merkle"]; ok {
		metadataToUpload["merkle"] = map[string]interface{}{}
	}

	contractUri, err := deployer.storage.Upload(
		metadataToUpload,
		deployer.helper.getAddress().String(),
		deployer.helper.GetSignerAddress().String(),
	)

	// fmt.Println(contractUri)
	// panic("Error!")

	if err != nil {
		return "", err
	}

	encodedType, err := deployer.getEncodedType(contractType)
	if err != nil {
		return "", err
	}

	contractAbi, err := deployer.getContractAbiByContractType(contractType)
	if err != nil {
		return "", err
	}

	deployArguments, err := deployer.getDeployArguments(contractType, metadata, contractUri)
	if err != nil {
		return "", err
	}

	encodedFunc, err := contractAbi.Pack("initialize", deployArguments...)

	opts, err := deployer.helper.getTxOptions(ctx)
	if err != nil {
		return "", err
	}

	tx, err := deployer.factory.DeployProxy(opts, encodedType, encodedFunc)
	if err != nil {
		return "", err
	}

	receipt, err := deployer.helper.awaitTx(tx.Hash())
	if err != nil {
		return "", err
	}

	txReceipt, err := deployer.GetProvider().TransactionReceipt(context.Background(), receipt.Hash())
	if err != nil {
		return "", err
	}

	for _, log := range txReceipt.Logs {
		event, err := deployer.factory.ParseProxyDeployed(*log)
		if err != nil {
			continue
		}

		return event.Proxy.String(), nil
	}

	return "", errors.New("No ProxyDeployed event found")
}

func (deployer *ContractDeployer) getEncodedType(contractType string) ([32]byte, error) {
	var remoteName string

	switch contractType {
	case "nft-collection":
		remoteName = "TokenERC721"
	case "edition":
		remoteName = "TokenERC1155"
	case "token":
		remoteName = "TokenERC20"
	case "nft-drop":
		remoteName = "DropERC721"
	case "edition-drop":
		remoteName = "DropERC1155"
	case "multiwrap":
		remoteName = "Multiwrap"
	case "marketplace":
		remoteName = "Marketplace"
	default:
		return [32]byte{}, fmt.Errorf("Unsupported contract type: %s", contractType)
	}

	contractBytes := []byte(remoteName)
	encodedType := [32]byte{}
	copy(encodedType[:], contractBytes)

	return encodedType, nil
}

func (deployer *ContractDeployer) getDeployArguments(contractType string, metadata interface{}, contractUri string) ([]interface{}, error) {
	trustedForwarders, err := deployer.getDefaultTrustForwarders()
	if err != nil {
		return nil, err
	}

	switch contractType {
	case "nft-collection":
		meta, ok := metadata.(*DeployNFTCollectionMetadata)
		if !ok {
			return nil, fmt.Errorf("Unsupported metadata type for contract type: %s", contractType)
		}

		return []interface{}{
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
			common.HexToAddress(meta.PlatformFeeRecipient),
		}, nil
	case "nft-drop":
		meta, ok := metadata.(*DeployNFTDropMetadata)
		if !ok {
			return nil, fmt.Errorf("Unsupported metadata type for contract type: %s", contractType)
		}

		return []interface{}{
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
			common.HexToAddress(meta.PlatformFeeRecipient),
		}, nil
	case "edition":
		meta, ok := metadata.(*DeployEditionMetadata)
		if !ok {
			return nil, fmt.Errorf("Unsupported metadata type for contract type: %s", contractType)
		}

		return []interface{}{
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
			common.HexToAddress(meta.PlatformFeeRecipient),
		}, nil
	case "edition-drop":
		meta, ok := metadata.(*DeployEditionDropMetadata)
		if !ok {
			return nil, fmt.Errorf("Unsupported metadata type for contract type: %s", contractType)
		}

		return []interface{}{
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
			common.HexToAddress(meta.PlatformFeeRecipient),
		}, nil
	case "token":
		meta, ok := metadata.(*DeployTokenMetadata)
		if !ok {
			return nil, fmt.Errorf("Unsupported metadata type for contract type: %s", contractType)
		}

		return []interface{}{
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.PlatformFeeRecipient),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
		}, nil
	case "multiwrap":
		meta, ok := metadata.(*DeployMultiwrapMetadata)
		if !ok {
			return nil, err
		}

		return []interface{}{
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
		}, nil
	case "marketplace":
		meta, ok := metadata.(*DeployMarketplaceMetadata)
		if !ok {
			return nil, err
		}

		return []interface{}{
			deployer.GetSignerAddress(),
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PlatformFeeRecipient),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
		}, nil
	default:
		return nil, fmt.Errorf("Unsupported contract type: %s", contractType)
	}
}

func (deployer *ContractDeployer) getDefaultTrustForwarders() ([]common.Address, error) {
	chainId, err := deployer.GetProvider().ChainID(context.Background())
	if err != nil {
		return []common.Address{}, err
	}

	biconomyForwarder, err := getContractAddressByChainId(ChainID(chainId.Int64()), "BiconomyForwarder")
	if err != nil {
		return []common.Address{}, err
	}

	if biconomyForwarder != zeroAddress {
		return []common.Address{
			common.HexToAddress(ozDefenderForwarderAddress),
			common.HexToAddress(biconomyForwarder),
		}, nil
	}

	return []common.Address{
		common.HexToAddress(ozDefenderForwarderAddress),
	}, nil
}

func (deployer *ContractDeployer) getContractAbiByContractType(contractType string) (*gethAbi.ABI, error) {
	var contractAbi string

	switch contractType {
	case "nft-collection":
		contractAbi = abi.TokenERC721ABI
	case "edition":
		contractAbi = abi.TokenERC1155ABI
	case "token":
		contractAbi = abi.TokenERC20ABI
	case "nft-drop":
		contractAbi = abi.DropERC721ABI
	case "edition-drop":
		contractAbi = abi.DropERC1155ABI
	case "multiwrap":
		contractAbi = abi.MultiwrapABI
	case "marketplace":
		contractAbi = abi.MarketplaceABI
	default:
		return nil, fmt.Errorf("Unsupported contract type: %s", contractType)
	}

	parsedAbi, err := gethAbi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	return &parsedAbi, nil
}
