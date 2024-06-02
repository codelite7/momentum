import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";
import { sealData } from "iron-session";
import { WorkOS } from "@workos-inc/node";
const workos = new WorkOS(process.env.WORKOS_API_KEY!);

export async function GET(request: NextRequest) {
  const code = request.nextUrl.searchParams.get("code");
  const state = request.nextUrl.searchParams.get("state");
  const returnPathname = state ? JSON.parse(atob(state)).returnPathname : null;
  const clientId = process.env.WORKOS_CLIENT_ID!;
  const password = process.env.WORKOS_COOKIE_PASSWORD!;

  if (code) {
    try {
      const { user, accessToken, refreshToken, impersonator } =
        await workos.userManagement.authenticateWithCode({
          code,
          clientId,
        });

      if (!accessToken || !refreshToken)
        throw new Error("response is missing tokens");

      // The refreshToken should never be accesible publicly,
      // hence why we encrypt it in the cookie session.
      // Alternatively you could persist the refresh token in a backend database
      const encryptedSession = await sealData(
        { accessToken, refreshToken, user, impersonator },
        { password },
      );

      const redirectUrl = new URL(process.env["WORKOS_REDIRECT_URI"]!);
      const isSecureProtocol = redirectUrl.protocol === "https:";

      // Store the session in a secure cookie
      cookies().set("wos-session", encryptedSession, {
        path: "/",
        httpOnly: true,
        secure: isSecureProtocol,
        sameSite: "lax",
      });
      // store the access token insecurely so it can be used for client side fetching
      cookies().set("accessToken", accessToken);
      const url = request.nextUrl.clone();

      // Cleanup params
      url.searchParams.delete("code");
      url.searchParams.delete("state");

      // Redirect to the requested path and store the session
      url.pathname = returnPathname ?? "/";

      const response = NextResponse.redirect(url);

      return response;
    } catch (error) {
      const errorRes = {
        error: error instanceof Error ? error.message : String(error),
      };

      console.error(errorRes);

      return errorResponse();
    }
  }

  return errorResponse();
}

function errorResponse() {
  return NextResponse.json(
    {
      error: {
        message: "Something went wrong",
        description: "Couldn’t sign in, please try again",
      },
    },
    { status: 500 },
  );
}
