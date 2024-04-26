import Sidebar from "@/components/layout/dashboard/sidebar";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";


export default function Layout({children}: {children: React.ReactNode}) {
    return (
      <section className="my-2 flex h-full flex-row">
        <Sidebar />
            <Card variant={"dashboard"}>{children}</Card>
      </section>
    );
}