import { useAtomValue } from "jotai/index";
import { Link } from "@nextui-org/link";
import { Button } from "@nextui-org/react";

import { threadAtom } from "@/state/atoms";

export default function ThreadHeader() {
  const thread = useAtomValue(threadAtom);

  return (
    // container
    <div className="w-full flex items-center justify-between border-b-1 border-default bg-card-background min-h-11 px-4">
      {/* name / links*/}
      <div>
        {/* parent thread link */}
        {thread && thread.parent && (
          <>
            <Link color="primary" href={`/thread/${thread.parent.thread.id}`}>
              {thread.parent.thread.name}
            </Link>
            &nbsp;/&nbsp;
          </>
        )}
        <span>{thread?.name}</span>
      </div>
      {/* back button */}
      {thread && thread.parent && (
        <Button
          as={Link}
          href={`/thread/${thread.parent.thread.id}`}
          radius="sm"
          size="sm"
          startContent={<i className="pi pi-arrow-left" />}
          variant="bordered"
        >
          Back
        </Button>
      )}
    </div>
  );
}
