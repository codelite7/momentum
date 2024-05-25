import "@/src/styles/globals.css";
import 'primeicons/primeicons.css';
import {Metadata, Viewport} from "next";
import clsx from "clsx";
import {_app} from "./_app";

import {siteConfig} from "@/config/site";
import {fontSatoshi} from "@/config/fonts";
import {getUser} from "@workos-inc/authkit-nextjs";
import LeftSidbar from "@/components/left-sidebar/left-sidbar";

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
    {media: "(prefers-color-scheme: light)", color: "white"},
    {media: "(prefers-color-scheme: dark)", color: "black"},
  ],
};

export default async function RootLayout({
                                           children,
                                         }: {
  children: React.ReactNode;
}) {
  const {user} = await getUser({ensureSignedIn: true});

  return (
    <html suppressHydrationWarning lang="en">
    <head/>
    <body
      className={clsx(
        "bg-background font-satoshi antialiased",
        fontSatoshi.variable,
      )}
    >
    <_app themeProps={{attribute: "class", defaultTheme: "dark"}}>
      <div className="flex border-1 border-sky-500 h-screen w-screen">
        <div className="flex flex-col h-full border-1 border-red-500">
          <LeftSidbar user={user}/>
        </div>
        <div className="w-full border-1 border-yellow-500">
          {children}
        </div>
        <div className="w-60 border-1 border-red-500">
          Right Sidebar
        </div>
      </div>
      {/*<div className="flex w-full h-full">*/}
      {/*  <div className="h-screen border-solid border-1 border-sky-500">*/}
      {/*    <LeftSidebar/>*/}
      {/*  </div>*/}
      {/*  <div className="h-full border-solid border-1 border-sky-500">*/}
      {/*    {children}*/}
      {/*  </div>*/}
      {/*</div>*/}
    </_app>
    </body>
    </html>
  );
}
