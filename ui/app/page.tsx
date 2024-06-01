import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { unsealData } from "iron-session";

import { getUser } from "@/client/user";
export default async function Page() {
  const cookie = await getSessionFromCookie();
  const { data } = await getUser();

  console.log("pagey", data);

  if (data.user.threads.edges.length == 1) {
    redirect(`/thread/${data.user.threads.edges[0].node.id}`);
  } else {
  }

  return (
    <>
      <div>{data.user.id}</div>
    </>
  );
}

async function getSessionFromCookie() {
  const cookie = cookies().get("wos-session");

  if (cookie) {
    return unsealData(cookie.value, {
      password: process.env["WORKOS_COOKIE_PASSWORD"]!,
    });
  }
}
