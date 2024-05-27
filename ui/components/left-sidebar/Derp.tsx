import {graphql, useFragment} from "react-relay";
import {DerpFragment$key} from "@/__generated__/DerpFragment.graphql";

const DerpFragment = graphql`
  fragment DerpFragment on Thread {
      id
      name
  }
`

type Props = {
    derp: DerpFragment$key;
}

export default function Derp({derp}: Props) {
    const data = useFragment(DerpFragment, derp)
    return (
        <div>{data.name}</div>
    )
}