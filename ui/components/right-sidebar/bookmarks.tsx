"use client"

import { graphql, useLazyLoadQuery } from "react-relay"
import {bookmarksQuery} from "@/__generated__/bookmarksQuery.graphql";
import {Card, CardBody, CardHeader, ScrollShadow } from "@nextui-org/react";

export const BookmarksQuery = graphql`
query bookmarksQuery($threadId: ID!) {
  bookmarks(
    where: {or: [{hasMessageWith: {hasThreadWith: {id: $threadId}}}], hasResponseWith: {hasMessageWith: {hasThreadWith: {id: $threadId}}}}
  ) {
    pageInfo {
      hasNextPage
      hasPreviousPage
    }
    edges {
      node {
        id
        message {
          id
          content
        }
      }
    }
  }
}
`
export default function Bookmarks({ params }: { params: { id: string } }) {
    const data = useLazyLoadQuery<bookmarksQuery>(BookmarksQuery, {threadId: params.id})
    const threadHasBookmarks = data.bookmarks?.edges?.length ?? 0 > 0
    return (
        <>
            {
                threadHasBookmarks ? (
                    <ScrollShadow>
                        {
                            data.bookmarks.edges?.map(edge => {
                                return (
                                    <span key={edge?.node?.id}>{edge?.node?.id}</span>
                                )
                            })
                        }
                    </ScrollShadow>
                ) : (
                    <Card className="border-1 border-primary">
                        <CardHeader className="flex justify-center">
                            <i className="pi pi-bookmark text-primary"  />
                        </CardHeader>
                        <CardBody>
                            <span className="text-sm text-primary-800 text-center">Bookmark important pieces of your conversation to collect them here</span>
                        </CardBody>
                    </Card>
                )
            }

        </>
    )
}