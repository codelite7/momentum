"use client"

import {useState, useEffect} from "react";
import {
    Button, Card, CardBody, CardFooter, CardHeader,
    Divider, Modal, ModalBody, ModalContent, ModalHeader, Tooltip, useDisclosure
} from '@nextui-org/react';
import ThreadButtons from "./thread-buttons";

export default function LeftSidebar() {
    const [collapsed, setCollapsed] = useState(false);
    const [width, setWidth] = useState("w-52");
    const {isOpen, onOpen, onOpenChange} = useDisclosure();
    useEffect(() => {
        let width = "w-52"
        if (collapsed) {
            width = "w-16"
        }
        console.log(`setting width: ${width}`)
        setWidth(width)
    }, [collapsed]);
    return (
        <>
            <Card className={`h-full ${width} rounded-none !transition-all !duration-1000`} disableAnimation>
                <CardBody className="h-full overflow-hidden p-0">
                    {collapsed ? (
                        <>
                            <div className="h-full p-2">
                                <Tooltip showArrow content="Expand" placement="right">
                                    <Button
                                        variant="light"
                                        isIconOnly
                                        className="w-full"
                                        color="primary"
                                        size="sm"
                                        onPress={() => {
                                            setCollapsed(!collapsed)
                                        }}
                                    >
                                        <img alt="stratifi" className="h-4 w-4" src="/logo.png"/>
                                    </Button>
                                </Tooltip>

                                <Tooltip showArrow content="Search" placement="right">
                                    <Button
                                        isIconOnly
                                        className="w-full mt-2"
                                        color="primary"
                                        size="sm"
                                        variant="bordered"
                                        onPress={() => {
                                        }}
                                    >
                                        <img alt="search" className="h-4 w-4" src="/search.svg"/>
                                    </Button>
                                </Tooltip>
                                <Tooltip showArrow content="New thread" placement="right">
                                    <Button
                                        isIconOnly
                                        className="w-full mt-2"
                                        color="primary"
                                        size="sm"
                                        onPress={() => {
                                        }}
                                    >
                                        <img alt="search" className="h-4 w-4" src="/layers.svg"/>
                                    </Button>
                                </Tooltip>
                            </div>

                        </>
                    ) : (
                        <>
                            <div className="p-2">
                                <div className="w-full flex flex-row justify-between items-center mb-2">
                                    <div className="flex items-center">
                                        <img
                                            alt="stratifi logo"
                                            className="mr-2"
                                            height="20"
                                            src="/logo.png"
                                            width="20"
                                        />
                                        <span className="tracking-widest">STRATIFI</span>
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
                                        <img alt="collapse" src="/chevron-left.svg"/>
                                    </Button>
                                </div>
                                {/* search box */}
                                <Button
                                    fullWidth
                                    className="flex flex-row start mb-2"
                                    color="primary"
                                    size="sm"
                                    variant="bordered"
                                >
                                    <div className="w-full flex start text-white rounded items-center">
                                        <img alt="search" className="mr-2" src="/search.svg"/>
                                        Search
                                    </div>
                                </Button>
                                {/* new thread button*/}
                                <Button
                                    fullWidth
                                    color="primary"
                                    size="sm"
                                    startContent={<img alt="new thread" src="/layers.svg"/>}
                                >
                                    New thread
                                </Button>
                            </div>
                            {/* thread buttons */}
                            <ThreadButtons/>
                        </>
                    )}
                </CardBody>
                <CardFooter className="bg-black rounded-none">
                    {collapsed ? (
                        <>
                            <Tooltip showArrow content="Account settings" placement="right">
                                <Button
                                    isIconOnly
                                    className="w-full"
                                    // className="min-w-8 w-8 h-8 rounded mr-4"
                                    color="primary"
                                    size="sm"
                                    onClick={() => isOpen}
                                >
                                    <img
                                        alt="avatar"
                                        className="h-4 w-4"
                                        src="/default-user-avatar.svg"
                                    />
                                </Button>
                            </Tooltip>
                        </>
                    ) : (
                        <Button
                            isIconOnly
                            className="min-w-8 w-8 h-8 rounded mr-4"
                            color="primary"
                            size="sm"
                            onClick={() => isOpen}
                        >
                            <img
                                alt="avatar"
                                className="h-4 w-4"
                                src="/default-user-avatar.svg"
                            />
                        </Button>
                    )}
                </CardFooter>
            </Card>
            <Modal isOpen={isOpen} size="5xl" onOpenChange={onOpenChange}>
                <ModalContent>
                    {(onClose) => (
                        <>
                            <ModalHeader className="flex items-center">
                                <i className="pi pi-sliders-v mr-2"/>
                                Settings
                            </ModalHeader>
                            <ModalBody>
                                <Card>
                                    <CardHeader>Account</CardHeader>
                                    <Divider/>
                                    <CardBody>
                                        <div className="flex flex-col">
                                            Email
                                            <br/>
                                            test@test.com
                                            {/*{user?.email}*/}
                                        </div>
                                    </CardBody>
                                </Card>
                                <Card>
                                    <CardBody>
                                        <div className="flex justify-between items-center">
                                            <div>System</div>
                                            <div/>
                                            <div>
                                                <Button
                                                    size="sm"
                                                    startContent={
                                                        <i className="pi pi-trash color-danger"/>
                                                    }
                                                    variant="light"
                                                >
                                                    Delete Account
                                                </Button>
                                                <Button
                                                    color="primary"
                                                    size="sm"
                                                    startContent={<i className="pi pi-sign-out"/>}
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
    )
}