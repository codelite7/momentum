"use client";

import { Avatar, Card, CardBody, CardHeader, Snippet } from "@nextui-org/react";
import { Link } from "@nextui-org/link";
import { useSetAtom } from "jotai/index";

import { Message, MessageMessageType } from "@/__generated__/graphql";
import { searchModalIsOpenAtom } from "@/state/atoms";

interface props {
  query: string;
  message: Message;
}

export default function MessageResult({ query, message }: props) {
  const setSearchModalIsOpen = useSetAtom(searchModalIsOpenAtom);
  const sentAtDate = new Date(message.createdAt);
  const sentToday =
    new Date(message.createdAt).setHours(0, 0, 0, 0) ==
    new Date().setHours(0, 0, 0, 0);
  const sentAtDisplay = sentToday
    ? sentAtDate.toLocaleTimeString()
    : sentAtDate.toLocaleString();

  return (
    <Card className="border-b-1 border-default">
      <CardHeader>
        <>
          {message.thread.parent?.thread && (
            <Link
              href={`/thread/${message.thread.parent.thread.id}`}
              onPress={() => setSearchModalIsOpen(false)}
            >
              {message.thread.parent.thread.name}&nbsp;/&nbsp;{" "}
            </Link>
          )}
          {
            <Link
              href={`/thread/${message.thread.id}`}
              onPress={() => setSearchModalIsOpen(false)}
            >
              {message.thread.name}
            </Link>
          }
        </>
      </CardHeader>
      <CardBody>
        {/* avatar, sender, time sent*/}
        <div className="flex">
          {message.messageType == MessageMessageType.Human ? (
            <>
              <Avatar
                className="rounded-br-none mr-2"
                imgProps={{ className: "h-6 w-6" }}
                radius="sm"
                size="md"
                src="/default-user-avatar.svg"
              />
              Me {sentAtDisplay}
            </>
          ) : (
            <>
              <Avatar
                isBordered
                className="bg-transparent rounded-bl-none mr-3"
                color="primary"
                imgProps={{ className: "h-8 w-8" }}
                radius="sm"
                size="md"
                src="/logo.png"
              />
              {sentAtDisplay}
            </>
          )}
        </div>
        {getMatchWithSnippet(query, message.content, 25, 25)}
      </CardBody>
    </Card>
  );
}

function getMatchWithSnippet(
  query: string,
  content: string,
  leading: number,
  trailing: number,
): JSX.Element {
  let lowerQuery = query.toLowerCase();
  let lowerContent = content.toLowerCase();
  let queryStartsAt = lowerContent.indexOf(lowerQuery);
  let queryEndsAt = queryStartsAt + query.length;
  let prefix = content.substring(queryStartsAt - leading, queryStartsAt);
  let suffix = content.substring(queryEndsAt, queryEndsAt + trailing);

  return (
    <span>
      ...{prefix}
      <Snippet
        hideCopyButton
        color="primary"
        size="sm"
        symbol=""
        variant="solid"
      >
        {` ${query} `}
      </Snippet>
      {suffix}...
    </span>
  );
}
