import { gql } from 'graphql-request'

export const bookmarkFragment = gql`
fragment Bookmark on Bookmark {
  id
  createdAt
  updatedAt
  user {
    id
    createdAt
    updatedAt
    email
  }
  thread {
    id
    createdAt
    updatedAt
    name
  }
  message {
    id
    createdAt
    updatedAt
    content
    thread {
      id
      createdAt
      updatedAt
      name
    }
  }
  response {
    id
    createdAt
    updatedAt
    content
    message {
      id
      createdAt
      updatedAt
      content
      thread {
        id
        createdAt
        updatedAt
        name
      }
    }
  }
}
`

export const bookmarks = gql`
  query bookmarks($after: Cursor, $first: Int, $before: Cursor, $last: Int, $orderBy: [BookmarkOrder!], $where: BookmarkWhereInput) {
    bookmarks(
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
        node {
          ...Bookmark
        }
        cursor
      }
    }
  }
`

export const bookmark = gql`
  query bookmark($id:ID!){
    bookmarks(where:{id:$id}){
      edges {
        node {
          ...Bookmark
        }
      }
    }
  }
`

export const createBookmark = gql`
  mutation createBookmark($input:CreateBookmarkInput!) {
    createBookmark(input:$input){
      id
    }
  }
`

export const deleteBookmark = gql`
  mutation deleteBookmark($id: ID!) {
    deleteBookmark(id: $id)
  }
`
