export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <section className="flex flex-row">
      {children}
    </section>
  );
}
