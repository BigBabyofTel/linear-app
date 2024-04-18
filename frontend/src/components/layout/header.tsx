import { useState } from "react";
import { Button } from "../ui/button";
import { ToggleMenuButton } from "./toggle-menu-button";

export function Header() {
  const [isMenuOpen, setIsMenuOpen] = useState<boolean>(false);
  return (
    <div className="flex w-full items-center border-b border-border">
      <header className="container flex items-center justify-between py-3">
        <h2 className="font-sans">Wuhuu </h2>
        <div className="hidden items-center justify-center gap-2 md:flex">
          <Button variant="ghost" size="sm">
            Sign Up
          </Button>
          <Button size="sm">Sign In</Button>
        </div>
        <ToggleMenuButton
          onToggleMenu={() => setIsMenuOpen(!isMenuOpen)}
          toggleMenu={isMenuOpen}
        />
      </header>
    </div>
  );
}
