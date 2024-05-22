import { gql } from 'graphql-request'
export const threadMessageFragment = gql`
fragment ThreadMessage on Message {
  id
  createdAt
  updatedAt
  content
  thread {
    id
    createdAt
    updatedAt
  }
  sentBy {
    id
    createdAt
    updatedAt
    email
  }
  response {
    id
    createdAt
    updatedAt
    content
    sentBy {
      id
      provider
      model
    }
  }
  bookmarks(where:{hasUserWith:{id:$userId}}) {
    edges {
      node {
        ...Bookmark
      }
    }
  }
}
`
export const messageFragment = gql`
fragment Message on Message {
  id
  createdAt
  updatedAt
  content
  sentBy {
    id
    createdAt
    updatedAt
    email
  }
  response {
    id
    createdAt
    updatedAt
    content
    sentBy {
      id
      provider
      model
    }
  }
}
`
export const messages = gql`
query messages($after: Cursor, $first: Int, $before: Cursor, $last: Int, $orderBy: [MessageOrder!], $where: MessageWhereInput) {
  messages(
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
        ...Message
      }
      cursor
    }
  }
}
`

export const message = gql`
query message($id:ID!){
  messages(where:{id:$id}){
    edges {
      node {
        ...Message
      }
    }
  }
}
`

export const createMessage = gql`
mutation createMessage($input:CreateMessageInput!) {
  createMessage(input:$input){
    id
  }
}
`

export const threadMessages = gql`
query threadMessages($after: Cursor, $first: Int, $before: Cursor, $last: Int, $orderBy: [MessageOrder!], $threadId:ID!, $userId:ID!) {
  messages(
    after: $after
    first: $first
    before: $before
    last: $last
    orderBy: $orderBy
    where: {hasThreadWith:{id:$threadId}}
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
        ...ThreadMessage
      }

      cursor
    }
  }
}
`

export const threadMessagesStartinFrom = gql`
query threadMessagesStartingFrom($first: Int, $threadId: ID!, $startFrom:Time!, $userId:ID!) {
  messages(
    first: $first
    orderBy: [
      {
        field:CREATED_AT,
        direction:ASC
      }
    ]
    where: {
      hasThreadWith: {id: $threadId},
      createdAtGTE:$startFrom
    }
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
        ...ThreadMessage
      }
      cursor
    }
  }
}
`

export const threadMessage = gql`
query threadMessage($id:ID!,$userId:ID!) {
  messages(
    where: {id:$id}
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
        ...ThreadMessage
      }
      cursor
    }
  }
}
`

export const mostRecentMessage = gql`
query mostRecentMessage($userId:ID!) {
  messages(first:1, orderBy:[{field:CREATED_AT, direction:DESC}]) {
    edges {
      node {
        ...ThreadMessage
      }
    }
  }
}`
