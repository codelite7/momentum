//
// import {
//   bookmarkFragment$data,
//   bookmarkFragment$key,
// } from "@/__generated__/bookmarkFragment.graphql";
//
// export const BookmarkFragment = graphql`
//   fragment bookmarkFragment on Bookmark {
//     id
//     message {
//       id
//       content
//       thread {
//         name
//       }
//     }
//     response {
//       id
//       content
//       message {
//         thread {
//           name
//         }
//       }
//     }
//   }
// `;

// const BookmarkDeleteMutation = graphql`
//   mutation bookmarkDeleteMutation($id: ID!) {
//     deleteBookmark(id: $id) @deleteRecord
//   }
// `;

import { Button, Card, CardBody, CardHeader, Tooltip } from "@nextui-org/react";
import LinesEllipsis from "react-lines-ellipsis";
import { useMutation } from "@apollo/client";
import { toast } from "sonner";

import { Bookmark } from "@/__generated__/graphql";
import {
  deleteBookmarkMutation,
  threadBookmarksQuery,
  threadByIdQuery,
} from "@/graphql-queries/queries";

type props = {
  showThreadName?: boolean;
  bookmark: Bookmark;
};

export default function Bookmark({ showThreadName, bookmark }: props) {
  const [deleteBookmark] = useMutation(deleteBookmarkMutation, {
    variables: {
      id: bookmark.id,
    },
    onError: (e) => {
      toast.error("Error deleting bookmark");
      console.error(e);
    },
    onCompleted: (data) => {
      toast.success("Deleted bookmark");
    },
    refetchQueries: [threadByIdQuery, threadBookmarksQuery],
  });

  return (
    <Card isHoverable className="border-1 border-default">
      <CardHeader>
        <div className="flex w-full justify-between">
          <div className="items-center gap-2 ">
            <i className="pi pi-bookmark text-primary" />
            {showThreadName && <span>{bookmark.message?.thread.name}</span>}
          </div>
          <div>
            <Tooltip content="Copy to clipboard">
              <Button
                isIconOnly
                className="bg-transparent hover:bg-default-200"
                size="sm"
                onPress={() =>
                  navigator.clipboard.writeText(bookmark.message?.content ?? "")
                }
              >
                <i className="pi pi-clipboard" />
              </Button>
            </Tooltip>
            <Tooltip content="Delete bookmark">
              <Button
                isIconOnly
                className="bg-transparent hover:bg-default-200"
                size="sm"
                onPress={() => {
                  deleteBookmark();
                }}
              >
                <i className="pi pi-times" />
              </Button>
            </Tooltip>
          </div>
        </div>
      </CardHeader>
      <CardBody>
        <LinesEllipsis maxLine={8} text={bookmark.message?.content} />
      </CardBody>
    </Card>
  );
}
