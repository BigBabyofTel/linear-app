import { Button } from "@/components/ui/button";
import { useTheme } from "next-themes";
import { Icons } from "./icons";

export function ThemeToggle() {
  const { setTheme, theme } = useTheme();

  function toggleTheme() {
    setTheme(theme === "dark" ? "light" : "dark");
  }

  return (
    <Button onClick={toggleTheme} variant="outline" size="icon">
      {theme === "dark" ? <Icons.Sun className="h-5 w-5" /> : <Icons.Moon className="h-5 w-5" />}
    </Button>
  );
}
