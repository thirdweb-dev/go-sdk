import fs from "fs";

const BASE_DOC_URL = "https://docs.thirdweb.com/go/";

// Setup for contract snippets
// Name must match contract name for TS and Python
// The first code block of first file in files will be the contract setup script
// docname corresponds with name of file page in docs repo
const contracts = [
  {
    name: "EditionDrop",
    files: ["edition_drop.md", "erc1155_standard.md"],
    docName: "edition_drop",
  },
  {
    name: "Edition",
    files: ["edition.md", "erc1155.md"],
    docName: "edition",
  },
  {
    name: "NFTCollection",
    files: ["nft_collection.md", "erc721_standard.md"],
    docName: "nft_collection",
  },
  {
    name: "NFTDrop",
    files: ["nft_drop.md", "erc721_standard.md"],
    docName: "nft_drop",
  },
  {
    name: "Token",
    files: ["token.md", "erc20_standard.md"],
    docName: "token",
  },
  {
    name: "Multiwrap",
    files: ["multiwrap.md", "erc721_standard.md"],
    docName: "multiwrap",
  },
  {
    name: "Marketplace",
    files: ["marketplace.md"],
    docName: "marketplace",
  },
  {
    name: "WalletAuthenticator",
    files: ["wallet_authenticator.md"],
    docName: "wallet_authenticator",
  },
  {
    name: "ContractEvents",
    files: ["contract_events.md"],
    docName: "contract_events",
  },
  {
    name: "ERC20",
    files: ["erc20.md"],
    docName: "erc20",
  },
  {
    name: "ERC721",
    files: ["erc721.md"],
    docName: "erc721",
  },
  {
    name: "ERC1155",
    files: ["erc1155.md"],
    docName: "erc1155",
  },
];

async function generateSnippets() {
  const snippets = {};

  contracts.map((contractData) => {
    const contractName = contractData.name;
    const contractReference = `${BASE_DOC_URL}${contractData.docName}`;

    const codeBlocks = fs
      .readFileSync(`./docs/${contractData.files[0]}`)
      .toString()
      .match(/\`\`\`(\n|.)+?\`\`\`/);
    const setupExample = codeBlocks?.length
      ? codeBlocks[0].replaceAll("```", "")
      : "";

    const contract = {
      name: contractName,
      summary: "",
      description: "",
      example: setupExample,
      methods: [],
      properties: [],
      reference: contractReference,
    };

    contractData.files.map((fileName) => {
      const functions = fs
        .readFileSync(`./docs/${fileName}`)
        .toString()
        .split(/### func \\\(\\.+\) \[.+\]/);
      functions.map((fn) => {
        const functionNames = fn.match(/(?<=func \(.+\) )\S+(?=\()/);
        const functionName = functionNames?.length ? functionNames[0] : "";

        const summaries = fn
          .split("#### Example")[0]
          .match(/(?<=#### )(.*)(?=\n)/);
        const summary = summaries?.length ? summaries[0] : "";

        const examples = fn.match(
          /(?<=#### Example\n\n\`\`\`\n)(\n|.)+?(?=\`\`\`)/
        );
        const example = examples?.length ? examples[0] : "";

        if (functionName) {
          const reference = `${contractReference}#func-${contractData.docName.replace(
            "_",
            ""
          )}-${functionName.toLowerCase()}`;

          const extensionConfig = fn.match(/(?<=@extension: )(.*)(?=\n)/);
          const extensions = [];
          if (extensionConfig?.length) {
            extensionConfig[0]
              .split(" | ")
              .forEach((extension) => extensions.push(extension));
          }

          const fnData = {
            name: functionName,
            summary,
            example,
            reference,
            extensions,
          };

          contract.methods.push(fnData);
        }
      });
    });

    snippets[contractName] = contract;
  });

  fs.writeFileSync(
    `./docs/snippets.json`,
    JSON.stringify(snippets, null, 2),
    (err) => {
      if (err) throw err;
    }
  );
}

async function generateFeatureSnippets() {
  const snippets = JSON.parse(fs.readFileSync("./docs/snippets.json"));

  const methods = [];
  for (const cls of Object.values(snippets)) {
    for (const method of cls.methods) {
      if (method.extensions.length > 0) {
        methods.push({
          name: method.name,
          summary: method.summary,
          examples: {
            go: method.example,
          },
          reference: {
            go: method.reference,
          },
          extensions: method.extensions,
        });
      }
    }
  }

  const features = {};
  for (const method of methods) {
    for (const extension of method.extensions) {
      const cleanedMethod = {
        name: method.name,
        summary: method.summary,
        examples: method.examples,
        reference: method.reference,
      };

      if (Object.keys(features).includes(extension)) {
        features[extension].push(cleanedMethod);
      } else {
        features[extension] = [cleanedMethod];
      }
    }
  }

  fs.writeFileSync(
    `./docs/feature_snippets.json`,
    JSON.stringify(features, null, 2),
    (err) => {
      if (err) throw err;
    }
  );
}

async function main() {
  generateSnippets();
  generateFeatureSnippets();
}

main();
