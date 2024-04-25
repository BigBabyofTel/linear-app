import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import axios from "axios";
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const API = axios.create({
  baseURL: "https://ip.lukabrx.dev/api/",
});
