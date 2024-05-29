/**
 * @generated SignedSource<<0a2829c38cda8d2088a2f2fa56fcb20f>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest, Query } from 'relay-runtime';
export type bookmarksQuery$variables = {
  threadId: string;
};
export type bookmarksQuery$data = {
  readonly bookmarks: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly id: string;
        readonly message: {
          readonly content: string;
          readonly id: string;
        } | null | undefined;
      } | null | undefined;
    } | null | undefined> | null | undefined;
    readonly pageInfo: {
      readonly hasNextPage: boolean;
      readonly hasPreviousPage: boolean;
    };
  };
};
export type bookmarksQuery = {
  response: bookmarksQuery$data;
  variables: bookmarksQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "threadId"
  }
],
v1 = [
  {
    "fields": [
      {
        "fields": [
          {
            "kind": "Variable",
            "name": "id",
            "variableName": "threadId"
          }
        ],
        "kind": "ObjectValue",
        "name": "hasThreadWith"
      }
    ],
    "kind": "ObjectValue",
    "name": "hasMessageWith"
  }
],
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v3 = [
  {
    "alias": null,
    "args": [
      {
        "fields": [
          {
            "fields": (v1/*: any*/),
            "kind": "ObjectValue",
            "name": "hasResponseWith"
          },
          {
            "items": [
              {
                "fields": (v1/*: any*/),
                "kind": "ObjectValue",
                "name": "or.0"
              }
            ],
            "kind": "ListValue",
            "name": "or"
          }
        ],
        "kind": "ObjectValue",
        "name": "where"
      }
    ],
    "concreteType": "BookmarkConnection",
    "kind": "LinkedField",
    "name": "bookmarks",
    "plural": false,
    "selections": [
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
            "name": "hasNextPage",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "hasPreviousPage",
            "storageKey": null
          }
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "concreteType": "BookmarkEdge",
        "kind": "LinkedField",
        "name": "edges",
        "plural": true,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "Bookmark",
            "kind": "LinkedField",
            "name": "node",
            "plural": false,
            "selections": [
              (v2/*: any*/),
              {
                "alias": null,
                "args": null,
                "concreteType": "Message",
                "kind": "LinkedField",
                "name": "message",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  {
                    "alias": null,
                    "args": null,
                    "kind": "ScalarField",
                    "name": "content",
                    "storageKey": null
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
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "bookmarksQuery",
    "selections": (v3/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "bookmarksQuery",
    "selections": (v3/*: any*/)
  },
  "params": {
    "cacheID": "4dd084d33e5f179cd5fbc29c934f11f0",
    "id": null,
    "metadata": {},
    "name": "bookmarksQuery",
    "operationKind": "query",
    "text": "query bookmarksQuery(\n  $threadId: ID!\n) {\n  bookmarks(where: {or: [{hasMessageWith: {hasThreadWith: {id: $threadId}}}], hasResponseWith: {hasMessageWith: {hasThreadWith: {id: $threadId}}}}) {\n    pageInfo {\n      hasNextPage\n      hasPreviousPage\n    }\n    edges {\n      node {\n        id\n        message {\n          id\n          content\n        }\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "7e23da39cf72c485b4ea49ae2f552fb8";

export default node;
