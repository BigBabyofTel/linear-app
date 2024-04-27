import Image from "next/image";
import { SignInForm } from "./sign-in-form";

// TODO improve login, now it only adds a token to the cookie
export default function SignUpPage() {
  return (
    <main className=" relative grid w-full flex-1 lg:bg-cover lg:bg-center lg:bg-no-repeat">
      <div className="absolute inset-0 z-0">
        <Image src="/background.jpg" alt="Background" layout="fill" objectFit="cover" className="brightness-[0.4]" />
      </div>
      <div className="relative z-10 flex items-center justify-center py-12 sm:p-0">
        <SignInForm />
      </div>
    </main>
  );
}
