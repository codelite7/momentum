"use client";

import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Textarea,
} from "@nextui-org/react";

import ModelSelect from "@/components/common/model-select";

function Divider() {
  return null;
}

interface props {
  showModelSelect: boolean;
}

export default function PromptInput({ showModelSelect }: props) {
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
