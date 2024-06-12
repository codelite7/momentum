import { ReactNode } from "react";
import { getUser } from "@workos-inc/authkit-nextjs";

import { ApolloWrapper } from "@/app/ApolloWrapper";

export type props = {
  children: ReactNode;
};

export default async function AccessToken({ children }: props) {
  const { accessToken } = await getUser({ ensureSignedIn: true });

  return <ApolloWrapper accessToken={accessToken}>{children}</ApolloWrapper>;
}
