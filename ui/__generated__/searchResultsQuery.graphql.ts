/**
 * @generated SignedSource<<a58a37e8f16dc2ad00b46039f857b0e3>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest, Query } from 'relay-runtime';
export type searchResultsQuery$variables = {
  query: string;
};
export type searchResultsQuery$data = {
  readonly messages: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly content: string;
        readonly createdAt: any;
        readonly id: string;
      } | null | undefined;
    } | null | undefined> | null | undefined;
    readonly totalCount: number;
  };
  readonly responses: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly content: string | null | undefined;
        readonly createdAt: any;
        readonly id: string;
      } | null | undefined;
    } | null | undefined> | null | undefined;
    readonly totalCount: number;
  };
};
export type searchResultsQuery = {
  response: searchResultsQuery$data;
  variables: searchResultsQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "query"
  }
],
v1 = [
  {
    "fields": [
      {
        "kind": "Variable",
        "name": "contentContainsFold",
        "variableName": "query"
      }
    ],
    "kind": "ObjectValue",
    "name": "where"
  }
],
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "totalCount",
  "storageKey": null
},
v3 = [
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
v4 = [
  {
    "alias": null,
    "args": (v1/*: any*/),
    "concreteType": "MessageConnection",
    "kind": "LinkedField",
    "name": "messages",
    "plural": false,
    "selections": [
      (v2/*: any*/),
      {
        "alias": null,
        "args": null,
        "concreteType": "MessageEdge",
        "kind": "LinkedField",
        "name": "edges",
        "plural": true,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "Message",
            "kind": "LinkedField",
            "name": "node",
            "plural": false,
            "selections": (v3/*: any*/),
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "storageKey": null
  },
  {
    "alias": null,
    "args": (v1/*: any*/),
    "concreteType": "ResponseConnection",
    "kind": "LinkedField",
    "name": "responses",
    "plural": false,
    "selections": [
      (v2/*: any*/),
      {
        "alias": null,
        "args": null,
        "concreteType": "ResponseEdge",
        "kind": "LinkedField",
        "name": "edges",
        "plural": true,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "Response",
            "kind": "LinkedField",
            "name": "node",
            "plural": false,
            "selections": (v3/*: any*/),
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "searchResultsQuery",
    "selections": (v4/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "searchResultsQuery",
    "selections": (v4/*: any*/)
  },
  "params": {
    "cacheID": "42414bc5838f4f836e56ee112f27abc8",
    "id": null,
    "metadata": {},
    "name": "searchResultsQuery",
    "operationKind": "query",
    "text": "query searchResultsQuery(\n  $query: String!\n) {\n  messages(where: {contentContainsFold: $query}) {\n    totalCount\n    edges {\n      node {\n        id\n        createdAt\n        content\n      }\n    }\n  }\n  responses(where: {contentContainsFold: $query}) {\n    totalCount\n    edges {\n      node {\n        id\n        createdAt\n        content\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "3345e42ec15257d38edf841051d507eb";

export default node;
