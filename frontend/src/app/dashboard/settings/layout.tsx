import Sidebar from "@/components/layout/dashboard/sidebar";
import { Card } from "@/components/ui/card";
import { ReactNode } from "react";

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <>
    <div className="flex">
      <Sidebar />
      <Card className="w-full">
        {children}
      </Card>
      </div>
    </>
  );
}
