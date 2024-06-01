"use server";

import { cookies } from "next/headers";
import { gql } from "@apollo/client";
import { getUser } from "@workos-inc/authkit-nextjs";

import { getClient } from "@/client/ApolloClient";

export async function setUserCookie() {
  // cookies().set("name", "lee");
  console.log("set name cookie");
}

async function ensureUserCookie() {
  const { user } = await getUser({ ensureSignedIn: true });
  const cookieStore = cookies();

  if (!cookieStore.get("user")) {
    const { data } = await getClient().query({
      query: gql`
        query user($email: String!) {
          user(email: $email) {
            id
            email
          }
        }
      `,
      variables: { email: user.email },
    });

    if (data.user) {
      cookieStore.set("user", data.user);
    }
  }
}

export { ensureUserCookie };
