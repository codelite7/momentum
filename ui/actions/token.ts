"use server";

import { getUser } from "@workos-inc/authkit-nextjs";
import { createRemoteJWKSet } from "jose";
import { WorkOS } from "@workos-inc/node";

const clientId = process.env.WORKOS_CLIENT_ID ?? "";

const workos = new WorkOS(process.env.WORKOS_API_KEY!);

// Set the JWKS URL. This is used to verify if the JWT is still valid
const JWKS = createRemoteJWKSet(
  new URL(workos.userManagement.getJwksUrl(clientId)),
);

export default async function getAccessToken(): Promise<string> {
  const { accessToken } = await getUser({ ensureSignedIn: true });

  return accessToken;
  // const validToken = await verifyAccessToken(accessToken);
  //
  // if (validToken) {
  //   console.log("token is valid");
  //
  //   return accessToken;
  // }
  // const session = await getSessionFromCookie();
  //
  // const response = await workos.userManagement.authenticateWithRefreshToken({
  //   clientId,
  //   refreshToken: session.refreshToken,
  // });
  // const data = {
  //   user: session.user,
  //   impersonator: session.impersonator,
  // };
  //
  // data.accessToken = response.accessToken;
  // data.refreshToken = response.refreshToken;
  // const encryptedSession = await sealData(data, {
  //   password: process.env.WORKOS_COOKIE_PASSWORD ?? "",
  // });
  //
  // // Update the cookie
  // const cookieStore = cookies();
  //
  // cookieStore.set("wos-session", encryptedSession, {
  //   path: "/",
  //   httpOnly: true,
  //   secure: true,
  //   sameSite: "lax",
  // });
  //
  // return response.accessToken;
}

// async function getSessionFromCookie(): Promise<any> {
//   const cookieStore = cookies();
//   const cookie = cookieStore.get("wos-session");
//
//   if (cookie) {
//     return unsealData(cookie, {
//       password: process.env.WORKOS_COOKIE_PASSWORD ?? "",
//     });
//   }
// }
//
// async function verifyAccessToken(accessToken: string) {
//   try {
//     console.log("verifying access token", accessToken);
//     await jwtVerify(accessToken, JWKS);
//
//     return true;
//   } catch (e) {
//     console.warn("Failed to verify session:", e);
//
//     return false;
//   }
// }
