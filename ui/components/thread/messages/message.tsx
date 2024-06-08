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
import { useMutation } from "@apollo/client";
import { toast } from "sonner";
import { Link } from "@nextui-org/link";

import { Bookmark, Message, MessageMessageType } from "@/__generated__/graphql";
import {
  createBookmarkMutation,
  deleteBookmarkMutation,
  threadBookmarksQuery,
} from "@/graphql-queries/queries";
import { getSentAtDisplay } from "@/common/helpers";

type props = {
  message: Message;
  animate?: boolean;
};

export default function Message({ message, animate }: props) {
  console.log("message bookmarks", message.bookmarks);
  const sentAtDisplay = getSentAtDisplay(message.createdAt);
  const [hovered, setHovered] = useState(false);
  const [text] = useTypewriter({
    words: [message.content],
    typeSpeed: 5,
  });
  let bookmarkss: Bookmark[] = [];

  message.bookmarks?.edges?.forEach((edge) => {
    if (edge?.node) {
      bookmarkss.push(edge.node as Bookmark);
    }
  });

  const [createBookmark] = useMutation(createBookmarkMutation, {
    variables: {
      messageId: message.id,
    },
    onError: (e) => {
      toast.error("Error bookmarking message");
      console.error(e);
    },
    onCompleted: (data) => {
      toast.success("Bookmarked message");
      bookmarkss.push(data.createBookmark as Bookmark);
    },
    refetchQueries: [threadBookmarksQuery],
  });
  const [deleteBookmark] = useMutation(deleteBookmarkMutation, {
    variables: {
      id: bookmarkss.length > 0 ? bookmarkss[0].id : "",
    },
    onError: (e) => {
      toast.error("Error deleting bookmark");
      console.error(e);
    },
    onCompleted: (data) => {
      toast.success("Deleted bookmark");
      bookmarkss = [];
    },
    refetchQueries: [threadBookmarksQuery],
  });

  return (
    <Card
      className="mb-6 bg-black hover:bg-default-100"
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
                  if (bookmarkss.length > 0) {
                    deleteBookmark();
                  } else {
                    createBookmark();
                  }
                }}
              >
                {bookmarkss.length > 0 ? (
                  <i className="pi pi-bookmark-fill text-primary" />
                ) : (
                  <i className="pi pi-bookmark" />
                )}
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
                onPress={() => navigator.clipboard.writeText(message.content)}
              >
                <i className="pi pi-clipboard" />
              </Button>
            </Tooltip>
            {message.child ? (
              <Tooltip
                showArrow
                content={message.child.name}
                placement="bottom"
              >
                <Button
                  isIconOnly
                  as={Link}
                  color="primary"
                  href={`/thread/${message.child.id}`}
                  size="sm"
                >
                  <i className="pi pi-arrow-right" />
                </Button>
              </Tooltip>
            ) : (
              <Tooltip
                showArrow
                content="Create child thread"
                placement="bottom"
              >
                <Button
                  isIconOnly
                  as={Link}
                  href={`/thread/new?parentId=${message.id}`}
                  size="sm"
                >
                  <i className="pi pi-arrow-right" />
                </Button>
              </Tooltip>
            )}
          </div>
        ) : (
          <div className="flex gap-1">
            {message.child && (
              <Tooltip
                showArrow
                content={message.child.name}
                placement="bottom"
              >
                <Button
                  isIconOnly
                  as={Link}
                  color="primary"
                  href={`/thread/${message.child.id}`}
                  size="sm"
                >
                  <i className="pi pi-arrow-right" />
                </Button>
              </Tooltip>
            )}
          </div>
        )}
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
