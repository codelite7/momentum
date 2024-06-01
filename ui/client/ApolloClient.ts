import { DocumentNode, HttpLink } from "@apollo/client";
import {
  registerApolloClient,
  ApolloClient,
  InMemoryCache,
} from "@apollo/experimental-nextjs-app-support";
import { cookies } from "next/headers";
import { unsealData } from "iron-session";

export const { getClient, query, PreloadQuery } = registerApolloClient(() => {
  return new ApolloClient({
    cache: new InMemoryCache(),
    link: new HttpLink({
      // this needs to be an absolute url, as relative urls cannot be used in SSR
      uri: "http://localhost:3001/query",
      // you can disable result caching here if you want to
      // (this does not work if you are rendering your page with `export const dynamic = "force-static"`)
      fetchOptions: { cache: "no-store" },
      credentials: "include",
    }),
  });
});

export async function makeRequest(query: DocumentNode, variables?: any) {
  const cookie = await getSessionFromCookie();
  const context = {
    headers: {
      authorization: `Bearer ${cookie.accessToken}`,
    },
  };

  return await getClient().query({ context, query, variables });
}

async function getSessionFromCookie() {
  const cookie = cookies().get("wos-session");

  if (cookie) {
    return unsealData(cookie.value, {
      password: process.env["WORKOS_COOKIE_PASSWORD"]!,
    });
  }
}
