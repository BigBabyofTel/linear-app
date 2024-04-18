import { Icons } from "../icons";
import { ThemeToggle } from "../theme-toggle";
import { buttonVariants } from "../ui/button";

export function Footer() {
  return (
    <div className="flex w-full items-center border-t border-border">
      <footer className="container flex items-center justify-between py-3">
        <p className="text-sm text-muted-foreground">
          Build by{" "}
          <a className="font-medium hover:underline" href="/">
            FullStack Community
          </a>
        </p>

        <div className="flex items-center justify-center gap-2">
          <ThemeToggle />
          <a
            href="/"
            className={buttonVariants({
              size: "icon",
              variant: "outline",
            })}
          >
            <Icons.Github className="size-5" />
          </a>
        </div>
      </footer>
    </div>
  );
}
