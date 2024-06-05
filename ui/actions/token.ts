"use server";

import { getUser } from "@workos-inc/authkit-nextjs";

export default async function getAccessToken(): Promise<string> {
  const { accessToken } = await getUser({ ensureSignedIn: true });

  return accessToken;
}
