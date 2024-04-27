import { API } from "@/lib/utils";
import { useQuery } from "@tanstack/react-query";

export type User = {
  id: number;
  email: string;
  name: string;
  createdAt: string;
};

export type UserRecord = Record<"user", User>;

async function fetcher(token: string) {
  const res = await API.get<UserRecord>("/user", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.data;
}

export function useSession(token: string) {
  return useQuery({
    queryKey: ["user", token],
    queryFn: () => fetcher(token),
  });
}
