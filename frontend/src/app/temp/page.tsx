// This is place where I show how to fetch stuff

import { UserRecord } from "@/hooks/useSession";
import { API } from "@/lib/utils";
import { cookies } from "next/headers";
import { TempClientFetch } from "./temp-client-fetch";

export default async function TestPage() {
  const cookieStore = cookies();
  const token = cookieStore.get("token");
  const res = await API.get<UserRecord>("/user", {
    headers: {
      Authorization: `Bearer ${token?.value}`,
    },
  });

  console.log(res.data.user);

  return (
    <main className="container grid gap-8 py-12">
      <p className="text-2xl font-medium">Hello, {res.data.user.name} 👋</p>
      <TempClientFetch token={token?.value!} />
    </main>
  );
}
