import { cn } from "@/lib/utils";
import { VariantProps, cva } from "class-variance-authority";
import { InputHTMLAttributes, forwardRef } from "react";

const inputVariants = cva("", {
  variants: {
    variant: {
      default:
        "flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-0 disabled:cursor-not-allowed disabled:opacity-50",
      empty: "appearance-none bg-transparent border-none w-full  mr-3 py-1 px-2 focus:outline-none",
    },
  },
  defaultVariants: {
    variant: "default",
  },
});

export type InputProps = InputHTMLAttributes<HTMLInputElement> &
  VariantProps<typeof inputVariants> & {
    isLoading?: boolean;
  };
const Input = forwardRef<HTMLInputElement, InputProps>(({ className, variant, type, ...props }, ref) => {
  return <input type={type} className={cn(inputVariants({ variant, className }))} ref={ref} {...props} />;
});
Input.displayName = "Input";

export { Input };
