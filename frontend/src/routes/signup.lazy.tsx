import { createLazyFileRoute } from "@tanstack/react-router";
import { Fragment } from "react/jsx-runtime";
import { SubmitHandler, useForm } from "react-hook-form";
import * as z from "zod";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

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
    watch,
    formState: { errors },
  } = useForm<SignupSchema>();
  const onSubmit: SubmitHandler<SignupSchema> = (data) => console.log(data);

  //watch input value by passing the name of it
  console.log(watch("name"));

  return (
    <Fragment>
      <form onSubmit={(e) => handleSubmit(onSubmit)(e)}>
        <div className="w-1/3 flex flex-col items-center mx-auto m-5">
          <Input {...register("name")} placeholder="Name" type="text" className="m-5" />
          <Input {...register("username")} placeholder="Username" type="text" className="m-5" />
          <Input {...register("age")} placeholder="Age" type="number" className="m-5" />
          <Input {...register("email")} placeholder="Email" type="email" className="m-5" />
          <Input {...register("password")} placeholder="Password" type="password" className="m-5" />
          <Button variant="default">Sign up</Button>
        </div>
      </form>
    </Fragment>
  );
}
