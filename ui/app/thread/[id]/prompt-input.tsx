"use client";

import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Textarea,
} from "@nextui-org/react";
import { useMutation } from "@apollo/client";
import { useState } from "react";
import { toast } from "sonner";
import { useRouter, useSearchParams } from "next/navigation";
import { useAtomValue } from "jotai";

import ModelSelect from "@/components/common/model-select";
import { sidebarThreadsQuery } from "@/components/left-sidebar/left-sidebar";
import {
  createMessageMutation,
  createThreadMutation,
  threadByIdQuery,
} from "@/graphql-queries/queries";
import { lastThreadMessageTypeAtom, modelSelectAtom } from "@/state/atoms";
import { MessageMessageType } from "@/__generated__/graphql";

function Divider() {
  return null;
}

interface props {
  threadId: string;
}

export default function PromptInput({ threadId }: props) {
  const searchParams = useSearchParams();
  const parentId = searchParams.get("parentId");
  const isNew = threadId == "";
  const lastMessageType = useAtomValue(lastThreadMessageTypeAtom);
  const router = useRouter();
  const [value, setValue] = useState("");
  const threadProvider = useAtomValue(modelSelectAtom);
  const [createThread, { data, loading, error }] = useMutation(
    createThreadMutation,
    {
      variables: {
        input: {
          name: `${value.substring(0, 15)}...`,
          parentID: parentId,
          provider: threadProvider,
        },
        messageInput: {
          threadID: "",
          content: value,
        },
      },
      onError: () => {
        toast.error("Error creating thread");
      },
      onCompleted: (data) => {
        toast.success("Created thread");
        router.push(`/thread/${data.createThread.id}`);
      },
      refetchQueries: [sidebarThreadsQuery],
    },
  );

  const [createMessage, createMessageResult] = useMutation(
    createMessageMutation,
    {
      variables: {
        input: {
          content: value,
          threadID: threadId,
        },
      },
      onError: (e) => {
        console.error(e);
        toast.error("Error sending message");
      },
      onCompleted: (data) => {
        setValue("");
      },
      refetchQueries: [threadByIdQuery],
    },
  );

  return (
    <Card className="w-full min-h-36 border-1 border-default">
      {!threadId && (
        <CardHeader className="bg-background border-b-1 border-default">
          <ModelSelect />
        </CardHeader>
      )}
      <Divider />
      <CardBody>
        <Textarea
          classNames={{
            inputWrapper:
              "!bg-transparent hover:!bg-transparent focus:!bg-transparent",
          }}
          disabled={!isNew && lastMessageType != MessageMessageType.Ai}
          minRows={2}
          placeholder="Prompt goes here"
          value={value}
          // TODO put this in a function if we can? I'm not sure if we can since we need the refs to createThread and createMessage
          onKeyPress={async (e) => {
            if (e.key === "Enter") {
              e.preventDefault();
              if (isNew) {
                createThread();
              } else {
                createMessage();
              }
            }
          }}
          onValueChange={setValue}
        />
      </CardBody>
      <CardFooter>
        <div className="w-full flex justify-end">
          <Button
            color="primary"
            size="sm"
            startContent={<i className="pi pi-arrow-down" />}
            type="submit"
            variant="bordered"
            onPress={() => {
              if (isNew) {
                createThread();
              } else {
                createMessage();
              }
            }}
          >
            Enter
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
}
