import Image from "next/image";
import { SignUpForm } from "./sign-up-form";

export default function SignInPage() {
  return (
    <main className=" relative grid w-full flex-1 lg:bg-cover lg:bg-center lg:bg-no-repeat">
      <div className="absolute inset-0 z-0">
        <Image src="/background.jpg" alt="Background" layout="fill" objectFit="cover" className="brightness-[0.4]" />
      </div>
      <div className="relative z-10 flex items-center justify-center py-12 sm:p-0">
        <SignUpForm />
      </div>
    </main>
  );
}
