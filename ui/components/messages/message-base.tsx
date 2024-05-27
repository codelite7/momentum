"use client"

import {Avatar, Button, Card, CardBody, CardHeader, Snippet, Tooltip} from "@nextui-org/react"
import { useState } from "react"

interface props {
    content: string
    sentAt: string
    user: any
}

export default function MessageBase({content, sentAt, user}: props) {
    const sentAtDate = new Date(sentAt)
    const sentToday = new Date(sentAt).setHours(0,0,0,0) == new Date().setHours(0,0,0,0)
    const sentAtDisplay = sentToday ? sentAtDate.toLocaleTimeString() : sentAtDate.toLocaleString()
    const [hovered, setHovered] = useState(false)
    return (
        <Card className="mb-6 bg-black hover:bg-default-100 mr-2" onMouseEnter={() => setHovered(true)} onMouseLeave={() => setHovered(false)}>
            <CardHeader className="justify-between">
                {/* avatar, sender, time sent*/}
                <div className="flex">
                    {
                        user ? (
                            <>
                                <Avatar
                                    className="rounded-br-none mr-2"
                                    radius="sm"
                                    size="md"
                                    src="/default-user-avatar.svg"
                                    imgProps={{className: "h-6 w-6"}}
                                />
                                Me {sentAtDisplay}
                            </>
                        ) : (
                            <>
                                <Avatar
                                    className="bg-transparent rounded-bl-none mr-3"
                                    color="primary"
                                    isBordered
                                    radius="sm"
                                    size="md"
                                    src="/logo.png"
                                    imgProps={{className: "h-8 w-8"}}
                                    // icon={<img src={"/logo.png"} alt="stratifi" className="h-6 w-6" />}
                                />
                                {sentAtDisplay}
                            </>
                        )
                    }
                </div>
                {/* context actions */}
                {
                    hovered ? (
                        <div className="flex gap-1">
                            {
                                !user && (
                                    <Tooltip showArrow placement="bottom" content="Generate new response">
                                        <Button
                                            size="sm"
                                            isIconOnly
                                        >
                                            <i className="pi pi-sync" />
                                        </Button>
                                    </Tooltip>

                                )
                            }
                            <Tooltip showArrow placement="bottom" content="Bookmark">
                            <Button
                                size="sm"
                                isIconOnly
                            >
                                <i className="pi pi-bookmark" />
                            </Button>
                            </Tooltip>
                            <Tooltip showArrow placement="bottom" content="Copy message to clipboard">
                            <Button
                                size="sm"
                                isIconOnly
                                onClick={() => navigator.clipboard.writeText(content)}
                            >
                                <i className="pi pi-clipboard" />
                            </Button>
                            </Tooltip>
                            <Tooltip showArrow placement="bottom" content="Create child thread">
                            <Button
                                size="sm"
                                isIconOnly
                            >
                                <i className="pi pi-arrow-right" />
                            </Button>
                            </Tooltip>
                        </div>
                    ) : null
                }

            </CardHeader>
            <CardBody>
                {content}
            </CardBody>
        </Card>
    )
}