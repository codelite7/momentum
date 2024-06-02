import { DocumentNode, HttpLink } from "@apollo/client";
import {
  registerApolloClient,
  ApolloClient,
  InMemoryCache,
} from "@apollo/experimental-nextjs-app-support";
import { cookies } from "next/headers";
import { unsealData } from "iron-session";

export const { getClient, PreloadQuery } = registerApolloClient(() => {
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

export async function query(query: DocumentNode, variables?: any) {
  const context = await getContext();

  return await getClient().query({ context, query, variables });
}

export async function mutation(mutation: DocumentNode, variables?: any) {
  const context = await getContext();

  return await getClient().mutate({ context, mutation, variables });
}

async function getContext() {
  const session = await getSessionFromCookie();

  return {
    headers: {
      authorization: `Bearer ${session?.accessToken}`,
    },
  };
}

async function getSessionFromCookie() {
  const cookie = cookies().get("wos-session");

  if (cookie) {
    return unsealData<any>(cookie.value, {
      password: process.env["WORKOS_COOKIE_PASSWORD"]!,
    });
  }
}
