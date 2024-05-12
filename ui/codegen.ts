import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: 'http://localhost:3000/query',
  documents: './src/gql_queries/*.ts',
  generates: {
    './graphql/generated.ts': {
      plugins: ['typescript', 'typescript-operations', 'typescript-graphql-request'],
      // config: {
      //   rawRequest: true
      // }
    }
  }
}
export default config
