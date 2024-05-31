"use client";

import { graphql, useLazyLoadQuery, usePaginationFragment } from "react-relay";
import { Card, CardBody, CardHeader } from "@nextui-org/react";
import { useParams } from "next/navigation";

import { bookmarksQuery } from "@/__generated__/bookmarksQuery.graphql";
import { bookmarksFragment$key } from "@/__generated__/bookmarksFragment.graphql";
import Bookmark from "@/components/right-sidebar/bookmark";
import InfiniteScroller from "@/components/common/scroll/infinite-scroller";

export const BookmarksQuery = graphql`
  query bookmarksQuery($first: Int, $after: Cursor, $threadId: ID!) {
    __id
    ...bookmarksFragment
      @arguments(first: $first, after: $after, threadId: $threadId)
  }
`;

const BookmarksFragment = graphql`
  fragment bookmarksFragment on Query
  @argumentDefinitions(
    first: { type: "Int", defaultValue: 1 }
    after: { type: "Cursor" }
    threadId: { type: "ID!" }
  )
  @refetchable(queryName: "BookmarksPaginationQuery") {
    bookmarks(
      first: $first
      after: $after
      where: { hasMessageWith: { hasThreadWith: { id: $threadId } } }
    ) @connection(key: "Thread_bookmarks") {
      __id
      totalCount
      pageInfo {
        endCursor
        hasNextPage
      }
      edges {
        node {
          id
          ...bookmarkFragment
        }
      }
    }
  }
`;

export default function Bookmarks() {
  const params = useParams<any>();

  const queryData = useLazyLoadQuery<bookmarksQuery>(BookmarksQuery, {
    threadId: params.id,
    first: 5,
  });

  console.log("query id ", queryData.__id);

  const { data, loadNext, hasNext } = usePaginationFragment<
    bookmarksQuery,
    bookmarksFragment$key
  >(BookmarksFragment, queryData);

  console.log("bookmarks connection id", data.bookmarks.__id);
  const bookmarks = data.bookmarks;

  const maybeLoadMore = (num: number) => {
    if (hasNext) {
      loadNext(num);
    }
  };

  console.log("bookmark edges length ", bookmarks.edges);

  return (
    <>
      {bookmarks.edges && bookmarks.edges.length > 0 ? (
        <InfiniteScroller
          hideScrollBar
          className="h-full"
          onScrollDown={() => maybeLoadMore(1)}
        >
          {bookmarks.edges?.map((edge) => {
            return (
              edge?.node && <Bookmark key={edge.node.id} bookmark={edge.node} />
            );
          })}
        </InfiniteScroller>
      ) : (
        <Card className="border-1 border-primary">
          <CardHeader className="flex justify-center">
            <i className="pi pi-bookmark text-primary" />
          </CardHeader>
          <CardBody>
            <span className="text-sm text-primary-800 text-center">
              Bookmark important pieces of your conversation to collect them
              here
            </span>
          </CardBody>
        </Card>
      )}
    </>
  );
}
