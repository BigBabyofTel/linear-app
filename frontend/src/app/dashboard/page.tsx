import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { CircleCheckBigIcon } from "lucide-react";

export default function Page() {
    return (
      <div className="mx-auto flex flex-wrap gap-2 p-3">
          <Button variant="default">Default</Button>
          <Button variant="default" size="icon">
            <CircleCheckBigIcon />
          </Button>
          <Button variant="blurry">Blurry</Button>
          <Button variant="ghost">Ghost</Button>
          <Button variant="secondary">Secondary</Button>
        </div>
    );
}