import { Footer } from "@/components/layout/footer";
import { Header } from "@/components/layout/header";
import { ThemeProvider } from "@/components/theme-provider";
import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { Toaster } from "sonner";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

export const Route = createRootRoute({
  component: () => (
    <>
      <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
        <QueryClientProvider client={queryClient}>
          <Header />
          <Outlet />
          <Footer />
          <Toaster richColors />
        </QueryClientProvider>
      </ThemeProvider>

      <TanStackRouterDevtools />
    </>
  ),
});
