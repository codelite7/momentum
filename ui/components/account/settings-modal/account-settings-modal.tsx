"use client";

import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Divider,
  Modal,
  ModalBody,
  ModalContent,
  ModalHeader,
} from "@nextui-org/react";
import { useShallow } from "zustand/react/shallow";

import useAccountSettingsModalStore from "@/stores/account-settings-modal-store";

export default function AccountSettingsModal() {
  const [isOpen, toggle] = useAccountSettingsModalStore(
    useShallow((state: any) => [state.isOpen, state.toggle]),
  );

  return (
    <Modal
      className="border-1 border-default"
      isOpen={isOpen}
      size="5xl"
      onClose={toggle}
    >
      <ModalContent>
        <ModalHeader className="flex items-center bg-background">
          <i className="pi pi-sliders-v mr-2" />
          Settings
        </ModalHeader>
        <ModalBody className="pb-4">
          <Card>
            <CardHeader>Account</CardHeader>
            <Divider />
            <CardBody>
              <div className="flex flex-col">
                Email
                <br />
                test@test.com
                {/*{user?.email}*/}
              </div>
            </CardBody>
          </Card>
          <Card>
            <CardBody>
              <div className="flex justify-between items-center">
                <div>System</div>
                <div />
                <div className="flex gap-1">
                  <Button
                    size="sm"
                    startContent={<i className="pi pi-trash text-danger" />}
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
          <Card className="bg-stone-100">
            <CardBody>
              <div className="text-center">
                <p>
                  &copy; 2024 <span className="tracking-widest">STRATIFI</span>
                </p>
              </div>
            </CardBody>
          </Card>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
}
