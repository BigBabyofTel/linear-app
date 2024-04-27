import { Card } from "@/components/ui/card";

export default function Page() {
    return (
      <div className="mx-auto p-3 flex gap-2 flex-wrap">
        <Card className="h-[300px] w-[300px]" variant="default">
          Default
        </Card>
        <Card className="h-[300px] w-[300px]" variant="dashboard">
          Dashboard
        </Card>
        <Card className="h-[300px] w-[300px]" variant="blurred">
          Blurred
        </Card>
        <Card className="h-[300px] w-[300px]" variant="clickable">
          Clickable
        </Card>
      </div>
    );
}