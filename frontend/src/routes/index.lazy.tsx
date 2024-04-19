import { Button } from "@/components/ui/button";
import { Modal } from "@/components/ui/modal";
import { TypewriterEffect } from "@/components/ui/typewriter-effect";
import { createLazyFileRoute } from "@tanstack/react-router";
import { useState } from "react";

export const Route = createLazyFileRoute("/")({
  component: Index,
});

const words = [
  {
    text: "Build",
  },
  {
    text: "awesome",
  },
  {
    text: "projects",
  },
  {
    text: "with",
  },
  {
    text: "Wuuhuu.",
    className: "text-blue-500 dark:text-blue-500",
  },
];

function Index() {
  const [open, setOpen] = useState(true);
  return (
    <>
      <main className="flex-1">
        <div className="mt-12 flex flex-col items-center justify-center md:h-[20rem] lg:h-[30rem]">
          <p className="mb-4 text-sm text-muted-foreground">Very Gucci</p>
          <TypewriterEffect words={words} />
          <div className="mt-10 flex flex-col space-x-0 space-y-4 md:flex-row md:space-x-4 md:space-y-0">
            <Button variant="ghost">Contact Us</Button>
            <Button onClick={() => setOpen(true)}>Sign Up</Button>
          </div>
        </div>
      </main>

      <Modal onClose={() => setOpen(false)} className="p-12" isOpen={open}>
        Hello World
      </Modal>
    </>
  );
}
