"use server";

import { gql } from "@apollo/client";

import { mutation, query } from "@/client/ApolloClient";

export async function deleteThread(id: string) {
  return mutation(
    gql`
      mutation deleteThread($id: ID!) {
        deleteThread(id: $id)
      }
    `,
    { id },
  );
}
export async function renameThread(id: string, name: string) {
  return mutation(
    gql`
      mutation updateThread($id: ID!, $input: UpdateThreadInput!) {
        updateThread(id: $id, input: $input) {
          id
          name
        }
      }
    `,
    {
      id: id,
      input: { name: name },
    },
  );
}

export async function createThread() {
  return mutation(
    gql`
      mutation createThread {
        createThread(input: { name: "New Thread" }) {
          id
        }
      }
    `,
    {},
  );
}

export async function getThread(id: string) {
  console.log(
    "fetching thread.........................................................................................................................................................................................................................................................................................................",
  );

  return query(
    gql`
      query thread($id: ID!) {
        thread(id: $id) {
          id
          createdAt
          name
          messages(first: 5) {
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
                response {
                  id
                  content
                }
              }
            }
          }
        }
      }
    `,
    { id },
  );
}
