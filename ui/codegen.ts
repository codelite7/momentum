import { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  schema: "http://localhost:3001/query",
  // this assumes that all your source files are in a top-level `src/` directory - you might need to adjust this to your file structure
  documents: [
    "app/**/*.{ts,tsx}",
    "client/**/*.{ts,tsx}",
    "components/**/*.{ts,tsx}",
    "graphql-queries/**/*.{ts,tsx}",
  ],
  generates: {
    "./__generated__/": {
      preset: "client",
      presetConfig: {
        gqlTagName: "gql",
      },
    },
  },
  ignoreNoDocuments: true,
};

export default config;
