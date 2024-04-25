import { useState } from "react";
import { Button } from "../ui/button";
import { ToggleMenuButton } from "./toggle-menu-button";

export function Header() {
  const [isMenuOpen, setIsMenuOpen] = useState<boolean>(false);
  return (
    <div className="flex w-full items-center border-b border-border">
      <header className="container flex items-center justify-between py-3">
        <h2 className="font-sans">
          <a href="./">Wuhuu</a>
        </h2>
        <div className="hidden items-center justify-center gap-2 md:flex">
          <Button variant="ghost" size="sm">
            <a href="./signup">Sign Up</a>
          </Button>
          <Button size="sm">
            <a href="./signin">Sign In</a>
          </Button>
        </div>
        <ToggleMenuButton onToggleMenu={() => setIsMenuOpen(!isMenuOpen)} toggleMenu={isMenuOpen} />
      </header>
    </div>
  );
}
