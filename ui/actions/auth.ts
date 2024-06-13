"use server";

import { signOut as workosSignOut } from "@workos-inc/authkit-nextjs";
import { createRemoteJWKSet, jwtVerify } from "jose";
import { WorkOS } from "@workos-inc/node";
const clientId = process.env.WORKOS_CLIENT_ID ?? "";
const workos = new WorkOS(process.env.WORKOS_API_KEY ?? "");
const JWKS = createRemoteJWKSet(
  new URL(workos.userManagement.getJwksUrl(clientId)),
);

export async function signOut() {
  await workosSignOut();
}

export async function verifyAccessToken(accessToken: string) {
  try {
    await jwtVerify(accessToken, JWKS);

    return true;
  } catch (e) {
    return false;
  }
}
