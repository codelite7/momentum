import { graphql, usePaginationFragment } from "react-relay";

import InfiniteScroller from "@/components/common/scroll/infinite-scroller";
import Message from "@/components/thread/messages/message";
import PromptInput from "@/app/thread/[id]/prompt-input";
import { messagesFragment$key } from "@/__generated__/messagesFragment.graphql";

type props = {
  thread: messagesFragment$key;
};

const MessagesFragment = graphql`
  fragment messagesFragment on Thread
  @argumentDefinitions(
    first: { type: "Int", defaultValue: 3 }
    after: { type: "Cursor" }
  )
  @refetchable(queryName: "MessagesPaginationQuery") {
    messages(first: $first, after: $after) @connection(key: "Thread_messages") {
      totalCount
      edges {
        node {
          id
          ...messageFragment
        }
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
`;

export default function Messages({ thread }: props) {
  const { data, loadNext, hasNext } = usePaginationFragment(
    MessagesFragment,
    thread,
  );
  const maybeLoadMore = (num: number) => {
    if (hasNext) {
      loadNext(num);
    }
  };
  const messages = data.messages;

  return (
    <>
      {messages.totalCount > 0 ? (
        <InfiniteScroller onScrollDown={() => maybeLoadMore(1)}>
          {messages.edges?.map((edge) => {
            return (
              edge?.node && <Message key={edge.node.id} message={edge.node} />
            );
          })}
        </InfiniteScroller>
      ) : (
        <span className="w-full text-center mb-2">
          <h1>What do you want to explore?</h1>
        </span>
      )}
      {/* prompt input */}
      <div className="mt-6">
        <PromptInput showModelSelect={messages.totalCount == 0} />
      </div>
    </>
  );
}
