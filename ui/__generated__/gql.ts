/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  mutation deleteThread($id: ID!) {\n    deleteThread(id: $id)\n  }\n": types.DeleteThreadDocument,
    "\n  mutation updateThread($id: ID!, $name: String!) {\n    updateThread(id: $id, input: { name: $name }) {\n      id\n      name\n    }\n  }\n": types.UpdateThreadDocument,
    "\n  query sidebarThreadsQuery {\n    user {\n      id\n      email\n      threads(first: 50, orderBy: { field: LAST_VIEWED_AT, direction: DESC }) {\n        totalCount\n        pageInfo {\n          hasNextPage\n          hasPreviousPage\n          startCursor\n          endCursor\n        }\n        edges {\n          node {\n            id\n            name\n          }\n        }\n      }\n    }\n  }\n": types.SidebarThreadsQueryDocument,
    "\n  query thread($id: ID!) {\n    thread(id: $id) {\n      id\n      createdAt\n      name\n      parent {\n        id\n        thread {\n          id\n          name\n        }\n      }\n      messages {\n        totalCount\n        pageInfo {\n          hasNextPage\n          hasPreviousPage\n          startCursor\n          endCursor\n        }\n        edges {\n          node {\n            id\n            createdAt\n            content\n            messageType\n            bookmarks {\n              edges {\n                node {\n                  id\n                }\n              }\n            }\n            child {\n              id\n              name\n            }\n          }\n        }\n      }\n    }\n  }\n": types.ThreadDocument,
    "\n  query threadByMessageIdQuery($messageId: ID!) {\n  threads(where: {hasMessagesWith: {id: $messageId}}) {\n    edges {\n      node {\n        id\n        createdAt\n        name\n        messages {\n          totalCount\n          pageInfo {\n            hasNextPage\n            hasPreviousPage\n            startCursor\n            endCursor\n          }\n          edges {\n            node {\n              id\n              createdAt\n              content\n              messageType\n              bookmarks {\n                edges {\n                  node {\n                    id\n                  }\n                }\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}\n": types.ThreadByMessageIdQueryDocument,
    "\n  mutation createThread(\n    $input: CreateThreadInput!\n    $messageInput: CreateMessageInput!\n  ) {\n    createThread(input: $input, messageInput: $messageInput) {\n      id\n      lastViewedAt\n      messages {\n        edges {\n          node {\n            id\n            createdAt\n            content\n          }\n        }\n      }\n    }\n  }\n": types.CreateThreadDocument,
    "\n  mutation createMessage($input:CreateMessageInput!) {\n    createMessage(input:$input){\n      id\n      content\n      createdAt\n    }\n  }\n": types.CreateMessageDocument,
    "\nquery mostRecentMessage($threadId: ID!, $after:Time!) {\n  messages(\n    where: {hasThreadWith: {id: $threadId}, createdAtGT:$after}\n    first: 1\n    orderBy: {field: CREATED_AT, direction: DESC}\n  ) {\n    edges {\n      node {\n        id\n        content\n        messageType\n      }\n    }\n  }\n}\n": types.MostRecentMessageDocument,
    "\nquery bookmarks($threadId:ID!){\n  bookmarks(where:{hasMessageWith:{hasThreadWith:{id:$threadId}}}) {\n    edges {\n      node {\n        id\n        createdAt\n        message {\n          id\n          createdAt\n          content\n        }\n      }\n    }\n  }\n}\n": types.BookmarksDocument,
    "\nmutation createBookmark($messageId:ID!){\n  createBookmark(input:{messageID:$messageId}){\n    id\n    createdAt\n    message {\n      id\n      createdAt\n      content\n    }\n  }\n}": types.CreateBookmarkDocument,
    "\nmutation deleteBookmark($id:ID!) {\n  deleteBookmark(id:$id)\n}": types.DeleteBookmarkDocument,
    "\n  query searchResultsQuery($query: String!) {\n    messages(where: { contentContainsFold: $query }) {\n      totalCount\n      edges {\n        node {\n          id\n          createdAt\n          content\n          messageType\n          thread {\n            id\n            name\n            parent {\n              id\n              thread {\n               id\n               name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n": types.SearchResultsQueryDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation deleteThread($id: ID!) {\n    deleteThread(id: $id)\n  }\n"): (typeof documents)["\n  mutation deleteThread($id: ID!) {\n    deleteThread(id: $id)\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation updateThread($id: ID!, $name: String!) {\n    updateThread(id: $id, input: { name: $name }) {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  mutation updateThread($id: ID!, $name: String!) {\n    updateThread(id: $id, input: { name: $name }) {\n      id\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query sidebarThreadsQuery {\n    user {\n      id\n      email\n      threads(first: 50, orderBy: { field: LAST_VIEWED_AT, direction: DESC }) {\n        totalCount\n        pageInfo {\n          hasNextPage\n          hasPreviousPage\n          startCursor\n          endCursor\n        }\n        edges {\n          node {\n            id\n            name\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query sidebarThreadsQuery {\n    user {\n      id\n      email\n      threads(first: 50, orderBy: { field: LAST_VIEWED_AT, direction: DESC }) {\n        totalCount\n        pageInfo {\n          hasNextPage\n          hasPreviousPage\n          startCursor\n          endCursor\n        }\n        edges {\n          node {\n            id\n            name\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query thread($id: ID!) {\n    thread(id: $id) {\n      id\n      createdAt\n      name\n      parent {\n        id\n        thread {\n          id\n          name\n        }\n      }\n      messages {\n        totalCount\n        pageInfo {\n          hasNextPage\n          hasPreviousPage\n          startCursor\n          endCursor\n        }\n        edges {\n          node {\n            id\n            createdAt\n            content\n            messageType\n            bookmarks {\n              edges {\n                node {\n                  id\n                }\n              }\n            }\n            child {\n              id\n              name\n            }\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query thread($id: ID!) {\n    thread(id: $id) {\n      id\n      createdAt\n      name\n      parent {\n        id\n        thread {\n          id\n          name\n        }\n      }\n      messages {\n        totalCount\n        pageInfo {\n          hasNextPage\n          hasPreviousPage\n          startCursor\n          endCursor\n        }\n        edges {\n          node {\n            id\n            createdAt\n            content\n            messageType\n            bookmarks {\n              edges {\n                node {\n                  id\n                }\n              }\n            }\n            child {\n              id\n              name\n            }\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query threadByMessageIdQuery($messageId: ID!) {\n  threads(where: {hasMessagesWith: {id: $messageId}}) {\n    edges {\n      node {\n        id\n        createdAt\n        name\n        messages {\n          totalCount\n          pageInfo {\n            hasNextPage\n            hasPreviousPage\n            startCursor\n            endCursor\n          }\n          edges {\n            node {\n              id\n              createdAt\n              content\n              messageType\n              bookmarks {\n                edges {\n                  node {\n                    id\n                  }\n                }\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"): (typeof documents)["\n  query threadByMessageIdQuery($messageId: ID!) {\n  threads(where: {hasMessagesWith: {id: $messageId}}) {\n    edges {\n      node {\n        id\n        createdAt\n        name\n        messages {\n          totalCount\n          pageInfo {\n            hasNextPage\n            hasPreviousPage\n            startCursor\n            endCursor\n          }\n          edges {\n            node {\n              id\n              createdAt\n              content\n              messageType\n              bookmarks {\n                edges {\n                  node {\n                    id\n                  }\n                }\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createThread(\n    $input: CreateThreadInput!\n    $messageInput: CreateMessageInput!\n  ) {\n    createThread(input: $input, messageInput: $messageInput) {\n      id\n      lastViewedAt\n      messages {\n        edges {\n          node {\n            id\n            createdAt\n            content\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation createThread(\n    $input: CreateThreadInput!\n    $messageInput: CreateMessageInput!\n  ) {\n    createThread(input: $input, messageInput: $messageInput) {\n      id\n      lastViewedAt\n      messages {\n        edges {\n          node {\n            id\n            createdAt\n            content\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createMessage($input:CreateMessageInput!) {\n    createMessage(input:$input){\n      id\n      content\n      createdAt\n    }\n  }\n"): (typeof documents)["\n  mutation createMessage($input:CreateMessageInput!) {\n    createMessage(input:$input){\n      id\n      content\n      createdAt\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery mostRecentMessage($threadId: ID!, $after:Time!) {\n  messages(\n    where: {hasThreadWith: {id: $threadId}, createdAtGT:$after}\n    first: 1\n    orderBy: {field: CREATED_AT, direction: DESC}\n  ) {\n    edges {\n      node {\n        id\n        content\n        messageType\n      }\n    }\n  }\n}\n"): (typeof documents)["\nquery mostRecentMessage($threadId: ID!, $after:Time!) {\n  messages(\n    where: {hasThreadWith: {id: $threadId}, createdAtGT:$after}\n    first: 1\n    orderBy: {field: CREATED_AT, direction: DESC}\n  ) {\n    edges {\n      node {\n        id\n        content\n        messageType\n      }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery bookmarks($threadId:ID!){\n  bookmarks(where:{hasMessageWith:{hasThreadWith:{id:$threadId}}}) {\n    edges {\n      node {\n        id\n        createdAt\n        message {\n          id\n          createdAt\n          content\n        }\n      }\n    }\n  }\n}\n"): (typeof documents)["\nquery bookmarks($threadId:ID!){\n  bookmarks(where:{hasMessageWith:{hasThreadWith:{id:$threadId}}}) {\n    edges {\n      node {\n        id\n        createdAt\n        message {\n          id\n          createdAt\n          content\n        }\n      }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nmutation createBookmark($messageId:ID!){\n  createBookmark(input:{messageID:$messageId}){\n    id\n    createdAt\n    message {\n      id\n      createdAt\n      content\n    }\n  }\n}"): (typeof documents)["\nmutation createBookmark($messageId:ID!){\n  createBookmark(input:{messageID:$messageId}){\n    id\n    createdAt\n    message {\n      id\n      createdAt\n      content\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nmutation deleteBookmark($id:ID!) {\n  deleteBookmark(id:$id)\n}"): (typeof documents)["\nmutation deleteBookmark($id:ID!) {\n  deleteBookmark(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query searchResultsQuery($query: String!) {\n    messages(where: { contentContainsFold: $query }) {\n      totalCount\n      edges {\n        node {\n          id\n          createdAt\n          content\n          messageType\n          thread {\n            id\n            name\n            parent {\n              id\n              thread {\n               id\n               name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query searchResultsQuery($query: String!) {\n    messages(where: { contentContainsFold: $query }) {\n      totalCount\n      edges {\n        node {\n          id\n          createdAt\n          content\n          messageType\n          thread {\n            id\n            name\n            parent {\n              id\n              thread {\n               id\n               name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;