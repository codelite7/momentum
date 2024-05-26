"use client"

import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import {RelayEnvironmentProvider} from "react-relay";
import {getCurrentEnvironment} from "@/src/relay/environment";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const environment = getCurrentEnvironment();

  return (
    <html lang="en">
    <RelayEnvironmentProvider environment={environment}>
      <body className={inter.className}>{children}</body>
    </RelayEnvironmentProvider>
    </html>
  );
}
