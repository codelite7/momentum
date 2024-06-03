"use client";

import { useSuspenseQuery } from "@apollo/client";

import PromptInput from "@/app/thread/[id]/prompt-input";
import { gql, useFragment } from "@/__generated__";
import Messages from "@/components/thread/messages/messages";

type props = {
  threadId: string;
};

const threadByIdQuery = gql(/* Graphql */ `
  query thread($id: ID!) {
    thread(id: $id) {
      ...threadByIdThreadFragment
    }
  }
`);

export const threadByIdThreadFragment = gql(/* Graphql */ `
  fragment threadByIdThreadFragment on Thread {
      id
      createdAt
      name
      messages(first: 5) {
        totalCount
        pageInfo {
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
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
`);
export default function Thread({ threadId }: props) {
  let { data, error } = useSuspenseQuery(threadByIdQuery, {
    variables: { id: threadId },
  });

  const result = useFragment(threadByIdThreadFragment, data.thread);

  return (
    <>
      <Messages thread={result} />
      <PromptInput threadId={threadId} />
    </>
  );
}
