import apiFetch from "./api-fetch";
import { Body } from "../models/document.model";
const indexName = import.meta.env.VITE_INDEX_NAME;

export async function getDocument(documentId: string) {
  return await apiFetch(`/v1/${indexName}/${documentId}`);
}

export async function searchDocument(query: string, currentPage: number) {
  const body: Body = {
    query,
    sort: "-date",
    pagination: {
      from: 20 * (currentPage - 1),
      size: 20,
    },
  };
  return await apiFetch(`/v1/${indexName}/search`, {
    body: JSON.stringify(body),
  });
}

export async function deleteDocument(documentId: string) {
  return await apiFetch(`/v1/${indexName}/${documentId}`, { method: "DELETE" });
}
