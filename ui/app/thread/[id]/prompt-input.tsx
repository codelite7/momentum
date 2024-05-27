"use client"

import {Button, Card, CardBody, CardFooter, CardHeader, Textarea} from "@nextui-org/react";
import ModelSelect from "@/components/common/model-select";

function Divider() {
    return null;
}

interface props {
    showModelSelect: boolean
}

export default function PromptInput({showModelSelect}: props) {
    return (
        <Card className="w-full">
            {
                showModelSelect && (
                    <CardHeader>
                        <ModelSelect/>
                    </CardHeader>

                )
            }
            <Divider/>
            <CardBody>
                <Textarea
                    minRows={1}
                    classNames={{
                        inputWrapper: "!bg-transparent hover:!bg-transparent focus:!bg-transparent",
                    }}
                    placeholder="Prompt goes here"
                />
            </CardBody>
            <CardFooter>
                <div className="w-full flex justify-end">
                    <Button
                        variant="bordered"
                        startContent={<i className="pi pi-arrow-down"/>}
                        color="primary"
                        type="submit"
                    >
                        Enter
                    </Button>
                </div>
            </CardFooter>
        </Card>
    )
}