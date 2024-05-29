import { graphql, useLazyLoadQuery, usePaginationFragment } from "react-relay";
import { Button } from "@nextui-org/react";

import { testQuery } from "@/__generated__/testQuery.graphql";
import { testPaginationFragment$key } from "@/__generated__/testPaginationFragment.graphql";
import InfiniteScroller from "@/components/common/scroll/infinite-scroller";

const TestQuery = graphql`
  query testQuery {
    ...testPaginationFragment
  }
`;

const TestPaginationFragment = graphql`
  fragment testPaginationFragment on Query
  @refetchable(queryName: "TestPaginationQuery")
  @argumentDefinitions(
    first: { type: "Int", defaultValue: 3 }
    after: { type: "Cursor" }
  ) {
    messages(first: $first, after: $after) @connection(key: "Query_messages") {
      totalCount
      edges {
        node {
          id
          ...testMessageFragment
        }
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
`;

export const TestMessageFragment = graphql`
  fragment testMessageFragment on Message {
    id
    createdAt
    content
  }
`;

export default function Test() {
  const queryData = useLazyLoadQuery<testQuery>(TestQuery, {});
  const { data, loadNext, hasNext } = usePaginationFragment<
    testQuery,
    testPaginationFragment$key
  >(TestPaginationFragment, queryData);

  const maybeLoadMore = (num: number) => {
    if (hasNext) {
      loadNext(num);
    }
  };

  return (
    <>
      <InfiniteScroller
        hideScrollBar
        onScrollDown={() => maybeLoadMore(1)}
        // onScrollUp={() => console.log("up")}
      >
        {data.messages.edges?.map(
          (edge) => edge?.node && <div key={edge.node.id}>{edge.node.id}</div>,
        )}
      </InfiniteScroller>
      <Button onPress={() => maybeLoadMore(1)}>Load More</Button>
    </>
  );
}
