import { RefObject, useEffect } from "react";

export function useClickOutside<T extends HTMLElement>(
  ref: RefObject<T>,
  handler: (e: Event) => void,
  secRef?: RefObject<HTMLElement>,
  disabled = false,
) {
  useEffect(() => {
    if (disabled) return;

    const listener = (e: Event) => {
      const el = ref?.current;
      if (!el || el.contains(e.target as Node) || null) {
        return;
      }

      if (secRef?.current?.contains(e.target as Node) || null) {
        return;
      }

      handler(e);
    };

    document.addEventListener("mousedown", listener);
    document.addEventListener("touchstart", listener);

    return () => {
      document.removeEventListener("mousedown", listener);
      document.removeEventListener("touchstart", listener);
    };
  }, [ref, handler, disabled, secRef]);
}