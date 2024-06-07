"use client";

import { Card, CardBody, CardHeader } from "@nextui-org/react";
import { useAtom, useAtomValue } from "jotai";
import { useQuery } from "@apollo/client";
import { toast } from "sonner";

import { Bookmark as BookmarkType } from "@/__generated__/graphql";
import InfiniteScroller from "@/components/common/scroll/infinite-scroller";
import Bookmark from "@/components/right-sidebar/bookmark";
import { threadAtom, threadBookmarksAtom } from "@/state/atoms";
import { threadBookmarksQuery } from "@/graphql-queries/queries";

export default function Bookmarks() {
  const thread = useAtomValue(threadAtom);
  const [bookmarks, setBookmarks] = useAtom(threadBookmarksAtom);

  const { loading } = useQuery(threadBookmarksQuery, {
    notifyOnNetworkStatusChange: true,
    variables: { threadId: thread!.id },
    onError: (e) => {
      toast.error("Error getting bookmarks");
      console.error(e);
    },
    onCompleted: (data) => {
      let bookmarks: BookmarkType[] = [];

      data.bookmarks?.edges?.forEach((edge) => {
        if (edge?.node) {
          bookmarks.push(edge.node as BookmarkType);
        }
      });
      setBookmarks(bookmarks);
    },
  });

  return (
    <>
      {loading ? (
        <></>
      ) : bookmarks.length > 0 ? (
        <InfiniteScroller hideScrollBar className="h-full">
          {bookmarks?.map((bookmark) => (
            <Bookmark key={bookmark.id} bookmark={bookmark} />
          ))}
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
