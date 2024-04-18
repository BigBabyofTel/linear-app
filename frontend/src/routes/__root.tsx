import { ThemeProvider } from '@/components/theme-provider'
import { createRootRoute, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

export const Route = createRootRoute({
  component: () => (
    <>
      <ThemeProvider defaultTheme='system' storageKey="vite-ui-theme">
      <Outlet />
      <TanStackRouterDevtools />
      </ThemeProvider>
    </>
  ),
})