/**
 * @generated SignedSource<<f67f679eaae2811a7fe75d0871f4f580>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { Fragment, ReaderFragment } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type testMessageFragment$data = {
  readonly content: string;
  readonly createdAt: any;
  readonly id: string;
  readonly " $fragmentType": "testMessageFragment";
};
export type testMessageFragment$key = {
  readonly " $data"?: testMessageFragment$data;
  readonly " $fragmentSpreads": FragmentRefs<"testMessageFragment">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "testMessageFragment",
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
      "name": "createdAt",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "content",
      "storageKey": null
    }
  ],
  "type": "Message",
  "abstractKey": null
};

(node as any).hash = "819f3d36504c241cabb28627ede07498";

export default node;
