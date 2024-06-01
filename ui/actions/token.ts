"use server";

import { cookies } from "next/headers";
import { unsealData } from "iron-session";

export default async function getAccessToken(): Promise<string> {
  const cookie = cookies().get("wos-session");

  if (cookie) {
    const data: any = await unsealData(cookie.value, {
      password: process.env["WORKOS_COOKIE_PASSWORD"]!,
    });

    return data.accessToken;
  }

  return "";
}
