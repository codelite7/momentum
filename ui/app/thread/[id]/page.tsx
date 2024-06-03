"use client";

import { RightSidebar } from "@/components/right-sidebar/right-sidebar";
import NewThread from "@/app/thread/[id]/new-thread";
import Thread from "@/components/thread/thread";

export default function Page({ params }: { params: { id: string } }) {
  const isNew = params.id == "new";

  return (
    <div className="h-full flex">
      <div
        className={`flex flex-col h-full w-full px-24 place-content-${isNew ? "center" : "end"}  pb-12 transition-all`}
      >
        {isNew ? <NewThread /> : <Thread threadId={params.id} />}
      </div>
      <div className="h-full">
        <RightSidebar />
      </div>
    </div>
  );
}
