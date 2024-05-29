/**
 * @generated SignedSource<<e57a69b9a9a75ec2ae2e7a03b23469f5>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest, Mutation } from 'relay-runtime';
export type deleteThreadButtonMutation$variables = {
  id: string;
};
export type deleteThreadButtonMutation$data = {
  readonly deleteThread: string;
};
export type deleteThreadButtonMutation = {
  response: deleteThreadButtonMutation$data;
  variables: deleteThreadButtonMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "id"
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "id"
  }
],
v2 = {
  "alias": null,
  "args": (v1/*: any*/),
  "kind": "ScalarField",
  "name": "deleteThread",
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "deleteThreadButtonMutation",
    "selections": [
      (v2/*: any*/)
    ],
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "deleteThreadButtonMutation",
    "selections": [
      (v2/*: any*/),
      {
        "alias": null,
        "args": (v1/*: any*/),
        "filters": null,
        "handle": "deleteRecord",
        "key": "",
        "kind": "ScalarHandle",
        "name": "deleteThread"
      }
    ]
  },
  "params": {
    "cacheID": "31924cde8d8ea139d7a8041a9279b13f",
    "id": null,
    "metadata": {},
    "name": "deleteThreadButtonMutation",
    "operationKind": "mutation",
    "text": "mutation deleteThreadButtonMutation(\n  $id: ID!\n) {\n  deleteThread(id: $id)\n}\n"
  }
};
})();

(node as any).hash = "20d99c02e27cc5e756bade2fa1f4fbac";

export default node;
