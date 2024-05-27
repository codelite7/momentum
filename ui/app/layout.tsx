import "@/styles/globals.css";
import { Metadata, Viewport } from "next";
import clsx from "clsx";

import { Providers } from "./providers";

import { siteConfig } from "@/config/site";
import localFont from "next/font/local";
import LeftSidebar from "@/components/left-sidebar/left-sidebar";
import {RightSidebar} from "@/components/right-sidebar/right-sidebar";

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s - ${siteConfig.name}`,
  },
  description: siteConfig.description,
  icons: {
    icon: "/favicon.ico",
  },
};

export const viewport: Viewport = {
  themeColor: [
    { media: "(prefers-color-scheme: light)", color: "white" },
    { media: "(prefers-color-scheme: dark)", color: "black" },
  ],
};
const myFont = localFont({src: '../config/satoshi/Satoshi-Variable.woff2'})

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html suppressHydrationWarning lang="en">
      <head />
      <body
        className={clsx(
          "min-h-screen bg-background antialiased overflow-hidden",
            myFont.className,
        )}
      >
        <Providers themeProps={{ attribute: "class", defaultTheme: "dark" }}>
          <div className="flex h-full w-full overflow-hidden">
            <div className="h-full"><LeftSidebar /></div>
            <div className="h-full w-full">{children}</div>
            <div className="h-full"><RightSidebar /></div>
          </div>
        </Providers>
      </body>
    </html>
  );
}
