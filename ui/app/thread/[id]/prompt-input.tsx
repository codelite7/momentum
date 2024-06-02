"use client";

import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Textarea,
} from "@nextui-org/react";
import { gql, useMutation } from "@apollo/client";
import { useState } from "react";
import { toast } from "sonner";

import ModelSelect from "@/components/common/model-select";

function Divider() {
  return null;
}

interface props {
  showModelSelect: boolean;
}

const createThreadMutation = gql`
  mutation createThread(
    $input: CreateThreadInput!
    $messageInput: CreateMessageInput!
  ) {
    createThread(input: $input, messageInput: $messageInput) {
      id
      lastViewedAt
      messages {
        edges {
          node {
            id
            createdAt
            content
          }
        }
      }
    }
  }
`;

export default function PromptInput({ showModelSelect }: props) {
  const [value, setValue] = useState("");
  const [createThread, { data, loading, error }] = useMutation(
    createThreadMutation,
    {
      onError: () => {
        toast.error("Error creating thread");
      },
      onCompleted: (data) => {
        toast.success("Created thread");
        console.log("create thread mutation response", data);
      },
    },
  );

  return (
    <Card className="w-full">
      {showModelSelect && (
        <CardHeader>
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
          minRows={1}
          placeholder="Prompt goes here"
          value={value}
          onKeyPress={async (e) => {
            if (e.key === "Enter") {
              e.preventDefault();
              createThread({
                variables: {
                  input: {
                    name: "New thread",
                  },
                  messageInput: {
                    threadId: "",
                    content: value,
                  },
                },
              });
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
          >
            Enter
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
}
