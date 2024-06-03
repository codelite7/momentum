"use client";

import * as React from "react";
import { NextUIProvider } from "@nextui-org/system";
import { useRouter } from "next/navigation";
import { ThemeProvider as NextThemesProvider } from "next-themes";
import { ThemeProviderProps } from "next-themes/dist/types";
import { RelayEnvironmentProvider } from "react-relay";
import { Provider as JotaiProvider } from "jotai";

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
          <JotaiProvider>{children}</JotaiProvider>
        </RelayEnvironmentProvider>
      </NextThemesProvider>
    </NextUIProvider>
  );
}
