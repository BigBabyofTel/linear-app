'use client'

import { CreateIssue } from "@/components/forms/create-issue";
import { Button } from "@/components/ui/button";
import { createLazyFileRoute } from "@tanstack/react-router";
import { useState } from "react";

function Test() {
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

export default Test;