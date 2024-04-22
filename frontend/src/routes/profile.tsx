import { Icons } from "@/components/icons";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { createFileRoute, Link } from "@tanstack/react-router";
import { useForm, SubmitHandler } from "react-hook-form";
import { z } from "zod";

export const Route = createFileRoute("/profile")({
  component: Profile,
});

const signupSchema = z.object({
  name: z.string().min(5, "Please use a min of 5 characters").trim(),
  username: z.string().min(5, "Please use a min of 5 characters").trim(),
  email: z.string().email(),
});

type SignupSchema = z.infer<typeof signupSchema>;

export function Profile() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignupSchema>({
    resolver: zodResolver(signupSchema),
  });

  const link = (
    <Link to="/signin" className="underline underline-offset-2">
      Sign in
    </Link>
  );
  const onSubmit: SubmitHandler<SignupSchema> = (data) => console.log(data);

  console.log(errors);
  return (
    <main className="grid flex-1 place-items-center">
      <form className="grid rounded-md border border-border shadow-sm" onSubmit={(e) => handleSubmit(onSubmit)(e)}>
        <div className="grid place-items-start border-b border-border bg-accent/30 px-12 py-6">
          <h2 className="text-xl font-semibold">Sign up</h2>
          <p className="text-sm text-muted-foreground">We won't sell your data to Bejing ( trust us )</p>
          <Button className="mt-10 w-full" variant="outline">
            <Icons.Github className="mr-2 size-4" />
            Sign up with Github
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
              value={"BigBabyofTel"}
            />
            {errors.username && <p className="text-sm text-destructive">{errors.username.message}</p>}
          </div>
          <div className="grid gap-1.5">
            <label className="text-sm font-medium">Email</label>
            <Input
              className={cn("md:min-w-[30rem]", errors.email && "ring-2 ring-destructive")}
              {...register("email")}
              placeholder="BigB@bossenterprises.com"
              type="email"
            />
            {errors.email && <p className="text-sm text-destructive">{errors.email.message}</p>}
          </div>
          <div className="grid gap-1.5 content-end">
            <div className="grid gap-1.5">
              <Button onClick={(e) => e.preventDefault()}>Change Password</Button>
            </div>
            <div className="grid gap-1.5">
              <Button onClick={(e) => e.preventDefault()}>Change Email</Button>
            </div>
            <div className="grid gap-1.5">
              <Button onClick={(e) => e.preventDefault()}>Remove connected accounts</Button>
            </div>
          </div>
        </div>
        <div className="border-t border-border px-12 py-4">
          <Button className="w-full" variant="default" size="lg">
            Sign up
          </Button>
          <p className="mt-4 text-center text-muted-foreground">Already a user? {link}</p>
        </div>
      </form>
    </main>
  );
}
