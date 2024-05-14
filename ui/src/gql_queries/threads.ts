import { gql } from 'graphql-request'
export const threadFragment = gql`
fragment Thread on Thread {
  id
  createdAt
  updatedAt
  name
  parent {
    id
    createdAt
    updatedAt
    name
  }
  children {
    edges {
      node {
        id
        createdAt
        updatedAt
        name
      }
    }
  }
}
`

export const threads = gql`
  query threads($after: Cursor, $first: Int, $before: Cursor, $last: Int, $orderBy: [ThreadOrder!], $where: ThreadWhereInput) {
    threads(
      after: $after
      first: $first
      before: $before
      last: $last
      orderBy: $orderBy
      where: $where
    ) {
      totalCount
      pageInfo {
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      edges {
        node{
          ...Thread
        }
        cursor
      }
    }
  }
`

export const thread = gql`
  query thread($id:ID!){
    threads(where:{id:$id}){
      edges {
        node{
          ...Thread
        }
      }
    }
  }
`

export const createThread = gql`
  mutation createThread($input:CreateThreadInput!) {
    createThread(input:$input){
      id
    }
  }
`
