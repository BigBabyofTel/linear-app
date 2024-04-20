import { createLazyFileRoute } from "@tanstack/react-router";
import { SubmitHandler, useForm } from "react-hook-form";
import * as z from "zod";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { zodResolver } from "@hookform/resolvers/zod";
import { cn } from "@/lib/utils";

export const Route = createLazyFileRoute("/signup")({
  component: Signup,
});

const signupSchema = z.object({
  name: z.string().min(5, "Please use a min of 5 characters").trim(),
  username: z.string().min(5, "Please use a min of 5 characters").trim(),
  age: z.number().min(18, "You must be 18 or older to sign up"),
  email: z.string().email(),
  password: z.string().min(5, "Please enter a password that is 5 characters or more"),
});

type SignupSchema = z.infer<typeof signupSchema>;

export function Signup() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignupSchema>({
    resolver: zodResolver(signupSchema),
  });

  const onSubmit: SubmitHandler<SignupSchema> = (data) => console.log(data);

  console.log(errors);
  return (
    <main className="grid flex-1 place-items-center">
      <form className="grid rounded-md border border-border shadow-sm" onSubmit={(e) => handleSubmit(onSubmit)(e)}>
        <div className="grid place-items-start border-b border-border bg-accent/30 px-12 py-6">
          <h2 className="text-xl font-semibold">Sign up</h2>
          <p className="text-sm text-muted-foreground">We won't sell your data to Bejing ( trust us )</p>
        </div>
        <div className="grid gap-3 px-12 py-6">
          <div className="grid gap-1.5">
            <label className="text-sm font-medium">Name</label>
            <Input
              className={cn("md:min-w-[30rem]", errors.name && "ring-2 ring-destructive")}
              {...register("name")}
              placeholder="Name"
              type="text"
            />
            {errors.name && <p className="text-sm text-destructive">{errors.name.message}</p>}
          </div>
          <Input className="md:min-w-[30rem]" {...register("username")} placeholder="Username" type="text" />
          <Input className="md:min-w-[30rem]" {...register("age")} placeholder="Age" type="number" />
          <Input className="md:min-w-[30rem]" {...register("email")} placeholder="Email" type="email" />
          <Input className="md:min-w-[30rem]" {...register("password")} placeholder="Password" type="password" />
        </div>
        <div className="flex justify-end border-t border-border px-12 py-4">
          <Button variant="default">Sign up</Button>
        </div>
      </form>
    </main>
  );
}
