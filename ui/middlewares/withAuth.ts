// middleware/withLogging.ts
import { NextFetchEvent, NextRequest } from "next/server";
import { authkitMiddleware } from "@workos-inc/authkit-nextjs";

import { MiddlewareFactory } from "./types";
export const withAuth: MiddlewareFactory = (next) => {
  return async (request: NextRequest, _next: NextFetchEvent) => {
    const next = authkitMiddleware();

    return next(request, _next);
  };
};
