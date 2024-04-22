import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { cn } from '@/lib/utils';
import { zodResolver } from '@hookform/resolvers/zod';
import { createFileRoute, Link } from '@tanstack/react-router'
import { useForm, SubmitHandler } from 'react-hook-form';
import { z } from 'zod';

export const Route = createFileRoute('/forgot-password')({
  component: ForgotPassword,
})


const signinSchema = z.object({
    email: z.string().email()
  })
  
  type signinSchema = z.infer<typeof signinSchema>
  
  
  export function ForgotPassword() {
    
    const {
      register,
      handleSubmit,
      formState: { errors },
    } = useForm<signinSchema>({
      resolver: zodResolver(signinSchema),
    });
   
    const link = <Link to="/signup" className="underline underline-offset-2">Sign up</Link>
    const onSubmit: SubmitHandler<signinSchema> = (data) => console.log(data);
  
    return (
      <main className="grid flex-1 place-items-center">
        <form className="grid rounded-md border border-border shadow-sm" onSubmit={(e) => handleSubmit(onSubmit)(e)}>
          <div className="grid place-items-start border-b border-border bg-accent/30 px-12 py-6">
            <h2 className="text-xl font-semibold">Forgot Password</h2>
            <p className="text-sm text-muted-foreground">We will try our best...</p>
          </div>
          <div className="grid gap-3 px-12 py-6">
            <div className="grid gap-1.5"></div>
            <div className="grid gap-1.5">
                {errors.email && <p className="text-sm text-destructive">There is a possibility that you don't have an account...</p> }
              <label className="text-sm font-medium">Email</label>
              <Input
                className={cn("md:min-w-[30rem]", errors.email && "ring-2 ring-destructive")}
                {...register("email")}
                placeholder="bigb@bossenterprises@boss.com"
                type="text"
              />
              {errors.email && <p className="text-sm text-destructive">{errors.email.message}</p>}
              {}
            </div>
          </div>
          <div className="border-t border-border px-12 py-4">
            <Button className="w-full" variant="default" size="lg">
              Reset Password
            </Button>
            <p className="mt-4 text-center text-muted-foreground">Don't have an account? {link}</p>
          </div>
        </form>
      </main>
    )
  }