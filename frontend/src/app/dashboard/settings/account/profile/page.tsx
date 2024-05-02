import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { changePassword } from "@/lib/actions";
import { Avatar, AvatarIcon } from "@nextui-org/avatar";
import { CalendarHeart } from "lucide-react";
import { Form } from "react-hook-form";

export default function Page() {
  return (
    <>
      <main className=" align-items-center m-2 flex flex-col p-5">
        <section className="flex flex-col text-left">
          <h1 className="flex justify-start text-left text-3xl">Profile</h1>
          <p className="text-sm text-muted-foreground">Manage your profile</p>
        </section>
        <section className="m-5 flex justify-center flex-col items-center">
          <Card className="w-[600px] bg-card-background m-4">
            <CardHeader className=" bg-">
              <div className="">
                <h2 className="p-2 mb-5">Profile picture</h2>
              </div>
              <div className="flex h-[200px] w-[200px] items-center justify-center rounded-full border-4 border-slate-600 text-center mx-auto" >
                <p className="text-5xl">T.A.B.</p>
              </div>
            </CardHeader>
            <CardContent className="inline-flex w-full flex-col text-left">
              <section className="m-2 flex justify-between">
                <label htmlFor="name">Name</label>
                <input id="name" value={"T. Augustus Baker"} />
              </section>
              <hr className="border-slate-300" />
              <section className="m-2 flex justify-between">
                <label htmlFor="">Username</label>
                <input id="username" value={"BigBabyofTel"} />
              </section>
              <hr className="border-slate-300" />
              <section className="m-2 flex justify-between">
                <label htmlFor="email">Email</label>
                <input id="email" value={"bigbaby@boss.com"} />
              </section>
            </CardContent>
          </Card>
          <Card className="w-[600px]">
            <CardHeader className="text-center">Change Password</CardHeader>
            <CardContent>
              <form action={changePassword}>
              <section className="flex justify-between p-5">
              <label htmlFor="current password">Current Password</label>
              <input type="password" name="currentPassword" />
              </section>
              <section className="flex justify-between p-5">
              <label htmlFor="new password">New Password</label>
              <input type="password" name="newPassword" />
              </section>
              <section className="flex justify-center p-5">
              <Button>Change</Button>
              </section>
              </form>
            </CardContent>
          </Card>
        </section>
      </main>
    </>
  );
}
