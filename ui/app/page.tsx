import { redirect } from "next/navigation";

import { getUser } from "@/client/user";
export default async function Page() {
  const { data } = await getUser();

  if (data.user.threads.edges.length == 1) {
    redirect(`/thread/${data.user.threads.edges[0].node.id}`);
  } else {
    redirect(`/thread/new`);
  }

  return <></>;
}
