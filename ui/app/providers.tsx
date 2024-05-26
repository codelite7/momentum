"use client";

import * as React from "react";
import {NextUIProvider} from "@nextui-org/system";
import {useRouter} from "next/navigation";
import {ThemeProvider as NextThemesProvider} from "next-themes";
import {ThemeProviderProps} from "next-themes/dist/types";
import {RelayEnvironmentProvider} from "react-relay";
import {getCurrentEnvironment} from "@/relay/environment";

export interface ProvidersProps {
    children: React.ReactNode;
    themeProps?: ThemeProviderProps;
}

export function Providers({children, themeProps}: ProvidersProps) {
    const router = useRouter();
    const environment = getCurrentEnvironment();

    return (
        <NextUIProvider navigate={router.push} className="h-screen w-screen overflow-hidden">
            <NextThemesProvider {...themeProps}>
                <RelayEnvironmentProvider environment={environment}>
                    {children}
                </RelayEnvironmentProvider>
            </NextThemesProvider>
        </NextUIProvider>
    );
}
