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
    <aside className="fixed inset-y-0 left-0 hidden w-72 content-start gap-4 px-6 py-8 md:grid">
      <div className="flex flex-row justify-between">
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
    </aside>
  );
};

export default Sidebar;
