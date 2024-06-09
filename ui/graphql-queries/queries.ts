import { gql } from "@/__generated__";

export const threadByIdQuery = gql(/* Graphql */ `
  query thread($id: ID!) {
    thread(id: $id) {
      id
      createdAt
      name
      parent {
        id
        thread {
          id
          name
        }
      }
      messages {
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
            bookmarks {
              edges {
                node {
                  id
                }
              }
            }
            child {
              id
              name
            }
          }
        }
      }
    }
  }
`);

export const threadByMessageIdQuery = gql(/* Graphql */ `
  query threadByMessageIdQuery($messageId: ID!) {
  threads(where: {hasMessagesWith: {id: $messageId}}) {
    edges {
      node {
        id
        createdAt
        name
        messages {
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
              bookmarks {
                edges {
                  node {
                    id
                  }
                }
              }
            }
          }
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

export const threadBookmarksQuery = gql(`
query bookmarks($threadId:ID!){
  bookmarks(where:{hasMessageWith:{hasThreadWith:{id:$threadId}}}) {
    edges {
      node {
        id
        createdAt
        message {
          id
          createdAt
          content
        }
      }
    }
  }
}
`);

export const createBookmarkMutation = gql(/* Graphql */ `
mutation createBookmark($messageId:ID!){
  createBookmark(input:{messageID:$messageId}){
    id
    createdAt
    message {
      id
      createdAt
      content
    }
  }
}`);

export const deleteBookmarkMutation = gql(/* Graphql */ `
mutation deleteBookmark($id:ID!) {
  deleteBookmark(id:$id)
}`);

export const searchResultsQuery = gql(/* Graphql */ `
  query searchResultsQuery($query: String!) {
    messages(where: { contentContainsFold: $query }) {
      totalCount
      edges {
        node {
          id
          createdAt
          content
          messageType
          thread {
            id
            name
            parent {
              id
              thread {
               id
               name
              }
            }
          }
        }
      }
    }
  }
`);
