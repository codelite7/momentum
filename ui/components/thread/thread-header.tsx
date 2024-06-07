import { useAtomValue } from "jotai/index";
import { Link } from "@nextui-org/link";
import { Button } from "@nextui-org/react";

import { parentThreadAtom, threadAtom } from "@/state/atoms";

export default function ThreadHeader() {
  const parentThread = useAtomValue(parentThreadAtom);
  const thread = useAtomValue(threadAtom);

  return (
    // container
    <div className="w-full flex items-center justify-between p-2">
      {/* name / links*/}
      <div>
        {/* parent thread link */}
        {parentThread && (
          <>
            <Link color="primary" href={`/thread/${parentThread.id}`}>
              {parentThread.name}
            </Link>
            &nbsp;/&nbsp;
          </>
        )}
        <span>{thread?.name}</span>
      </div>
      {/* back button */}
      {parentThread && (
        <Button
          as={Link}
          href={`/thread/${parentThread.id}`}
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
