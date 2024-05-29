/**
 * @generated SignedSource<<f0e835d53960aea82a8a93a6ab3b7460>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { Fragment, ReaderFragment } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type ThreadFragment$data = {
  readonly createdAt: any;
  readonly id: string;
  readonly name: string;
  readonly " $fragmentSpreads": FragmentRefs<"MessagesFragment">;
  readonly " $fragmentType": "ThreadFragment";
};
export type ThreadFragment$key = {
  readonly " $data"?: ThreadFragment$data;
  readonly " $fragmentSpreads": FragmentRefs<"ThreadFragment">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "ThreadFragment",
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
      "name": "name",
      "storageKey": null
    },
    {
      "args": null,
      "kind": "FragmentSpread",
      "name": "MessagesFragment"
    }
  ],
  "type": "Thread",
  "abstractKey": null
};

(node as any).hash = "0bbef03e120a6683e66d5820197c8c40";

export default node;
