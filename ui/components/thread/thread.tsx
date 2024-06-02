"use client";

import Messages from "@/components/thread/messages/messages";

type props = {
  thread: any;
};

export default function Thread({ thread }: props) {
  return (
    <div className="w-full h-full flex flex-col place-content-center px-52 pb-4">
      <Messages thread={thread} />
    </div>
  );
}
