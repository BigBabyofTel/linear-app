import { cn } from "@/lib/utils";
import { ReactNode, useEffect, useRef } from "react";
import { Button } from "./button";
import { Icons } from "../icons";
import { useClickOutside } from "@/hooks/use-click-outside";

export type ModalProps = {
  children: ReactNode;
  isOpen: boolean;
  onClose: () => void;
  className?: string;
};
export function Modal({ children, isOpen, className, onClose }: ModalProps) {
  const modalRef = useRef<HTMLDivElement | null>(null);

  useClickOutside(modalRef, onClose);

  useEffect(() => {
    function handleKeyDown(event: KeyboardEvent) {
      if (event.key === "Escape") {
        onClose();
      }
    }

    if (isOpen) {
      window.addEventListener("keydown", handleKeyDown);
    }

    return () => {
      window.removeEventListener("keydown", handleKeyDown);
    };
  }, [isOpen, onClose]);

  return (
    isOpen && (
      <div className="fixed inset-0 flex items-center justify-center bg-background/60 backdrop-blur-sm">
        <div ref={modalRef} className={cn("relative rounded-md border border-border bg-card shadow-lg", className)}>
          <Button className="absolute right-2 top-2" variant="ghost" size="iconSm" onClick={onClose}>
            <Icons.Close className="size-4" />
          </Button>
          {children}
        </div>
      </div>
    )
  );
}
