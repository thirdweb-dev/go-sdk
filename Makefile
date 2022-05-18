.PHONY: abi test

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC20.json --out internal/abi/token_erc20.go --type TokenERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out internal/abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out internal/abi/token_erc1155.go --type TokenERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721.json --out internal/abi/drop_erc721.go --type DropERC721

	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC20.json --out internal/abi/ierc20.go --type IERC20

cmd: FORCE
	cd cmd/thirdweb && go build -o ../../bin/thirdweb && cd -

test-nft-collection:
	./bin/thirdweb nft getAll -a 0x4a945C93b79998d55E4F6103D000478B5D03aB92
	./bin/thirdweb nft getOwned -a 0x4a945C93b79998d55E4F6103D000478B5D03aB92 -k 4916d58e7ece81883cc5dd9ac8ce292460be4c5e6b0b92495c3d00f85fdb7a74
	./bin/thirdweb nft mint -a 0x4a945C93b79998d55E4F6103D000478B5D03aB92 -k 4916d58e7ece81883cc5dd9ac8ce292460be4c5e6b0b92495c3d00f85fdb7a74

test-edition:
	./bin/thirdweb edition getAll -a 0x543B4DC47C4DB12c4Cb8611f604D6c338e3D4B2E
	./bin/thirdweb edition getOwned -a 0x543B4DC47C4DB12c4Cb8611f604D6c338e3D4B2E -k 4916d58e7ece81883cc5dd9ac8ce292460be4c5e6b0b92495c3d00f85fdb7a74
	./bin/thirdweb edition mint -a 0x543B4DC47C4DB12c4Cb8611f604D6c338e3D4B2E -k 4916d58e7ece81883cc5dd9ac8ce292460be4c5e6b0b92495c3d00f85fdb7a74

test-nft-drop:
	./bin/thirdweb nftdrop getAll -a 0x3B6c493757fD7f19681aa3004E573B5d4BE62b19
	./bin/thirdweb nftdrop mint -a 0x3B6c493757fD7f19681aa3004E573B5d4BE62b19 -k 4916d58e7ece81883cc5dd9ac8ce292460be4c5e6b0b92495c3d00f85fdb7a74

test:
	make cmd
	make test-nft-collection
	make test-edition
	make test-nft-drop

FORCE: ;