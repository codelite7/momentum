"use client";

import { graphql, useFragment } from "react-relay";

import Messages from "@/components/thread/messages/Messages";
import { ThreadFragment$key } from "@/__generated__/ThreadFragment.graphql";

const ThreadFragment = graphql`
  fragment ThreadFragment on Thread {
    id
    createdAt
    name
    ...MessagesFragment
  }
`;

type props = {
  thread: ThreadFragment$key;
};

export default function Thread({ thread }: props) {
  const data = useFragment(ThreadFragment, thread);

  return (
    <div className="w-full h-full flex flex-col place-content-center px-52 pb-4">
      <Messages thread={data} />
    </div>
  );
}
