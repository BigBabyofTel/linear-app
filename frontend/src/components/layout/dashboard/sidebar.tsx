"use client";
import { ThemeToggle } from "@/components/theme-toggle";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

const Sidebar = () => {
  const router = useRouter();
  return (
    <div className="flex h-screen w-[23%] flex-col items-center rounded-2xl border-8 p-3 shadow-inner shadow-white">
      <h1 className="p-3 text-3xl">Wuhu</h1>
      <section className="m-2 border-8">
        <ThemeToggle />
      </section>
      <section className="fixed top-[73%] flex h-60 w-[16%] flex-col justify-evenly">
        <Button className="rounded" variant="blurry" onClick={() => router.push("/dashboard/workplace")}>
          Workplace
        </Button>
        <Button className="rounded" variant="blurry" onClick={() => router.push("/dashboard/settings")}>
          Settings
        </Button>
        <Button className="rounded" variant="blurry" onClick={() => router.push("/dashboard/settings/account")}>
          Account
        </Button>
        <Button className="rounded" variant="blurry" onClick={() => router.push("/dashboard/settings/account/profile")}>
          Profile
        </Button>
      </section>
    </div>
  );
};

export default Sidebar;
