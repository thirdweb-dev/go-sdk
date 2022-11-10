.PHONY: abi docs publish local-test

SHELL := /bin/bash

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC20.json --out abi/token_erc20.go --type TokenERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out abi/token_erc1155.go --type TokenERC1155
	# If you want to generate drop contracts, you'll have to delete a struct
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721.json --out abi/drop_erc721.go --type DropERC721
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC1155.json --out abi/drop_erc1155.go --type DropERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/Multiwrap.json --out abi/multiwrap.go --type Multiwrap
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/Marketplace.json --out abi/marketplace.go --type Marketplace

	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TWFactory.json --out abi/twfactory.go --type TWFactory
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC20.json --out abi/ierc20.go --type IERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC721.json --out abi/ierc721.go --type IERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC1155.json --out abi/ierc1155.go --type IERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC165.json --out abi/ierc165.go --type IERC165

docs:
	rm -rf docs
	mkdir docs
	gomarkdoc --output docs/doc.md --repository.default-branch main ./thirdweb
	node ./scripts/generate-docs.mjs
	rm ./docs/doc.md ./docs/start.md ./docs/delete.md
	node ./scripts/generate-snippets.mjs

cmd: FORCE
	cd cmd/thirdweb && go build -o ../../bin/thirdweb && cd -

test-nft-read:
	./bin/thirdweb nft getAll -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nft getOwned -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-write:
	./bin/thirdweb nft mint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nft mintLink -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-sigmint:
	./bin/thirdweb nft sigmint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-read:
	./bin/thirdweb edition getAll -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb edition getOwned -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-write:
	./bin/thirdweb edition mint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-sigmint:
	./bin/thirdweb edition sigmint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb edition sigmint-tokenid -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-read:
	./bin/thirdweb token get -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-write:
	./bin/thirdweb token mint -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb token mintBatch -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-read:
	./bin/thirdweb multiwrap getAll -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	# ./bin/thirdweb multiwrap getContents -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-write:
	./bin/thirdweb multiwrap wrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC} -n ${GO_NFT_COLLECTION} -e ${GO_EDITION} -t ${GO_TOKEN}
	./bin/thirdweb multiwrap unwrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-read:
	./bin/thirdweb nftdrop getAll -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nftdrop getActive -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-write:
	./bin/thirdweb nftdrop createBatch -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nftdrop claim -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-read:
	./bin/thirdweb editiondrop getAll -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-write:
	./bin/thirdweb editiondrop createBatch -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb editiondrop claim -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-storage:
	./bin/thirdweb storage upload
	./bin/thirdweb storage uploadBatch
	./bin/thirdweb storage uploadImage
	./bin/thirdweb storage uploadImageLink

test-custom:
	./bin/thirdweb custom set -a ${GO_CUSTOM} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-deploy:
	./bin/thirdweb deploy nft -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb deploy edition -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb deploy token -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb deploy nftdrop -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb deploy editiondrop -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb deploy multiwrap -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb deploy marketplace -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-encoder:
	./bin/thirdweb marketplace encode -a ${GO_MARKETPLACE} -k ${GO_PRIVATE_KEY} -u ${GO_AVAX_RPC}

test-cmd:
	make cmd
	make test-nft-read
	make test-nft-write
	make test-edition-read
	make test-edition-write
	make test-token-read
	make test-token-write
	make test-drop-read
	make test-drop-write
	make test-edition-drop-read
	make test-edition-drop-write
	make test-multiwrap-read
	make test-multiwrap-write
	make test-storage

stop-docker:
	docker stop hardhat-node
	docker rm hardhat-node

test: FORCE
	docker build . -t hardhat-mainnet-fork
	docker start hardhat-node || docker run --name hardhat-node -d -p 8545:8545 -e SDK_ALCHEMY_KEY=${SDK_ALCHEMY_KEY} hardhat-mainnet-fork
	sudo bash ./scripts/test/await-hardhat.sh
	go clean -testcache
	go test -v ./thirdweb
	docker stop hardhat-node
	docker rm hardhat-node

local-test:
  # Needs to be run along with npx hardhat node from this repo, and needs to be a mainnet fork hardhat
	go clean -testcache
	go test -v ./thirdweb

publish:
	# Make sure to pass the TAG variable to this command ex: `make publish TAG=v2.0.0`
	# Publish following https://go.dev/doc/modules/publishing
	go mod tidy
	git tag $(TAG)
	git push origin $(TAG)
	GOPROXY=proxy.golang.org go list -m github.com/thirdweb-dev/go-sdk@$(TAG)

FORCE: ;