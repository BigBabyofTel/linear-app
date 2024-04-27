"use client";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Field } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { API } from "@/lib/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";
import type { GoError } from "@/types/errors.types";
const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+{}\[\]:;<>,.?\/~`|\\-]).{8,}$/;

const signUpSchema = z
  .object({
    name: z.string().min(2, "Please enter a valid name"),
    email: z.string().email("Please enter a valid email"),
    password: z
      .string()
      .min(8, "Password must be at least 8 characters long")
      .regex(
        passwordRegex,
        "Password must have at least one uppercase letter, one lowercase letter, one number, and one special character",
      ),
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });

type SignUpFormData = z.infer<typeof signUpSchema>;
type SignUpPayload = Omit<SignUpFormData, "confirmPassword">;

export function SignUpForm() {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
    setError,
  } = useForm<SignUpFormData>({
    resolver: zodResolver(signUpSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: (data: SignUpPayload) => API.post("/auth/register", data),
    onSuccess: () => {
      toast.success("Account created successfully, please check your email to verify your account.");
      reset();
      void router.push("/sign-in");
    },
    onError: (e) => {
      const error = e as GoError;

      if (error.response?.status === 422) {
        setError("email", { message: "Email is already taken" });
        toast.error("Email is already taken");
        return;
      }

      toast.error("An error occurred, please try again later.");
    },
  });

  function onSubmit(data: SignUpFormData) {
    const payload: SignUpPayload = {
      name: data.name,
      email: data.email,
      password: data.password,
    };

    mutate(payload);
  }

  return (
    <Card className="mx-auto grid gap-6 rounded-md p-0 shadow-sm sm:w-[30rem] sm:border sm:border-border sm:bg-card md:w-[35rem]">
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid gap-2 border-b border-border bg-accent/50 py-6 text-left sm:px-8">
          <h2 className="text-2xl font-bold">Sign Up</h2>
          <p className="block text-sm text-muted-foreground sm:hidden">Embrace the future of work.</p>
          <p className="hidden text-sm text-muted-foreground sm:block">Embrace the future of work. Sign up now.</p>
        </div>
        <div className="grid gap-2 pb-8 pt-4 text-left sm:px-8">
          <Field label="Name" labelFor="name" error={errors.name?.message}>
            <Input id="name" {...register("name")} type="text" placeholder="Your Name" />
          </Field>
          <Field label="Email" labelFor="email" error={errors.email?.message}>
            <Input id="email" {...register("email")} type="email" placeholder="Your Email" />
          </Field>
          <Field label="Password" labelFor="password" error={errors.password?.message}>
            <Input id="password" {...register("password")} type="password" placeholder="Create a Password" />
          </Field>
          <Field label="Confirm Password" labelFor="confirmPassword" error={errors.confirmPassword?.message}>
            <Input
              id="confirmPassword"
              {...register("confirmPassword")}
              type="password"
              placeholder="Confirm Your Password"
            />
          </Field>
          <Button disabled={isPending} isLoading={isPending} type="submit" className="mt-4 w-full">
            Sign Up
          </Button>
        </div>
      </form>
    </Card>
  );
}
