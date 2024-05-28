import { ReactNode } from "react";

import { RightSidebar } from "@/components/right-sidebar/right-sidebar";

export default function ThreadLayout({ children }: { children: ReactNode }) {
  return (
    <div className="h-full flex">
      <div className="h-full w-full">{children}</div>
      <div className="h-full">
        <RightSidebar />
      </div>
    </div>
  );
}
