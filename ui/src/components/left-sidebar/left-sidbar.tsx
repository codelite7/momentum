"use client";
import { Button } from "@nextui-org/button";
import { useState } from "react";
import {
  Modal,
  ModalBody,
  ModalContent,
  ModalHeader,
  useDisclosure,
  Card,
  CardBody,
  Divider,
  CardHeader,
} from "@nextui-org/react";

import ThreadButtons from "@/components/left-sidebar/thread-buttons";

interface LeftSidbearProps {
  user: any;
}

export default function LeftSidbar({ user }: LeftSidbearProps) {
  const [collapsed, setCollapsed] = useState(false);
  const [width, setWidth] = useState("w-52");
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <div
        className={`h-full bg-forward-900 flex flex-col justify-items-center justify-between ${width} transition-all`}
      >
        {collapsed ? (
          <div className="flex flex-col justify-center items-center">
            {/* Logo and collapse control button*/}
            <div className="flex flex-row justify-between items-center pl-2 pt-2 mb-2">
              <div className="flex items-center">
                <img
                  alt="stratifi logo"
                  className="mr-2"
                  height="20"
                  src="/logo.png"
                  width="20"
                />
              </div>
              <Button
                isIconOnly
                className="min-w-6 w-6 h-6 rounded"
                size="sm"
                variant="light"
                onPress={() => {
                  setWidth("w-52");
                  setCollapsed(false);
                }}
              >
                <img alt="collapse" src="/chevron-right.svg" />
              </Button>
            </div>
            <div className="w-full pl-2 pr-2">
              <Button
                isIconOnly
                className="w-full"
                color="primary"
                size="sm"
                variant="bordered"
                onPress={() => {}}
              >
                <img alt="search" className="h-4 w-4" src="/search.svg" />
              </Button>
            </div>
            <div className="w-full pl-2 pr-2 mt-2">
              <Button
                isIconOnly
                className="w-full"
                color="primary"
                size="sm"
                onPress={() => {}}
              >
                <img alt="search" className="h-4 w-4" src="/layers.svg" />
              </Button>
            </div>
          </div>
        ) : (
          <div className="h-full overflow-hidden">
            {/*logo and collapse button*/}
            <div className="w-full flex flex-row justify-between items-center pl-2 pt-2 mb-2">
              <div className="flex items-center">
                <img
                  alt="stratifi logo"
                  className="mr-2"
                  height="20"
                  src="/logo.png"
                  width="20"
                />
                STRATIFI
              </div>
              <Button
                isIconOnly
                className="min-w-6 w-6 h-6 rounded"
                variant="light"
                onPress={() => {
                  setCollapsed(true);
                  setWidth("w-16");
                }}
              >
                <img alt="collapse" src="/chevron-left.svg" />
              </Button>
            </div>
            {/* search box */}
            <div className="pl-2 pr-2">
              <Button
                fullWidth
                className="flex flex-row start mb-2"
                color="primary"
                size="sm"
                variant="bordered"
              >
                <div className="w-full flex start text-white rounded items-center">
                  <img alt="search" className="mr-2" src="/search.svg" />
                  Search
                </div>
              </Button>
            </div>
            {/* new thread button*/}
            <div className="pl-2 pr-2">
              <Button
                fullWidth
                className="mb-4"
                color="primary"
                size="sm"
                startContent={<img alt="new thread" src="/layers.svg" />}
              >
                New thread
              </Button>
            </div>
            {/* thread buttons */}
            <ThreadButtons />
          </div>
        )}
        {/* footer */}
        <div
          className="w-full p-4 bg-black flex place-items-center overflow-hidden cursor-pointer"
          onClick={onOpen}
        >
          {collapsed ? (
            <div className="flex place-items-center">
              <Button
                isIconOnly
                className="min-w-8 w-8 h-8 rounded"
                color="primary"
                size="sm"
              >
                <img
                  alt="avatar"
                  className="h-4 w-4"
                  src="/default-user-avatar.svg"
                />
              </Button>
            </div>
          ) : (
            <div className="w-full overflow-hidden flex items-center">
              <Button
                isIconOnly
                className="min-w-8 w-8 h-8 rounded mr-4"
                color="primary"
                size="sm"
              >
                <img
                  alt="avatar"
                  className="h-4 w-4"
                  src="/default-user-avatar.svg"
                />
              </Button>
              <span className="w-full overflow-hidden text-ellipsis">
                {user?.email}
              </span>
            </div>
          )}
        </div>
      </div>
      <Modal isOpen={isOpen} size="5xl" onOpenChange={onOpenChange}>
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className="flex items-center">
                <i className="pi pi-sliders-v mr-2" />
                Settings
              </ModalHeader>
              <ModalBody>
                <Card>
                  <CardHeader>Account</CardHeader>
                  <Divider />
                  <CardBody>
                    <div className="flex flex-col">
                      Email
                      <br />
                      {user?.email}
                    </div>
                  </CardBody>
                </Card>
                <Card>
                  <CardBody>
                    <div className="flex justify-between items-center">
                      <div>System</div>
                      <div />
                      <div>
                        <Button
                          size="sm"
                          startContent={
                            <i className="pi pi-trash color-danger" />
                          }
                          variant="light"
                        >
                          Delete Account
                        </Button>
                        <Button
                          color="primary"
                          size="sm"
                          startContent={<i className="pi pi-sign-out" />}
                        >
                          Sign out
                        </Button>
                      </div>
                    </div>
                  </CardBody>
                </Card>
                <Card>
                  <CardBody>
                    <div className="text-center">
                      <p>&copy;2024 STRATIFI</p>
                    </div>
                  </CardBody>
                </Card>
              </ModalBody>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}
