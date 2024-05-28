import { graphql, useFragment } from "react-relay";

import { responseFragment$key } from "@/__generated__/responseFragment.graphql";
import MessageBase from "@/components/thread/messages/message-base";

export const ResponseFragment = graphql`
  fragment responseFragment on Response {
    id
    createdAt
    content
  }
`;

type props = {
  response: responseFragment$key;
};
export default function Response({ response }: props) {
  const data = useFragment(ResponseFragment, response);

  return (
    <MessageBase
      content={data.content}
      sentAt={data.createdAt}
      user={undefined}
    />
  );
}
