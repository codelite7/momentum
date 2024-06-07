"use client";

// you can optionally enhance the `DefaultContext` like this to add some type safety to it.
import { setContext } from "@apollo/client/link/context";
import {
  ApolloClient,
  ApolloNextAppProvider,
  InMemoryCache,
} from "@apollo/experimental-nextjs-app-support";
import {
  createHttpLink,
  from,
  makeVar,
  ReactiveVar,
  useApolloClient,
} from "@apollo/client";
import { onError } from "@apollo/client/link/error";
import { useRouter } from "next/navigation";

import { Thread } from "@/__generated__/graphql";
import getAccessToken from "@/actions/token";

declare module "@apollo/client" {
  export interface DefaultContext {
    token?: string;
  }
}

type props = {
  children: React.ReactNode;
  accessTokenn: string;
};

export function ApolloWrapper({ accessTokenn, children }: props) {
  const router = useRouter();

  function makeClient() {
    const errorLink = onError(
      ({ graphQLErrors, networkError, operation, forward }) => {
        if (graphQLErrors) {
          console.error("[GraphQL error on operation]:", operation);
          graphQLErrors.forEach(({ message, locations, path }) =>
            console.error(
              `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`,
            ),
          );
        } else if (networkError) {
          if (networkError.statusCode == 401) {
            router.refresh();
            operation.setContext(async (request) => {
              const accessToken = await getAccessToken();

              return {
                headers: {
                  authorization: accessToken,
                },
              };
            });

            return forward(operation);
          }
          console.error(`[Network error]: ${networkError}`);
        }
      },
    );
    const httpLink = createHttpLink({
      // this needs to be an absolute url, as relative urls cannot be used in SSR
      uri: "http://localhost:3001/query",
      // you can disable result caching here if you want to
      // (this does not work if you are rendering your page with `export const dynamic = "force-static"`)
      fetchOptions: { cache: "no-store" },
    });
    const authLink = setContext(async (_, { headers, token }) => {
      token = token ? token : accessTokenn;

      return {
        headers: {
          ...headers,
          ...(token ? { authorization: `Bearer ${accessTokenn}` } : {}),
        },
      };
    });

    return new ApolloClient({
      cache: new InMemoryCache({
        typePolicies: {
          Query: {
            fields: {
              thread: {
                read() {
                  return threadVar();
                },
              },
            },
          },
        },
      }),
      link: from([errorLink, authLink, httpLink]),
    });
  }

  return (
    <ApolloNextAppProvider makeClient={makeClient}>
      <UpdateAuth accessTokenn={accessTokenn}>{children}</UpdateAuth>
    </ApolloNextAppProvider>
  );
}

function UpdateAuth({ accessTokenn, children }: props) {
  const apolloClient = useApolloClient();

  // just synchronously update the `apolloClient.defaultContext` before any child component can be rendered
  // so the value is available for any query started in a child
  apolloClient.defaultContext.token = accessTokenn;

  return <>{children}</>;
}

export const threadVar: ReactiveVar<Thread | undefined> = makeVar<
  Thread | undefined
>(undefined);
