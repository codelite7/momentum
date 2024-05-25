import {Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Button, useDisclosure} from "@nextui-org/react";

export function userSettingsModal() {
  const {isOpen, onOpen, onOpenChange} = useDisclosure();
  return (
    <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
      <ModalContent>
        User Settings
      </ModalContent>
    </Modal>
  )
}
