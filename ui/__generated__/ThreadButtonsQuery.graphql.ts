/**
 * @generated SignedSource<<8e62c1c614f623eda07cb1ff990d9dcb>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest, Query } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type ThreadButtonsQuery$variables = Record<PropertyKey, never>;
export type ThreadButtonsQuery$data = {
  readonly " $fragmentSpreads": FragmentRefs<"ThreadButtonsPaginationFragment">;
};
export type ThreadButtonsQuery = {
  response: ThreadButtonsQuery$data;
  variables: ThreadButtonsQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "kind": "Literal",
    "name": "first",
    "value": 50
  }
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "ThreadButtonsQuery",
    "selections": [
      {
        "args": null,
        "kind": "FragmentSpread",
        "name": "ThreadButtonsPaginationFragment"
      }
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "ThreadButtonsQuery",
    "selections": [
      {
        "alias": null,
        "args": (v0/*: any*/),
        "concreteType": "ThreadConnection",
        "kind": "LinkedField",
        "name": "threads",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "totalCount",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "concreteType": "ThreadEdge",
            "kind": "LinkedField",
            "name": "edges",
            "plural": true,
            "selections": [
              {
                "alias": null,
                "args": null,
                "concreteType": "Thread",
                "kind": "LinkedField",
                "name": "node",
                "plural": false,
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
                  },
                  {
                    "alias": null,
                    "args": null,
                    "kind": "ScalarField",
                    "name": "__typename",
                    "storageKey": null
                  }
                ],
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "cursor",
                "storageKey": null
              }
            ],
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "concreteType": "PageInfo",
            "kind": "LinkedField",
            "name": "pageInfo",
            "plural": false,
            "selections": [
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "endCursor",
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "hasNextPage",
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": "threads(first:50)"
      },
      {
        "alias": null,
        "args": (v0/*: any*/),
        "filters": null,
        "handle": "connection",
        "key": "Query_threads",
        "kind": "LinkedHandle",
        "name": "threads"
      }
    ]
  },
  "params": {
    "cacheID": "300a81916a90bd5ca56a7aa04cc72f71",
    "id": null,
    "metadata": {},
    "name": "ThreadButtonsQuery",
    "operationKind": "query",
    "text": "query ThreadButtonsQuery {\n  ...ThreadButtonsPaginationFragment\n}\n\nfragment ThreadButtonFragment on Thread {\n  id\n  name\n}\n\nfragment ThreadButtonsPaginationFragment on Query {\n  threads(first: 50) {\n    totalCount\n    edges {\n      node {\n        id\n        ...ThreadButtonFragment\n        __typename\n      }\n      cursor\n    }\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "15f131e68dca889c897aeef837d0d2d4";

export default node;
