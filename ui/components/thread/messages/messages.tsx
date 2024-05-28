"use client";

import { graphql, useFragment } from "react-relay";
import { ScrollShadow } from "@nextui-org/react";

import MessageBase from "@/components/thread/messages/message-base";
import { threadFragment$key } from "@/__generated__/threadFragment.graphql";
import {threadQuery$data} from "@/__generated__/threadQuery.graphql";

export const MessagesFragment = graphql`
  fragment messagesFragment on Thread {
    messages {
        edges {
            node {
                id
            }
        }
    }
  }
`;


interface props {
  messages: threadQuery$data;
}

export default function Messages({ messages }: props) {

  return (
    // <ScrollShadow hideScrollBar>
    //   {data.messages.edges?.map((edge) => {
    //     return (
    //       <>
    //         <MessageBase
    //           key={edge?.node?.id}
    //           content={edge?.node?.content ?? ""}
    //           sentAt={edge?.node?.createdAt ?? ""}
    //           user={{ name: "Me" }}
    //         />
    //         <MessageBase
    //           key={edge?.node?.response?.id}
    //           content={edge?.node?.response?.content ?? ""}
    //           sentAt={edge?.node?.response?.createdAt ?? ""}
    //           user={undefined}
    //         />
    //       </>
    //     );
    //   })}
    // </ScrollShadow>
  );
}
