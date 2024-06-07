"use client";

import { useQuery } from "@apollo/client";
import { useAtom } from "jotai";

import { Thread } from "@/__generated__/graphql";
import { threadByIdQuery } from "@/graphql-queries/queries";
import Messages from "@/components/thread/messages/messages";
import { threadAtom } from "@/state/atoms";
import PromptInput from "@/app/thread/[id]/prompt-input";

type props = {
  threadId: string;
};

export default function Thread({ threadId }: props) {
  const [thread, setThread] = useAtom(threadAtom);

  const { data, refetch } = useQuery(threadByIdQuery, {
    notifyOnNetworkStatusChange: true,
    variables: { id: threadId },
    onError: (e) => console.error(e),
    onCompleted: (e) => {
      // not in onComplete hook so that refetch triggers re-render
      setThread(data?.thread as Thread);
    },
  });

  return (
    <>
      {thread && thread.messages && (
        <>
          <div className="h-full p-2 flex flex-col">
            <Messages refetch={refetch} />
            <PromptInput threadId={threadId} />
          </div>
        </>
      )}
    </>
  );
}
