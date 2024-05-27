import {Button, Popover, PopoverTrigger, PopoverContent, Card, CardBody} from "@nextui-org/react";
import {useState} from "react";
import {useRouter} from "next/navigation";

interface ThreadButtonProps {
    edge: any
}

export default function ThreadButton({edge}: ThreadButtonProps) {
    const [showButton, setShowButton] = useState(false);
    const [popoverOpen, setPopoverOpen] = useState(false);
    const [hovered, setHovered] = useState(false);
    const router = useRouter()
    return (
        <div className="flex items-center mb-2 mr-2 ml-2" onMouseEnter={() => setHovered(true)} onMouseLeave={() => setHovered(false)}>
            {
                hovered || popoverOpen ? (
                    <>
                        {/*<Button*/}
                        {/*    size="sm"*/}
                        {/*    className="flex flex-row start rounded-r-none w-full"*/}
                        {/*    onClick={() => router.push(`/thread/${edge?.node?.id}`)}*/}
                        {/*>*/}
                        {/*    <div className="w-full flex flex-row start">*/}
                        {/*        {edge?.node?.name}*/}
                        {/*    </div>*/}
                        {/*</Button>*/}
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
                            <Button
                                size="sm"
                                className="flex flex-row start w-full"
                                onClick={() => setPopoverOpen(true)}
                            >
                                <div className="flex start items-center justify-between w-full">
                                    <div>
                                        {edge?.node?.name}
                                    </div>
                                    <i className="pi pi-ellipsis-v"></i>
                                </div>
                            </Button>
                            {/*<Button*/}
                            {/*    isIconOnly*/}
                            {/*    size="sm"*/}
                            {/*    className="rounded-l-none"*/}
                            {/*    onClick={() => setPopoverOpen(true)}*/}
                            {/*>*/}
                            {/*    <i className="pi pi-ellipsis-v"></i>*/}
                            {/*</Button>*/}
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
                    </>
                ) : (
                    <Button
                        size="sm"
                        fullWidth
                        className="flex flex-row start rounded-r-none w-full bg-transparent"
                        onClick={() => router.push(`/thread/${edge?.node?.id}`)}
                    >
                        <div className="w-full flex flex-row start">
                            {edge?.node?.name}
                        </div>
                    </Button>
                )
            }

        </div>
    )
}