"use client";

import { graphql } from "react-relay";

import homeUserQueryNode, {
  homeUserQuery,
} from "@/__generated__/homeUserQuery.graphql";
import { SerializablePreloadedQuery } from "@/relay/loadSerializableQuery";

const HomeUserQuery = graphql`
  query homeUserQuery($email: String!) {
    user(email: $email) {
      id
      email
      threads(first: 1, orderBy: { field: LAST_VIEWED_AT, direction: DESC }) {
        edges {
          node {
            id
          }
        }
      }
    }
  }
`;

type props = {
  preloadedQuery: SerializablePreloadedQuery<
    typeof homeUserQueryNode,
    homeUserQuery
  >;
};

export default function Home() {
  // const environment = useRelayEnvironment();
  // const queryRef = useSerializablePreloadedQuery(environment, preloadedQuery);
  // const data = usePreloadedQuery<homeUserQuery>(HomeUserQuery, queryRef);
  // const [commitMutation, isMutationInFlight] = useMutation(
  //   HomeCreateThreadMutation,
  // );

  // if (data.user.threads.edges?.length == 0) {
  //   console.log("no threads");
  //   // use has no threads, create one, then navigate to it
  //   // commitMutation({
  //   //   variables: {
  //   //     userId: "us01HZ8QX3T767Y529T21CTGGH9M",
  //   //   },
  //   //   onCompleted: () => {},
  //   // });
  // } else {
  //   // navigate to the latest thread
  // }

  return <div>hi</div>;
}
