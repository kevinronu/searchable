import { render, screen } from "@testing-library/vue";
import HomeView from "../views/HomeView.vue";
import { describe, it, expect } from "vitest";

describe("HomeView", () => {
  it("should have an input with testid 'search-input'", () => {
    render(HomeView);
    const input = screen.getByTestId("search-input");
    expect(input).not.toBeNull();
  });
});
