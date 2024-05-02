"use client";
import { CreateIssue } from "@/components/forms/create-issue";
import { ThemeToggle } from "@/components/theme-toggle";
import { Button } from "@/components/ui/button";
import { PenBoxIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { useState } from "react";

const Sidebar = () => {
  const router = useRouter();
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div className="flex h-screen w-[23%] flex-col items-center gap-2 rounded-2xl border p-3 min-w-[260px]">
      <div className="flex flex-row justify-between w-[80%]">
          <select name="workspace">
          <option value="1">Workspace 1</option>
          <option value="2">Workspace 2</option>
        </select>
        <Button className="py-2 px-1" variant="secondary" onClick={() => setIsOpen(true)}><PenBoxIcon /></Button>
        <CreateIssue isOpen={isOpen} closeFn={() => setIsOpen(false)} />
      </div>
      <ThemeToggle />
      <section className="flex flex-col gap-1 -z-30">
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
