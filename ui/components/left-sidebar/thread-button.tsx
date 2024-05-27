import {Button, Popover, PopoverTrigger, PopoverContent, Card, CardBody} from "@nextui-org/react";
import {useState} from "react";
import {useRouter} from "next/navigation";

interface ThreadButtonProps {
    edge: any
}

export default function ThreadButton({edge}: ThreadButtonProps) {
    const [popoverOpen, setPopoverOpen] = useState(false);
    const [hovered, setHovered] = useState(false);
    const router = useRouter()
    return (
        <div className="flex items-center mb-2 mr-2 ml-2" onMouseEnter={() => setHovered(true)} onMouseLeave={() => setHovered(false)}>
            {
                hovered || popoverOpen ? (
                    <>
                        <Button
                            size="sm"
                            className="flex flex-row start w-full"
                            onClick={() => router.push(`/thread/${edge?.node?.id}`)}
                        >
                            <div className="flex start items-center justify-between w-full">
                                <div>
                                    {edge?.node?.name}
                                </div>
                                <Popover
                                    placement="bottom-start"
                                    classNames={{
                                        content: [
                                            "rounded border border-default-200 p-0 m-0"
                                        ]
                                    }}
                                    onOpenChange={(isOpen) => {
                                        setPopoverOpen(isOpen)
                                    }}
                                >
                                    <PopoverTrigger>
                                        <i className="pi pi-ellipsis-v"></i>
                                    </PopoverTrigger>
                                    <PopoverContent>
                                    <div className="flex flex-col">
                                            <div
                                                className="flex flex-row items-center hover:bg-default-200 cursor-pointer pl-4 pr-4 pt-2 pb-2">
                                                <div className="mr-4">
                                                    <img src="/edit.svg" alt="edit thread name"/>
                                                </div>
                                                Rename
                                            </div>
                                            <div
                                                className="flex flex-row items-center hover:bg-default-200 cursor-pointer pl-4 pr-4 pt-2 pb-2"
                                                onClick={() => {
                                                    console.log(`deleting thread: ${edge?.node?.id}`)
                                                }}>
                                                <div className="mr-4">
                                                    <img src="/trash.svg" alt="delete thread"/>
                                                </div>
                                                Delete
                                            </div>
                                        </div>
                                    </PopoverContent>
                                </Popover>
                            </div>
                        </Button>
                    </>
                ) : (
                    <Button
                        size="sm"
                        className="flex flex-row start w-full bg-transparent"
                        onClick={() => router.push(`/thread/${edge?.node?.id}`)}
                    >
                        <div className="flex start items-center justify-between w-full">
                            <div>
                                {edge?.node?.name}
                            </div>
                        </div>
                    </Button>
                )
            }

        </div>
    )
}