"use client"

import DeleteAccount from "@/components/settings/deleteaccount";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { changePassword } from "@/lib/actions";
import { useState } from "react";



export default function Page() {
  const [open, setOpen] = useState(false);

  return (
    <>
      <main className=" align-items-center m-2 flex flex-col p-5">
        <section className="flex flex-col text-left">
          <h1 className="flex justify-start text-left text-3xl">Profile</h1>
          <p className="text-sm text-muted-foreground">Manage your profile</p>
        </section>
        <section className="m-5 flex flex-col items-center justify-center">
          <Card className="bg-card-background m-4 w-[600px]">
            <CardHeader className=" bg-">
              <div className="">
                <h2 className="mb-5 p-2">Profile picture</h2>
              </div>
              <div className="mx-auto flex h-[200px] w-[200px] items-center justify-center rounded-full border-4 border-slate-600 text-center">
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
          <Card className="w-[600px] m-4">
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
          <Card className="w-[600px] m-4 flex flex-col items-center">
            <CardHeader>Delete Account</CardHeader>
            <CardContent>
              <Button onClick={() => setOpen(true)}><DeleteAccount isOpen={open} closeFn={() => setOpen(false)} />Delete</Button>
            </CardContent>
          </Card>
        </section>
      </main>
    </>
  );
}
