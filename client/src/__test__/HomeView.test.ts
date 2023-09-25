import { render, screen } from "@testing-library/vue";
import HomeView from "../views/HomeView.vue";
import { describe, it, expect } from "vitest";
import userEvent from "@testing-library/user-event";
import router from "../router";

describe("HomeView", () => {
  it("should have an input with testid 'search-input'", () => {
    render(HomeView);
    const input = screen.getByTestId("search-input");
    expect(input).not.toBeNull();
  });

  it("should have an button with testid 'search-button'", () => {
    render(HomeView);
    const input = screen.getByTestId("search-button");
    expect(input).not.toBeNull();
  });

  it("should change the URL when typing into the input and click the button", async () => {
    render(HomeView, {
      global: {
        plugins: [router],
      },
    });

    const input = screen.getByTestId("search-input");
    const button = screen.getByTestId("search-button");

    await userEvent.type(input, "john");
    await userEvent.click(button);

    await router.isReady();

    expect(router.currentRoute.value.path).toContain(
      `/search/${encodeURIComponent("john")}/page/1`
    );
  });
});
