import Sidebar from "@/components/layout/dashboard/sidebar";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";


export default function Layout({children}: {children: React.ReactNode}) {
    return (
      <section className="my-2 flex h-full w-full flex-row">
        <Sidebar />
        <Card
          variant={"dashboard"}
          className="w-full bg-no-repeat bg-cover"
        >
          {children}
        </Card>
      </section>
    );
}