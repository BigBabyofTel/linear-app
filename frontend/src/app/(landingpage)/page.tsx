'use client'

import { Button } from "@/components/ui/button";
import { Modal } from "@/components/ui/modal";
import { API } from "@/lib/utils";
import { TypewriterEffect } from "@/components/ui/typewriter-effect";
import { useState } from "react";
import { useQuery } from "@tanstack/react-query";

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

async function fetchNothing() {
  const response = await API.get("/");
  return response.data;
}

const Home = () => {

  const [open, setOpen] = useState(false);
    const { data } = useQuery({
      queryKey: ["br"],
      queryFn: fetchNothing,
    });
  
  return (
    <>
      <main className="flex-1">
        <div className="mt-12 flex flex-col items-center justify-center md:h-[20rem] lg:h-[30rem]">
          <p className="mb-4 text-sm text-muted-foreground">Very Gucci</p>
          <TypewriterEffect words={words} />
          <div className="mt-10 flex flex-col space-x-0 space-y-4 md:flex-row md:space-x-4 md:space-y-0">
            <Button variant="ghost" onClick={() => setOpen(true)}>
              Contact Us
            </Button>
            <Button>
              <a href="./signup">Sign Up</a>
            </Button>
          </div>
        </div>
      </main>

      <Modal onClose={() => setOpen(false)} className="p-12" isOpen={open}>
        {data?.status}
      </Modal>
    </>
  );
};

export default Home;
