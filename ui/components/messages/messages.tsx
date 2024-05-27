"use client"

import {messagesQuery} from "@/__generated__/messagesQuery.graphql";
import {graphql, useLazyLoadQuery} from "react-relay";
import MessageBase from "@/components/messages/message-base";
import { ScrollShadow } from "@nextui-org/react";

export const MessagesQuery = graphql`
  query messagesQuery($threadId:ID!){
  messages(where:{hasThreadWith:{id:$threadId}}) {
    pageInfo {
      hasNextPage
      hasPreviousPage
    }
    edges {
      node {
        id
        createdAt
        content
        response {
          id
          createdAt
          content
        }
      }
    }
  }
}
`

interface props {
    threadId: string,
}

export default function Messages({threadId}: props) {
    const data = useLazyLoadQuery<messagesQuery>(MessagesQuery, {threadId})
    return (
        <ScrollShadow hideScrollBar>
            {
                data.messages.edges?.map(edge => {
                    return (
                        <>
                            <MessageBase key={edge?.node?.id} content={edge?.node?.content ?? ''} sentAt={edge?.node?.createdAt ?? ''} user={{name: 'Me'}} />
                            <MessageBase key={edge?.node?.response?.id} content={edge?.node?.response?.content ?? ''} sentAt={edge?.node?.response?.createdAt ?? ''} user={undefined} />
                        </>
                    )
                })
            }
        </ScrollShadow>
    )
}