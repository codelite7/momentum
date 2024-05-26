/**
 * @generated SignedSource<<8eb1bf8938e8b5662e146bf2dc3b065f>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest, Query } from 'relay-runtime';
export type threadsQuery$variables = Record<PropertyKey, never>;
export type threadsQuery$data = {
  readonly threads: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly createdAt: any;
        readonly id: string;
        readonly messages: {
          readonly edges: ReadonlyArray<{
            readonly node: {
              readonly content: string;
              readonly createdAt: any;
              readonly id: string;
              readonly response: {
                readonly content: string | null | undefined;
                readonly id: string;
              } | null | undefined;
            } | null | undefined;
          } | null | undefined> | null | undefined;
        };
        readonly name: string;
      } | null | undefined;
    } | null | undefined> | null | undefined;
    readonly totalCount: number;
  };
};
export type threadsQuery = {
  response: threadsQuery$data;
  variables: threadsQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "createdAt",
  "storageKey": null
},
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "content",
  "storageKey": null
},
v3 = [
  {
    "alias": null,
    "args": null,
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
              (v0/*: any*/),
              (v1/*: any*/),
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "name",
                "storageKey": null
              },
              {
                "alias": null,
                "args": [
                  {
                    "kind": "Literal",
                    "name": "first",
                    "value": 5
                  }
                ],
                "concreteType": "MessageConnection",
                "kind": "LinkedField",
                "name": "messages",
                "plural": false,
                "selections": [
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
                        "selections": [
                          (v0/*: any*/),
                          (v1/*: any*/),
                          (v2/*: any*/),
                          {
                            "alias": null,
                            "args": null,
                            "concreteType": "Response",
                            "kind": "LinkedField",
                            "name": "response",
                            "plural": false,
                            "selections": [
                              (v0/*: any*/),
                              (v2/*: any*/)
                            ],
                            "storageKey": null
                          }
                        ],
                        "storageKey": null
                      }
                    ],
                    "storageKey": null
                  }
                ],
                "storageKey": "messages(first:5)"
              }
            ],
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
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "threadsQuery",
    "selections": (v3/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "threadsQuery",
    "selections": (v3/*: any*/)
  },
  "params": {
    "cacheID": "c0122979c9730ad427b0745d8515996f",
    "id": null,
    "metadata": {},
    "name": "threadsQuery",
    "operationKind": "query",
    "text": "query threadsQuery {\n  threads {\n    totalCount\n    edges {\n      node {\n        id\n        createdAt\n        name\n        messages(first: 5) {\n          edges {\n            node {\n              id\n              createdAt\n              content\n              response {\n                id\n                content\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "a3bd33dd084bac1e4db02a1560f2ce11";

export default node;
