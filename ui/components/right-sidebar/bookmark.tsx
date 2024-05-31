import { graphql, useFragment, useMutation } from "react-relay";
import { Button, Card, CardBody, CardHeader, Tooltip } from "@nextui-org/react";
import LinesEllipsis from "react-lines-ellipsis";
import { toast } from "sonner";

import {
  bookmarkFragment$data,
  bookmarkFragment$key,
} from "@/__generated__/bookmarkFragment.graphql";

export const BookmarkFragment = graphql`
  fragment bookmarkFragment on Bookmark {
    id
    message {
      id
      content
      thread {
        name
      }
    }
    response {
      id
      content
      message {
        thread {
          name
        }
      }
    }
  }
`;

const BookmarkDeleteMutation = graphql`
  mutation bookmarkDeleteMutation($id: ID!) {
    deleteBookmark(id: $id) @deleteRecord
  }
`;

type props = {
  bookmark: bookmarkFragment$key;
  showThreadName?: boolean;
};

export default function Bookmark({ bookmark, showThreadName }: props) {
  const data = useFragment<bookmarkFragment$key>(BookmarkFragment, bookmark);
  const [commitMutation, isMutationInFlight] = useMutation(
    BookmarkDeleteMutation,
  );

  return (
    <Card isHoverable className="border-1 border-default">
      <CardHeader>
        <div className="flex w-full justify-between">
          <div className="items-center gap-2 ">
            <i className="pi pi-bookmark text-primary" />
            {showThreadName && (
              <span>
                {data.message
                  ? data.message.thread.name
                  : data.response?.message.thread.name}
              </span>
            )}
          </div>
          <div>
            <Tooltip content="Copy to clipboard">
              <Button
                isIconOnly
                className="bg-transparent hover:bg-default-200"
                size="sm"
                onPress={() =>
                  navigator.clipboard.writeText(getClipboardContent(data))
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
                  commitMutation({
                    variables: {
                      id: data.id,
                    },
                    onError: (e) => {
                      toast.error("Error deleting bookmark");
                    },
                    onCompleted: (response) => {
                      toast.success("Deleted bookmark");
                    },
                  });
                }}
              >
                <i className="pi pi-times" />
              </Button>
            </Tooltip>
          </div>
        </div>
      </CardHeader>
      <CardBody>
        {data.message?.content && (
          <LinesEllipsis maxLine={8} text={data.message.content} />
        )}
        {data.response?.content && (
          <LinesEllipsis maxLine={8} text={data.response.content} />
        )}
      </CardBody>
    </Card>
  );
}

function getClipboardContent(bookmark: bookmarkFragment$data) {
  if (bookmark.message?.content) {
    return bookmark.message.content;
  } else if (bookmark.response?.content) {
    return bookmark.response.content;
  }

  return "";
}
