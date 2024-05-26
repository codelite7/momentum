"use client"

import {graphql, useLazyLoadQuery} from "react-relay";
import {threadsQuery} from "@/__generated__/threadsQuery.graphql";

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
        <div>
            Threads: {data.threads.totalCount}
        </div>
    )
}