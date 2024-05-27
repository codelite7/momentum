import {graphql, useFragment, useLazyLoadQuery, usePaginationFragment} from "react-relay";
import { Button } from "@nextui-org/react";
import { startTransition } from "react";

import Thread from "@/components/left-sidebar/Thread";
import {ThreadsQuery as ThreadsQueryType, ThreadsQuery$data} from "@/__generated__/ThreadsQuery.graphql";
import {
  ThreadsPaginationFragment$data,
  ThreadsPaginationFragment$key
} from "@/__generated__/ThreadsPaginationFragment.graphql";
const ThreadsQuery = graphql`
  query ThreadsQuery {
    ...ThreadsPaginationFragment
  }
`;

const ThreadsPaginationFragment = graphql`
  fragment ThreadsPaginationFragment on Query
  @refetchable(queryName: "ThreadsPaginationQuery")
  @argumentDefinitions(
    first: { type: "Int", defaultValue: 1 }
    after: { type: "Cursor" }
  ) {
    threads(first: $first, after: $after) @connection(key: "Query_threads") {
      totalCount
      edges {
        node {
          id
          ...ThreadFragment
        }
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
`;

export default function Threads() {
  const queryData = useLazyLoadQuery<ThreadsQueryType>(ThreadsQuery, {});
  const { loadNext } = usePaginationFragment(
    ThreadsPaginationFragment,
    queryData,
  );
  const data = useFragment<ThreadsPaginationFragment$key>(ThreadsPaginationFragment, queryData)
  const onLoadMore = () =>
    startTransition(() => {
      loadNext(1);
    });

  return (
    <>
      {data.threads.edges?.map((edge) => (
        <Thread key={edge?.node?.id} thread={edge?.node} />
      ))}
      <Button fullWidth onPress={onLoadMore}>Load More</Button>
    </>
  );
}
