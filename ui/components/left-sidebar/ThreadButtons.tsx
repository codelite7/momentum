"use client";

import { useAtomValue } from "jotai";

import ThreadButton from "./ThreadButton";

import InfiniteScroller from "@/components/common/scroll/infinite-scroller";
import { threadAtom } from "@/state/atoms";

type props = {
  threads: any;
};

export default function ThreadButtons({ threads }: props) {
  const thread = useAtomValue(threadAtom);
  const filteredThreads = threads?.edges?.filter((edge: any) => {
    // if we have a thread, filter to only other threads
    if (thread) {
      return edge.node.id != thread.id;
    }

    // else no thread, don't filter at all
    return true;
  });

  console.log("filtered threads", filteredThreads);

  if (thread) {
    filteredThreads.unshift({ node: thread });
  }

  return (
    <InfiniteScroller
      hideScrollBar
      // onScrollDown={() => maybeLoadMore(50)}
      // onScrollUp={() => console.log("up")}
    >
      {filteredThreads?.map(
        (edge: any) =>
          edge?.node && <ThreadButton key={edge.node?.id} thread={edge.node} />,
      )}
    </InfiniteScroller>
  );
}
