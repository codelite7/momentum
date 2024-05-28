"use client";

import Thread from "@/components/thread/thread";

export default function Page({ params }: { params: { id: string } }) {
  return <Thread id={params.id} />;
}
