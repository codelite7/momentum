"use client"

import {Avatar, Button, Card, CardBody, CardHeader, Snippet, Tooltip} from "@nextui-org/react"
import { useState } from "react"

interface props {
    query: string
    content: string
    sentAt: string
    user: any
}

export default function MessageResult({query, content, sentAt, user}: props) {
    const sentAtDate = new Date(sentAt)
    const sentToday = new Date(sentAt).setHours(0,0,0,0) == new Date().setHours(0,0,0,0)
    const sentAtDisplay = sentToday ? sentAtDate.toLocaleTimeString() : sentAtDate.toLocaleString()
    return (
        <Card className="">
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
                                />
                                {sentAtDisplay}
                            </>
                        )
                    }
                </div>
            </CardHeader>
            <CardBody>
                {getMatchWithSnippet(query, content, 25, 25)}
            </CardBody>
        </Card>
    )
}

function getMatchWithSnippet(query:string, content: string, leading: number, trailing: number): JSX.Element {
    let lowerQuery = query.toLowerCase()
    let lowerContent = content.toLowerCase()
    let queryStartsAt = lowerContent.indexOf(lowerQuery)
    let queryEndsAt = queryStartsAt + query.length
    let prefix = content.substring(queryStartsAt - leading, queryStartsAt)
    let suffix = content.substring(queryEndsAt, queryEndsAt + trailing)
    return (
        <span>
            ...{prefix}
            <Snippet hideCopyButton color="primary" variant="solid" size="sm" symbol="">
                {` ${query} `}
            </Snippet>
            {suffix}...
        </span>
    )
}