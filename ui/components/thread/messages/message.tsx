import { graphql, useFragment } from "react-relay";

import { messageFragment$key } from "@/__generated__/messageFragment.graphql";
import MessageBase from "@/components/thread/messages/message-base";
import Response from "@/components/thread/response";

const MessageFragment = graphql`
  fragment messageFragment on Message {
    id
    createdAt
    content
    response {
      ...responseFragment
    }
  }
`;

type props = {
  message: messageFragment$key;
  threadId: string;
};

export default function Message({ message, threadId }: props) {
  const data = useFragment(MessageFragment, message);

  return (
    <>
      <MessageBase
        content={data.content}
        id={data.id}
        sentAt={data.createdAt}
        threadId={threadId}
        user={{ name: "Me" }}
      />
      {data.response && <Response response={data.response} />}
    </>
  );
}
