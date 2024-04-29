import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "@/lib/utils";
import { ButtonHTMLAttributes, forwardRef } from "react";
import { Icons } from "../icons";

const buttonVariants = cva(
  "inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background border transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
  {
    variants: {
      variant: {
        default: "bg-primary hover:bg-primary/80 text-primary-foreground border-black/10 hover:border-black/40 dark:border-white/10 dark:hover:border-white/20",
        destructive: "bg-destructive border-black/30 hover:border-black/60 dark:border-white/20 dark:hover:border-white-40 text-destructive-foreground hover:bg-destructive/90",
        outline: "border border-foreground dark:border-foreground/40 bg-background hover:bg-accent hover:text-accent-foreground",
        secondary:
          "bg-secondary text-secondary-foreground border-black/20 hover:border-black/40 dark:border-white/10 dark:hover:border-white/20",
        link: "text-black dark:text-white border-transparent underline-offset-4 hover:underline",
        blurry: "bg-white/10 dark:border-white/10 backdrop-blur-[2px] text-foreground border-black/10 hover:border-black/20 dark:hover:border-white/20",
      },
      size: {
        default: "rounded-full h-8 px-4 py-2",
        sm: "h-9 rounded-full px-3",
        lg: "h-11 rounded-full px-8",
        icon: "h-10 w-10 rounded-full shadow-md p-2",
        iconSm: "rounded-full size-8 p-0",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  },
);

export type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> &
  VariantProps<typeof buttonVariants> & {
    isLoading?: boolean;
  };

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, children, isLoading, ...props }, ref) => {
    return (
      <button className={cn(buttonVariants({ variant, size, className }))} ref={ref} {...props}>
        {isLoading && <Icons.Spinner className="size-4 animate-spin" />}
        {children}
      </button>
    );
  },
);
Button.displayName = "Button";

export { Button, buttonVariants };