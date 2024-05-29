/**
 * @generated SignedSource<<97cf0fb20807171d67f7fa4861aa2c3a>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { Fragment, ReaderFragment } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type MessageFragment$data = {
  readonly content: string;
  readonly createdAt: any;
  readonly id: string;
  readonly " $fragmentType": "MessageFragment";
};
export type MessageFragment$key = {
  readonly " $data"?: MessageFragment$data;
  readonly " $fragmentSpreads": FragmentRefs<"MessageFragment">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "MessageFragment",
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

(node as any).hash = "c4e760dbc94042700f9191595046d2b1";

export default node;
