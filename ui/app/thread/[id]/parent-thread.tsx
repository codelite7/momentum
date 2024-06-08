import { useAtom, useAtomValue } from "jotai/index";
import { useQuery } from "@apollo/client";
import { Avatar, Card, CardBody, CardHeader } from "@nextui-org/react";
import * as ScrollArea from "@radix-ui/react-scroll-area";

import { parentThreadAtom, parentThreadMessagesAtom } from "@/state/atoms";
import { threadByMessageIdQuery } from "@/graphql-queries/queries";
import { Thread } from "@/__generated__/graphql";
import { getSentAtDisplay } from "@/common/helpers";

interface ParentThreadProps {
  parentId?: string;
}

export default function ParentThread({ parentId }: ParentThreadProps) {
  const [thread, setThread] = useAtom(parentThreadAtom);
  const messages = useAtomValue(parentThreadMessagesAtom);

  const { data } = useQuery(threadByMessageIdQuery, {
    notifyOnNetworkStatusChange: true,
    variables: { messageId: parentId ?? "" },
    onError: (e) => console.error(e),
    onCompleted: (data) => {
      if (
        data.threads.edges &&
        data.threads.edges.length > 0 &&
        data.threads.edges[0]?.node
      ) {
        setThread(data.threads.edges[0].node as Thread);
      }
    },
  });

  return (
    <>
      {thread && (
        <Card className="w-68 h-full">
          <CardHeader>
            <div className="w-full flex flex-col gap-2">
              <div className="w-full">Parent Summary</div>
              <div className="flex items-center gap-2">
                <Avatar
                  isBordered
                  className="bg-transparent rounded-bl-none mr-3 w-6 h-6"
                  color="primary"
                  imgProps={{ className: "h-5 w-5" }}
                  radius="sm"
                  size="sm"
                  src="/logo.png"
                />
                <span>
                  {thread?.createdAt && getSentAtDisplay(thread.createdAt)}
                </span>
              </div>
            </div>
          </CardHeader>
          <CardBody>
            <ScrollArea.Root>
              <ScrollArea.Viewport>
                <div className="flex flex-col gap-2">
                  {messages.map((message) => (
                    <div key={message.id}>{message.content}</div>
                  ))}
                </div>
              </ScrollArea.Viewport>
            </ScrollArea.Root>
          </CardBody>
        </Card>
      )}
    </>
  );
}
