import {graphql, useLazyLoadQuery} from "react-relay"
import {searchResultsQuery} from "@/__generated__/searchResultsQuery.graphql";
import {memo, useEffect, useMemo, useState} from "react";
import {Card, CardBody, ScrollShadow} from "@nextui-org/react";
import MessageResult from "@/components/search/message-result";

const SearchResultsQuery = graphql`
query searchResultsQuery($query:String!) {
  messages(where: {contentContainsFold:$query}) {
    totalCount
    edges {
      node {
        id
        createdAt
        content
      }
    }
  }
  responses(where: {contentContainsFold:$query}) {
    totalCount
    edges {
      node {
        id
        createdAt
        content
      }
    }
  }
}
`

interface props {
    query: string
}

export default memo(function SearchResults({query}: props) {
    const data = useLazyLoadQuery<searchResultsQuery>(SearchResultsQuery, {query})
    const totalResults = data.messages.totalCount + data.responses.totalCount
    return (
        <>
            {
                totalResults > 0 ? (
                    <>
                        <span>Matches: {totalResults}</span>
                        <ScrollShadow hideScrollBar className="max-h-96">
                            {
                                data.messages.edges?.map(edge => {
                                    return (
                                        <MessageResult query={query} content={edge?.node?.content ?? ''}
                                                       sentAt={edge?.node?.createdAt} user={{}}/>
                                    )
                                })

                            }
                            {
                                data.responses.edges?.map(edge => {
                                    return (
                                        <MessageResult query={query} content={edge?.node?.content ?? ''}
                                                       sentAt={edge?.node?.createdAt} user={undefined}/>
                                    )
                                })
                            }
                        </ScrollShadow>
                    </>
                ) : (
                    <div className="w-full text-center">
                        No results
                    </div>
                )
            }
        </>
    )
})