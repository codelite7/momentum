"use client";
import { graphql } from "react-relay";

const ThreadButtonsQuery = graphql`
  query ThreadButtonsQuery {
    ...ThreadButtonsPaginationFragment
  }
`;

const ThreadButtonsPaginationFragment = graphql`
  fragment ThreadButtonsPaginationFragment on Query
  @refetchable(queryName: "ThreadButtonsPaginationQuery")
  @argumentDefinitions(
    first: { type: "Int", defaultValue: 50 }
    after: { type: "Cursor" }
  ) {
    threads(first: $first, after: $after) @connection(key: "Query_threads") {
      totalCount
      edges {
        node {
          id
          ...ThreadButtonFragment
        }
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
`;

export default function ThreadButtons() {
  return <></>;
  // const loaderRef = useRef(null);
  // const queryData = useLazyLoadQuery<ThreadButtonsQueryType>(
  //   ThreadButtonsQuery,
  //   {},
  // );
  // const { loadNext, hasNext } = usePaginationFragment(
  //   ThreadButtonsPaginationFragment,
  //   queryData,
  // );
  // const data = useFragment<ThreadButtonsPaginationFragment$key>(
  //   ThreadButtonsPaginationFragment,
  //   queryData,
  // );
  // const maybeLoadMore = (num: number) => {
  //   if (hasNext) {
  //     loadNext(num);
  //   }
  // };
  //
  // return (
  //   <InfiniteScroller
  //     hideScrollBar
  //     onScrollDown={() => maybeLoadMore(50)}
  //     // onScrollUp={() => console.log("up")}
  //   >
  //     {data.threads.edges?.map(
  //       (edge) =>
  //         edge?.node && <ThreadButton key={edge.node?.id} thread={edge.node} />,
  //     )}
  //   </InfiniteScroller>
  // );
}
