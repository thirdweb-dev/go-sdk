.PHONY: abi test docs publish

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC20.json --out internal/abi/token_erc20.go --type TokenERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out internal/abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out internal/abi/token_erc1155.go --type TokenERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721.json --out internal/abi/drop_erc721.go --type DropERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC1155.json --out internal/abi/drop_erc1155.go --type DropERC1155

	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC20.json --out internal/abi/ierc20.go --type IERC20

docs:
	rm -rf docs
	mkdir docs
	gomarkdoc --output docs/doc.md --repository.default-branch main ./pkg/thirdweb
	node ./scripts/generate-docs.mjs
	rm ./docs/doc.md ./docs/start.md ./docs/finish.md

cmd: FORCE
	cd cmd/thirdweb && go build -o ../../bin/thirdweb && cd -

test-nft-read:
	./bin/thirdweb nft getAll -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nft getOwned -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-write:
	./bin/thirdweb nft mint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nft mintLink -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-read:
	./bin/thirdweb edition getAll -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb edition getOwned -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-write:
	./bin/thirdweb edition mint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-read:
	./bin/thirdweb token get -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-write:
	./bin/thirdweb token mint -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb token mintBatch -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-read:
	./bin/thirdweb nftdrop getAll -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

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

test:
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
	make test-storage

publish:
	# Make sure to pass the TAG variable to this command ex: `make publish TAG=v2.0.0`
	# Publish following https://go.dev/doc/modules/publishing
	go mod tidy
	git tag $(TAG)
	git push origin $(TAG)
	GOPROXY=proxy.golang.org go list -m github.com/thirdweb-dev/go-sdk@$(TAG)

FORCE: ;