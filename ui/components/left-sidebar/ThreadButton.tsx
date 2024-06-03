"use client";

import {
  Button,
  Popover,
  PopoverTrigger,
  PopoverContent,
  Input,
} from "@nextui-org/react";
import { useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useMutation } from "@apollo/client";
import { toast } from "sonner";

import { gql } from "@/__generated__";
import { sidebarThreadsQuery } from "@/components/left-sidebar/left-sidebar";

type props = {
  thread: any;
};

const deleteThreadMutation = gql(/* GraphQL */ `
  mutation deleteThread($id: ID!) {
    deleteThread(id: $id)
  }
`);

const renameThreadMutation = gql(/* GraphQL */ `
  mutation updateThread($id: ID!, $name: String!) {
    updateThread(id: $id, input: { name: $name }) {
      id
      name
    }
  }
`);

export default function ThreadButton({ thread }: props) {
  const params = useParams<{ id: string }>();
  const router = useRouter();
  const [deleteThread] = useMutation(deleteThreadMutation);
  const [renameThread] = useMutation(renameThreadMutation);
  const [popoverOpen, setPopoverOpen] = useState(false);
  const [hovered, setHovered] = useState(false);
  const [renaming, setRenaming] = useState(false);
  const [renameValue, setRenameValue] = useState(thread.name);

  return (
    <div
      className="flex items-center mb-2 mr-2 ml-2"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      {(hovered && !renaming) || popoverOpen ? (
        <div
          className="flex flex-row start w-full hover:bg-default-200 cursor-pointer pl-2 text-sm rounded-md"
          onClick={() => router.push(`/thread/${thread.id}`)}
        >
          <div className="flex start items-center justify-between w-full">
            <div>{thread.name}</div>
            <Popover
              classNames={{
                content: ["rounded border border-default-200"],
              }}
              isOpen={popoverOpen}
              placement="bottom-start"
              onOpenChange={(open) => setPopoverOpen(open)}
            >
              <PopoverTrigger>
                <Button
                  className="flex flex-row start min-w-9 px-0 hover:bg-default-100 rounded-l-none"
                  size="sm"
                  onClick={() => router.push(`/thread/${thread.id}`)}
                >
                  <div className="flex start items-center justify-between h-full">
                    <i className="pi pi-ellipsis-v" />
                  </div>
                </Button>
                {/*<div className="w-full h-full bg-red-500 place-items-center">*/}
                {/*  <i className="pi pi-ellipsis-v" />*/}
                {/*</div>*/}
              </PopoverTrigger>
              <PopoverContent className="p-0">
                <div className="flex flex-col w-full h-full">
                  <Button
                    fullWidth
                    className="bg-transparent hover:bg-default-200"
                    radius="none"
                    onPress={() => {
                      setRenaming(true);
                      setPopoverOpen(false);
                      setHovered(false);
                    }}
                  >
                    <div className="flex w-full items-center gap-4">
                      <i className="pi pi-pencil" />
                      Rename
                    </div>
                  </Button>
                  <Button
                    fullWidth
                    className="bg-transparent hover:bg-default-200"
                    radius="none"
                    onPress={() => {
                      deleteThread({
                        variables: { id: thread.id },
                        onError: (e) => {
                          console.error(e);
                          toast.error("Error deleting thread");
                        },
                        onCompleted: () => {
                          setHovered(false);
                          setPopoverOpen(false);
                          toast.success("Deleted thread");
                          if (thread.id == params.id) {
                            router.push("/");
                          }
                        },
                        refetchQueries: [sidebarThreadsQuery],
                      });
                    }}
                  >
                    <div className="flex w-full items-center gap-4">
                      <i className="pi pi-trash text-danger" />
                      Delete
                    </div>
                  </Button>
                </div>
              </PopoverContent>
            </Popover>
          </div>
        </div>
      ) : renaming ? (
        <div className="flex flex-row start w-full pt-1">
          <Input
            autoFocus
            className="focus:border-0"
            size="sm"
            value={renameValue}
            onChange={(e) => setRenameValue(e.target.value)}
            onFocusChange={(focused) => {
              if (!focused) {
                setRenameValue(thread.name);
                setRenaming(false);
              }
            }}
            onKeyPress={(e) => {
              if (e.key === "Enter") {
                renameThread({
                  variables: { id: thread.id, name: renameValue },
                  onError: (e) => {
                    console.error(e);
                    toast.error("Error renaming thread");
                  },
                  onCompleted: (data) => {
                    toast.success("Renamed thread");
                    setRenaming(false);
                  },
                });
              }
            }}
          />
        </div>
      ) : (
        <div
          className="flex flex-row start w-full h-8 pl-2 text-sm"
          onClick={() => router.push(`/thread/${thread.id}`)}
        >
          <div className="flex start items-center justify-between w-full h-full">
            <div>{thread.name}</div>
          </div>
        </div>
      )}
    </div>
  );
}
