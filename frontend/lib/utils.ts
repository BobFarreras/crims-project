import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

// Utilitat per combinar classes de Tailwind sense conflictes
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}
