"use client";

import { toast } from "sonner";
import { graphql, useMutation } from "react-relay";

import { ThreadButtonFragment$data } from "@/__generated__/ThreadButtonFragment.graphql";

interface props {
  thread: ThreadButtonFragment$data;
}

const DeleteThreadButtonMutation = graphql`
  mutation deleteThreadButtonMutation($id: ID!) {
    deleteThread(id: $id) @deleteRecord
  }
`;

export default function DeleteThreadButton({ thread }: props) {
  const [commitMutation, isMutationInFlight] = useMutation(
    DeleteThreadButtonMutation,
  );

  return (
    <div
      className="flex flex-row items-center hover:bg-default-200 cursor-pointer pl-4 pr-4 pt-2 pb-2"
      onClick={() => {
        commitMutation({
          variables: {
            id: thread.id,
          },
          onError: (e) => {
            toast.error("Error deleting thread");
          },
          onCompleted: (response) => {
            toast.success("Deleted thread");
          },
        });
      }}
    >
      <div className="mr-4">
        <img alt="delete thread" src="/trash.svg" />
      </div>
      Delete
    </div>
  );
}
