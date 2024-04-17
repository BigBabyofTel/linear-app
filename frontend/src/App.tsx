import { ThemeProvider } from './components/theme-provider'
import { ThemeToggle } from './components/theme-toggle'

function App() {

  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <div className="flex flex-col items-center justify-center min-h-screen space-y-4">
        <h1 className="text-4xl font-bold">Hello Vite + React!</h1>
        
        <ThemeToggle />
      </div>
  </ThemeProvider>
  )
}

export default App
