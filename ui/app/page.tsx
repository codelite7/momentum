import { gql } from "@apollo/client";
import { redirect } from "next/navigation";
import { getUser } from "@workos-inc/authkit-nextjs";
import { cookies } from "next/headers";
import { unsealData } from "iron-session";

import { getClient } from "@/client/ApolloClient";
export default async function Page() {
  const cookie = await getSessionFromCookie();

  console.log("cookie", cookie);
  const { user } = await getUser();

  const vars = { email: user.email };

  let userId = global.cacheUser.get(user!.email);

  if (userId) {
    console.log("cached user id");
  }
  // console.log("cookies", cookies().getAll());
  const { data } = await getClient().query({
    query: gql`
      query homeUserQuery {
        user {
          id
          email
          threads(
            first: 1
            orderBy: { field: LAST_VIEWED_AT, direction: DESC }
          ) {
            edges {
              node {
                id
              }
            }
          }
        }
      }
    `,
    context: {
      headers: {
        authorization: `Bearer ${cookie.accessToken}`,
      },
    },
  });

  console.log("page", data);
  global.cacheUser.set(user!.email, data.user.id);

  if (data.user.threads.edges.length == 1) {
    redirect(`/thread/${data.user.threads.edges[0].node.id}`);
  } else {
    // const { data } = await getClient().mutate({
    //   mutation: gql`
    //     mutation CreateThread($userId: ID!) {
    //       createThread(input: { createdByID: $userId, name: "New Thread" }) {
    //         id
    //       }
    //     }
    //   `,
    // });
  }

  return <></>;
}

async function getSessionFromCookie() {
  const cookie = cookies().get("wos-session");

  if (cookie) {
    return unsealData(cookie.value, {
      password: process.env["WORKOS_COOKIE_PASSWORD"]!,
    });
  }
}
