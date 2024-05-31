"use client";

import {
  graphql,
  useFragment,
  useLazyLoadQuery,
  useMutation,
} from "react-relay";
import { Button } from "@nextui-org/react";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

import { threadQuery } from "@/__generated__/threadQuery.graphql";
import { threadFragment$key } from "@/__generated__/threadFragment.graphql";
import Messages from "@/components/thread/messages/messages";

const ThreadQuery = graphql`
  query threadQuery($id: ID!) {
    thread(id: $id) {
      ...threadFragment
    }
  }
`;

const ThreadFragment = graphql`
  fragment threadFragment on Thread {
    id
    createdAt
    name
    ...messagesFragment
  }
`;

type props = {
  id: string;
};

const ThreadDeleteThreadButtonMutation = graphql`
  mutation threadDeleteThreadButtonMutation($id: ID!) {
    deleteThread(id: $id) @deleteRecord
  }
`;

export default function Thread({ id }: props) {
  const router = useRouter();
  const data = useLazyLoadQuery<threadQuery>(ThreadQuery, { id });
  const thread = useFragment<threadFragment$key>(ThreadFragment, data.thread);
  const [commitMutation, isMutationInFlight] = useMutation(
    ThreadDeleteThreadButtonMutation,
  );

  return (
    <div className="w-full h-full flex flex-col place-content-center px-52 pb-4">
      {thread && (
        <>
          <div className="flex gap-2 items-center mt-2 mb-2">
            Thread name: {thread.name}
            <Button
              color="primary"
              size="sm"
              onPress={() => {
                commitMutation({
                  variables: {
                    id: thread.id,
                  },
                  onError: (e) => {
                    toast.error("Error deleting thread");
                  },
                  onCompleted: (response) => {
                    toast.success("Deleted thread");
                    router.push("/");
                  },
                });
              }}
            >
              Delete Thread
            </Button>
          </div>
          <Messages thread={thread} threadId={thread.id} />
        </>
      )}
    </div>
  );
}
