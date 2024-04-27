import { type ReactNode } from "react";

type FieldProps = {
  labelFor: string;
  label: string;
  error?: string;
  children: ReactNode;
};
export function Field({ labelFor, label, error, children }: FieldProps) {
  return (
    <div className="grid gap-2">
      <label className="text-sm" htmlFor={labelFor}>
        {label}
      </label>
      {children}
      {error && <p className="text-sm text-destructive">{error}</p>}
    </div>
  );
}
