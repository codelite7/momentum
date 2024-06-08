"use client";

import { useQuery } from "@apollo/client";
import { useAtom } from "jotai";

import { Thread } from "@/__generated__/graphql";
import { threadByIdQuery } from "@/graphql-queries/queries";
import Messages from "@/components/thread/messages/messages";
import { threadAtom } from "@/state/atoms";
import PromptInput from "@/app/thread/[id]/prompt-input";
import ThreadHeader from "@/components/thread/thread-header";
import ParentThread from "@/app/thread/[id]/parent-thread";

type props = {
  threadId: string;
  parentId?: string;
};

export default function Thread({ threadId, parentId }: props) {
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
          <ThreadHeader />
          <div className="w-full h-full flex gap-2 px-2 overflow-hidden pb-2">
            {thread.parent && <ParentThread parentId={thread.parent.id} />}

            {/* prompt input */}
            <div className="flex flex-col w-full h-full place-content-center p-2">
              <Messages refetch={refetch} />
              <PromptInput threadId={threadId} />
            </div>
          </div>
        </>
      )}
    </>
  );
}
