"use client";

import { ApolloQueryResult, useQuery } from "@apollo/client";
import { toast } from "sonner";
import { Spinner } from "@nextui-org/react";
import { useAtom, useAtomValue } from "jotai";

import {
  Message as MessageType,
  MessageMessageType,
  MostRecentMessageQuery,
  ThreadQuery,
  ThreadQueryVariables,
} from "@/__generated__/graphql";
import { getMostRecentMessage } from "@/graphql-queries/queries";
import { lastThreadMessageAtom, threadAtom } from "@/state/atoms";

interface WaitForResponseProps {}

interface WaitForResponseProps {
  refetch?:
    | ((
        variables?: Partial<ThreadQueryVariables>,
      ) => Promise<ApolloQueryResult<ThreadQuery>>)
    | undefined;
}

export default function WaitForResponse({ refetch }: WaitForResponseProps) {
  const [thread, setThread] = useAtom(threadAtom);
  const lastMessage = useAtomValue(lastThreadMessageAtom);
  const { stopPolling } = useQuery(getMostRecentMessage, {
    variables: {
      threadId: thread!.id,
      after: lastMessage!.createdAt,
    },
    // this makes onCompleted get called each time the request is made, without this it only gets called once
    notifyOnNetworkStatusChange: true,
    onError: (e) => console.error(e),
    onCompleted: (data) => {
      const message = getMessageFromData(data);

      if (
        message.messageType == MessageMessageType.Ai &&
        message.content.length > 0
      ) {
        console.log("got response from ai", data);
        stopPolling();
        if (refetch) {
          refetch();
        } else {
          toast.error("We've encountered an error");
        }
      }
    },
    pollInterval: 1000,
  });

  return <Spinner />;
}

function getMessageFromData(
  data: MostRecentMessageQuery | undefined,
): MessageType {
  // forcing ai message type while we're waiting for a response
  let message = {
    messageType: MessageMessageType.Ai,
    content: "",
  } as MessageType;
  const edges = data?.messages?.edges;

  if (
    edges &&
    edges.length > 0 &&
    edges[0] &&
    edges[0].node &&
    edges[0].node.messageType == MessageMessageType.Ai
  ) {
    return edges[0].node as MessageType;
  }

  return message as MessageType;
}
