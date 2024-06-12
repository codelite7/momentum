"use server";

import { cookies } from "next/headers";

export async function exampleAction() {
  // Set cookie
  cookies().set("name", "Delba");
}
