export type SearchResult = {
  total_found: number;
  documents: Document[];
};

export type Document = {
  id: string;
  subject: string;
  from: string;
  to: string[];
  cc: string[];
  body: string;
  date: Date;
};

type Pagination = {
  from: number;
  size: number;
};

type DateRange = {
  from: Date | null;
  to: Date | null;
};

export type Body = {
  query: string;
  sort: string;
  pagination: Pagination;
  date_range: DateRange;
};

export type Filters = {
  sort: string;
  from: string | null;
  to: string | null;
};

type HttpMethod =
  | "GET"
  | "POST"
  | "PUT"
  | "PATCH"
  | "DELETE"
  | "OPTIONS"
  | "HEAD";

export type RequestOptions = {
  method?: HttpMethod;
  headers?: Record<string, string>;
  body?: Body;
};
