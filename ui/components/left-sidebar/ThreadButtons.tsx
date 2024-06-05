import ThreadButton from "./ThreadButton";

import InfiniteScroller from "@/components/common/scroll/infinite-scroller";

type props = {
  threads: any;
};
export default function ThreadButtons({ threads }: props) {
  return (
    <InfiniteScroller
      hideScrollBar
      // onScrollDown={() => maybeLoadMore(50)}
      // onScrollUp={() => console.log("up")}
    >
      {threads.edges?.map(
        (edge: any) =>
          edge?.node && <ThreadButton key={edge.node?.id} thread={edge.node} />,
      )}
    </InfiniteScroller>
  );
}
