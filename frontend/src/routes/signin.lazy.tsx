import { Icons } from '@/components/icons'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { cn } from '@/lib/utils'
import { createLazyFileRoute, Link } from '@tanstack/react-router'
import { zodResolver } from '@hookform/resolvers/zod'
import * as z from "zod"
import { useForm, SubmitHandler } from 'react-hook-form'

export const Route = createLazyFileRoute('/signin')({
  component: Signin,
})

const signinSchema = z.object({
  username: z.string().trim(),
  password: z.string()
})

type signinSchema = z.infer<typeof signinSchema>


export function Signin() {
  
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<signinSchema>({
    resolver: zodResolver(signinSchema),
  });
 
  const link = <Link to="/signup" className="underline underline-offset-2">Sign up</Link>
  const forgot = <Link to="/forgot-password" className="underline underline-offset-2">Forgot password ?</Link>
  const onSubmit: SubmitHandler<signinSchema> = (data) => console.log(data);

  return (
    <main className="grid flex-1 place-items-center">
      <form className="grid rounded-md border border-border shadow-sm" onSubmit={(e) => handleSubmit(onSubmit)(e)}>
        <div className="grid place-items-start border-b border-border bg-accent/30 px-12 py-6">
          <h2 className="text-xl font-semibold">Sign in</h2>
          <p className="text-sm text-muted-foreground">Welcome back (blind homies)</p>
          <Button className="mt-10 w-full" variant="outline">
            <Icons.Github className="mr-2 size-4" />
            Sign in with Github
          </Button>
        </div>
        <div className="grid gap-3 px-12 py-6">
          <div className="grid gap-1.5"></div>
          <div className="grid gap-1.5">
            <label className="text-sm font-medium">Username</label>
            <Input
              className={cn("md:min-w-[30rem]", errors.username && "ring-2 ring-destructive")}
              {...register("username")}
              placeholder="BigBabyofTel"
              type="text"
            />
            {errors.username && <p className="text-sm text-destructive">{errors.username.message}</p>}
          </div>
          <div className="grid gap-1.5">
            <label className="text-sm font-medium">Password</label>
            <Input
              className={cn("md:min-w-[30rem]", errors.password && "ring-2 ring-destructive")}
              {...register("password")}
              placeholder="password"
              type="email"
            />
            {errors.password && <p className="text-sm text-destructive">{errors.password.message}</p>}
          </div>
          <p className='text-right'>{forgot}</p>
        </div>
        <div className="border-t border-border px-12 py-4">
          <Button className="w-full" variant="default" size="lg">
            Sign in
          </Button>
          <p className="mt-4 text-center text-muted-foreground">Don't have an account? {link}</p>
        </div>
      </form>
    </main>
  )
}