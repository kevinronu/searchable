const indexName = import.meta.env.VITE_INDEX_NAME;
import { Body, Filters } from "../models/document.model";
import apiFetch from "./api-fetch";

export async function getDocument(documentId: string) {
  return await apiFetch(`/v1/${indexName}/${documentId}`);
}

export async function searchDocument(
  query: string,
  currentPage: number,
  filters?: Filters
) {
  const { sort, from, to } = filters || { sort: "-date", from: null, to: null };

  let dateFrom = null;
  let dateTo = null;
  if (from) {
    dateFrom = new Date(from);
  }
  if (to) {
    dateTo = new Date(to);
  }

  const body: Body = {
    query,
    sort,
    pagination: {
      from: 20 * (currentPage - 1),
      size: 20,
    },
    date_range: {
      from: dateFrom,
      to: dateTo,
    },
  };
  return await apiFetch(`/v1/${indexName}/search`, {
    body: JSON.stringify(body),
  });
}

export async function deleteDocument(documentId: string) {
  return await apiFetch(`/v1/${indexName}/${documentId}`, { method: "DELETE" });
}
