"use client";

import { useEffect, useRef } from "react";
import { useAtomValue } from "jotai";
import { ApolloQueryResult } from "@apollo/client";

import InfiniteScroller from "@/components/common/scroll/infinite-scroller";
import Message from "@/components/thread/messages/message";
import {
  Message as MessageType,
  MessageMessageType,
  ThreadQuery,
  ThreadQueryVariables,
} from "@/__generated__/graphql";
import { lastThreadMessageTypeAtom, threadMessagesAtom } from "@/state/atoms";
import WaitForResponse from "@/components/thread/messages/waitForResponse";

interface MessagesProps {
  refetch?: (
    variables?: Partial<ThreadQueryVariables>,
  ) => Promise<ApolloQueryResult<ThreadQuery>>;
}

export default function Messages({ refetch }: MessagesProps) {
  const messages = useAtomValue(threadMessagesAtom);
  const lastMessageType = useAtomValue(lastThreadMessageTypeAtom);
  const messagesEndRef = useRef(null);

  // scroll to bottom after render
  useEffect(() => {
    console.log("messages", messages);
    // @ts-ignore
    messagesEndRef?.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <>
      <InfiniteScroller hideScrollBar>
        {messages.map((message: MessageType) => {
          return <Message key={message.id} message={message} />;
        })}
        {lastMessageType == MessageMessageType.Human && (
          <WaitForResponse refetch={refetch} />
        )}
        <div ref={messagesEndRef} />
      </InfiniteScroller>
    </>
  );
}
