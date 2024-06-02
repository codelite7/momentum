import { createHttpLink, DocumentNode, from } from "@apollo/client";
import {
  registerApolloClient,
  ApolloClient,
  InMemoryCache,
} from "@apollo/experimental-nextjs-app-support";
import { cookies } from "next/headers";
import { setContext } from "@apollo/client/link/context";
import { onError } from "@apollo/client/link/error";

export const { getClient, PreloadQuery } = registerApolloClient(() => {
  const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors)
      graphQLErrors.forEach(({ message, locations, path }) =>
        console.log(
          `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`,
        ),
      );
    if (networkError) console.error(`[Network error]: ${networkError}`);
  });
  const httpLink = createHttpLink({
    // this needs to be an absolute url, as relative urls cannot be used in SSR
    uri: "http://localhost:3001/query",
    // you can disable result caching here if you want to
    // (this does not work if you are rendering your page with `export const dynamic = "force-static"`)
    fetchOptions: { cache: "no-store" },
  });
  const authLink = setContext((_, { headers }) => {
    // get the authentication token from local storage if it exists
    const token = cookies().get("accessToken")?.value;

    // return the headers to the context so httpLink can read them
    return {
      headers: {
        ...headers,
        authorization: token ? `Bearer ${token}` : "",
      },
    };
  });

  return new ApolloClient({
    cache: new InMemoryCache(),
    link: from([errorLink, authLink, httpLink]),
  });
});

export async function query(query: DocumentNode, variables?: any) {
  return await getClient().query({ query, variables });
}

export async function mutation(mutation: DocumentNode, variables?: any) {
  // const context = await getContext();

  return await getClient().mutate({ mutation, variables });
}

// async function getContext() {
//   const session = await getSessionFromCookie();
//
//   return {
//     headers: {
//       authorization: `Bearer ${session?.accessToken}`,
//     },
//   };
// }

// async function getSessionFromCookie() {
//   const cookie = cookies().get("wos-session");
//
//   if (cookie) {
//     return unsealData<any>(cookie.value, {
//       password: process.env["WORKOS_COOKIE_PASSWORD"]!,
//     });
//   }
// }
