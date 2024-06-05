"use client";

import { useState } from "react";
import {
  Avatar,
  Button,
  Card,
  CardBody,
  CardHeader,
  Tooltip,
} from "@nextui-org/react";
import { Cursor, useTypewriter } from "react-simple-typewriter";

import { Message, MessageMessageType } from "@/__generated__/graphql";

type props = {
  message: Message;
  animate?: boolean;
};

export default function Message({ message, animate }: props) {
  const sentAtDisplay = getSentAtDisplay(message.createdAt);
  const [hovered, setHovered] = useState(false);
  const [text] = useTypewriter({
    words: [message.content],
    typeSpeed: 5,
  });
  // const [commitMutation, isMutationInFlight] = useMutation(
  //   MessageBaseCreateBookmarkMutation,
  // );

  return (
    <Card
      className="mb-6 bg-black hover:bg-default-100 mr-2"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <CardHeader className="justify-between">
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
                radius="sm"
                size="md"
                src="/logo.png"
                imgProps={{ className: "h-8 w-8" }}
                // icon={<img src={"/logo.png"} alt="stratifi" className="h-6 w-6" />}
              />
              {sentAtDisplay}
            </>
          )}
        </div>
        {/* context actions */}
        {hovered ? (
          <div className="flex gap-1">
            {message.messageType == MessageMessageType.Ai && (
              <Tooltip
                showArrow
                content="Generate new response"
                placement="bottom"
              >
                <Button isIconOnly size="sm">
                  <i className="pi pi-sync" />
                </Button>
              </Tooltip>
            )}
            <Tooltip showArrow content="Bookmark" placement="bottom">
              <Button
                isIconOnly
                size="sm"
                onPress={() => {
                  // commitMutation({
                  //   variables: {
                  //     userId: "us01HZ5TYGYW36GGJ7Z9VMS5Y5TE",
                  //     messageId: isResponse ? null : id,
                  //     responseId: isResponse ? id : null,
                  //     connections: [connectionID],
                  //   },
                  //   onError: () => {
                  //     toast.error("Error bookmarking message");
                  //   },
                  //   onCompleted: () => {
                  //     toast.success("Bookmarked message");
                  //   },
                  // });
                }}
              >
                <i className="pi pi-bookmark" />
              </Button>
            </Tooltip>
            <Tooltip
              showArrow
              content="Copy message to clipboard"
              placement="bottom"
            >
              <Button
                isIconOnly
                size="sm"
                onClick={() => navigator.clipboard.writeText(message.content)}
              >
                <i className="pi pi-clipboard" />
              </Button>
            </Tooltip>
            <Tooltip showArrow content="Create child thread" placement="bottom">
              <Button isIconOnly size="sm">
                <i className="pi pi-arrow-right" />
              </Button>
            </Tooltip>
          </div>
        ) : null}
      </CardHeader>
      <CardBody>
        {animate ? (
          <>
            {message.content ? (
              <>
                <span>{message.content}</span>
              </>
            ) : (
              <Cursor cursorColor="red" />
            )}
          </>
        ) : (
          message.content
        )}
      </CardBody>
    </Card>
  );
}

function getSentAtDisplay(createdAt: string): string {
  let sentAtDisplay = "";

  if (createdAt) {
    const sentAtDate = new Date(createdAt);
    const sentToday =
      new Date(createdAt).setHours(0, 0, 0, 0) ==
      new Date().setHours(0, 0, 0, 0);

    sentAtDisplay = sentToday
      ? sentAtDate.toLocaleTimeString()
      : sentAtDate.toLocaleString();
  }

  return sentAtDisplay;
}
