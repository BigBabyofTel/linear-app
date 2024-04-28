import Sidebar from "@/components/layout/dashboard/sidebar";
import { Card } from "@/components/ui/card";
import { ReactNode } from "react";

export default function Layout({ children }: { children: ReactNode }) {
  return <section className="my-2 flex h-full w-full flex-row">{children}</section>;
}
