import { Response } from "@/__generated__/graphql";
import MessageBase from "@/components/thread/messages/message-base";

type props = {
  threadId: string;
  response: Response;
};
export default function Respons({ threadId, response }: props) {
  return (
    <MessageBase
      isResponse
      content={response.content ?? ""}
      id={response.id}
      sentAt={response.createdAt}
      threadId={threadId}
      user={undefined}
    />
  );
}
