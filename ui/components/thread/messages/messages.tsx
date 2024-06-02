import InfiniteScroller from "@/components/common/scroll/infinite-scroller";
import Message from "@/components/thread/messages/message";
import PromptInput from "@/app/thread/[id]/prompt-input";

type props = {
  thread: any;
};

export default function Messages({ thread }: props) {
  const messages = thread?.messages;

  return (
    <>
      {messages?.totalCount > 0 ? (
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
      ) : (
        <span className="w-full text-center mb-2">
          <h1>What do you want to explore?</h1>
        </span>
      )}
      {/* prompt input */}
      <div className="mt-6">
        <PromptInput showModelSelect={!messages} />
      </div>
    </>
  );
}
