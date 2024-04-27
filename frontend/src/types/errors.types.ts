import { AxiosError } from "axios";

type CustomAxiosErrorResponse = {
  error: string;
};

export type GoError = AxiosError<CustomAxiosErrorResponse>;
