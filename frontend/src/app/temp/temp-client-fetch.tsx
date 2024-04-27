"use client";

import { useSession } from "@/hooks/useSession";

type TempClientFetchProps = {
  token: string;
};
export function TempClientFetch({ token }: TempClientFetchProps) {
  const { data, isLoading } = useSession(token);

  return <div>{isLoading ? <p>Loading...</p> : <p>Hello, {data?.user.name} 👋</p>}</div>;
}
