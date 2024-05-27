"use client"
import {graphql, useLazyLoadQuery} from "react-relay";
import {threadButtonsQuery} from "@/__generated__/threadButtonsQuery.graphql";
import ThreadButton from "@/components/left-sidebar/thread-button";
import { ScrollShadow } from "@nextui-org/react";

const ThreadButtonsQuery = graphql`
    query threadButtonsQuery {
        threads {
            edges {
                node {
                    id
                    name
                }
            }
        }
    }
`
export default function threadButtons() {
    const data = useLazyLoadQuery<threadButtonsQuery>(ThreadButtonsQuery, {})

    return (
        <ScrollShadow hideScrollBar>
            {
                data.threads.edges?.map(edge => {
                    return (
                        <ThreadButton key={edge?.node?.id} edge={edge} />
                    )
                })
            }
        </ScrollShadow>
    )
}