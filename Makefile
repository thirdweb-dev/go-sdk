.PHONY: abi test docs

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC20.json --out internal/abi/token_erc20.go --type TokenERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out internal/abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out internal/abi/token_erc1155.go --type TokenERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721.json --out internal/abi/drop_erc721.go --type DropERC721

	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC20.json --out internal/abi/ierc20.go --type IERC20

docs:
	rm -rf docs
	mkdir docs
	gomarkdoc --output docs/doc.md ./pkg/thirdweb
	node ./scripts/generate-docs.mjs
	rm ./docs/doc.md ./docs/start.md ./docs/finish.md

cmd: FORCE
	cd cmd/thirdweb && go build -o ../../bin/thirdweb && cd -

test-nft-read:
	./bin/thirdweb nft getAll -a ${GO_NFT_COLLECTION} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nft getOwned -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-write:
	./bin/thirdweb nft mint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-read:
	./bin/thirdweb edition getAll -a ${GO_EDITION} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb edition getOwned -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-write:
	./bin/thirdweb edition mint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-read:
	./bin/thirdweb nftdrop getAll -a ${GO_NFT_DROP} -u ${GO_ALCHEMY_RPC}

test-drop-write:
	./bin/thirdweb nftdrop createBatch -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/thirdweb nftdrop claim -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-storage:
	./bin/thirdweb storage upload
	./bin/thirdweb storage uploadBatch

test:
	make cmd
	make test-nft-read
	make test-nft-write
	make test-edition-read
	make test-edition-write
	make test-drop-read
	make test-drop-write
	make test-storage

FORCE: ;