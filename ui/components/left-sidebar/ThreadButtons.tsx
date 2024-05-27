"use client";

import {
  graphql,
  useFragment,
  useLazyLoadQuery,
  usePaginationFragment,
} from "react-relay";

import {
    ThreadButtonsQuery as ThreadButtonsQueryType,
    ThreadButtonsQuery$data
} from "@/__generated__/ThreadButtonsQuery.graphql";
import ThreadButton from "./ThreadButton";
import Derp from "@/components/left-sidebar/Derp";

// working
// const ThreadButtonsQuery = graphql`
//   query ThreadButtonsQuery($first:Int=3)
//   {
//     threads(first: $first) {
//       edges {
//         node {
//           id
//           name
//         }
//       }
//     }
//   }
// `;

const ThreadButtonsQuery = graphql`
  query ThreadButtonsQuery {
    ...ThreadButtons_threads
  }
`;

const ThreadButtonsThreadFragment = graphql`
  fragment ThreadButtonsThreadFragment on Thread {
    id
    name
  }
`;

const ThreadButtonsPaginationFragment = graphql`
  fragment ThreadButtons_threads on Query
  @argumentDefinitions(
    first: { type: "Int", defaultValue: 10 }
    after: { type: "Cursor" }
  )
  @refetchable(queryName: "ThreadButtonsPaginationQuery") {
    threads(first: $first, after: $after) @connection(key: "Query_threads") {
      totalCount
      edges {
        node {
          ...ThreadButtonsThreadFragment
        }
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
`;

export default function ThreadButtons() {
  const queryData = useLazyLoadQuery<ThreadButtonsQueryType>(
    ThreadButtonsQuery,
    {},
  );
  const { data, loadNext, hasNext } = usePaginationFragment(
    ThreadButtonsPaginationFragment,
    queryData,
  );

  console.log("data ", data);

  return (
    <>
      {
        data.threads.edges.map((edge) => (
            <Derp derp={edge.node} key={edge.node.id} />
        ))
      }
      </>
  );
}

// export default function threadButtons() {
//   const threads = useLazyLoadQuery()
//   const data = useLazyLoadQuery<ThreadButtonsQueryType>(ThreadButtonsQuery, {first: 1});
//   const threads = data.threads.edges;
//
//   return (
//     <ScrollShadow hideScrollBar>
//       {threads?.map((thread) => {
//         return <ThreadButton key={thread?.node?.id} thread={thread?.node} />;
//       })}
//     </ScrollShadow>
//   );
// }
