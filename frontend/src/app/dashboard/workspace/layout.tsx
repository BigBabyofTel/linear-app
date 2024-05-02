import Sidebar from "@/components/layout/dashboard/sidebar";

export default function Layout({children}: {children: React.ReactNode}) {
    return (
        <section className="flex gap-1">
            <Sidebar />
            {children}
        </section>
    );
}