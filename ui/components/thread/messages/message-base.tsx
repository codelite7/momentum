"use client";

import {
  Avatar,
  Button,
  Card,
  CardBody,
  CardHeader,
  Tooltip,
} from "@nextui-org/react";
import { useState } from "react";
import { ConnectionHandler, graphql, useMutation } from "react-relay";
import { toast } from "sonner";

interface props {
  content: string;
  sentAt: string;
  user: any;
  isResponse?: boolean;
  id: string;
  threadId: string;
}

const MessageBaseCreateBookmarkMutation = graphql`
  mutation messageBaseCreateBookmarkMutation(
    $userId: ID!
    $messageId: ID
    $responseID: ID
    $connections: [ID!]!
  ) {
    createBookmark(
      input: { userID: $userId, messageID: $messageId, responseID: $responseID }
    ) {
      edges @prependEdge(connections: $connections) {
        node {
          id
          ...bookmarkFragment
        }
      }
    }
  }
`;

export default function MessageBase({
  content,
  sentAt,
  user,
  isResponse,
  id,
  threadId,
}: props) {
  const sentAtDate = new Date(sentAt);
  const sentToday =
    new Date(sentAt).setHours(0, 0, 0, 0) == new Date().setHours(0, 0, 0, 0);
  const sentAtDisplay = sentToday
    ? sentAtDate.toLocaleTimeString()
    : sentAtDate.toLocaleString();
  const [hovered, setHovered] = useState(false);
  const [commitMutation, isMutationInFlight] = useMutation(
    MessageBaseCreateBookmarkMutation,
  );

  return (
    <Card
      className="mb-6 bg-black hover:bg-default-100 mr-2"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <CardHeader className="justify-between">
        {/* avatar, sender, time sent*/}
        <div className="flex">
          {user ? (
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
            {!user && (
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
                  const connectionID = ConnectionHandler.getConnectionID(
                    "root",
                    "Thread_bookmarks",
                    [
                      {
                        where: {
                          hasMessageWith: { hasThreadWith: { id: threadId } },
                        },
                      },
                    ],
                  );

                  console.log("connection id: ", connectionID);

                  commitMutation({
                    variables: {
                      userId: "us01HZ5TYGYW36GGJ7Z9VMS5Y5TE",
                      messageId: isResponse ? null : id,
                      responseId: isResponse ? id : null,
                      connections: [connectionID],
                    },
                    onError: () => {
                      toast.error("Error bookmarking message");
                    },
                    onCompleted: () => {
                      toast.success("Bookmarked message");
                    },
                  });
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
                onClick={() => navigator.clipboard.writeText(content)}
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
      <CardBody>{content}</CardBody>
    </Card>
  );
}
