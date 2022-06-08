package thirdweb

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"

	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
)

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

func (deployer *ContractDeployer) DeployNFTCollection(metadata *NFTCollectionContractMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract("nft-collection", metadata)
}

func (deployer *ContractDeployer) DeployEditionCollection(metadata *EditionContractMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract("edition", metadata)
}

func (deployer *ContractDeployer) DeployToken(metadata *TokenContractMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract("token", metadata)
}

func (deployer *ContractDeployer) DeployNFTDropCollection(metadata *NFTDropContractMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract("nft-drop", metadata)
}

func (deployer *ContractDeployer) DeployEditionDropCollection(metadata *EditionDropContractMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract("edition-drop", metadata)
}

func (deployer *ContractDeployer) DeployMultiwrap(metadata *MultiwrapContractMetadata) (string, error) {
	metadata.fillDefaults()
	return deployer.deployContract("multiwrap", metadata)
}

func (deployer *ContractDeployer) deployContract(contractType string, metadata interface{}) (string, error) {
	contractUri, err := deployer.storage.Upload(
		metadata, deployer.helper.getAddress().String(),
		deployer.helper.GetSignerAddress().String(),
	)
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

	opts, err := deployer.helper.getTxOptions()
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
		meta, ok := metadata.(*NFTCollectionContractMetadata)
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
		meta, ok := metadata.(*NFTDropContractMetadata)
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
		meta, ok := metadata.(*EditionContractMetadata)
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
		meta, ok := metadata.(*EditionDropContractMetadata)
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
		meta, ok := metadata.(*TokenContractMetadata)
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
		meta, ok := metadata.(*MultiwrapContractMetadata)
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
	default:
		return nil, fmt.Errorf("Unsupported contract type: %s", contractType)
	}

	parsedAbi, err := gethAbi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	return &parsedAbi, nil
}
