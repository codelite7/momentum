"use client"
import {graphql, useLazyLoadQuery} from "react-relay";
import {threadButtonsQuery} from "@/__generated__/threadButtonsQuery.graphql";
import { Button } from "@nextui-org/react";

const ThreadButtonsQuery = graphql`
    query threadButtonsQuery {
        threads(first: 10) {
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
        <div className="w-full h-full overflow-x-hidden overflow-y-auto">
            {
                data.threads.edges?.map(edge => {
                    return (
                        <Button
                            fullWidth
                            size="sm"
                            className="mb-2"
                        >
                            {edge?.node?.name}
                        </Button>
                    )
                })
            }
        </div>
    )
}