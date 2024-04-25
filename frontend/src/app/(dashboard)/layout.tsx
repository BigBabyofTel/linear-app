import { Sidebar } from "lucide-react";

export default function Layout({children}: {children: React.ReactNode}) {
    return (
        <section>
            <Sidebar />
            {children}
        </section>
    );
}