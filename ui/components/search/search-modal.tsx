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
        body: "px-0 mt-4",
        header: "px-0 mt-4 border-b-1",
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
        <ModalBody>
          <Card>
            <CardBody>
              {query ? (
                <DynamicSearchResults query={query} />
              ) : (
                <div className="flex w-full place-content-center">
                  <span>
                    Type to search{" "}
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
