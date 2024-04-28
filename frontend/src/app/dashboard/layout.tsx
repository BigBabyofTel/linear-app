import Sidebar from "@/components/layout/dashboard/sidebar";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { ReactNode } from "react";


export default function Layout({children}: {children: ReactNode}) {
    return (
      <section className="my-2 flex h-full w-full flex-row">
        <Sidebar/>
        <Card
          variant={"dashboard"}
          className="bg-[url('https://img.freepik.com/free-photo/ultra-detailed-nebula-abstract-wallpaper-4_1562-749.jpg?size=626&ext=jpg&ga=GA1.1.2130352719.1714217947&semt=sph')] w-full bg-no-repeat bg-cover"
        >
{children}
        </Card>
      </section>
    );
}