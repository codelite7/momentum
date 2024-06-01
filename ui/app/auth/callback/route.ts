import { getUser, handleAuth } from "@workos-inc/authkit-nextjs";
import { cookies } from "next/headers";
import { NextRequest } from "next/server";

export const GET = async (request: NextRequest) => {
  const webosHandler = handleAuth();

  const next = webosHandler(request);

  const { user } = await getUser();

  cookies().set("user", JSON.stringify(user));

  return next;
};
