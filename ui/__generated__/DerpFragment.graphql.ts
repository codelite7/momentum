/**
 * @generated SignedSource<<b335dee8d169e772853ad6cb2610fea2>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { Fragment, ReaderFragment } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type DerpFragment$data = {
  readonly id: string;
  readonly name: string;
  readonly " $fragmentType": "DerpFragment";
};
export type DerpFragment$key = {
  readonly " $data"?: DerpFragment$data;
  readonly " $fragmentSpreads": FragmentRefs<"DerpFragment">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "DerpFragment",
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

(node as any).hash = "ee4c741fcf47b980a1b04b2ad25b4569";

export default node;
