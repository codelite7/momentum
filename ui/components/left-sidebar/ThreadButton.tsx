"use client";

import {
  Button,
  Popover,
  PopoverTrigger,
  PopoverContent,
} from "@nextui-org/react";
import { useState } from "react";
import { useRouter } from "next/navigation";

import DeleteThreadButton from "@/components/left-sidebar/delete-thread-button";

export default function ThreadButton({ thread }: any) {
  const [popoverOpen, setPopoverOpen] = useState(false);
  const [hovered, setHovered] = useState(false);
  const router = useRouter();

  return (
    <div
      className="flex items-center mb-2 mr-2 ml-2"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      {hovered || popoverOpen ? (
        <>
          <Button
            className="flex flex-row start w-full"
            size="sm"
            onClick={() => router.push(`/thread/${thread.id}`)}
          >
            <div className="flex start items-center justify-between w-full">
              <div>{thread.name}</div>
              <Popover
                classNames={{
                  content: ["rounded border border-default-200 p-0 m-0"],
                }}
                placement="bottom-start"
                onOpenChange={(isOpen) => {
                  setPopoverOpen(isOpen);
                }}
              >
                <PopoverTrigger>
                  <i className="pi pi-ellipsis-v" />
                </PopoverTrigger>
                <PopoverContent>
                  <div className="flex flex-col">
                    <div className="flex flex-row items-center hover:bg-default-200 cursor-pointer pl-4 pr-4 pt-2 pb-2">
                      <div className="mr-4">
                        <img alt="edit thread name" src="/edit.svg" />
                      </div>
                      Rename
                    </div>
                    <DeleteThreadButton id={thread.id} />
                  </div>
                </PopoverContent>
              </Popover>
            </div>
          </Button>
        </>
      ) : (
        <Button
          className="flex flex-row start w-full bg-transparent"
          size="sm"
          onClick={() => router.push(`/thread/${thread.id}`)}
        >
          <div className="flex start items-center justify-between w-full">
            <div>{thread.name}</div>
          </div>
        </Button>
      )}
    </div>
  );
}
