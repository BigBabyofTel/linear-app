import { cn } from "@/lib/utils";

type ToggleMenuButtonProps = {
  onToggleMenu: () => void;
  toggleMenu: boolean;
};

export function ToggleMenuButton({
  onToggleMenu,
  toggleMenu,
}: ToggleMenuButtonProps) {
  return (
    <div
      onClick={onToggleMenu}
      className="relative ml-4 h-7 w-5 cursor-pointer md:hidden"
    >
      <span
        className={cn(
          "absolute h-[1.8px] w-full origin-center rounded-full bg-slate-800 transition-transform duration-500 ease-in-out dark:bg-slate-100",
          toggleMenu ? "top-1/2 rotate-45" : "top-1/4 -rotate-0",
        )}
      />
      <span
        className={cn(
          "absolute h-[1.8px] w-full origin-center rounded-full bg-slate-800 transition-transform duration-300 ease-in-out dark:bg-slate-100",
          toggleMenu ? "top-1/2 opacity-0 duration-500" : "top-1/2 opacity-100",
        )}
      />
      <span
        className={cn(
          "absolute h-[1.8px] w-full origin-center rounded-full bg-slate-800 transition-transform duration-500 ease-in-out dark:bg-slate-100",
          toggleMenu ? "top-1/2 -rotate-45" : "top-3/4 -rotate-0",
        )}
      />
    </div>
  );
}
