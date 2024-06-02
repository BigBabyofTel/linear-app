import { Icons } from "@/components/icons";
import { API } from "@/lib/utils";
import { cookies } from "next/headers";

type Issue = {
  id: string;
  title: string;
  description: any;
  status: string;
  createdAt: string;
  priority: string;
  dueDate?: string;
};

type Metadata = {
  currentPage: number;
  firstPage: number;
  lastPage: number;
  totalRecords: number;
};

type Response = Record<"issues", Issue[] & Metadata>;
type IssueGroup = Record<string, Issue[]>;

export default async function Page() {
  const cookieStore = cookies();
  let res;

  try {
    res = await API.get<Response>("/issues/user/all", {
      headers: {
        Authorization: `Bearer ${cookieStore.get("token")?.value}`,
      },
    });
  } catch (error) {
    // Errors should be handled properly
    res = { status: error.response?.status || 500, data: { issues: [] } };
  }

  const groupedIssuesByStatus = res.data.issues.reduce<IssueGroup>((acc, issue) => {
    const key = issue.status;
    if (!acc[key]) {
      acc[key] = [];
    }
    acc[key].push(issue);
    return acc;
  }, {} as IssueGroup);

  const issueCountsByStatus = Object.keys(groupedIssuesByStatus).reduce(
    (acc, status) => {
      acc[status] = groupedIssuesByStatus[status].length;
      return acc;
    },
    {} as Record<string, number>,
  );

  const statusOrder = ["In Progress", "Todo", "Backlog"];

  const sortedStatusKeys = Object.keys(groupedIssuesByStatus).sort((a, b) => {
    const indexA = statusOrder.indexOf(a);
    const indexB = statusOrder.indexOf(b);
    if (indexA !== -1 && indexB !== -1) {
      return indexA - indexB;
    }
    if (indexA !== -1) return -1;
    if (indexB !== -1) return 1;
    // Sort alphabetically if not in predefined list
    return a.localeCompare(b);
  });

  return (
    <main className="scroll m-4 grid h-dvh max-h-dvh content-start overflow-y-scroll rounded-md border border-border shadow-sm md:ml-72">
      <section className="h-fit border-b border-border">
        <div className="px-4 py-2 lg:px-8">
          <h3 className="text-sm font-medium">My Issues</h3>
        </div>
      </section>
      <section className="h-fit border-b border-border">
        <div className="flex justify-between px-4 py-2 lg:px-8">
          <h3 className="text-sm font-medium">Filter</h3>
          <h3 className="text-sm font-medium">Display</h3>
        </div>
      </section>

      {sortedStatusKeys.length > 0 ? (
        sortedStatusKeys.map((status) => (
          <section key={status} className="grid">
            <div className="flex w-full items-center justify-start gap-4 bg-accent px-4 py-2 lg:px-8">
              <Icons.Sun className="size-5" />
              <h2 className="text-xs font-medium">
                {status} ({issueCountsByStatus[status]})
              </h2>
            </div>
            {groupedIssuesByStatus[status].map((issue: Issue) => (
              <p key={issue.id} className="flex items-center justify-between border-b px-4 py-2 lg:px-8">
                <div className="flex items-center justify-center gap-5">
                  {/** Priority Icon */}
                  <Icons.Moon className="size-5" />
                  <p className="text-xs text-muted-foreground">LUK-82</p>
                  {/** Status icon */}
                  <Icons.Moon className="size-5" />
                  <p className="font-medium">{issue.title}</p>
                </div>
                <div>
                  {/** Should be issue day */}
                  <p className="text-xs text-muted-foreground">{"-"}</p>
                </div>
              </p>
            ))}
          </section>
        ))
      ) : (
        <section className="grid h-full place-items-center">
          <h2 className="mt-4 text-lg font-medium text-muted-foreground">No issues found</h2>
        </section>
      )}
    </main>
  );
}