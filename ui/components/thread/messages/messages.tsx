import InfiniteScroller from "@/components/common/scroll/infinite-scroller";
import Message from "@/components/thread/messages/message";
import { ThreadByIdThreadFragmentFragment } from "@/__generated__/graphql";

type props = {
  thread: ThreadByIdThreadFragmentFragment;
};

export default function Messages({ thread }: props) {
  const messages = thread.messages;

  return (
    <>
      <InfiniteScroller
        hideScrollBar
        onScrollDown={() => {
          if (thread.messages.pageInfo.hasNextPage) {
          }
        }}
      >
        {messages.edges?.map((edge: any) => {
          return (
            edge?.node && (
              <Message
                key={edge.node.id}
                message={edge.node}
                threadId={thread.id}
              />
            )
          );
        })}
      </InfiniteScroller>
    </>
  );
}
