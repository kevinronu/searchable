import { render, screen } from "@testing-library/vue";
import HomeView from "../views/HomeView.vue";
import { describe, it, expect } from "vitest";
import userEvent from "@testing-library/user-event";
import router from "../router";

describe("HomeView", () => {
  beforeEach(async () => {
    router.push("/");
    await router.isReady();
  });

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

  it("should doesn't change the URL when the input are empty and click the button", async () => {
    render(HomeView, {
      global: {
        plugins: [router],
      },
    });

    const push = vi.spyOn(router, "push");

    const button = screen.getByTestId("search-button");

    await userEvent.click(button);

    await router.isReady();

    expect(push).not.toHaveBeenCalled();

    expect(router.currentRoute.value.path).toBe(`/`);
  });

  it("should change the URL when typing into the input and click the button", async () => {
    render(HomeView, {
      global: {
        plugins: [router],
      },
    });

    const push = vi.spyOn(router, "push");

    const input = screen.getByTestId("search-input");
    const button = screen.getByTestId("search-button");

    await userEvent.type(input, "John Doe");
    await userEvent.click(button);

    await router.isReady();

    expect(push).toHaveBeenCalledWith({
      path: `/search/${encodeURIComponent("John Doe")}/page/1`,
    });

    expect(router.currentRoute.value.path).toBe(
      `/search/${encodeURIComponent("John Doe")}/page/1`
    );
  });
});
