"use client";

import { graphql, useLazyLoadQuery } from "react-relay";

import { threadQuery } from "@/__generated__/threadQuery.graphql";
import PromptInput from "@/app/thread/[id]/prompt-input";
import InfiniteScroller from "@/components/common/scroll/infinite-scroller";
import Message from "@/components/thread/message";

const ThreadQuery = graphql`
  query threadQuery($id: ID!) {
    thread(id: $id) {
      id
      createdAt
      name
      messages {
        edges {
          node {
            id
            ...messageFragment
          }
        }
      }
    }
  }
`;

type props = {
  id: string;
};

export default function Thread({ id }: props) {
  const data = useLazyLoadQuery<threadQuery>(ThreadQuery, { id });
  const threadHasMessages = data.thread.messages.edges ?? [].length > 0;

  return (
    <div className="w-full h-full flex flex-col place-content-center px-52 pb-4">
      {threadHasMessages ? (
        <InfiniteScroller>
          {data.thread.messages.edges?.map((edge) => {
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
        <PromptInput showModelSelect={!threadHasMessages} />
      </div>
    </div>
  );
  // return <Messages thread={data.thread} />;
}
