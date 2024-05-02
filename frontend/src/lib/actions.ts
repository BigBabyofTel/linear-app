"use server";
import { createSafeActionClient } from "next-safe-action";

import { z } from "zod";
import { API } from "./utils";
import { cookies } from "next/headers";
const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+{}\[\]:;<>,.?\/~`|\\-]).{8,}$/;

const signInSchema = z.object({
  email: z.string().email("Please enter a valid email"),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters long")
    .regex(
      passwordRegex,
      "Password must have at least one uppercase letter, one lowercase letter, one number, and one special character",
    ),
});

const action = createSafeActionClient();
export const signInAction = action(signInSchema, async ({ email, password }) => {
  const cookieStore = cookies();
  const res = await API.post("/auth/login", { email, password });

  cookieStore.set("token", res.data.accessToken);
  return {
    success: true,
  };
});

export async function changePassword() {
  try {
  const res = await API.patch('/password', {
    headers: {'Content-Type': 'application/json'},
    body: {
      "currentPassword": "string",
    "newPassword": "string"
    }
  })
  } catch (e) {
    console.log(e)
  }
}
