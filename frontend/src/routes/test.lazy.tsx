import { CreateIssue } from "@/components/forms/create-issue";
import { Button } from "@/components/ui/button";
import { createLazyFileRoute } from "@tanstack/react-router";
import { useState } from "react";

export const Route = createLazyFileRoute("/test")({
  component: Index,
});

function Index() {
  const [open, setOpen] = useState(false);
  return (
    <>
      <main className="container my-10 flex-1">
        <Button onClick={() => setOpen(true)}>Open Modal</Button>
      </main>

      <CreateIssue isOpen={open} closeFn={() => setOpen(false)} />
    </>
  );
}
