"use client";

import { RightSidebar } from "@/components/right-sidebar/right-sidebar";
import NewThread from "@/app/thread/[id]/new-thread";
import Thread from "@/components/thread/thread";

export default function Page({
  params,
  searchParams,
}: {
  params: { id: string };
  searchParams: { parentId: string };
}) {
  const isNew = params.id == "new";
  const parentId = searchParams.parentId;

  return (
    <div className="h-full flex">
      <div
        className={`flex flex-col h-full w-full place-content-${isNew ? "center" : "end"}  pb-12 transition-all`}
      >
        {isNew ? (
          <NewThread parentId={parentId} />
        ) : (
          <Thread parentId={parentId} threadId={params.id} />
        )}
      </div>
      <div className="h-full">
        <RightSidebar />
      </div>
    </div>
  );
}
