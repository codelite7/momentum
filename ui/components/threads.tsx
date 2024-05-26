"use client"
import {graphql, useLazyLoadQuery} from "react-relay";
import {threadsQuery} from "@/__generated__/threadsQuery.graphql";
import {Suspense} from "react";

const ThreadsQuery = graphql`
  query threadsQuery {
    threads {
        totalCount
        edges {
          node {
            id
            createdAt
            name
            messages(first:5){
              edges {
                node {
                  id
                  createdAt
                  content
                  response {
                    id
                    content
                  }
                }
              }
            }
          }
        }
    }
  }
`

export default function Threads() {
    const data = useLazyLoadQuery<threadsQuery>(ThreadsQuery, {})

    return (
        <div className="h-full overflow-y-auto">
            Threads: <span>{JSON.stringify(data.threads, null, 2)}</span>
            {/*Threads: <span>{data.threads.totalCount}</span>*/}
        </div>

    )
}