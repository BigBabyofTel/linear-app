"use client";
import "../globals.css";
import { Header } from "@/components/layout/header";
import { Footer } from "@/components/layout/footer";


export default function LandingPageLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
      <>
          <Header />
          {children}
          <Footer />
      </>
  );
}
