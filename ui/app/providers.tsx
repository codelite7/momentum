"use client";

import * as React from "react";
import { NextUIProvider } from "@nextui-org/system";
import { useRouter } from "next/navigation";
import { ThemeProvider as NextThemesProvider } from "next-themes";
import { ThemeProviderProps } from "next-themes/dist/types";
import { RelayEnvironmentProvider } from "react-relay";
import { Provider as JotaiProvider } from "jotai";
import {
  ApolloClient,
  ApolloNextAppProvider,
  InMemoryCache,
} from "@apollo/experimental-nextjs-app-support";
import { HttpLink } from "@apollo/client";

import { getCurrentEnvironment } from "@/relay/environment";

export interface ProvidersProps {
  children: React.ReactNode;
  themeProps?: ThemeProviderProps;
}

export function Providers({ children, themeProps }: ProvidersProps) {
  const router = useRouter();
  const environment = getCurrentEnvironment();

  return (
    <NextUIProvider
      className="h-screen w-screen overflow-hidden"
      navigate={router.push}
    >
      <NextThemesProvider {...themeProps}>
        <RelayEnvironmentProvider environment={environment}>
          <JotaiProvider>
            <ApolloNextAppProvider makeClient={makeClient}>
              {children}
            </ApolloNextAppProvider>
          </JotaiProvider>
        </RelayEnvironmentProvider>
      </NextThemesProvider>
    </NextUIProvider>
  );
}

function makeClient() {
  const httpLink = new HttpLink({
    // this needs to be an absolute url, as relative urls cannot be used in SSR
    uri: "https://example.com/api/graphql",
    // you can disable result caching here if you want to
    // (this does not work if you are rendering your page with `export const dynamic = "force-static"`)
    fetchOptions: { cache: "no-store" },
    // you can override the default `fetchOptions` on a per query basis
    // via the `context` property on the options passed as a second argument
    // to an Apollo Client data fetching hook, e.g.:
    // const { data } = useSuspenseQuery(MY_QUERY, { context: { fetchOptions: { cache: "force-cache" }}});
  });

  // use the `ApolloClient` from "@apollo/experimental-nextjs-app-support"
  return new ApolloClient({
    // use the `InMemoryCache` from "@apollo/experimental-nextjs-app-support"
    cache: new InMemoryCache(),
    link: httpLink,
  });
}
