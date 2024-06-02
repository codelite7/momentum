import Thread from "@/components/thread/thread";
import { getThread } from "@/client/thread";

export default async function Page({ params }: { params: { id: string } }) {
  let thread: any = undefined;

  if (params.id != "new") {
    const { data } = await getThread(params.id);

    thread = data.thread;
  }

  return <Thread thread={thread} />;
}
