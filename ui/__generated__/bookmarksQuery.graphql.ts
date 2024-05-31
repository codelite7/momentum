/**
 * @generated SignedSource<<3623d6580b47049fbf4614d6c3f1752f>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest, Query } from 'relay-runtime';
import type { FragmentRefs } from "relay-runtime";
export type bookmarksQuery$variables = {
  after?: any | null | undefined;
  first?: number | null | undefined;
  threadId: string;
};
export type bookmarksQuery$data = {
  readonly __id: string;
  readonly " $fragmentSpreads": FragmentRefs<"bookmarksFragment">;
};
export type bookmarksQuery = {
  response: bookmarksQuery$data;
  variables: bookmarksQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "after"
},
v1 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "first"
},
v2 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "threadId"
},
v3 = {
  "kind": "Variable",
  "name": "after",
  "variableName": "after"
},
v4 = {
  "kind": "Variable",
  "name": "first",
  "variableName": "first"
},
v5 = {
  "kind": "ClientExtension",
  "selections": [
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "__id",
      "storageKey": null
    }
  ]
},
v6 = [
  (v3/*: any*/),
  (v4/*: any*/),
  {
    "fields": [
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
    "kind": "ObjectValue",
    "name": "where"
  }
],
v7 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v8 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "content",
  "storageKey": null
},
v9 = {
  "alias": null,
  "args": null,
  "concreteType": "Thread",
  "kind": "LinkedField",
  "name": "thread",
  "plural": false,
  "selections": [
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "name",
      "storageKey": null
    },
    (v7/*: any*/)
  ],
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": [
      (v0/*: any*/),
      (v1/*: any*/),
      (v2/*: any*/)
    ],
    "kind": "Fragment",
    "metadata": null,
    "name": "bookmarksQuery",
    "selections": [
      {
        "args": [
          (v3/*: any*/),
          (v4/*: any*/),
          {
            "kind": "Variable",
            "name": "threadId",
            "variableName": "threadId"
          }
        ],
        "kind": "FragmentSpread",
        "name": "bookmarksFragment"
      },
      (v5/*: any*/)
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [
      (v1/*: any*/),
      (v0/*: any*/),
      (v2/*: any*/)
    ],
    "kind": "Operation",
    "name": "bookmarksQuery",
    "selections": [
      {
        "alias": null,
        "args": (v6/*: any*/),
        "concreteType": "BookmarkConnection",
        "kind": "LinkedField",
        "name": "bookmarks",
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
                  (v7/*: any*/),
                  {
                    "alias": null,
                    "args": null,
                    "concreteType": "Message",
                    "kind": "LinkedField",
                    "name": "message",
                    "plural": false,
                    "selections": [
                      (v7/*: any*/),
                      (v8/*: any*/),
                      (v9/*: any*/)
                    ],
                    "storageKey": null
                  },
                  {
                    "alias": null,
                    "args": null,
                    "concreteType": "Response",
                    "kind": "LinkedField",
                    "name": "response",
                    "plural": false,
                    "selections": [
                      (v7/*: any*/),
                      (v8/*: any*/),
                      {
                        "alias": null,
                        "args": null,
                        "concreteType": "Message",
                        "kind": "LinkedField",
                        "name": "message",
                        "plural": false,
                        "selections": [
                          (v9/*: any*/),
                          (v7/*: any*/)
                        ],
                        "storageKey": null
                      }
                    ],
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
          (v5/*: any*/)
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": (v6/*: any*/),
        "filters": [
          "where"
        ],
        "handle": "connection",
        "key": "Thread_bookmarks",
        "kind": "LinkedHandle",
        "name": "bookmarks"
      },
      (v5/*: any*/)
    ]
  },
  "params": {
    "cacheID": "4ff080c8a5cfa922b80f97dac8ddb2d4",
    "id": null,
    "metadata": {},
    "name": "bookmarksQuery",
    "operationKind": "query",
    "text": "query bookmarksQuery(\n  $first: Int\n  $after: Cursor\n  $threadId: ID!\n) {\n  ...bookmarksFragment_2iXj1n\n}\n\nfragment bookmarkFragment on Bookmark {\n  id\n  message {\n    id\n    content\n    thread {\n      name\n      id\n    }\n  }\n  response {\n    id\n    content\n    message {\n      thread {\n        name\n        id\n      }\n      id\n    }\n  }\n}\n\nfragment bookmarksFragment_2iXj1n on Query {\n  bookmarks(first: $first, after: $after, where: {hasMessageWith: {hasThreadWith: {id: $threadId}}}) {\n    totalCount\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n    edges {\n      node {\n        id\n        ...bookmarkFragment\n        __typename\n      }\n      cursor\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "f98692c930f148e25be5c977ee19b047";

export default node;
