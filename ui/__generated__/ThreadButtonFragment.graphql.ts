/**
 * @generated SignedSource<<cdf6bb8a948841b56beb07d7df35f49a>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { Fragment, ReaderFragment } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type ThreadButtonFragment$data = {
  readonly id: string;
  readonly name: string;
  readonly " $fragmentType": "ThreadButtonFragment";
};
export type ThreadButtonFragment$key = {
  readonly " $data"?: ThreadButtonFragment$data;
  readonly " $fragmentSpreads": FragmentRefs<"ThreadButtonFragment">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "ThreadButtonFragment",
  "selections": [
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "id",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "name",
      "storageKey": null
    }
  ],
  "type": "Thread",
  "abstractKey": null
};

(node as any).hash = "0dfa01c5ed6eb39bedcd52de7becd3b1";

export default node;
