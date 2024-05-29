"use client";

import { graphql, useFragment, useLazyLoadQuery } from "react-relay";

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

export default function Thread({ id }: props) {
  const data = useLazyLoadQuery<threadQuery>(ThreadQuery, { id });
  const thread = useFragment<threadFragment$key>(ThreadFragment, data.thread);

  return (
    <div className="w-full h-full flex flex-col place-content-center px-52 pb-4">
      <Messages thread={thread} />
    </div>
  );
}
