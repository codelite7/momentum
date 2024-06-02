"use client";

import { Button } from "@nextui-org/react";

interface props {
  thread: any;
}

export default function DeleteThreadButton({ thread }: props) {
  return (
    <Button
      fullWidth
      className="bg-transparent hover:bg-default-200"
      radius="none"
      startContent={<i className="pi pi-trash" />}
      onPress={() => {
        console.log("deleting thread", thread.id);
      }}
    >
      Delete
    </Button>
  );
}
