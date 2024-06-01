import { gql } from "@apollo/client";

import { makeRequest } from "@/client/ApolloClient";

export async function getUser() {
  return makeRequest(
    gql`
      query homeUserQuery {
        user {
          id
          email
          threads(
            first: 1
            orderBy: { field: LAST_VIEWED_AT, direction: DESC }
          ) {
            edges {
              node {
                id
              }
            }
          }
        }
      }
    `,
    {},
  );
}
