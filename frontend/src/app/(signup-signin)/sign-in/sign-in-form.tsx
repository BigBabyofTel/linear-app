"use client";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Field } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { signInAction } from "@/lib/actions";
import { useRouter } from "next/navigation";

export function SignInForm() {
  const router = useRouter();
  return (
    <Card className="mx-auto grid gap-6 rounded-md p-0 shadow-sm sm:w-[30rem] sm:border sm:border-border sm:bg-card md:w-[35rem]">
      <div>
        <div className="grid gap-2 border-b border-border bg-accent/50 py-6 text-left sm:px-8">
          <h2 className="text-2xl font-bold">Sign In</h2>
          <p className="block text-sm text-muted-foreground sm:hidden">Embrace the future of work.</p>
          <p className="hidden text-sm text-muted-foreground sm:block">Embrace the future of work. Sign up now.</p>
        </div>
        <form
          onSubmit={async (e) => {
            e.preventDefault();
            const formData = new FormData(e.currentTarget);
            const email = formData.get("email") as string;
            const password = formData.get("password") as string;
            const res = await signInAction({ email, password });

            if (res.data?.success) {
              router.push("/temp");
            }
          }}
          className="grid gap-2 pb-8 pt-4 text-left sm:px-8"
        >
          <Field label="Email" labelFor="email" error={""}>
            <Input name="email" id="email" type="email" placeholder="Your Email" />
          </Field>
          <Field label="Password" labelFor="password" error={""}>
            <Input name="password" id="password" type="password" placeholder="Create a Password" />
          </Field>
          <Button type="submit" className="mt-4 w-full">
            Sign Up
          </Button>
        </form>
      </div>
    </Card>
  );
}
