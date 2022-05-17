.PHONY: abi

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out internal/abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out internal/abi/token_erc1155.go --type TokenERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721.json --out internal/abi/drop_erc721.go --type DropERC721

cmd: FORCE
	cd cmd/thirdweb && go build -o ../../bin/thirdweb && cd -

FORCE: ;