"use client";
import { CreateIssue } from "@/components/forms/create-issue";
import { ThemeToggle } from "@/components/theme-toggle";
import { Button } from "@/components/ui/button";
import { PenBoxIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { useState } from "react";

type SidebarProps = {
  authToken: string;
};

const Sidebar: React.FC<SidebarProps> = ({ authToken }) => {
  const router = useRouter();
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div className="flex h-screen w-[23%] min-w-[260px] flex-col items-center gap-2 rounded-2xl border p-3">
      <div className="flex w-[80%] flex-row justify-between">
        <select name="workspace">
          <option value="1">Workspace 1</option>
          <option value="2">Workspace 2</option>
        </select>
        <Button className="px-1 py-2" variant="secondary" onClick={() => setIsOpen(true)}>
          <PenBoxIcon />
        </Button>
        <CreateIssue authToken={authToken} isOpen={isOpen} closeFn={() => setIsOpen(false)} />
      </div>
      <ThemeToggle />
      <section className="-z-30 flex flex-col gap-1">
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
