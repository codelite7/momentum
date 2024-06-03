import { Message } from "@/__generated__/graphql";
import MessageBase from "@/components/thread/messages/message-base";
import Respons from "@/components/thread/response";

type props = {
  message: Message;
  threadId: string;
};

export default function Message({ message, threadId }: props) {
  return (
    <>
      <MessageBase
        content={message.content}
        id={message.id}
        sentAt={message.createdAt}
        threadId={threadId}
        user={{ name: "Me" }}
      />
      {message.response && (
        <Respons response={message.response} threadId={threadId} />
      )}
    </>
  );
}
