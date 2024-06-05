import { gql } from "@/__generated__";

export const threadByIdQuery = gql(/* Graphql */ `
  query thread($id: ID!) {
    thread(id: $id) {
      id
      createdAt
      name
      messages(first: 50) {
        totalCount
        pageInfo {
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
        edges {
          node {
            id
            createdAt
            content
            messageType
          }
        }
      }
    }
  }
`);

export const createThreadMutation = gql(/* Graphql */ `
  mutation createThread(
    $input: CreateThreadInput!
    $messageInput: CreateMessageInput!
  ) {
    createThread(input: $input, messageInput: $messageInput) {
      id
      lastViewedAt
      messages {
        edges {
          node {
            id
            createdAt
            content
          }
        }
      }
    }
  }
`);

export const createMessageMutation = gql(/* Graphql */ `
  mutation createMessage($input:CreateMessageInput!) {
    createMessage(input:$input){
      id
      content
      createdAt
    }
  }
`);

export const getMostRecentMessage = gql(`
query mostRecentMessage($threadId: ID!, $after:Time!) {
  messages(
    where: {hasThreadWith: {id: $threadId}, createdAtGT:$after}
    first: 1
    orderBy: {field: CREATED_AT, direction: DESC}
  ) {
    edges {
      node {
        id
        content
        messageType
      }
    }
  }
}
`);
