import {graphql, useFragment} from "react-relay";

import { ThreadFragment$key } from "@/__generated__/ThreadFragment.graphql";

type props = {
  thread?: ThreadFragment$key | null;
};

const ThreadFragment = graphql`
  fragment ThreadFragment on Thread {
    id
    name
  }
`;

export default function Thread({ thread }: props) {
  const data = useFragment(ThreadFragment, thread)

  return <div>{data?.name}</div>;
}
