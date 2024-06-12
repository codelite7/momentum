"use client";

import { useEffect, useRef } from "react";
import { useAtomValue } from "jotai";
import { ApolloQueryResult } from "@apollo/client";

import {
  Message as MessageType,
  MessageMessageType,
  ThreadQuery,
  ThreadQueryVariables,
} from "@/__generated__/graphql";
import { lastThreadMessageTypeAtom, threadMessagesAtom } from "@/state/atoms";
import Message from "@/components/thread/messages/message";
import WaitForResponse from "@/components/thread/messages/waitForResponse";
import InfiniteScroller from "@/components/common/scroll/infinite-scroller";

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
    // @ts-ignore
    messagesEndRef?.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <InfiniteScroller>
      {messages.map((message: MessageType) => {
        return <Message key={message.id} message={message} refetch={refetch} />;
      })}
      {lastMessageType == MessageMessageType.Human && (
        <WaitForResponse refetch={refetch} />
      )}
      <div ref={messagesEndRef} />
    </InfiniteScroller>
  );
}
