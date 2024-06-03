"use client";

// you can optionally enhance the `DefaultContext` like this to add some type safety to it.
import { setContext } from "@apollo/client/link/context";
import {
  ApolloClient,
  ApolloNextAppProvider,
  InMemoryCache,
} from "@apollo/experimental-nextjs-app-support";
import { createHttpLink, from, useApolloClient } from "@apollo/client";
import { onError } from "@apollo/client/link/error";

declare module "@apollo/client" {
  export interface DefaultContext {
    token?: string;
  }
}

type props = {
  children: React.ReactNode;
  accessToken: string;
};

export function ApolloWrapper({ accessToken, children }: props) {
  function makeClient() {
    const errorLink = onError(({ graphQLErrors, networkError, operation }) => {
      if (graphQLErrors) {
        console.error("[GraphQL error on operation]:", operation);
        graphQLErrors.forEach(({ message, locations, path }) =>
          console.error(
            `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`,
          ),
        );
        if (networkError) console.error(`[Network error]: ${networkError}`);
      }
    });
    const httpLink = createHttpLink({
      // this needs to be an absolute url, as relative urls cannot be used in SSR
      uri: "http://localhost:3001/query",
      // you can disable result caching here if you want to
      // (this does not work if you are rendering your page with `export const dynamic = "force-static"`)
      fetchOptions: { cache: "no-store" },
    });
    const authLink = setContext(async (_, { headers, token }) => {
      token = token ? token : accessToken;
      console.log("authLink access token ", token);

      return {
        headers: {
          ...headers,
          ...(token ? { authorization: `Bearer ${token}` } : {}),
        },
      };
    });

    return new ApolloClient({
      cache: new InMemoryCache(),
      link: from([errorLink, authLink, httpLink]),
    });
  }

  return (
    <ApolloNextAppProvider makeClient={makeClient}>
      <UpdateAuth accessToken={accessToken}>{children}</UpdateAuth>
    </ApolloNextAppProvider>
  );
}

function UpdateAuth({ accessToken, children }: props) {
  const apolloClient = useApolloClient();

  // just synchronously update the `apolloClient.defaultContext` before any child component can be rendered
  // so the value is available for any query started in a child
  apolloClient.defaultContext.token = accessToken;

  return <>{children}</>;
}
