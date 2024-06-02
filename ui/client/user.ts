import { gql } from "@apollo/client";

import { query } from "@/client/ApolloClient";

export async function getUser() {
  return query(
    gql`
      query homeUserQuery {
        user {
          id
          email
          threads(
            first: 50
            orderBy: { field: LAST_VIEWED_AT, direction: DESC }
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
                id
                name
              }
            }
          }
        }
      }
    `,
    {},
  );
}
