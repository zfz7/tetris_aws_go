import { render, screen } from "@testing-library/react";
import { expect, test, describe } from "vitest";
import App from "./App";
import { act } from "react";

describe("App Component", () => {
  test("renders authenticated user", async () => {
    render(<App />);
    await act(async () => {
      await Promise.resolve();
    });
    expect(screen.getByText(/welcome back, john doe!/i)).toBeInTheDocument();
    expect(screen.getByText(/sign out/i)).toBeInTheDocument();
    expect(screen.getByText(/tetris template/i)).toBeInTheDocument();
    expect(screen.getByText(/server communication/i)).toBeInTheDocument();
  });
});
