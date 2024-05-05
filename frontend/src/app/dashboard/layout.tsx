import Sidebar from "@/components/layout/dashboard/sidebar";
import { cookies } from "next/headers";
export default function Layout({ children }: { children: React.ReactNode }) {
  const cookieStore = cookies();
  const authToken = cookieStore.get("token")?.value;
  return (
    <>
      <Sidebar authToken={authToken!} />
      {children}
    </>
  );
}
