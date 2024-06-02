import "@/styles/globals.css";
import { Metadata, Viewport } from "next";
import clsx from "clsx";
import localFont from "next/font/local";
import { Toaster } from "sonner";
import { getUser as getAuthUser } from "@workos-inc/authkit-nextjs";

import { Providers } from "./providers";

import { siteConfig } from "@/config/site";
import LeftSidebar from "@/components/left-sidebar/left-sidebar";
import AccountSettingsModal from "@/components/account/settings-modal/account-settings-modal";
import { getUser } from "@/client/user";
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
  await getAuthUser({ ensureSignedIn: true });
  const { data } = await getUser();

  console.log("user", data);

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
            <LeftSidebar user={data.user} />
            <div className="h-full w-full">{children}</div>
          </div>
          <AccountSettingsModal />
          {/*<SearchModal />*/}
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
      </body>
    </html>
  );
}
