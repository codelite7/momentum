"use server";

import { getUser } from "@workos-inc/authkit-nextjs";

export default async function getAccessToken(): Promise<string> {
  const { accessToken } = await getUser({ ensureSignedIn: true });

  return accessToken;
}

// export default async function getAccessToken(): Promise<string> {
//   const cookie = cookies().get("wos-session");
//
//   if (cookie) {
//     const data: any = await unsealData(cookie.value, {
//       password: process.env["WORKOS_COOKIE_PASSWORD"]!,
//     });
//
//     return data.accessToken;
//   }
//
//   return "";
// }
