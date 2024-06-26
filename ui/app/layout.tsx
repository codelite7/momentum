import "@/styles/globals.css";
import { Metadata, Viewport } from "next";
import clsx from "clsx";
import localFont from "next/font/local";
import { Toaster } from "sonner";
import { getUser } from "@workos-inc/authkit-nextjs";

import { Providers } from "./providers";

import { siteConfig } from "@/config/site";
import AccountSettingsModal from "@/components/account/settings-modal/account-settings-modal";
import LeftSidebar from "@/components/left-sidebar/left-sidebar";
import SearchModal from "@/components/search/search-modal";
import AccessToken from "@/components/common/accessToken";

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
const myFont = localFont({ src: "../config/satoshi/Satoshi-Variable.woff2" });

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { accessToken } = await getUser({ ensureSignedIn: true });

  return (
    <html suppressHydrationWarning lang="en">
      <head />
      <body
        className={clsx(
          "min-h-screen bg-background antialiased overflow-hidden",
          myFont.className,
        )}
      >
        {/* typescript and jsx don't know about nextjs's magic that allows async jsx on server components*/}
        {/* @ts-expect-error Server Component */}
        <AccessToken>
          <Providers themeProps={{ attribute: "class", defaultTheme: "dark" }}>
            <div className="flex h-full w-full overflow-hidden">
              <LeftSidebar />
              <div className="h-full w-full">{children}</div>
            </div>
            <AccountSettingsModal />
            <SearchModal />
            <Toaster
              toastOptions={{
                classNames: {
                  error: "bg-danger border-0",
                  success: "bg-success  border-0",
                  warning: "bg-warning  border-0",
                  info: "bg-primary  border-0",
                },
              }}
            />
          </Providers>
        </AccessToken>
      </body>
    </html>
  );
}
