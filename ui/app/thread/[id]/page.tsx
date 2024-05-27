"use client"

import PromptInput from "@/app/thread/[id]/prompt-input";
import Messages, {MessagesQuery} from "@/components/messages/messages";
import {messagesQuery} from "@/__generated__/messagesQuery.graphql";
import { useLazyLoadQuery } from "react-relay";

export default function Page({ params }: { params: { id: string } }) {
    const data = useLazyLoadQuery<messagesQuery>(MessagesQuery, {threadId: params.id})
    const threadHasMessages = data.messages?.edges?.length ?? 0 > 0
    return (
        <>
            <div className="w-full h-full flex flex-col place-content-center px-52 pb-4">
                <Messages threadId={params.id} />
                {
                    !threadHasMessages && (
                        <span className="w-full text-center mb-2"><h1>What do you want to explore?</h1></span>
                    )
                }
                <div className="mt-6">
                    <PromptInput showModelSelect={!threadHasMessages}/>
                </div>
            </div>
        </>

    )
}