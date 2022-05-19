import fs from "fs";

const docPath = "./docs/doc.md"

const delimiters = {
  "## type [ERC1155]": "erc1155.md",
  "## type [ERC721]": "erc721.md",
  "## type [Edition]": "edition.md",
  "## type [NFTCollection]": "nft_collection.md",
  "## type [NFTDrop]": "nft_drop.md",
  "## type [ProviderHandler]": "provider.md",
  "## type [SDKOptions]": "sdk.md",
  "## type [WrappedToken]": "finish.md",
}

async function main() {
  const lines = fs.readFileSync(docPath).toString().split("\n");
  let file = ""
  let name = "start.md"

  for (const line of lines) {
    let matched = false;

    for (const delimiter of Object.keys(delimiters)) {
      if (line.indexOf(delimiter) == 0 && !matched) {
        matched = true;
        fs.writeFile(`./docs/${name}`, file, (err) => {
          if (err) throw err;
        })
        name = delimiters[delimiter]
        file = ""
      }
    }

    if (!matched) {
      file += "\n" + line
    }
  }


  fs.writeFile(`./docs/${name}`, file, (err) => {
    if (err) throw err;
  })
}

main()
