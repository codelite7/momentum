/**
 * @generated SignedSource<<88fc461366cb1fd054d80a6583557379>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { Fragment, ReaderFragment } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type responseFragment$data = {
  readonly content: string | null | undefined;
  readonly createdAt: any;
  readonly id: string;
  readonly " $fragmentType": "responseFragment";
};
export type responseFragment$key = {
  readonly " $data"?: responseFragment$data;
  readonly " $fragmentSpreads": FragmentRefs<"responseFragment">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "responseFragment",
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
  "type": "Response",
  "abstractKey": null
};

(node as any).hash = "6a902504e9969ab894dfecd9b6f04963";

export default node;
