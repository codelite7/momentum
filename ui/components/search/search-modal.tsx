"use client";

import {
  Card,
  CardBody,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalHeader,
  Spinner,
} from "@nextui-org/react";
import { useDebouncedCallback } from "use-debounce";
import { useState } from "react";
import dynamic from "next/dynamic";
import { useAtom } from "jotai";

import { searchModalIsOpenAtom } from "@/state/atoms";

const DynamicSearchResults = dynamic(
  () => import("@/components/search/search-results"),
  {
    loading: () => {
      return (
        <div className="flex place-content-center">
          <Spinner color="primary" size="lg" />
        </div>
      );
    },
  },
);

export default function SearchModal() {
  const [isOpen, setIsOpen] = useAtom(searchModalIsOpenAtom);
  const [query, setQuery] = useState("");
  const debounced = useDebouncedCallback((query) => {
    setQuery(query);
  }, 1000);

  return (
    <Modal
      classNames={{
        base: "border-1 border-default",
        body: "px-0 mt-4",
        header: "px-0 mt-4 border-b-1 border-default bg-background",
      }}
      isOpen={isOpen}
      size="4xl"
      onClose={() => {
        debounced.cancel();
        setQuery("");
        setIsOpen(false);
      }}
    >
      <ModalContent>
        <ModalHeader>
          <div className="flex w-full items-center p-4">
            <Input
              autoFocus
              color="primary"
              size="sm"
              variant="bordered"
              onValueChange={debounced}
            />
          </div>
        </ModalHeader>
        <ModalBody className="m-0 px-4 py-6">
          <Card>
            <CardBody className="p-0">
              {query ? (
                <DynamicSearchResults query={query} />
              ) : (
                <div className="flex w-full h-full place-content-center bg-stone-100 px-4 py-2">
                  <span>
                    Type to search&nbsp;
                    <span className="tracking-widest">STRATIFI</span>
                  </span>
                </div>
              )}
            </CardBody>
          </Card>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
}
