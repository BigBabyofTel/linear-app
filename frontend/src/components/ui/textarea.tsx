import { cn } from "@/lib/utils";
import { VariantProps, cva } from "class-variance-authority";
import { TextareaHTMLAttributes, forwardRef } from "react";

const textAreaVariants = cva("", {
  variants: {
    variant: {
      default:
        "flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
      empty: "appearance-none bg-transparent border-none w-full  mr-3 py-1 px-2 focus:outline-none resize-none",
    },
  },
  defaultVariants: {
    variant: "default",
  },
});

export type InputProps = TextareaHTMLAttributes<HTMLTextAreaElement> &
  VariantProps<typeof textAreaVariants> & {
    isLoading?: boolean;
  };
const Textarea = forwardRef<HTMLTextAreaElement, InputProps>(({ className, variant, ...props }, ref) => {
  return <textarea className={cn(textAreaVariants({ variant, className }))} ref={ref} {...props} />;
});
Textarea.displayName = "Textarea";

export { Textarea };
