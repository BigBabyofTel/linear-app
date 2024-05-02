"use server";
import { createSafeActionClient } from "next-safe-action";

import { z } from "zod";
import { API } from "./utils";
import { cookies } from "next/headers";
import { AxiosError } from "axios";

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

export async function changePassword(formData: FormData, bearerToken: string) {
  try {
    const res = await API.patch("/password", {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${bearerToken}`,
      },
      body: {
        currentPassword: formData.get("currentPassword"),
        newPassword: formData.get("newPassword"),
      },
    });
    if (!res) {
      console.log(AxiosError);
    }
    console.log("Password updated successfully");
  } catch (error) {
    if (error instanceof AxiosError) {
      console.error("Network error:", error.response?.data?.message || error.message);
    } else {
      console.log(error);
    }
  }
}

export async function deleteAccount() {
  try {
    API.delete("/");
  } catch (e) {
    console.log(e);
  }
}
